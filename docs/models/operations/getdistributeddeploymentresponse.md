# GetDistributedDeploymentResponse


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ContentType`                                                               | *string*                                                                    | :heavy_check_mark:                                                          | N/A                                                                         |
| `DistributedSummaries`                                                      | [*shared.DistributedSummaries](../../models/shared/distributedsummaries.md) | :heavy_minus_sign:                                                          | a list of DistributedSummary objects                                        |
| `StatusCode`                                                                | *int*                                                                       | :heavy_check_mark:                                                          | N/A                                                                         |
| `RawResponse`                                                               | [*http.Response](https://pkg.go.dev/net/http#Response)                      | :heavy_minus_sign:                                                          | N/A                                                                         |