# ConfigGroup

### Available Operations

* [Create](#create) - Create ConfigGroup
* [Delete](#delete) - Delete ConfigGroup
* [Get](#get) - Get a specific ConfigGroup object
* [Update](#update) - Update ConfigGroup

## Create

Create ConfigGroup

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
    res, err := s.ConfigGroup.Create(ctx, shared.ConfigGroup{
        ConfigVersion: "voluptatibus",
        Description: cribl.String("molestias"),
        EstimatedIngestRate: cribl.Int64(889794),
        Git: &shared.ConfigGroupGit{
            Commit: cribl.String("sapiente"),
            LocalChanges: cribl.Int64(764562),
            Log: []shared.Commit{
                shared.Commit{
                    AuthorEmail: cribl.String("rerum"),
                    AuthorName: cribl.String("tempora"),
                    Date: "quis",
                    Hash: "inventore",
                    Message: "fugit",
                    Short: "cumque",
                },
            },
        },
        ID: "1032648d-c2f6-4151-99eb-fd0e9fe6c632",
        Inherits: cribl.String("cumque"),
        IsFleet: cribl.Bool(false),
        IsSearch: cribl.Bool(false),
        Name: cribl.String("Philip O'Kon"),
        OnPrem: cribl.Bool(false),
        Provisioned: cribl.Bool(false),
        SourceGroupID: cribl.String("consequatur"),
        Tags: cribl.String("quasi"),
        WorkerCount: cribl.Int64(90233),
        WorkerRemoteAccess: cribl.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [shared.ConfigGroup](../../models/shared/configgroup.md) | :heavy_check_mark:                                       | The request object to use for the request.               |


### Response

**[*operations.CreateConfigGroupResponse](../../models/operations/createconfiggroupresponse.md), error**


## Delete

Delete ConfigGroup

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
    id := "ducimus"

    ctx := context.Background()
    res, err := s.ConfigGroup.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigGroup != nil {
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

**[*operations.DeleteConfigGroupResponse](../../models/operations/deleteconfiggroupresponse.md), error**


## Get

Get a specific ConfigGroup object

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
    id := "natus"
    fields := "occaecati"

    ctx := context.Background()
    res, err := s.ConfigGroup.Get(ctx, id, fields)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |
| `id`                                                                                    | *string*                                                                                | :heavy_check_mark:                                                                      | Unique ID                                                                               |
| `fields`                                                                                | **string*                                                                               | :heavy_minus_sign:                                                                      | query string additional fields to add to results: git.commit, git.localChanges, git.log |


### Response

**[*operations.GetConfigGroupResponse](../../models/operations/getconfiggroupresponse.md), error**


## Update

Update ConfigGroup

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
    id := "suscipit"
    configGroup := &shared.ConfigGroup{
        ConfigVersion: "adipisci",
        Description: cribl.String("quasi"),
        EstimatedIngestRate: cribl.Int64(169025),
        Git: &shared.ConfigGroupGit{
            Commit: cribl.String("doloribus"),
            LocalChanges: cribl.Int64(859581),
            Log: []shared.Commit{
                shared.Commit{
                    AuthorEmail: cribl.String("ipsa"),
                    AuthorName: cribl.String("tempora"),
                    Date: "nihil",
                    Hash: "molestiae",
                    Message: "dicta",
                    Short: "iusto",
                },
                shared.Commit{
                    AuthorEmail: cribl.String("esse"),
                    AuthorName: cribl.String("praesentium"),
                    Date: "maiores",
                    Hash: "reiciendis",
                    Message: "vel",
                    Short: "architecto",
                },
                shared.Commit{
                    AuthorEmail: cribl.String("fugiat"),
                    AuthorName: cribl.String("doloremque"),
                    Date: "dicta",
                    Hash: "odio",
                    Message: "tempora",
                    Short: "esse",
                },
                shared.Commit{
                    AuthorEmail: cribl.String("ex"),
                    AuthorName: cribl.String("consectetur"),
                    Date: "aliquid",
                    Hash: "ipsa",
                    Message: "laborum",
                    Short: "sunt",
                },
            },
        },
        ID: "5db6a660-659a-41ad-aaab-5851d6c645b0",
        Inherits: cribl.String("molestias"),
        IsFleet: cribl.Bool(false),
        IsSearch: cribl.Bool(false),
        Name: cribl.String("Gene Brekke"),
        OnPrem: cribl.Bool(false),
        Provisioned: cribl.Bool(false),
        SourceGroupID: cribl.String("veritatis"),
        Tags: cribl.String("rerum"),
        WorkerCount: cribl.Int64(665678),
        WorkerRemoteAccess: cribl.Bool(false),
    }

    ctx := context.Background()
    res, err := s.ConfigGroup.Update(ctx, id, configGroup)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |
| `id`                                                      | *string*                                                  | :heavy_check_mark:                                        | Unique ID                                                 |
| `configGroup`                                             | [*shared.ConfigGroup](../../models/shared/configgroup.md) | :heavy_minus_sign:                                        | ConfigGroup object to be updated                          |


### Response

**[*operations.UpdateConfigGroupResponse](../../models/operations/updateconfiggroupresponse.md), error**

