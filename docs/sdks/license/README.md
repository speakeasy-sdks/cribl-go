# License

### Available Operations

* [Create](#create) - Create License
* [Delete](#delete) - Delete License
* [Get](#get) - Get License by ID

## Create

Create License

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
    res, err := s.License.Create(ctx, shared.License{
        Cls: shared.LicenseClsFree,
        Email: "Lurline.Welch@yahoo.com",
        Exp: 370195,
        FPh: 729612,
        FPhg: 153942,
        GUID: "omnis",
        Iat: 120646,
        ID: "22030d83-f5ae-4b77-99d2-2e8c1f849382",
        Iss: "nemo",
        License: "delectus",
        Limits: map[string]interface{}{
            "porro": "quaerat",
            "magni": "cumque",
            "quos": "in",
            "commodi": "maxime",
        },
        Quota: 151486,
        Title: "Miss",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.License != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.License](../../models/shared/license.md)      | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateLicenseResponse](../../models/operations/createlicenseresponse.md), error**


## Delete

Delete License

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
    res, err := s.License.Delete(ctx, operations.DeleteLicenseRequest{
        ID: "2dfb4cfc-1c76-4230-b841-fb1bd23fdb14",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.License != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteLicenseRequest](../../models/operations/deletelicenserequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.DeleteLicenseResponse](../../models/operations/deletelicenseresponse.md), error**


## Get

Get License by ID

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
    res, err := s.License.Get(ctx, operations.GetLicenseRequest{
        ID: "db6be5a6-8599-48e2-aae2-0da16fc2b271",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.License != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetLicenseRequest](../../models/operations/getlicenserequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetLicenseResponse](../../models/operations/getlicenseresponse.md), error**

