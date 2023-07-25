# PostInternalSystemMetricsResponse


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ContentType`                                                     | *string*                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `MetricsResponse`                                                 | [*shared.MetricsResponse](../../models/shared/metricsresponse.md) | :heavy_minus_sign:                                                | a list of any objects                                             |
| `StatusCode`                                                      | *int*                                                             | :heavy_check_mark:                                                | N/A                                                               |
| `RawResponse`                                                     | [*http.Response](https://pkg.go.dev/net/http#Response)            | :heavy_minus_sign:                                                | N/A                                                               |