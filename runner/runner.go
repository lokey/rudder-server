package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"

	"github.com/rudderlabs/rudder-server/info"
	"github.com/rudderlabs/rudder-server/warehouse/datalake"

	"github.com/bugsnag/bugsnag-go/v2"
	_ "go.uber.org/automaxprocs"
	"golang.org/x/sync/errgroup"

	"github.com/rudderlabs/rudder-server/admin"
	"github.com/rudderlabs/rudder-server/admin/profiler"
	"github.com/rudderlabs/rudder-server/app"
	"github.com/rudderlabs/rudder-server/app/apphandlers"
	"github.com/rudderlabs/rudder-server/config"
	backendconfig "github.com/rudderlabs/rudder-server/config/backend-config"
	eventschema "github.com/rudderlabs/rudder-server/event-schema"
	"github.com/rudderlabs/rudder-server/gateway"
	"github.com/rudderlabs/rudder-server/gateway/webhook"
	"github.com/rudderlabs/rudder-server/jobsdb"
	"github.com/rudderlabs/rudder-server/processor"
	"github.com/rudderlabs/rudder-server/processor/integrations"
	"github.com/rudderlabs/rudder-server/processor/stash"
	"github.com/rudderlabs/rudder-server/processor/transformer"
	ratelimiter "github.com/rudderlabs/rudder-server/rate-limiter"
	"github.com/rudderlabs/rudder-server/router"
	"github.com/rudderlabs/rudder-server/router/batchrouter"
	"github.com/rudderlabs/rudder-server/router/batchrouter/asyncdestinationmanager"
	"github.com/rudderlabs/rudder-server/router/customdestinationmanager"
	oauth "github.com/rudderlabs/rudder-server/router/oauthResponseHandler"
	routertransformer "github.com/rudderlabs/rudder-server/router/transformer"
	batchrouterutils "github.com/rudderlabs/rudder-server/router/utils"
	"github.com/rudderlabs/rudder-server/services/alert"
	"github.com/rudderlabs/rudder-server/services/archiver"
	"github.com/rudderlabs/rudder-server/services/controlplane/features"
	"github.com/rudderlabs/rudder-server/services/db"
	destinationdebugger "github.com/rudderlabs/rudder-server/services/debugger/destination"
	sourcedebugger "github.com/rudderlabs/rudder-server/services/debugger/source"
	transformationdebugger "github.com/rudderlabs/rudder-server/services/debugger/transformation"
	"github.com/rudderlabs/rudder-server/services/dedup"
	destinationconnectiontester "github.com/rudderlabs/rudder-server/services/destination-connection-tester"
	"github.com/rudderlabs/rudder-server/services/diagnostics"
	"github.com/rudderlabs/rudder-server/services/multitenant"
	"github.com/rudderlabs/rudder-server/services/pgnotifier"
	"github.com/rudderlabs/rudder-server/services/stats"
	"github.com/rudderlabs/rudder-server/services/streammanager/kafka"
	"github.com/rudderlabs/rudder-server/utils/logger"
	"github.com/rudderlabs/rudder-server/utils/misc"
	"github.com/rudderlabs/rudder-server/utils/types/deployment"
	"github.com/rudderlabs/rudder-server/warehouse"
	azuresynapse "github.com/rudderlabs/rudder-server/warehouse/azure-synapse"
	"github.com/rudderlabs/rudder-server/warehouse/bigquery"
	"github.com/rudderlabs/rudder-server/warehouse/clickhouse"
	"github.com/rudderlabs/rudder-server/warehouse/deltalake"
	"github.com/rudderlabs/rudder-server/warehouse/mssql"
	"github.com/rudderlabs/rudder-server/warehouse/postgres"
	"github.com/rudderlabs/rudder-server/warehouse/redshift"
	"github.com/rudderlabs/rudder-server/warehouse/snowflake"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
	"github.com/rudderlabs/rudder-server/warehouse/validations"
)

// ReleaseInfo holds the release information
type ReleaseInfo struct {
	Version         string
	Major           string
	Minor           string
	Patch           string
	Commit          string
	BuildDate       string
	BuiltBy         string
	GitURL          string
	EnterpriseToken string
}

// Runner is responsible for running the application
type Runner struct {
	releaseInfo               ReleaseInfo
	application               app.App
	warehouseMode             string
	enableSuppressUserFeature bool
	logger                    logger.Logger
	appHandler                apphandlers.AppHandler
	readTimeout               time.Duration
	readHeaderTimeout         time.Duration
	writeTimeout              time.Duration
	idleTimeout               time.Duration
	gracefulShutdownTimeout   time.Duration
	maxHeaderBytes            int
}

