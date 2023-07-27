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
    id := "architecto"

    ctx := context.Background()
    res, err := s.PipelineID.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pipelines != nil {
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

**[*operations.DeletePipelineIDResponse](../../models/operations/deletepipelineidresponse.md), error**


## Get

Get Pipeline by ID

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
    id := "praesentium"

    ctx := context.Background()
    res, err := s.PipelineID.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pipelines != nil {
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

**[*operations.GetPipelineIDResponse](../../models/operations/getpipelineidresponse.md), error**


## Update

Update Pipeline

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
    id := "eveniet"
    pipeline := &shared.Pipeline{
        Conf: shared.PipelineConf{
            AsyncFuncTimeout: cribl.Int64(221781),
            Description: cribl.String("expedita"),
            Functions: []shared.PipelineFunctionConf{
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("iste"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("illo"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("minus"),
                    ID: "8d975e0e-8419-4d8f-84f1-44f3e07edcc4",
                },
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("dolorum"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("deserunt"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("ad"),
                    ID: "f3cabd90-5a97-42e0-9672-8227b2d30947",
                },
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("accusantium"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("distinctio"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("sapiente"),
                    ID: "7a4fa87c-f535-4a6f-ae54-ebf60c321f02",
                },
            },
            Groups: map[string]shared.PipelineConfGroups{
                "rerum": shared.PipelineConfGroups{
                    Description: cribl.String("in"),
                    Disabled: cribl.Bool(false),
                    Name: "Paulette Dibbert",
                },
            },
            Output: cribl.String("dignissimos"),
            Streamtags: []string{
                "itaque",
                "vitae",
                "est",
                "accusantium",
            },
        },
        ID: "cc8df79f-0a39-46d9-8c36-4b7c15dfbace",
    }

    ctx := context.Background()
    res, err := s.PipelineID.Update(ctx, id, pipeline)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pipelines != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |
| `pipeline`                                            | [*shared.Pipeline](../../models/shared/pipeline.md)   | :heavy_minus_sign:                                    | Pipeline object to be updated                         |


### Response

**[*operations.UpdatePipelineIDResponse](../../models/operations/updatepipelineidresponse.md), error**

