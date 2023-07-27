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
        APIKey: cribl.String("a"),
        Description: cribl.String("consequatur"),
        ID: "ea2216cb-e071-4bc1-a3e2-79a3b084da99",
        Password: cribl.String("consequuntur"),
        SecretKey: cribl.String("veniam"),
        SecretType: shared.SecretTypeKeypair,
        Tags: cribl.String("repellendus"),
        Username: cribl.String("Amely.Gorczany5"),
        Value: cribl.String("quas"),
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

    ctx := context.Background()
    res, err := s.RestSecret.Delete(ctx, operations.DeleteRestSecretRequest{
        ID: "47a742d8-4496-4cbd-aecf-6b99bc63562e",
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

    ctx := context.Background()
    res, err := s.RestSecret.Get(ctx, operations.GetRestSecretRequest{
        ID: "bfdf55c2-94c0-460b-86a1-287764eef6d0",
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

    ctx := context.Background()
    res, err := s.RestSecret.Update(ctx, operations.UpdateRestSecretRequest{
        RestSecret: &shared.RestSecret{
            APIKey: cribl.String("porro"),
            Description: cribl.String("suscipit"),
            ID: "d6ed9c73-dd63-4457-9509-a8e870d3c5a1",
            Password: cribl.String("maiores"),
            SecretKey: cribl.String("perspiciatis"),
            SecretType: shared.SecretTypeCredentials,
            Tags: cribl.String("magni"),
            Username: cribl.String("Dorothy.Considine69"),
            Value: cribl.String("iure"),
        },
        ID: "6a1f30c7-3df5-4b67-9989-0f42a4bb438d",
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

