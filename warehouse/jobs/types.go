package jobs

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/rudderlabs/rudder-server/services/pgnotifier"
	"github.com/rudderlabs/rudder-server/utils/logger"
)

var pkgLogger logger.Logger

// StartJobReqPayload For processing requests payload in handlers.go
type StartJobReqPayload struct {
	SourceID      string `json:"source_id"`
	Type          string `json:"type"`
	Channel       string `json:"channel"`
	DestinationID string `json:"destination_id"`
	StartTime     string `json:"start_time"`
	JobRunID      string `json:"job_run_id"`
	TaskRunID     string `json:"task_run_id"`
	AsyncJobType  string `json:"async_job_type"`
}

type AsyncJobWhT struct {
	dbHandle   *sql.DB
	enabled    bool
	pgnotifier *pgnotifier.PgNotifierT
	context    context.Context
}

type WhJobsMetaData struct {
	JobRunID  string `json:"job_run_id"`
	TaskRunID string `json:"task_run_id"`
	JobType   string `json:"jobtype"`
	StartTime string `json:"start_time"`
}

// AsyncJobPayloadT For creating job payload to wh_async_jobs table
type AsyncJobPayloadT struct {
	Id            string          `json:"id"`
	SourceID      string          `json:"source_id"`
	DestinationID string          `json:"destination_id"`
	TableName     string          `json:"tablename"`
	AsyncJobType  string          `json:"async_job_type"`
	MetaData      json.RawMessage `json:"metadata"`
}

const (
	WhJobWaiting   string = "waiting"
	WhJobExecuting string = "executing"
	WhJobSucceeded string = "succeeded"
	WhJobAborted   string = "aborted"
	WhJobFailed    string = "failed"
	AsyncJobType   string = "async_job"
)

type PGNotifierOutput struct {
	Id string `json:"id"`
}

type WhAddJobResponse struct {
	JobIds []int64 `json:"jobids"`
	Err    error   `json:"error"`
}

type WhStatusResponse struct {
	Status string
	Err    string
}

type WhAsyncJobRunner interface {
	startAsyncJobRunner(context.Context)
	getTableNamesBy(context.Context, string, string)
	getPendingAsyncJobs(context.Context) ([]AsyncJobPayloadT, error)
	getStatusAsyncJob(*StartJobReqPayload) (string, error)
	updateMultipleAsyncJobs(*[]AsyncJobPayloadT, string, string)
}

type AsyncJobStatus struct {
	Id     string
	Status string
	Error  error
}

const (
	MaxBatchSizeToProcess int = 10
	MaxCleanUpRetries     int = 5
	MaxQueryRetries       int = 3
	RetryTimeInterval         = 10 * time.Second
	MaxAttemptsPerJob     int = 3
	WhAsyncJobTimeOut         = 10 * time.Second
)
