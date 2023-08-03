# ListLogFileListRequest


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Allow`                                                                 | **string*                                                               | :heavy_minus_sign:                                                      | query array[string] optional List of allowed-filename wildcard patterns |
| `Depth`                                                                 | **int64*                                                                | :heavy_minus_sign:                                                      | Search depth for "manual" mode                                          |
| `Mode`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | Discovery Mode (default is "auto")                                      |
| `Path`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | Search directory for "manual" mode                                      |