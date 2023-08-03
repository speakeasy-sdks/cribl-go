# Diag

## Overview

Actions related to diagnostics

### Available Operations

* [Delete](#delete) - Remove diag bundle
* [Get](#get) - Returns a diag bundle as a tar.gz archive
* [GetSystemInfo](#getsysteminfo) - Get basic system information
* [ListExistingDiagBundles](#listexistingdiagbundles) - Get list of existing diag bundles
* [Send](#send) - Send a diag bundle (tar.gz archive) to Cribl

## Delete

Remove diag bundle

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
    path := "officiis"

    ctx := context.Background()
    res, err := s.Diag.Delete(ctx, path)
    if err != nil {
        log.Fatal(err)
    }

    if res.RemoveDiagResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `path`                                                | *string*                                              | :heavy_check_mark:                                    | Diag bundle path                                      |


### Response

**[*operations.DeleteDiagBundleResponse](../../models/operations/deletediagbundleresponse.md), error**


## Get

Returns a diag bundle as a tar.gz archive

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
    res, err := s.Diag.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDiagBundle200ApplicationTarPlusGzipBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetDiagBundleResponse](../../models/operations/getdiagbundleresponse.md), error**


## GetSystemInfo

Get basic system information

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
    res, err := s.Diag.GetSystemInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.SystemInfoObjects != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetSystemInfoResponse](../../models/operations/getsysteminforesponse.md), error**


## ListExistingDiagBundles

Get list of existing diag bundles

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
    res, err := s.Diag.ListExistingDiagBundles(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ExistingDiag != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListExistingDiagBundlesResponse](../../models/operations/listexistingdiagbundlesresponse.md), error**


## Send

Send a diag bundle (tar.gz archive) to Cribl

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
    res, err := s.Diag.Send(ctx, shared.SendDiagBundle{
        IncludeMetrics: cribl.Bool(false),
        MaxIncludeJobs: cribl.Int64(103298),
        Path: cribl.String("fuga"),
        RenameJs: cribl.Bool(false),
        SendToCribl: cribl.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RemoveDiagResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [shared.SendDiagBundle](../../models/shared/senddiagbundle.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |


### Response

**[*operations.SendDiagBundleResponse](../../models/operations/senddiagbundleresponse.md), error**

