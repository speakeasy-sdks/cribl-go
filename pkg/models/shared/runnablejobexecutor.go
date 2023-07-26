// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type RunnableJobExecutorExecutorExecutorSpecificSettings struct {
}

type RunnableJobExecutorExecutor struct {
	Conf *RunnableJobExecutorExecutorExecutorSpecificSettings `json:"conf,omitempty"`
	// Determines whether or not to write task results to disk.
	StoreTaskResults *bool `json:"storeTaskResults,omitempty"`
	// The type of executor to run.
	Type string `json:"type"`
}

func (o *RunnableJobExecutorExecutor) GetConf() *RunnableJobExecutorExecutorExecutorSpecificSettings {
	if o == nil {
		return nil
	}
	return o.Conf
}

func (o *RunnableJobExecutorExecutor) GetStoreTaskResults() *bool {
	if o == nil {
		return nil
	}
	return o.StoreTaskResults
}

func (o *RunnableJobExecutorExecutor) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

// RunnableJobExecutorRunLogLevel - Level at which to set task logging.
type RunnableJobExecutorRunLogLevel string

const (
	RunnableJobExecutorRunLogLevelError RunnableJobExecutorRunLogLevel = "error"
	RunnableJobExecutorRunLogLevelWarn  RunnableJobExecutorRunLogLevel = "warn"
	RunnableJobExecutorRunLogLevelInfo  RunnableJobExecutorRunLogLevel = "info"
	RunnableJobExecutorRunLogLevelDebug RunnableJobExecutorRunLogLevel = "debug"
	RunnableJobExecutorRunLogLevelSilly RunnableJobExecutorRunLogLevel = "silly"
)

func (e RunnableJobExecutorRunLogLevel) ToPointer() *RunnableJobExecutorRunLogLevel {
	return &e
}

func (e *RunnableJobExecutorRunLogLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "error":
		fallthrough
	case "warn":
		fallthrough
	case "info":
		fallthrough
	case "debug":
		fallthrough
	case "silly":
		*e = RunnableJobExecutorRunLogLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RunnableJobExecutorRunLogLevel: %v", v)
	}
}

type RunnableJobExecutorRun struct {
	// Maximum time the job is allowed to run (e.g., 30, 45s or 15m). Units are seconds, if not specified. Enter 0 for unlimited time.
	JobTimeout *string `json:"jobTimeout,omitempty"`
	// Level at which to set task logging.
	LogLevel *RunnableJobExecutorRunLogLevel `json:"logLevel,omitempty"`
	// Max number of times a task can be rescheduled.
	MaxTaskReschedule *int64 `json:"maxTaskReschedule,omitempty"`
	// Reschedule tasks that failed with non-fatal errors.
	RescheduleDroppedTasks *bool `json:"rescheduleDroppedTasks,omitempty"`
}

func (o *RunnableJobExecutorRun) GetJobTimeout() *string {
	if o == nil {
		return nil
	}
	return o.JobTimeout
}

func (o *RunnableJobExecutorRun) GetLogLevel() *RunnableJobExecutorRunLogLevel {
	if o == nil {
		return nil
	}
	return o.LogLevel
}

func (o *RunnableJobExecutorRun) GetMaxTaskReschedule() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxTaskReschedule
}

func (o *RunnableJobExecutorRun) GetRescheduleDroppedTasks() *bool {
	if o == nil {
		return nil
	}
	return o.RescheduleDroppedTasks
}

// RunnableJobExecutorScheduleRunSettingsLogLevel - Level at which to set task logging.
type RunnableJobExecutorScheduleRunSettingsLogLevel string

const (
	RunnableJobExecutorScheduleRunSettingsLogLevelError RunnableJobExecutorScheduleRunSettingsLogLevel = "error"
	RunnableJobExecutorScheduleRunSettingsLogLevelWarn  RunnableJobExecutorScheduleRunSettingsLogLevel = "warn"
	RunnableJobExecutorScheduleRunSettingsLogLevelInfo  RunnableJobExecutorScheduleRunSettingsLogLevel = "info"
	RunnableJobExecutorScheduleRunSettingsLogLevelDebug RunnableJobExecutorScheduleRunSettingsLogLevel = "debug"
	RunnableJobExecutorScheduleRunSettingsLogLevelSilly RunnableJobExecutorScheduleRunSettingsLogLevel = "silly"
)