// New creates and initializes a new Runner
func New(releaseInfo ReleaseInfo) *Runner {
	getConfigDuration := func(defaultValueInTimescaleUnits int64, timeScale time.Duration, keys ...string) time.Duration {
		for i, key := range keys {
			if config.IsSet(key) || i == len(keys)-1 {
				return config.GetDuration(key, defaultValueInTimescaleUnits, timeScale)
			}
		}
		return 0
	}
	return &Runner{
		releaseInfo:               releaseInfo,
		logger:                    logger.NewLogger().Child("runner"),
		warehouseMode:             config.GetString("Warehouse.mode", "embedded"),
		enableSuppressUserFeature: config.GetBool("Gateway.enableSuppressUserFeature", true),
		readTimeout:               getConfigDuration(0, time.Second, "ReadTimeOut", "ReadTimeOutInSec"),
		readHeaderTimeout:         getConfigDuration(0, time.Second, "ReadHeaderTimeout", "ReadHeaderTimeoutInSec"),
		writeTimeout:              getConfigDuration(10, time.Second, "WriteTimeout", "WriteTimeOutInSec"),
		idleTimeout:               getConfigDuration(720, time.Second, "IdleTimeout", "IdleTimeoutInSec"),
		gracefulShutdownTimeout:   config.GetDuration("GracefulShutdownTimeout", 15, time.Second),
		maxHeaderBytes:            config.GetInt("MaxHeaderBytes", 524288),
	}
}

// Run runs the application and returns the exit code
func (r *Runner) Run(ctx context.Context, args []string) int {
	runAllInit()

	options := app.LoadOptions(args)
	if options.VersionFlag {
		r.printVersion()
		return 0
	}

	options.EnterpriseToken = r.releaseInfo.EnterpriseToken

	r.application = app.New(options)

	// application & backend setup should be done before starting any new goroutines.
	r.application.Setup()

	appTypeStr := strings.ToUpper(config.GetString("APP_TYPE", app.EMBEDDED))
	r.appHandler = apphandlers.GetAppHandler(r.application, appTypeStr, r.versionHandler)

	versionDetail := r.versionInfo()
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       config.GetString("BUGSNAG_KEY", ""),
		ReleaseStage: config.GetString("GO_ENV", "development"),
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{"main", "github.com/rudderlabs/rudder-server"},
		// more configuration options
		AppType:      r.appHandler.GetAppType(),
		AppVersion:   versionDetail["Version"].(string),
		PanicHandler: func() {},
	})
	ctx = bugsnag.StartSession(ctx)
	defer misc.BugsnagNotify(ctx, "Core")()

	deploymentType, err := deployment.GetFromEnv()
	if err != nil {
		r.logger.Errorf("failed to get deployment type: %w", err)
		return 1
	}
	// TODO: remove as soon as we update the configuration with statsExcludedTags where necessary
	if !config.IsSet("statsExcludedTags") && deploymentType == deployment.MultiTenantType && (!config.IsSet("WORKSPACE_NAMESPACE") || strings.Contains(config.GetString("WORKSPACE_NAMESPACE", ""), "free")) {
		config.Set("statsExcludedTags", []string{"workspaceId", "sourceID", "destId"})
	}
	stats.Default.Start(ctx)
	stats.Default.NewTaggedStat("rudder_server_config",
		stats.GaugeType,
		stats.Tags{
			"version":            r.releaseInfo.Version,
			"major":              r.releaseInfo.Major,
			"minor":              r.releaseInfo.Minor,
			"patch":              r.releaseInfo.Patch,
			"commit":             r.releaseInfo.Commit,
			"buildDate":          r.releaseInfo.BuildDate,
			"builtBy":            r.releaseInfo.BuiltBy,
			"gitUrl":             r.releaseInfo.GitURL,
			"TransformerVersion": transformer.GetVersion(),
			"DatabricksVersion":  misc.GetDatabricksVersion(),
		}).Gauge(1)

	configEnvHandler := r.application.Features().ConfigEnv.Setup()

	if config.GetString("Warehouse.mode", "") != "slave" {
		if err := backendconfig.Setup(configEnvHandler); err != nil {
			r.logger.Errorf("Unable to setup backend config: %s", err)
			return 1
		}

		backendconfig.DefaultBackendConfig.StartWithIDs(ctx, "")
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		if err := admin.StartServer(ctx); err != nil {
			return fmt.Errorf("admin server routine: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		p := &profiler.Profiler{}
		if err := p.StartServer(ctx); err != nil {
			return fmt.Errorf("profiler server routine: %w", err)
		}
		return nil
	})

	misc.AppStartTime = time.Now().Unix()
	if r.canStartServer() {
		r.appHandler.HandleRecovery(options)
		g.Go(misc.WithBugsnag(func() (err error) {
			if err := r.appHandler.StartRudderCore(ctx, options); err != nil {
				return fmt.Errorf("rudder core: %w", err)
			}
			return nil
		}))
		g.Go(misc.WithBugsnag(func() error {
			backendconfig.DefaultBackendConfig.WaitForConfig(ctx)

			c := features.NewClient(
				config.GetString("CONFIG_BACKEND_URL", "https://api.rudderlabs.com"),
				backendconfig.DefaultBackendConfig.Identity(),
			)

			err := c.Send(ctx, info.ServerComponent.Name, info.ServerComponent.Features)
			if err != nil {
				r.logger.Errorf("error sending server features: %v", err)
			}

			// we don't want to exit if we can't send server features
			return nil
		}))
	}

	// initialize warehouse service after core to handle non-normal recovery modes
	if appTypeStr != app.GATEWAY && r.canStartWarehouse() {
		g.Go(misc.WithBugsnagForWarehouse(func() error {
			if err := r.startWarehouseService(ctx, r.application); err != nil {
				return fmt.Errorf("warehouse service routine: %w", err)
			}
			return nil
		}))
	}

	shutdownDone := make(chan struct{})
	go func() {
		err := g.Wait()
		if err != nil {
			r.logger.Errorf("Terminal error: %v", err)
		}

		r.logger.Info("Attempting to shutdown gracefully")
		backendconfig.DefaultBackendConfig.Stop()
		close(shutdownDone)
	}()

	<-ctx.Done()
	ctxDoneTime := time.Now()

	select {
	case <-shutdownDone:
		r.application.Stop()
		r.logger.Infof(
			"Graceful terminal after %s, with %d go-routines",
			time.Since(ctxDoneTime),
			runtime.NumGoroutine(),
		)
		// clearing zap Log buffer to std output
		logger.Sync()
		stats.Default.Stop()
	case <-time.After(r.gracefulShutdownTimeout):
		// Assume graceful shutdown failed, log remain goroutines and force kill
		r.logger.Errorf(
			"Graceful termination failed after %s, goroutine dump:\n",
			time.Since(ctxDoneTime),
		)

		fmt.Print("\n\n")
		_ = pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
		fmt.Print("\n\n")

		r.application.Stop()
		logger.Sync()
		stats.Default.Stop()
		if config.GetBool("RUDDER_GRACEFUL_SHUTDOWN_TIMEOUT_EXIT", true) {
			return 1
		}
	}

	return 0
}

