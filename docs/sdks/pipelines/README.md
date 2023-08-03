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
    id := "dolore"

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
    id := "tenetur"

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
            AsyncFuncTimeout: cribl.Int64(493407),
            Description: cribl.String("esse"),
            Functions: []shared.PipelineFunctionConf{
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("laudantium"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("esse"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("eveniet"),
                    ID: "e3e4be75-2c65-4b34-818e-3bb91c8d975e",
                },
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("aperiam"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("voluptates"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("laudantium"),
                    ID: "419d8f84-f144-4f3e-87ed-cc4aa5f3cabd",
                },
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("omnis"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("aut"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("ipsam"),
                    ID: "a972e056-7282-427b-ad30-9470bf7a4fa8",
                },
            },
            Groups: map[string]shared.PipelineConfGroups{
                "quod": shared.PipelineConfGroups{
                    Description: cribl.String("voluptatibus"),
                    Disabled: cribl.Bool(false),
                    Name: "Sheila Hermann",
                },
                "doloribus": shared.PipelineConfGroups{
                    Description: cribl.String("animi"),
                    Disabled: cribl.Bool(false),
                    Name: "Lewis Fritsch",
                },
            },
            Output: cribl.String("maiores"),
            Streamtags: []string{
                "voluptatem",
                "optio",
            },
        },
        ID: "321f023b-75d2-4367-be1a-0cc8df79f0a3",
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
    id := "provident"
    pipeline := &shared.Pipeline{
        Conf: shared.PipelineConf{
            AsyncFuncTimeout: cribl.Int64(404774),
            Description: cribl.String("repellendus"),
            Functions: []shared.PipelineFunctionConf{
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("alias"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("impedit"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("sequi"),
                    ID: "64b7c15d-fbac-4e18-8b1c-4ee2c8c6ce61",
                },
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("veritatis"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("maiores"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("itaque"),
                    ID: "eb1c7cbd-b6ee-4c74-b78b-a25317747dc9",
                },
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("illo"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("exercitationem"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("laborum"),
                    ID: "d2caf5dd-6723-4dc0-b5ae-2f3a6b700878",
                },
            },
            Groups: map[string]shared.PipelineConfGroups{
                "quis": shared.PipelineConfGroups{
                    Description: cribl.String("iure"),
                    Disabled: cribl.Bool(false),
                    Name: "Gail Fay",
                },
                "est": shared.PipelineConfGroups{
                    Description: cribl.String("iure"),
                    Disabled: cribl.Bool(false),
                    Name: "Eduardo Larkin",
                },
            },
            Output: cribl.String("enim"),
            Streamtags: []string{
                "minima",
                "tempora",
            },
        },
        ID: "080d40bc-acc6-4cbd-ab5f-3ec909304f92",
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

