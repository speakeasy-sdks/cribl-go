# Packs

### Available Operations

* [Clone](#clone) - Clone Pack
* [Export](#export) - Export Pack
* [Install](#install) - Install Pack
* [ListPacks](#listpacks) - Get info on packs
* [Uninstall](#uninstall) - Uninstall Pack from the system
* [Upgrade](#upgrade) - Upgrade Pack
* [Upload](#upload) - Upload Pack

## Clone

Clone Pack

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
    res, err := s.Packs.Clone(ctx, shared.PackClone{
        Dest: cribl.String("officia"),
        DstGroups: []string{
            "cumque",
            "natus",
            "quaerat",
        },
        Force: cribl.Bool(false),
        Packs: []string{
            "quia",
            "officiis",
            "mollitia",
            "cumque",
        },
        SrcGroup: "quis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PackInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.PackClone](../../models/shared/packclone.md)  | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.ClonePackResponse](../../models/operations/clonepackresponse.md), error**


## Export

Export Pack

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
    id := "enim"
    mode := "eum"
    filename := "nemo"

    ctx := context.Background()
    res, err := s.Packs.Export(ctx, id, mode, filename)
    if err != nil {
        log.Fatal(err)
    }

    if res.PackInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |
| `id`                                                      | *string*                                                  | :heavy_check_mark:                                        | Pack name                                                 |
| `mode`                                                    | *string*                                                  | :heavy_check_mark:                                        | Export mode, one of "merge_safe", "merge", "default_only" |
| `filename`                                                | **string*                                                 | :heavy_minus_sign:                                        | Filename of the exported Pack                             |


### Response

**[*operations.ExportPackResponse](../../models/operations/exportpackresponse.md), error**


## Install

Install Pack

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
    res, err := s.Packs.Install(ctx, shared.CrudEntityBase{
        ID: "d307cfee-8120-46e2-813f-a4a41c480d3f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PackInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [shared.CrudEntityBase](../../models/shared/crudentitybase.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |


### Response

**[*operations.InstallPackResponse](../../models/operations/installpackresponse.md), error**


## ListPacks

Get info on packs

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
    res, err := s.Packs.ListPacks(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.PackInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListPacksResponse](../../models/operations/listpacksresponse.md), error**


## Uninstall

Uninstall Pack from the system

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
    id := "aspernatur"

    ctx := context.Background()
    res, err := s.Packs.Uninstall(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.PackInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Pack name                                             |


### Response

**[*operations.UninstallPackResponse](../../models/operations/uninstallpackresponse.md), error**


## Upgrade

Upgrade Pack

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
    id := "inventore"
    minor := "sequi"
    source := "fugit"
    spec := "fuga"

    ctx := context.Background()
    res, err := s.Packs.Upgrade(ctx, id, minor, source, spec)
    if err != nil {
        log.Fatal(err)
    }

    if res.PackInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |
| `id`                                                        | *string*                                                    | :heavy_check_mark:                                          | Pack name                                                   |
| `minor`                                                     | **string*                                                   | :heavy_minus_sign:                                          | body boolean optional Only upgrade to minor/patch versions  |
| `source`                                                    | **string*                                                   | :heavy_minus_sign:                                          | body string required Pack source                            |
| `spec`                                                      | **string*                                                   | :heavy_minus_sign:                                          | body string optional Specify a branch, tag or a semver spec |


### Response

**[*operations.UpgradePackResponse](../../models/operations/upgradepackresponse.md), error**


## Upload

Upload Pack

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
    filename := "hic"

    ctx := context.Background()
    res, err := s.Packs.Upload(ctx, filename)
    if err != nil {
        log.Fatal(err)
    }

    if res.PackInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `filename`                                            | **string*                                             | :heavy_minus_sign:                                    | the file to upload                                    |


### Response

**[*operations.UploadPackResponse](../../models/operations/uploadpackresponse.md), error**

