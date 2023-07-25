# GetJobInfosRequest


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `CollectorID`                                                 | **string*                                                     | :heavy_minus_sign:                                            | Filter by collector id, e.g. "collectorId=Prometheus-in"      |
| `ID`                                                          | **string*                                                     | :heavy_minus_sign:                                            | Filter by job id, e.g. "id=1608713335.3&id=1608713326.1"      |
| `Limit`                                                       | **int64*                                                      | :heavy_minus_sign:                                            | Maximum number of items to return                             |
| `Offset`                                                      | **int64*                                                      | :heavy_minus_sign:                                            | Pagination offset                                             |
| `RunType`                                                     | **string*                                                     | :heavy_minus_sign:                                            | Filter by job run type, one of "adhoc", "scheduled", "system" |
| `State`                                                       | **string*                                                     | :heavy_minus_sign:                                            | Filter by current job state, e.g. "running"                   |