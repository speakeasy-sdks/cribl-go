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
    res, err := s.ConfigGroup.Create(ctx, shared.ConfigGroup{
        ConfigVersion: "optio",
        Description: cribl.String("occaecati"),
        EstimatedIngestRate: cribl.Int64(364544),
        Git: &shared.ConfigGroupGit{
            Commit: cribl.String("voluptate"),
            LocalChanges: cribl.Int64(501063),
            Log: []shared.Commit{
                shared.Commit{
                    AuthorEmail: cribl.String("voluptas"),
                    AuthorName: cribl.String("numquam"),
                    Date: "nemo",
                    Hash: "quos",
                    Message: "eius",
                    Short: "aspernatur",
                },
                shared.Commit{
                    AuthorEmail: cribl.String("ducimus"),
                    AuthorName: cribl.String("nesciunt"),
                    Date: "fuga",
                    Hash: "laudantium",
                    Message: "incidunt",
                    Short: "quasi",
                },
                shared.Commit{
                    AuthorEmail: cribl.String("rem"),
                    AuthorName: cribl.String("fugiat"),
                    Date: "dicta",
                    Hash: "nisi",
                    Message: "consequuntur",
                    Short: "consectetur",
                },
            },
        },
        ID: "09fb0929-921a-4efb-9f58-c4d86e68e4be",
        Inherits: cribl.String("voluptatem"),
        IsFleet: cribl.Bool(false),
        IsSearch: cribl.Bool(false),
        Name: cribl.String("Mrs. Gina Abbott"),
        OnPrem: cribl.Bool(false),
        Provisioned: cribl.Bool(false),
        SourceGroupID: cribl.String("enim"),
        Tags: cribl.String("sint"),
        WorkerCount: cribl.Int64(858778),
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
    res, err := s.ConfigGroup.Delete(ctx, operations.DeleteConfigGroupRequest{
        ID: "a757a59e-cfef-466e-b1ca-a3383c2beb47",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteConfigGroupRequest](../../models/operations/deleteconfiggrouprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


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
    res, err := s.ConfigGroup.Get(ctx, operations.GetConfigGroupRequest{
        Fields: cribl.String("voluptate"),
        ID: "373c8d72-f64d-41db-9f2c-4310661e9634",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetConfigGroupRequest](../../models/operations/getconfiggrouprequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
    res, err := s.ConfigGroup.Update(ctx, operations.UpdateConfigGroupRequest{
        ConfigGroup: &shared.ConfigGroup{
            ConfigVersion: "perspiciatis",
            Description: cribl.String("earum"),
            EstimatedIngestRate: cribl.Int64(117525),
            Git: &shared.ConfigGroupGit{
                Commit: cribl.String("impedit"),
                LocalChanges: cribl.Int64(975884),
                Log: []shared.Commit{
                    shared.Commit{
                        AuthorEmail: cribl.String("itaque"),
                        AuthorName: cribl.String("alias"),
                        Date: "nisi",
                        Hash: "itaque",
                        Message: "velit",
                        Short: "laborum",
                    },
                    shared.Commit{
                        AuthorEmail: cribl.String("non"),
                        AuthorName: cribl.String("dolor"),
                        Date: "iusto",
                        Hash: "sit",
                        Message: "doloremque",
                        Short: "consequatur",
                    },
                    shared.Commit{
                        AuthorEmail: cribl.String("officia"),
                        AuthorName: cribl.String("recusandae"),
                        Date: "ea",
                        Hash: "quidem",
                        Message: "voluptas",
                        Short: "facilis",
                    },
                },
            },
            ID: "c9b8f759-eac5-45a9-b41d-311352965bb8",
            Inherits: cribl.String("dolorum"),
            IsFleet: cribl.Bool(false),
            IsSearch: cribl.Bool(false),
            Name: cribl.String("Beverly Abbott"),
            OnPrem: cribl.Bool(false),
            Provisioned: cribl.Bool(false),
            SourceGroupID: cribl.String("quae"),
            Tags: cribl.String("quae"),
            WorkerCount: cribl.Int64(264333),
            WorkerRemoteAccess: cribl.Bool(false),
        },
        ID: "35e139db-c225-49b1-abda-8c070e1084cb",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateConfigGroupRequest](../../models/operations/updateconfiggrouprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateConfigGroupResponse](../../models/operations/updateconfiggroupresponse.md), error**

