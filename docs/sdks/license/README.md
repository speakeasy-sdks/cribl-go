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
    res, err := s.License.Create(ctx, shared.License{
        Cls: shared.LicenseClsTrial,
        Email: "Levi.Hyatt@yahoo.com",
        Exp: 221319,
        FPh: 344010,
        FPhg: 184738,
        GUID: "tenetur",
        Iat: 477352,
        ID: "4533994d-78de-43b6-a938-9f5abb7f6625",
        Iss: "ullam",
        License: "doloremque",
        Limits: map[string]interface{}{
            "qui": "praesentium",
            "adipisci": "totam",
            "qui": "deserunt",
        },
        Quota: 754784,
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
    res, err := s.License.Delete(ctx, operations.DeleteLicenseRequest{
        ID: "83afd231-5bba-4650-964e-06f5bf6ae591",
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
    res, err := s.License.Get(ctx, operations.GetLicenseRequest{
        ID: "bc8bdef3-612b-463c-a05f-da840774a68a",
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