func (e RunnableJobExecutorScheduleRunSettingsLogLevel) ToPointer() *RunnableJobExecutorScheduleRunSettingsLogLevel {
	return &e
}

func (e *RunnableJobExecutorScheduleRunSettingsLogLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "error":
		fallthrough
	case "warn":
		fallthrough
	case "info":
		fallthrough
	case "debug":
		fallthrough
	case "silly":
		*e = RunnableJobExecutorScheduleRunSettingsLogLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RunnableJobExecutorScheduleRunSettingsLogLevel: %v", v)
	}
}

type RunnableJobExecutorScheduleRunSettings struct {
	// Earliest time, for the given Range Timezone.
	Earliest *int64 `json:"earliest,omitempty"`
	// A filter for tokens in the provided collect path and/or the events being collected
	Expression *string `json:"expression,omitempty"`
	// Maximum time the job is allowed to run (e.g., 30, 45s or 15m). Units are seconds, if not specified. Enter 0 for unlimited time.
	JobTimeout *string `json:"jobTimeout,omitempty"`
	// Latest time, for the given Range Timezone.
	Latest *int64 `json:"latest,omitempty"`
	// Level at which to set task logging.
	LogLevel *RunnableJobExecutorScheduleRunSettingsLogLevel `json:"logLevel,omitempty"`
	// Max number of times a task can be rescheduled.
	MaxTaskReschedule *int64 `json:"maxTaskReschedule,omitempty"`
	// Limits the bundle size for files above the Lower task bundle size. E.g., bundle five 2MB files into one 10MB task bundle. Files greater than this size will be assigned to individual tasks.
	MaxTaskSize *string `json:"maxTaskSize,omitempty"`
	// Limits the bundle size for small tasks. E.g., bundle five 200KB files into one 1M task.
	MinTaskSize *string `json:"minTaskSize,omitempty"`
	// Job run mode. Preview will either return up to N matching results, or will run until capture time T is reached. Discovery will gather the list of files to turn into streaming tasks, without running the data collection job. Full Run will run the collection job.
	Mode *string `json:"mode,omitempty"`
	// Reschedule tasks that failed with non-fatal errors.
	RescheduleDroppedTasks *bool `json:"rescheduleDroppedTasks,omitempty"`
	// Time range for scheduled job.
	TimeRangeType *string `json:"timeRangeType,omitempty"`
	// Timezone to use for Earliest and Latest times (defaults to UTC).
	TimestampTimezone *string `json:"timestampTimezone,omitempty"`
}

func (o *RunnableJobExecutorScheduleRunSettings) GetEarliest() *int64 {
	if o == nil {
		return nil
	}
	return o.Earliest
}

func (o *RunnableJobExecutorScheduleRunSettings) GetExpression() *string {
	if o == nil {
		return nil
	}
	return o.Expression
}

func (o *RunnableJobExecutorScheduleRunSettings) GetJobTimeout() *string {
	if o == nil {
		return nil
	}
	return o.JobTimeout
}

func (o *RunnableJobExecutorScheduleRunSettings) GetLatest() *int64 {
	if o == nil {
		return nil
	}
	return o.Latest
}

func (o *RunnableJobExecutorScheduleRunSettings) GetLogLevel() *RunnableJobExecutorScheduleRunSettingsLogLevel {
	if o == nil {
		return nil
	}
	return o.LogLevel
}

func (o *RunnableJobExecutorScheduleRunSettings) GetMaxTaskReschedule() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxTaskReschedule
}

func (o *RunnableJobExecutorScheduleRunSettings) GetMaxTaskSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxTaskSize
}

func (o *RunnableJobExecutorScheduleRunSettings) GetMinTaskSize() *string {
	if o == nil {
		return nil
	}
	return o.MinTaskSize
}

func (o *RunnableJobExecutorScheduleRunSettings) GetMode() *string {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *RunnableJobExecutorScheduleRunSettings) GetRescheduleDroppedTasks() *bool {
	if o == nil {
		return nil
	}
	return o.RescheduleDroppedTasks
}

func (o *RunnableJobExecutorScheduleRunSettings) GetTimeRangeType() *string {
	if o == nil {
		return nil
	}
	return o.TimeRangeType
}

func (o *RunnableJobExecutorScheduleRunSettings) GetTimestampTimezone() *string {
	if o == nil {
		return nil
	}
	return o.TimestampTimezone
}

