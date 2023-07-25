# Pack

### Available Operations

* [Clone](#clone) - Clone Pack
* [Export](#export) - Export Pack
* [Install](#install) - Install Pack
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
    res, err := s.Pack.Clone(ctx, shared.PackClone{
        Dest: cribl.String("dignissimos"),
        DstGroups: []string{
            "laudantium",
            "dolore",
        },
        Force: cribl.Bool(false),
        Packs: []string{
            "repudiandae",
            "aspernatur",
        },
        SrcGroup: "quod",
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
    res, err := s.Pack.Export(ctx, operations.ExportPackRequest{
        Filename: cribl.String("dolorem"),
        ID: "309db053-6d9e-475c-a006-f5392c11a25a",
        Mode: "rem",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ExportPackRequest](../../models/operations/exportpackrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


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
    res, err := s.Pack.Install(ctx, shared.CrudEntityBase{
        ID: "bf92f974-28ad-49a9-b8bf-8221125359d9",
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


## Uninstall

Uninstall Pack from the system

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
    res, err := s.Pack.Uninstall(ctx, operations.UninstallPackRequest{
        ID: "8387f7a7-9cd7-42cd-a484-da21729f2ac4",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UninstallPackRequest](../../models/operations/uninstallpackrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


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
    res, err := s.Pack.Upgrade(ctx, operations.UpgradePackRequest{
        ID: "1ef5725f-1169-4ac1-a41d-8a23c23e34f2",
        Minor: cribl.String("quibusdam"),
        Source: cribl.String("repellat"),
        Spec: cribl.String("mollitia"),
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.UpgradePackRequest](../../models/operations/upgradepackrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


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
    res, err := s.Pack.Upload(ctx, operations.UploadPackRequest{
        Filename: cribl.String("quaerat"),
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.UploadPackRequest](../../models/operations/uploadpackrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.UploadPackResponse](../../models/operations/uploadpackresponse.md), error**

