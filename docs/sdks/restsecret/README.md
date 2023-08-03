# RestSecret

### Available Operations

* [Create](#create) - Create RestSecret
* [Delete](#delete) - Delete RestSecret
* [Get](#get) - Get RestSecret by ID
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
    res, err := s.RestSecret.Create(ctx, shared.RestSecret{
        APIKey: cribl.String("qui"),
        Description: cribl.String("accusantium"),
        ID: "688f77c1-ffc7-41dc-a163-f2a3c80a97ff",
        Password: cribl.String("velit"),
        SecretKey: cribl.String("adipisci"),
        SecretType: shared.SecretTypeText,
        Tags: cribl.String("optio"),
        Username: cribl.String("Rossie_Strosin37"),
        Value: cribl.String("esse"),
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
    id := "laborum"

    ctx := context.Background()
    res, err := s.RestSecret.Delete(ctx, id)
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
    id := "perspiciatis"

    ctx := context.Background()
    res, err := s.RestSecret.Get(ctx, id)
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
    id := "voluptates"
    restSecret := &shared.RestSecret{
        APIKey: cribl.String("eum"),
        Description: cribl.String("quasi"),
        ID: "876c6ab2-1d29-4dfc-94d6-fecd79939006",
        Password: cribl.String("laboriosam"),
        SecretKey: cribl.String("laborum"),
        SecretType: shared.SecretTypeKeypair,
        Tags: cribl.String("assumenda"),
        Username: cribl.String("Brayan3"),
        Value: cribl.String("alias"),
    }

    ctx := context.Background()
    res, err := s.RestSecret.Update(ctx, id, restSecret)
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

