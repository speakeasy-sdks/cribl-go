# Security

### Available Operations

* [GetKMSConfig](#getkmsconfig) - Get Cribl KMS config
* [GetKMSHealth](#getkmshealth) - Get Cribl KMS health
* [Update](#update) - Update Cribl KMS config

## GetKMSConfig

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
    res, err := s.Security.GetKMSConfig(ctx)
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


## GetKMSHealth

Get Cribl KMS health

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
    res, err := s.Security.GetKMSHealth(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.KMSHealth != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetKMSHealthResponse](../../models/operations/getkmshealthresponse.md), error**


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
    res, err := s.Security.Update(ctx, shared.IKMSProviderConfig{
        EnableHealthCheck: false,
        Engine: &shared.VaultKMSEngineConfig{
            Mount: cribl.String("quae"),
            SecretPath: cribl.String("quas"),
            Type: shared.VaultKMSEngineConfigTypeKv2,
        },
        HealthCheckEndpoint: cribl.String("placeat"),
        Namespace: cribl.String("enim"),
        Provider: shared.SecretProviderLocal,
        SecretDir: cribl.String("sapiente"),
        Service: &shared.IAWSKMSServiceConfig{
            KmsKeyArn: "saepe",
            Region: "delectus",
        },
        URL: cribl.String("officia"),
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

