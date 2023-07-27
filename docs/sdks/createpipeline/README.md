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
            AsyncFuncTimeout: cribl.Int64(959143),
            Description: cribl.String("officiis"),
            Functions: []shared.PipelineFunctionConf{
                shared.PipelineFunctionConf{
                    Conf: shared.PipelineFunctionConfFunctionSpecificConfigs{},
                    Description: cribl.String("fuga"),
                    Disabled: cribl.Bool(false),
                    Filter: cribl.String("pariatur"),
                    Final: cribl.Bool(false),
                    GroupID: cribl.String("debitis"),
                    ID: "008e6f8c-5f35-40d8-8db5-a34181430104",
                },
            },
            Groups: map[string]shared.PipelineConfGroups{
                "ab": shared.PipelineConfGroups{
                    Description: cribl.String("laudantium"),
                    Disabled: cribl.Bool(false),
                    Name: "Rosa Stiedemann",
                },
            },
            Output: cribl.String("ipsa"),
            Streamtags: []string{
                "eveniet",
                "impedit",
                "officiis",
            },
        },
        ID: "7e253b66-8451-4c6c-ae20-5e16deab3fec",
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

