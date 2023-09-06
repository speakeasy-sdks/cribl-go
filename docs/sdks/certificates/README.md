# Certificates

## Overview

Actions related to certificates

### Available Operations

* [Create](#create) - Create Certificate
* [Delete](#delete) - Delete Certificate
* [Get](#get) - Get Certificate by ID
* [ListCertificates](#listcertificates) - Get a list of Certificate objects
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
    res, err := s.Certificates.Create(ctx, shared.Certificate{
        Ca: cribl.String("incidunt"),
        Cert: "qui",
        Description: cribl.String("cupiditate"),
        ID: "cdb1a842-2bb6-479d-a322-715bf0cbb1e3",
        InUse: []string{
            "veritatis",
        },
        Passphrase: cribl.String("nobis"),
        PrivKey: "quos",
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
    id := "tempore"

    ctx := context.Background()
    res, err := s.Certificates.Delete(ctx, id)
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
    id := "cupiditate"

    ctx := context.Background()
    res, err := s.Certificates.Get(ctx, id)
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


## ListCertificates

Get a list of Certificate objects

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
    res, err := s.Certificates.ListCertificates(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Certificates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListCertificatesResponse](../../models/operations/listcertificatesresponse.md), error**


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
    id := "aperiam"
    certificate := &shared.Certificate{
        Ca: cribl.String("delectus"),
        Cert: "dolorem",
        Description: cribl.String("dolore"),
        ID: "43a1108e-0adc-4f4b-9218-79fce953f73e",
        InUse: []string{
            "tenetur",
        },
        Passphrase: cribl.String("dignissimos"),
        PrivKey: "hic",
    }

    ctx := context.Background()
    res, err := s.Certificates.Update(ctx, id, certificate)
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

