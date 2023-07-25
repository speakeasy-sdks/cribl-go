# GetAuthorizationsResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `AuthPolicyEntries`                                                   | [*shared.AuthPolicyEntries](../../models/shared/authpolicyentries.md) | :heavy_minus_sign:                                                    | a list of AuthPolicyEntry objects                                     |
| `ContentType`                                                         | *string*                                                              | :heavy_check_mark:                                                    | N/A                                                                   |
| `StatusCode`                                                          | *int*                                                                 | :heavy_check_mark:                                                    | N/A                                                                   |
| `RawResponse`                                                         | [*http.Response](https://pkg.go.dev/net/http#Response)                | :heavy_minus_sign:                                                    | N/A                                                                   |