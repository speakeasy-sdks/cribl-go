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
    res, err := s.Pack.Clone(ctx, shared.PackClone{
        Dest: cribl.String("consectetur"),
        DstGroups: []string{
            "nesciunt",
            "earum",
        },
        Force: cribl.Bool(false),
        Packs: []string{
            "dolor",
            "placeat",
        },
        SrcGroup: "quos",
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

    ctx := context.Background()
    res, err := s.Pack.Export(ctx, operations.ExportPackRequest{
        Filename: cribl.String("sed"),
        ID: "09379aa6-9cd5-4fbc-b79d-a18a7822bf95",
        Mode: "quos",
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
    res, err := s.Pack.Install(ctx, shared.CrudEntityBase{
        ID: "94e6861a-db55-4f9e-9d75-1c9fe8f7502b",
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
    res, err := s.Pack.Uninstall(ctx, operations.UninstallPackRequest{
        ID: "fdc34508-41f1-4764-8563-79f3fb27e21f",
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
    res, err := s.Pack.Upgrade(ctx, operations.UpgradePackRequest{
        ID: "862657b3-6fc6-4b9f-987c-e525c67641a8",
        Minor: cribl.String("velit"),
        Source: cribl.String("quasi"),
        Spec: cribl.String("sed"),
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
    res, err := s.Pack.Upload(ctx, operations.UploadPackRequest{
        Filename: cribl.String("officiis"),
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

