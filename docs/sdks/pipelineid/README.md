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
        ID: "46082bfb-dc41-4ff5-94e2-ae4fb5cb35d1",
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
        ID: "7638f1ed-b783-459e-8c5c-b860f8cd580b",
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
                AsyncFuncTimeout: cribl.Int64(657520),
                Description: cribl.String("molestiae"),
                Functions: []shared.PipelineFunctionConf{
                    shared.PipelineFunctionConf{
                        Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                        Description: cribl.String("quas"),
                        Disabled: cribl.Bool(false),
                        Filter: cribl.String("sunt"),
                        Final: cribl.Bool(false),
                        GroupID: cribl.String("aperiam"),
                        ID: "e4fe4447-297c-4d3b-9dd3-bbce247b7684",
                    },
                },
                Groups: map[string]shared.PipelineConfGroups{
                    "reiciendis": shared.PipelineConfGroups{
                        Description: cribl.String("asperiores"),
                        Disabled: cribl.Bool(false),
                        Name: "Sandra Brekke",
                    },
                    "temporibus": shared.PipelineConfGroups{
                        Description: cribl.String("in"),
                        Disabled: cribl.Bool(false),
                        Name: "Vicky Wolf",
                    },
                    "facere": shared.PipelineConfGroups{
                        Description: cribl.String("aut"),
                        Disabled: cribl.Bool(false),
                        Name: "Jean Krajcik",
                    },
                    "blanditiis": shared.PipelineConfGroups{
                        Description: cribl.String("quaerat"),
                        Disabled: cribl.Bool(false),
                        Name: "Virginia Mitchell",
                    },
                },
                Output: cribl.String("rerum"),
                Streamtags: []string{
                    "tempora",
                    "quidem",
                },
            },
            ID: "d3c43159-d33e-4595-bc00-1139863aa41e",
        },
        ID: "6c31cc2f-1fcb-451c-9a41-ffbe9cbd795e",
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

