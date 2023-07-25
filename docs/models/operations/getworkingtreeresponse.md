# GetWorkingTreeResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ContentType`                                                       | *string*                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `GitStatusResults`                                                  | [*shared.GitStatusResults](../../models/shared/gitstatusresults.md) | :heavy_minus_sign:                                                  | a list of GitStatusResult objects                                   |
| `StatusCode`                                                        | *int*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `RawResponse`                                                       | [*http.Response](https://pkg.go.dev/net/http#Response)              | :heavy_minus_sign:                                                  | N/A                                                                 |