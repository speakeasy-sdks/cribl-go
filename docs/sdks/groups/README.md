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
        ConfigVersion: "debitis",
        Description: cribl.String("neque"),
        EstimatedIngestRate: cribl.Int64(677115),
        Git: &shared.ConfigGroupGit{
            Commit: cribl.String("nostrum"),
            LocalChanges: cribl.Int64(639028),
            Log: []shared.Commit{
                shared.Commit{
                    AuthorEmail: cribl.String("dolorum"),
                    AuthorName: cribl.String("corrupti"),
                    Date: "accusamus",
                    Hash: "tempora",
                    Message: "atque",
                    Short: "fugit",
                },
            },
        },
        ID: "4d0ab407-5088-4e51-8620-65e904f3b119",
        Inherits: cribl.String("labore"),
        IsFleet: cribl.Bool(false),
        IsSearch: cribl.Bool(false),
        Name: cribl.String("Alberto Osinski"),
        OnPrem: cribl.Bool(false),
        Provisioned: cribl.Bool(false),
        SourceGroupID: cribl.String("laboriosam"),
        Tags: cribl.String("alias"),
        WorkerCount: cribl.Int64(227084),
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
    id := "deserunt"

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
    id := "voluptate"
    deployRequest := &shared.DeployRequest{
        Version: "unde",
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
    id := "reiciendis"
    fields := "provident"

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
    id := "repellendus"
    deployRequest := &shared.DeployRequest{
        Version: "delectus",
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
    fields := "voluptates"
    product := "perferendis"

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
    id := "est"
    configGroup := &shared.ConfigGroup{
        ConfigVersion: "quidem",
        Description: cribl.String("reprehenderit"),
        EstimatedIngestRate: cribl.Int64(813679),
        Git: &shared.ConfigGroupGit{
            Commit: cribl.String("fuga"),
            LocalChanges: cribl.Int64(509807),
            Log: []shared.Commit{
                shared.Commit{
                    AuthorEmail: cribl.String("mollitia"),
                    AuthorName: cribl.String("veniam"),
                    Date: "voluptatem",
                    Hash: "quisquam",
                    Message: "repudiandae",
                    Short: "quasi",
                },
            },
        },
        ID: "87f86bc1-73d6-489e-ae95-26f8d986e881",
        Inherits: cribl.String("recusandae"),
        IsFleet: cribl.Bool(false),
        IsSearch: cribl.Bool(false),
        Name: cribl.String("Mack Grant DVM"),
        OnPrem: cribl.Bool(false),
        Provisioned: cribl.Bool(false),
        SourceGroupID: cribl.String("dicta"),
        Tags: cribl.String("accusantium"),
        WorkerCount: cribl.Int64(106429),
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

