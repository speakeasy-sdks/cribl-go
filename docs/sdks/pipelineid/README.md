# PipelineID

### Available Operations

* [Delete](#delete) - Delete Pipeline
* [Get](#get) - Get Pipeline by ID
* [Update](#update) - Update Pipeline

## Delete

Delete Pipeline

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
    res, err := s.PipelineID.Delete(ctx, operations.DeletePipelineIDRequest{
        ID: "8140b64f-f8ae-4170-af03-b5f37e4aa868",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pipelines != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeletePipelineIDRequest](../../models/operations/deletepipelineidrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeletePipelineIDResponse](../../models/operations/deletepipelineidresponse.md), error**


## Get

Get Pipeline by ID

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
    res, err := s.PipelineID.Get(ctx, operations.GetPipelineIDRequest{
        ID: "55596673-2aa5-4dcb-a682-cb70f8cfd5fb",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pipelines != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetPipelineIDRequest](../../models/operations/getpipelineidrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetPipelineIDResponse](../../models/operations/getpipelineidresponse.md), error**


## Update

Update Pipeline

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
    res, err := s.PipelineID.Update(ctx, operations.UpdatePipelineIDRequest{
        Pipeline: &shared.Pipeline{
            Conf: shared.PipelineConf{
                AsyncFuncTimeout: cribl.Int64(388851),
                Description: cribl.String("earum"),
                Functions: []shared.PipelineFunctionConf{
                    shared.PipelineFunctionConf{
                        Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                        Description: cribl.String("dicta"),
                        Disabled: cribl.Bool(false),
                        Filter: cribl.String("quidem"),
                        Final: cribl.Bool(false),
                        GroupID: cribl.String("omnis"),
                        ID: "a9f74846-e2c3-4309-9b05-36d9e75ca006",
                    },
                    shared.PipelineFunctionConf{
                        Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                        Description: cribl.String("earum"),
                        Disabled: cribl.Bool(false),
                        Filter: cribl.String("quis"),
                        Final: cribl.Bool(false),
                        GroupID: cribl.String("dolorem"),
                        ID: "92c11a25-a8bf-492f-9742-8ad9a9f8bf82",
                    },
                    shared.PipelineFunctionConf{
                        Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                        Description: cribl.String("odit"),
                        Disabled: cribl.Bool(false),
                        Filter: cribl.String("illo"),
                        Final: cribl.Bool(false),
                        GroupID: cribl.String("architecto"),
                        ID: "25359d98-387f-47a7-9cd7-2cd2484da217",
                    },
                },
                Groups: map[string]shared.PipelineConfGroups{
                    "omnis": shared.PipelineConfGroups{
                        Description: cribl.String("reiciendis"),
                        Disabled: cribl.Bool(false),
                        Name: "Marguerite Roob PhD",
                    },
                },
                Output: cribl.String("tenetur"),
                Streamtags: []string{
                    "nihil",
                    "quia",
                },
            },
            ID: "5f1169ac-1e41-4d8a-a3c2-3e34f2dfa4a1",
        },
        ID: "97f6de92-2151-4fe1-b120-99853e9f543d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pipelines != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdatePipelineIDRequest](../../models/operations/updatepipelineidrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdatePipelineIDResponse](../../models/operations/updatepipelineidresponse.md), error**

