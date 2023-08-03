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
        Cls: shared.LicenseClsTrial,
        Email: "Michele52@yahoo.com",
        Exp: 978154,
        FPh: 35160,
        FPhg: 331452,
        GUID: "saepe",
        Iat: 206063,
        ID: "d48fdaf3-13a1-4f5f-9942-59c0b36f25ea",
        Iss: "sint",
        License: "ut",
        Limits: map[string]interface{}{
            "tenetur": "adipisci",
            "libero": "in",
        },
        Quota: 329651,
        Title: "Ms.",
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
    id := "minus"

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
    id := "ab"

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

