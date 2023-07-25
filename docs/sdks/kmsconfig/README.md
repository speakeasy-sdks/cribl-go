# KMSConfig

### Available Operations

* [Get](#get) - Get Cribl KMS config
* [Update](#update) - Update Cribl KMS config

## Get

Get Cribl KMS config

### Example Usage

```go
package main

import(
	"context"
	"log"
	"cribl"
	"cribl/pkg/models/shared"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.KMSConfig.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.KMSConfigs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetKMSConfigResponse](../../models/operations/getkmsconfigresponse.md), error**


## Update

Update Cribl KMS config

### Example Usage

```go
package main

import(
	"context"
	"log"
	"cribl"
	"cribl/pkg/models/shared"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.KMSConfig.Update(ctx, shared.IKMSProviderConfig{
        EnableHealthCheck: false,
        Engine: &shared.VaultKMSEngineConfig{
            Mount: cribl.String("sit"),
            SecretPath: cribl.String("voluptatum"),
            Type: shared.VaultKMSEngineConfigTypeKv2,
        },
        HealthCheckEndpoint: cribl.String("tempora"),
        Namespace: cribl.String("delectus"),
        Provider: shared.SecretProviderVault,
        SecretDir: cribl.String("dolorum"),
        Service: &shared.IAWSKMSServiceConfig{
            KmsKeyArn: "libero",
            Region: "ratione",
        },
        URL: cribl.String("molestiae"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IKMSProviderConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.IKMSProviderConfig](../../models/shared/ikmsproviderconfig.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.UpdateKMSConfigResponse](../../models/operations/updatekmsconfigresponse.md), error**

