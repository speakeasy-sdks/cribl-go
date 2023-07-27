# CreatePipeline

### Available Operations

* [Post](#post) - Create Pipeline

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
    res, err := s.CreatePipeline.Post(ctx, shared.Pipeline{
        Conf: shared.PipelineConf{
            AsyncFuncTimeout: cribl.Int64(885208),
            Description: cribl.String("eos"),
            Functions: []shared.PipelineFunctionConf{
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("odio"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("praesentium"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("odit"),
                    ID: "259e3ea4-b519-47f9-a443-da7ce52b895c",
                },
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("enim"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("neque"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("in"),
                    ID: "c6454efb-0b34-4896-83ca-5acfbe2fd570",
                },
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("odio"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("minima"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("in"),
                    ID: "7929177d-eac6-446e-8b57-3409e3eb1e5a",
                },
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("dolores"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("nam"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("dicta"),
                    ID: "2eb07f11-6db9-4954-9fc9-5fa88970e189",
                },
            },
            Groups: map[string]shared.PipelineConfGroups{
                "tempore": shared.PipelineConfGroups{
                    Description: cribl.String("libero"),
                    Disabled: cribl.Bool(false),
                    Name: "Sharon Windler",
                },
                "ipsum": shared.PipelineConfGroups{
                    Description: cribl.String("adipisci"),
                    Disabled: cribl.Bool(false),
                    Name: "Doug Baumbach",
                },
                "libero": shared.PipelineConfGroups{
                    Description: cribl.String("architecto"),
                    Disabled: cribl.Bool(false),
                    Name: "Fernando Roob",
                },
                "magnam": shared.PipelineConfGroups{
                    Description: cribl.String("itaque"),
                    Disabled: cribl.Bool(false),
                    Name: "Ollie Harris",
                },
            },
            Output: cribl.String("laudantium"),
            Streamtags: []string{
                "pariatur",
            },
        },
        ID: "3513bb6f-48b6-456b-8db3-5ff2e4b27537",
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

