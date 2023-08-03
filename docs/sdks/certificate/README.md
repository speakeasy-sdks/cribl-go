# Certificate

### Available Operations

* [Create](#create) - Create Certificate
* [Delete](#delete) - Delete Certificate
* [Get](#get) - Get Certificate by ID
* [Update](#update) - Update Certificate

## Create

Create Certificate

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
    res, err := s.Certificate.Create(ctx, shared.Certificate{
        Ca: cribl.String("eveniet"),
        Cert: "occaecati",
        Description: cribl.String("consequuntur"),
        ID: "2a57a15b-e3e0-4608-87e2-b6e3ab8845f0",
        InUse: []string{
            "perspiciatis",
            "nihil",
        },
        Passphrase: cribl.String("mollitia"),
        PrivKey: "voluptas",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Certificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [shared.Certificate](../../models/shared/certificate.md) | :heavy_check_mark:                                       | The request object to use for the request.               |


### Response

**[*operations.CreateCertificateResponse](../../models/operations/createcertificateresponse.md), error**


## Delete

Delete Certificate

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
    id := "alias"

    ctx := context.Background()
    res, err := s.Certificate.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Certificate != nil {
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

**[*operations.DeleteCertificateResponse](../../models/operations/deletecertificateresponse.md), error**


## Get

Get Certificate by ID

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
    id := "maiores"

    ctx := context.Background()
    res, err := s.Certificate.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Certificate != nil {
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

**[*operations.GetCertificateResponse](../../models/operations/getcertificateresponse.md), error**


## Update

Update Certificate

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
    id := "reiciendis"
    certificate := &shared.Certificate{
        Ca: cribl.String("dolores"),
        Cert: "id",
        Description: cribl.String("minima"),
        ID: "4a31e947-64a3-4e86-9e79-56f9251a5a9d",
        InUse: []string{
            "ex",
            "aliquid",
            "accusantium",
        },
        Passphrase: cribl.String("repellat"),
        PrivKey: "doloribus",
    }

    ctx := context.Background()
    res, err := s.Certificate.Update(ctx, id, certificate)
    if err != nil {
        log.Fatal(err)
    }

    if res.Certificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |
| `id`                                                      | *string*                                                  | :heavy_check_mark:                                        | Unique ID                                                 |
| `certificate`                                             | [*shared.Certificate](../../models/shared/certificate.md) | :heavy_minus_sign:                                        | Certificate object to be updated                          |


### Response

**[*operations.UpdateCertificateResponse](../../models/operations/updatecertificateresponse.md), error**