func runAllInit() {
	admin.Init()
	misc.Init()
	db.Init()
	diagnostics.Init()
	backendconfig.Init()
	warehouseutils.Init()
	bigquery.Init()
	clickhouse.Init()
	archiver.Init()
	destinationdebugger.Init()
	pgnotifier.Init()
	jobsdb.Init()
	jobsdb.Init2()
	jobsdb.Init3()
	destinationconnectiontester.Init()
	warehouse.Init()
	warehouse.Init2()
	warehouse.Init3()
	warehouse.Init4()
	warehouse.Init5()
	warehouse.Init6()
	validations.Init()
	datalake.Init()
	azuresynapse.Init()
	mssql.Init()
	postgres.Init()
	redshift.Init()
	snowflake.Init()
	deltalake.Init()
	transformer.Init()
	webhook.Init()
	batchrouter.Init()
	batchrouter.Init2()
	asyncdestinationmanager.Init()
	batchrouterutils.Init()
	dedup.Init()
	eventschema.Init()
	eventschema.Init2()
	stash.Init()
	transformationdebugger.Init()
	processor.Init()
	kafka.Init()
	customdestinationmanager.Init()
	routertransformer.Init()
	router.Init()
	router.InitRouterAdmin()
	ratelimiter.Init()
	sourcedebugger.Init()
	gateway.Init()
	apphandlers.Init()
	apphandlers.Init2()
	integrations.Init()
	alert.Init()
	multitenant.Init()
	oauth.Init()
}

func (r *Runner) versionInfo() map[string]interface{} {
	return map[string]interface{}{
		"Version":            r.releaseInfo.Version,
		"Major":              r.releaseInfo.Major,
		"Minor":              r.releaseInfo.Minor,
		"Patch":              r.releaseInfo.Patch,
		"Commit":             r.releaseInfo.Commit,
		"BuildDate":          r.releaseInfo.BuildDate,
		"BuiltBy":            r.releaseInfo.BuiltBy,
		"GitUrl":             r.releaseInfo.GitURL,
		"TransformerVersion": transformer.GetVersion(),
		"DatabricksVersion":  misc.GetDatabricksVersion(),
	}
}

func (r *Runner) versionHandler(w http.ResponseWriter, _ *http.Request) {
	version := r.versionInfo()
	versionFormatted, _ := json.Marshal(&version)
	_, _ = w.Write(versionFormatted)
}

func (r *Runner) printVersion() {
	version := r.versionInfo()
	versionFormatted, _ := json.MarshalIndent(&version, "", " ")
	fmt.Printf("Version Info %s\n", versionFormatted)
}

func (*Runner) startWarehouseService(ctx context.Context, application app.App) error {
	return warehouse.Start(ctx, application)
}

func (r *Runner) canStartServer() bool {
	r.logger.Info("warehousemode ", r.warehouseMode)
	return r.warehouseMode == config.EmbeddedMode || r.warehouseMode == config.OffMode || r.warehouseMode == config.PooledWHSlaveMode
}

func (r *Runner) canStartWarehouse() bool {
	return r.warehouseMode != config.OffMode
}
