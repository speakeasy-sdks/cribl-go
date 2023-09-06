# Licenses

## Overview

Actions related to licenses. The <code>/licenses</code> endpoints do not apply to Cribl.Cloud deployments.

### Available Operations

* [Create](#create) - Create License
* [Delete](#delete) - Delete License
* [Get](#get) - Get License by ID
* [ListLicenseUsageMetrics](#listlicenseusagemetrics) - Get license usage metrics, aggregated by day, up to last 90 days
* [ListLicenses](#listlicenses) - Get a list of License objects

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
    res, err := s.Licenses.Create(ctx, shared.License{
        Cls: shared.LicenseClsProd,
        Email: "Davion.Toy@yahoo.com",
        Exp: 367927,
        FPh: 928219,
        FPhg: 456520,
        GUID: "provident",
        Iat: 337477,
        ID: "6f9251a5-a9da-4660-bf57-bfaad4f9efc1",
        Iss: "rerum",
        License: "tempora",
        Limits: map[string]interface{}{
            "quis": "inventore",
        },
        Quota: 147685,
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
    id := "quae"

    ctx := context.Background()
    res, err := s.Licenses.Delete(ctx, id)
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
    id := "perferendis"

    ctx := context.Background()
    res, err := s.Licenses.Get(ctx, id)
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


## ListLicenseUsageMetrics

Get license usage metrics, aggregated by day, up to last 90 days

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
    res, err := s.Licenses.ListLicenseUsageMetrics(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.DailyMetrics != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListLicenseUsageMetricsResponse](../../models/operations/listlicenseusagemetricsresponse.md), error**


## ListLicenses

Get a list of License objects

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
    res, err := s.Licenses.ListLicenses(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Licenses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListLicensesResponse](../../models/operations/listlicensesresponse.md), error**

