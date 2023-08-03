# Groups

### Available Operations

* [Create](#create) - Create ConfigGroup
* [Delete](#delete) - Delete ConfigGroup
* [Deploy](#deploy) - Deploy commits for a Fleet or Worker Group
* [Get](#get) - Get a specific ConfigGroup object
* [GetConfigVersion](#getconfigversion) - Get effective bundle version for given Group
* [ListGroups](#listgroups) - Get a list of ConfigGroup objects
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
    res, err := s.Groups.Create(ctx, shared.ConfigGroup{
        ConfigVersion: "atque",
        Description: cribl.String("beatae"),
        EstimatedIngestRate: cribl.Int64(868255),
        Git: &shared.ConfigGroupGit{
            Commit: cribl.String("labore"),
            LocalChanges: cribl.Int64(794988),
            Log: []shared.Commit{
                shared.Commit{
                    AuthorEmail: cribl.String("voluptatem"),
                    AuthorName: cribl.String("perferendis"),
                    Date: "rerum",
                    Hash: "ea",
                    Message: "aperiam",
                    Short: "dignissimos",
                },
                shared.Commit{
                    AuthorEmail: cribl.String("repellat"),
                    AuthorName: cribl.String("velit"),
                    Date: "porro",
                    Hash: "provident",
                    Message: "consectetur",
                    Short: "eligendi",
                },
            },
        },
        ID: "73b9da3f-2ced-4a7e-a3f2-257411faf4b7",
        Inherits: cribl.String("exercitationem"),
        IsFleet: cribl.Bool(false),
        IsSearch: cribl.Bool(false),
        Name: cribl.String("Marjorie Waelchi"),
        OnPrem: cribl.Bool(false),
        Provisioned: cribl.Bool(false),
        SourceGroupID: cribl.String("explicabo"),
        Tags: cribl.String("accusamus"),
        WorkerCount: cribl.Int64(525809),
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
    id := "aperiam"

    ctx := context.Background()
    res, err := s.Groups.Delete(ctx, id)
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


## Deploy

Deploy commits for a Fleet or Worker Group

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
    id := "odit"
    deployRequest := &shared.DeployRequest{
        Version: "deleniti",
    }

    ctx := context.Background()
    res, err := s.Groups.Deploy(ctx, id, deployRequest)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                     | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ctx`                                                         | [context.Context](https://pkg.go.dev/context#Context)         | :heavy_check_mark:                                            | The context to use for the request.                           |
| `id`                                                          | *string*                                                      | :heavy_check_mark:                                            | Unique ID                                                     |
| `deployRequest`                                               | [*shared.DeployRequest](../../models/shared/deployrequest.md) | :heavy_minus_sign:                                            | DeployRequest object                                          |


### Response

**[*operations.DeployFleetOrWorkerGroupResponse](../../models/operations/deployfleetorworkergroupresponse.md), error**


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
    id := "enim"
    fields := "voluptate"

    ctx := context.Background()
    res, err := s.Groups.Get(ctx, id, fields)
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


## GetConfigVersion

Get effective bundle version for given Group

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
    id := "similique"
    deployRequest := &shared.DeployRequest{
        Version: "minima",
    }

    ctx := context.Background()
    res, err := s.Groups.GetConfigVersion(ctx, id, deployRequest)
    if err != nil {
        log.Fatal(err)
    }

    if res.GroupBundle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                     | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ctx`                                                         | [context.Context](https://pkg.go.dev/context#Context)         | :heavy_check_mark:                                            | The context to use for the request.                           |
| `id`                                                          | *string*                                                      | :heavy_check_mark:                                            | Group ID                                                      |
| `deployRequest`                                               | [*shared.DeployRequest](../../models/shared/deployrequest.md) | :heavy_minus_sign:                                            | DeployRequest object                                          |


### Response

**[*operations.GetGroupBundleResponse](../../models/operations/getgroupbundleresponse.md), error**


## ListGroups

Get a list of ConfigGroup objects

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
    fields := "libero"
    product := "magnam"

    ctx := context.Background()
    res, err := s.Groups.ListGroups(ctx, fields, product)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `fields`                                                                   | **string*                                                                  | :heavy_minus_sign:                                                         | additional fields to add to results: git.commit, git.localChanges, git.log |
| `product`                                                                  | **string*                                                                  | :heavy_minus_sign:                                                         | filter to specific product: "stream" or "edge"                             |


### Response

**[*operations.ListGroupsResponse](../../models/operations/listgroupsresponse.md), error**


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
    id := "sit"
    configGroup := &shared.ConfigGroup{
        ConfigVersion: "modi",
        Description: cribl.String("eum"),
        EstimatedIngestRate: cribl.Int64(199529),
        Git: &shared.ConfigGroupGit{
            Commit: cribl.String("mollitia"),
            LocalChanges: cribl.Int64(492632),
            Log: []shared.Commit{
                shared.Commit{
                    AuthorEmail: cribl.String("nostrum"),
                    AuthorName: cribl.String("molestiae"),
                    Date: "veniam",
                    Hash: "reiciendis",
                    Message: "ab",
                    Short: "modi",
                },
                shared.Commit{
                    AuthorEmail: cribl.String("aut"),
                    AuthorName: cribl.String("aut"),
                    Date: "eveniet",
                    Hash: "odio",
                    Message: "commodi",
                    Short: "numquam",
                },
                shared.Commit{
                    AuthorEmail: cribl.String("dolorum"),
                    AuthorName: cribl.String("possimus"),
                    Date: "voluptate",
                    Hash: "consectetur",
                    Message: "nesciunt",
                    Short: "quaerat",
                },
                shared.Commit{
                    AuthorEmail: cribl.String("itaque"),
                    AuthorName: cribl.String("minus"),
                    Date: "sunt",
                    Hash: "distinctio",
                    Message: "iusto",
                    Short: "quas",
                },
            },
        },
        ID: "1b36a080-88d1-400e-bada-200ef0422eb2",
        Inherits: cribl.String("beatae"),
        IsFleet: cribl.Bool(false),
        IsSearch: cribl.Bool(false),
        Name: cribl.String("Pauline Rowe"),
        OnPrem: cribl.Bool(false),
        Provisioned: cribl.Bool(false),
        SourceGroupID: cribl.String("officia"),
        Tags: cribl.String("libero"),
        WorkerCount: cribl.Int64(520678),
        WorkerRemoteAccess: cribl.Bool(false),
    }

    ctx := context.Background()
    res, err := s.Groups.Update(ctx, id, configGroup)
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

