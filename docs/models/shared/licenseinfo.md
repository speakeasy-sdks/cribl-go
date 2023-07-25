# LicenseInfo


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `EffectiveClass`                                          | [LicenseType](../../models/shared/licensetype.md)         | :heavy_check_mark:                                        | N/A                                                       |
| `Email`                                                   | **string*                                                 | :heavy_minus_sign:                                        | N/A                                                       |
| `Expiration`                                              | *int64*                                                   | :heavy_check_mark:                                        | N/A                                                       |
| `GUID`                                                    | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       |
| `IsRunningInSaaS`                                         | *bool*                                                    | :heavy_check_mark:                                        | N/A                                                       |
| `IsSplunkApp`                                             | **bool*                                                   | :heavy_minus_sign:                                        | N/A                                                       |
| `LicenseEnforceReason`                                    | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       |
| `Limits`                                                  | map[string]*interface{}*                                  | :heavy_check_mark:                                        | N/A                                                       |
| `PhoneHome`                                               | *bool*                                                    | :heavy_check_mark:                                        | N/A                                                       |
| `PhoneHomeGrace`                                          | *int64*                                                   | :heavy_check_mark:                                        | N/A                                                       |
| `Quota`                                                   | *int64*                                                   | :heavy_check_mark:                                        | N/A                                                       |
| `Type`                                                    | [LicenseInfoType](../../models/shared/licenseinfotype.md) | :heavy_check_mark:                                        | N/A                                                       |