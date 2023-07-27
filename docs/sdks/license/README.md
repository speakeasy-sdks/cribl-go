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
        Cls: shared.LicenseClsProd,
        Email: "Joy33@gmail.com",
        Exp: 906232,
        FPh: 206063,
        FPhg: 816365,
        GUID: "aliquam",
        Iat: 552581,
        ID: "fdaf313a-1f5f-4d94-a59c-0b36f25ea944",
        Iss: "tenetur",
        License: "adipisci",
        Limits: map[string]interface{}{
            "in": "minima",
            "ex": "minus",
            "ab": "beatae",
        },
        Quota: 943143,
        Title: "Mrs.",
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
    id := "quisquam"

    ctx := context.Background()
    res, err := s.License.Delete(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


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
    id := "dolor"

    ctx := context.Background()
    res, err := s.License.Get(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


### Response

**[*operations.GetLicenseResponse](../../models/operations/getlicenseresponse.md), error**

