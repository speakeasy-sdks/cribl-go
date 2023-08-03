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
        APIKey: cribl.String("dolore"),
        Description: cribl.String("nam"),
        ID: "af91e506-ef89-40a5-8b47-5f16f56d385a",
        Password: cribl.String("sequi"),
        SecretKey: cribl.String("maxime"),
        SecretType: shared.SecretTypeText,
        Tags: cribl.String("laborum"),
        Username: cribl.String("Niko12"),
        Value: cribl.String("rerum"),
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
    id := "occaecati"

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
    id := "provident"

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
    id := "necessitatibus"
    restSecret := &shared.RestSecret{
        APIKey: cribl.String("fugit"),
        Description: cribl.String("autem"),
        ID: "ced8f9fd-b941-40f6-bbbf-817837b01afd",
        Password: cribl.String("quibusdam"),
        SecretKey: cribl.String("quam"),
        SecretType: shared.SecretTypeKeypair,
        Tags: cribl.String("rem"),
        Username: cribl.String("Herman12"),
        Value: cribl.String("blanditiis"),
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

