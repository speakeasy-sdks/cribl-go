# ListLogFilesContentsRequest


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Et`                                                                          | **int64*                                                                      | :heavy_minus_sign:                                                            | Epoch timestamp of the earliest event (includes rolled files present on disk) |
| `Files`                                                                       | **string*                                                                     | :heavy_minus_sign:                                                            | query string[] optional file or files to query                                |
| `Filter`                                                                      | **string*                                                                     | :heavy_minus_sign:                                                            | Filter                                                                        |
| `GroupID`                                                                     | **string*                                                                     | :heavy_minus_sign:                                                            | id of the group to query                                                      |
| `Limit`                                                                       | **int64*                                                                      | :heavy_minus_sign:                                                            | Maximum number of log lines to retrieve starting from offset.                 |
| `Lt`                                                                          | **int64*                                                                      | :heavy_minus_sign:                                                            | Epoch timestamp of the latest event (includes rolled files present on disk)   |
| `Type`                                                                        | *string*                                                                      | :heavy_check_mark:                                                            | type of logs request single multi group                                       |