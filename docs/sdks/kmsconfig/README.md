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
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
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
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
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
            Mount: cribl.String("modi"),
            SecretPath: cribl.String("itaque"),
            Type: shared.VaultKMSEngineConfigTypeKv2,
        },
        HealthCheckEndpoint: cribl.String("maxime"),
        Namespace: cribl.String("modi"),
        Provider: shared.SecretProviderLocal,
        SecretDir: cribl.String("assumenda"),
        Service: &shared.IAWSKMSServiceConfig{
            KmsKeyArn: "vero",
            Region: "doloribus",
        },
        URL: cribl.String("impedit"),
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

