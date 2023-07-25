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
    res, err := s.RestSecret.Create(ctx, shared.RestSecret{
        APIKey: cribl.String("doloribus"),
        Description: cribl.String("magnam"),
        ID: "39e39266-cbd9-45f7-aa2b-24113695d1e6",
        Password: cribl.String("aliquid"),
        SecretKey: cribl.String("excepturi"),
        SecretType: shared.SecretTypeKeypair,
        Tags: cribl.String("repellat"),
        Username: cribl.String("Raina33"),
        Value: cribl.String("perspiciatis"),
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
	"cribl"
	"cribl/pkg/models/shared"
	"cribl/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RestSecret.Delete(ctx, operations.DeleteRestSecretRequest{
        ID: "6217c297-7676-4334-a540-38bfb5971e98",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteRestSecretRequest](../../models/operations/deleterestsecretrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


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
	"cribl"
	"cribl/pkg/models/shared"
	"cribl/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RestSecret.Get(ctx, operations.GetRestSecretRequest{
        ID: "19055738-9ced-4bac-bfda-39594d66bc2a",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetRestSecretRequest](../../models/operations/getrestsecretrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


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
	"cribl"
	"cribl/pkg/models/shared"
	"cribl/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RestSecret.Update(ctx, operations.UpdateRestSecretRequest{
        RestSecret: &shared.RestSecret{
            APIKey: cribl.String("officiis"),
            Description: cribl.String("aliquam"),
            ID: "80632b99-54b6-4fa2-a063-69828553cb10",
            Password: cribl.String("accusantium"),
            SecretKey: cribl.String("doloremque"),
            SecretType: shared.SecretTypeKeypair,
            Tags: cribl.String("tempore"),
            Username: cribl.String("Sophie59"),
            Value: cribl.String("dolores"),
        },
        ID: "1ec2053b-7493-466a-88ee-0f2bf19588d4",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateRestSecretRequest](../../models/operations/updaterestsecretrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateRestSecretResponse](../../models/operations/updaterestsecretresponse.md), error**

