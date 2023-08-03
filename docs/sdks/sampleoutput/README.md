# SampleOutput

### Available Operations

* [Post](#post) - Send sample data to an output to validate configuration or test connectivity

## Post

Send sample data to an output to validate configuration or test connectivity

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
    id := "magni"
    outputTestRequest := &shared.OutputTestRequest{
        Events: []shared.CriblEvent{
            shared.CriblEvent{
                Raw: "est",
            },
        },
    }

    ctx := context.Background()
    res, err := s.SampleOutput.Post(ctx, id, outputTestRequest)
    if err != nil {
        log.Fatal(err)
    }

    if res.OutputTestResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |
| `id`                                                                  | *string*                                                              | :heavy_check_mark:                                                    | Output Id                                                             |
| `outputTestRequest`                                                   | [*shared.OutputTestRequest](../../models/shared/outputtestrequest.md) | :heavy_minus_sign:                                                    | OutputTestRequest object                                              |


### Response

**[*operations.PostSampleOutputResponse](../../models/operations/postsampleoutputresponse.md), error**

