# Pipelines

## Overview

Actions related to pipelines

### Available Operations

* [Delete](#delete) - Delete Pipeline
* [Get](#get) - Get Pipeline by ID
* [ListPipelineObject](#listpipelineobject) - Get a list of Pipeline objects
* [Post](#post) - Create Pipeline
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
    id := "corporis"

    ctx := context.Background()
    res, err := s.Pipelines.Delete(ctx, id)
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
    id := "officiis"

    ctx := context.Background()
    res, err := s.Pipelines.Get(ctx, id)
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


## ListPipelineObject

Get a list of Pipeline objects

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
    res, err := s.Pipelines.ListPipelineObject(ctx)
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


### Response

**[*operations.ListPipelineObjectResponse](../../models/operations/listpipelineobjectresponse.md), error**


## Post

Create Pipeline

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
    res, err := s.Pipelines.Post(ctx, shared.Pipeline{
        Conf: shared.PipelineConf{
            AsyncFuncTimeout: cribl.Int64(972912),
            Description: cribl.String("cum"),
            Functions: []shared.PipelineFunctionConf{
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("at"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("alias"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("quia"),
                    ID: "bae0be2d-7822-459e-bea4-b5197f92443d",
                },
            },
            Groups: map[string]shared.PipelineConfGroups{
                "officia": shared.PipelineConfGroups{
                    Description: cribl.String("dignissimos"),
                    Disabled: cribl.Bool(false),
                    Name: "Santiago Herzog",
                },
            },
            Output: cribl.String("voluptatum"),
            Streamtags: []string{
                "cupiditate",
            },
        },
        ID: "5c537c64-54ef-4b0b-b489-6c3ca5acfbe2",
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Pipeline](../../models/shared/pipeline.md)    | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostCreatePipelineResponse](../../models/operations/postcreatepipelineresponse.md), error**


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
    id := "repellat"
    pipeline := &shared.Pipeline{
        Conf: shared.PipelineConf{
            AsyncFuncTimeout: cribl.Int64(865946),
            Description: cribl.String("nemo"),
            Functions: []shared.PipelineFunctionConf{
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("reprehenderit"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("aperiam"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("odio"),
                    ID: "57792917-7dea-4c64-aecb-573409e3eb1e",
                },
            },
            Groups: map[string]shared.PipelineConfGroups{
                "veniam": shared.PipelineConfGroups{
                    Description: cribl.String("animi"),
                    Disabled: cribl.Bool(false),
                    Name: "Juana Buckridge",
                },
            },
            Output: cribl.String("nobis"),
            Streamtags: []string{
                "ipsa",
            },
        },
        ID: "7f116db9-9545-4fc9-9fa8-8970e189dbb3",
    }

    ctx := context.Background()
    res, err := s.Pipelines.Update(ctx, id, pipeline)
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

