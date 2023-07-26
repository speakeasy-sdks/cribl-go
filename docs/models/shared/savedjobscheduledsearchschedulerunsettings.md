# SavedJobScheduledSearchScheduleRunSettings


## Fields

| Field                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Earliest`                                                                                                                                                                                                                                                           | **int64*                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                   | Earliest time, for the given Range Timezone.                                                                                                                                                                                                                         |
| `Expression`                                                                                                                                                                                                                                                         | **string*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                   | A filter for tokens in the provided collect path and/or the events being collected                                                                                                                                                                                   |
| `JobTimeout`                                                                                                                                                                                                                                                         | **string*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                   | Maximum time the job is allowed to run (e.g., 30, 45s or 15m). Units are seconds, if not specified. Enter 0 for unlimited time.                                                                                                                                      |
| `Latest`                                                                                                                                                                                                                                                             | **int64*                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                   | Latest time, for the given Range Timezone.                                                                                                                                                                                                                           |
| `LogLevel`                                                                                                                                                                                                                                                           | [*SavedJobScheduledSearchScheduleRunSettingsLogLevel](../../models/shared/savedjobscheduledsearchschedulerunsettingsloglevel.md)                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                   | Level at which to set task logging.                                                                                                                                                                                                                                  |
| `MaxTaskReschedule`                                                                                                                                                                                                                                                  | **int64*                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                   | Max number of times a task can be rescheduled.                                                                                                                                                                                                                       |
| `MaxTaskSize`                                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                   | Limits the bundle size for files above the Lower task bundle size. E.g., bundle five 2MB files into one 10MB task bundle. Files greater than this size will be assigned to individual tasks.                                                                         |
| `MinTaskSize`                                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                   | Limits the bundle size for small tasks. E.g., bundle five 200KB files into one 1M task.                                                                                                                                                                              |
| `Mode`                                                                                                                                                                                                                                                               | **string*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                   | Job run mode. Preview will either return up to N matching results, or will run until capture time T is reached. Discovery will gather the list of files to turn into streaming tasks, without running the data collection job. Full Run will run the collection job. |
| `RescheduleDroppedTasks`                                                                                                                                                                                                                                             | **bool*                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                   | Reschedule tasks that failed with non-fatal errors.                                                                                                                                                                                                                  |
| `TimeRangeType`                                                                                                                                                                                                                                                      | **string*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                   | Time range for scheduled job.                                                                                                                                                                                                                                        |
| `TimestampTimezone`                                                                                                                                                                                                                                                  | **string*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                   | Timezone to use for Earliest and Latest times (defaults to UTC).                                                                                                                                                                                                     |