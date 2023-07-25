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
    res, err := s.Certificate.Create(ctx, shared.Certificate{
        Ca: cribl.String("aliquam"),
        Cert: "ad",
        Description: cribl.String("repellat"),
        ID: "0597a60f-f2a5-44a3-9e94-764a3e865e79",
        InUse: []string{
            "eum",
            "reiciendis",
        },
        Passphrase: cribl.String("provident"),
        PrivKey: "aspernatur",
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
    res, err := s.Certificate.Delete(ctx, operations.DeleteCertificateRequest{
        ID: "51a5a9da-660f-4f57-bfaa-d4f9efc1b451",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteCertificateRequest](../../models/operations/deletecertificaterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


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
    res, err := s.Certificate.Get(ctx, operations.GetCertificateRequest{
        ID: "2c103264-8dc2-4f61-9199-ebfd0e9fe6c6",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetCertificateRequest](../../models/operations/getcertificaterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
    res, err := s.Certificate.Update(ctx, operations.UpdateCertificateRequest{
        Certificate: &shared.Certificate{
            Ca: cribl.String("dolorem"),
            Cert: "fugit",
            Description: cribl.String("cumque"),
            ID: "a3aed011-7996-4312-bde0-4771778ff61d",
            InUse: []string{
                "dicta",
            },
            Passphrase: cribl.String("odio"),
            PrivKey: "tempora",
        },
        ID: "76360a15-db6a-4660-a59a-1adeaab5851d",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateCertificateRequest](../../models/operations/updatecertificaterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateCertificateResponse](../../models/operations/updatecertificateresponse.md), error**

