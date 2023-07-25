# IKMSProviderConfig

IKMSProviderConfig object


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `EnableHealthCheck`                                                  | *bool*                                                               | :heavy_check_mark:                                                   | N/A                                                                  |
| `Engine`                                                             | [*VaultKMSEngineConfig](../../models/shared/vaultkmsengineconfig.md) | :heavy_minus_sign:                                                   | N/A                                                                  |
| `HealthCheckEndpoint`                                                | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `Namespace`                                                          | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `Provider`                                                           | [SecretProvider](../../models/shared/secretprovider.md)              | :heavy_check_mark:                                                   | N/A                                                                  |
| `SecretDir`                                                          | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `Service`                                                            | [*IAWSKMSServiceConfig](../../models/shared/iawskmsserviceconfig.md) | :heavy_minus_sign:                                                   | N/A                                                                  |
| `URL`                                                                | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |