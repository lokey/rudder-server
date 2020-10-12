package warehouse

import (
	"fmt"
	"strings"
	"time"

	"github.com/rudderlabs/rudder-server/services/stats"
	"github.com/rudderlabs/rudder-server/utils/logger"
	"github.com/rudderlabs/rudder-server/utils/misc"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
)

func (job *UploadJobT) timerStat(name string) stats.RudderStats {
	return warehouseutils.DestStat(stats.TimerType, name, job.warehouse.Identifier)
}

func (job *UploadJobT) counterStat(name string) stats.RudderStats {
	return warehouseutils.DestStat(stats.CountType, name, job.warehouse.Identifier)
}

func (jobRun *JobRunT) timerStat(name string) stats.RudderStats {
	return warehouseutils.DestStat(stats.TimerType, name, jobRun.whIdentifier)
}

func (jobRun *JobRunT) counterStat(name string) stats.RudderStats {
	return warehouseutils.DestStat(stats.CountType, name, jobRun.whIdentifier)
}

func (job *UploadJobT) generateUploadSuccessMetrics() {
	// Total loaded events in the upload
	numUploadedEvents, err := getTotalEventsUploaded(job.upload.ID)
	if err != nil {
		logger.Errorf("[WH]: Failed to generate load metrics: %s", job.warehouse.Identifier)
		return
	}
	job.counterStat("num_loaded_events").Count(int(numUploadedEvents))

	// Total staged events in the upload
	numStagedEvents, err := getTotalEventsStaged(job.upload.StartStagingFileID, job.upload.EndStagingFileID)
	if err != nil {
		logger.Errorf("[WH]: Failed to generate stage metrics: %s", job.warehouse.Identifier)
		return
	}
	job.counterStat("num_staged_events").Count(int(numStagedEvents))

	// Delay for the oldest event in the batch
	firstEventAt, err := getFirstStagedEventAt(job.upload.StartStagingFileID)
	if err != nil {
		logger.Errorf("[WH]: Failed to generate delay metrics: %s", job.warehouse.Identifier)
		return
	}
	loadDelayInS := time.Now().Unix() - firstEventAt.Local().Unix()
	job.counterStat("load_delay_in_sec").Count(int(loadDelayInS))

	job.counterStat("upload_success").Count(1)
}

func (job *UploadJobT) generateUploadAbortedMetrics() {
	// Total staged events in the upload
	numStagedEvents, err := getTotalEventsStaged(job.upload.StartStagingFileID, job.upload.EndStagingFileID)
	if err != nil {
		logger.Errorf("[WH]: Failed to generate stage metrics: %s", job.warehouse.Identifier)
		return
	}
	job.counterStat("num_staged_events").Count(int(numStagedEvents))

	job.counterStat("upload_aborted").Count(1)
}

func (job *UploadJobT) recordTableLoad(tableName string, numEvents int64) {
	// add metric to record total loaded rows to standard tables
	// adding metric for all event tables might result in too many metrics
	tablesToRecordEventsMetric := []string{"tracks", "users", "identifies", "pages", "screens", "aliases", "groups", "rudder_discards"}
	if misc.Contains(tablesToRecordEventsMetric, strings.ToLower(tableName)) {
		// TODO: Add table name as tag
		job.counterStat(fmt.Sprintf(`%s_table_records_uploaded`, strings.ToLower(tableName))).Count(int(numEvents))
	}
}
