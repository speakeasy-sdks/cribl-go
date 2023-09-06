# Secrets

### Available Operations

* [Create](#create) - Create RestSecret
* [Delete](#delete) - Delete RestSecret
* [Get](#get) - Get RestSecret by ID
* [ListRestSecrets](#listrestsecrets) - Get a list of RestSecret objects
* [Update](#update) - Update RestSecret

## Create

Create RestSecret

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
    res, err := s.Secrets.Create(ctx, shared.RestSecret{
        APIKey: cribl.String("magni"),
        Description: cribl.String("nihil"),
        ID: "5a60a04c-495c-4c69-9171-b51c1bdb1cf4",
        Password: cribl.String("expedita"),
        SecretKey: cribl.String("corrupti"),
        SecretType: shared.SecretTypeKeypair,
        Tags: cribl.String("atque"),
        Username: cribl.String("Sheldon_Ritchie77"),
        Value: cribl.String("incidunt"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RestSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                              | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ctx`                                                  | [context.Context](https://pkg.go.dev/context#Context)  | :heavy_check_mark:                                     | The context to use for the request.                    |
| `request`                                              | [shared.RestSecret](../../models/shared/restsecret.md) | :heavy_check_mark:                                     | The request object to use for the request.             |


### Response

**[*operations.CreateRestSecretResponse](../../models/operations/createrestsecretresponse.md), error**


## Delete

Delete RestSecret

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )
    id := "quod"

    ctx := context.Background()
    res, err := s.Secrets.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.RestSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


### Response

**[*operations.DeleteRestSecretResponse](../../models/operations/deleterestsecretresponse.md), error**


## Get

Get RestSecret by ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )
    id := "minus"

    ctx := context.Background()
    res, err := s.Secrets.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.RestSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


### Response

**[*operations.GetRestSecretResponse](../../models/operations/getrestsecretresponse.md), error**


## ListRestSecrets

Get a list of RestSecret objects

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
    res, err := s.Secrets.ListRestSecrets(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.RestSecrets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListRestSecretsResponse](../../models/operations/listrestsecretsresponse.md), error**


## Update

Update RestSecret

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )
    id := "porro"
    restSecret := &shared.RestSecret{
        APIKey: cribl.String("id"),
        Description: cribl.String("excepturi"),
        ID: "9bc7fc0b-2dce-4108-b3e4-2b006d678878",
        Password: cribl.String("rerum"),
        SecretKey: cribl.String("deserunt"),
        SecretType: shared.SecretTypeKeypair,
        Tags: cribl.String("nostrum"),
        Username: cribl.String("Karen.Bradtke52"),
        Value: cribl.String("magni"),
    }

    ctx := context.Background()
    res, err := s.Secrets.Update(ctx, id, restSecret)
    if err != nil {
        log.Fatal(err)
    }

    if res.RestSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                               | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ctx`                                                   | [context.Context](https://pkg.go.dev/context#Context)   | :heavy_check_mark:                                      | The context to use for the request.                     |
| `id`                                                    | *string*                                                | :heavy_check_mark:                                      | Unique ID                                               |
| `restSecret`                                            | [*shared.RestSecret](../../models/shared/restsecret.md) | :heavy_minus_sign:                                      | RestSecret object to be updated                         |


### Response

**[*operations.UpdateRestSecretResponse](../../models/operations/updaterestsecretresponse.md), error**