// RunnableJobExecutorSchedule - Configuration for a scheduled job.
type RunnableJobExecutorSchedule struct {
	// A cron schedule on which to run this job.
	CronSchedule *string `json:"cronSchedule,omitempty"`
	// Determines whether or not this schedule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The maximum number of instances that may be running of this scheduled job at any given time.
	MaxConcurrentRuns *int64                                  `json:"maxConcurrentRuns,omitempty"`
	ResumeMissed      interface{}                             `json:"resumeMissed,omitempty"`
	Run               *RunnableJobExecutorScheduleRunSettings `json:"run,omitempty"`
	// Skippable jobs can be delayed, up to their next run time, if the system is hitting concurrency limits.
	Skippable *bool `json:"skippable,omitempty"`
}

func (o *RunnableJobExecutorSchedule) GetCronSchedule() *string {
	if o == nil {
		return nil
	}
	return o.CronSchedule
}

func (o *RunnableJobExecutorSchedule) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *RunnableJobExecutorSchedule) GetMaxConcurrentRuns() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxConcurrentRuns
}

func (o *RunnableJobExecutorSchedule) GetResumeMissed() interface{} {
	if o == nil {
		return nil
	}
	return o.ResumeMissed
}

func (o *RunnableJobExecutorSchedule) GetRun() *RunnableJobExecutorScheduleRunSettings {
	if o == nil {
		return nil
	}
	return o.Run
}

func (o *RunnableJobExecutorSchedule) GetSkippable() *bool {
	if o == nil {
		return nil
	}
	return o.Skippable
}

// RunnableJobExecutorJobType - Job type.
type RunnableJobExecutorJobType string

const (
	RunnableJobExecutorJobTypeCollection      RunnableJobExecutorJobType = "collection"
	RunnableJobExecutorJobTypeExecutor        RunnableJobExecutorJobType = "executor"
	RunnableJobExecutorJobTypeScheduledSearch RunnableJobExecutorJobType = "scheduledSearch"
)

func (e RunnableJobExecutorJobType) ToPointer() *RunnableJobExecutorJobType {
	return &e
}

func (e *RunnableJobExecutorJobType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "collection":
		fallthrough
	case "executor":
		fallthrough
	case "scheduledSearch":
		*e = RunnableJobExecutorJobType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RunnableJobExecutorJobType: %v", v)
	}
}

type RunnableJobExecutor struct {
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string                     `json:"environment,omitempty"`
	Executor    RunnableJobExecutorExecutor `json:"executor"`
	// Unique ID for this Job.
	ID *string `json:"id,omitempty"`
	// List of fields to remove from Discover results. Wildcards (e.g.: aws*) are allowed. This is useful when discovery returns sensitive fields that should not be exposed in the Jobs user interface.
	RemoveFields []string `json:"removeFields,omitempty"`
	// Resumes the ad hoc job if a failure condition causes Stream to restart during job execution.
	ResumeOnBoot *bool                  `json:"resumeOnBoot,omitempty"`
	Run          RunnableJobExecutorRun `json:"run"`
	// Configuration for a scheduled job.
	Schedule *RunnableJobExecutorSchedule `json:"schedule,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Time to keep the job's artifacts on disk after job completion. This also affects how long a job is listed in the Job Inspector.
	TTL *string `json:"ttl,omitempty"`
	// Job type.
	Type *RunnableJobExecutorJobType `json:"type,omitempty"`
}

func (o *RunnableJobExecutor) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *RunnableJobExecutor) GetExecutor() RunnableJobExecutorExecutor {
	if o == nil {
		return RunnableJobExecutorExecutor{}
	}
	return o.Executor
}

func (o *RunnableJobExecutor) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RunnableJobExecutor) GetRemoveFields() []string {
	if o == nil {
		return nil
	}
	return o.RemoveFields
}

func (o *RunnableJobExecutor) GetResumeOnBoot() *bool {
	if o == nil {
		return nil
	}
	return o.ResumeOnBoot
}

func (o *RunnableJobExecutor) GetRun() RunnableJobExecutorRun {
	if o == nil {
		return RunnableJobExecutorRun{}
	}
	return o.Run
}

func (o *RunnableJobExecutor) GetSchedule() *RunnableJobExecutorSchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}

func (o *RunnableJobExecutor) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *RunnableJobExecutor) GetTTL() *string {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *RunnableJobExecutor) GetType() *RunnableJobExecutorJobType {
	if o == nil {
		return nil
	}
	return o.Type
}
