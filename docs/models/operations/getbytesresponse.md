# GetBytesResponse


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ContentType`                                             | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       |
| `SampleFiles`                                             | [*shared.SampleFiles](../../models/shared/samplefiles.md) | :heavy_minus_sign:                                        | a list of SampleFile objects                              |
| `StatusCode`                                              | *int*                                                     | :heavy_check_mark:                                        | N/A                                                       |
| `RawResponse`                                             | [*http.Response](https://pkg.go.dev/net/http#Response)    | :heavy_minus_sign:                                        | N/A                                                       |