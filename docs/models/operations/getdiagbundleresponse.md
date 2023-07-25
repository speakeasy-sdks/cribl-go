# GetDiagBundleResponse


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ContentType`                                                    | *string*                                                         | :heavy_check_mark:                                               | N/A                                                              |
| `StatusCode`                                                     | *int*                                                            | :heavy_check_mark:                                               | N/A                                                              |
| `RawResponse`                                                    | [*http.Response](https://pkg.go.dev/net/http#Response)           | :heavy_minus_sign:                                               | N/A                                                              |
| `GetDiagBundle200ApplicationTarPlusGzipBinaryString`             | *[]byte*                                                         | :heavy_minus_sign:                                               | A tar.gz file consisting all configuration files and recent logs |