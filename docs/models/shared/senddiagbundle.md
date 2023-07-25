# SendDiagBundle

SendDiagBundle object


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `IncludeMetrics`                                                        | **bool*                                                                 | :heavy_minus_sign:                                                      | Toggle to No to exclude metrics from the diag bundle.                   |
| `MaxIncludeJobs`                                                        | **int64*                                                                | :heavy_minus_sign:                                                      | Number of jobs including all tasks that will be included in the bundle. |
| `Path`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | Existing diag bundle that will be send to Cribl Support. Max 100MB.     |
| `RenameJs`                                                              | **bool*                                                                 | :heavy_minus_sign:                                                      | Append .txt to JavaScript files.                                        |
| `SendToCribl`                                                           | **bool*                                                                 | :heavy_minus_sign:                                                      | Send diag bundle to Cribl Support                                       |