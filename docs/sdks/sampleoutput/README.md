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
    res, err := s.SampleOutput.Post(ctx, operations.PostSampleOutputRequest{
        OutputTestRequest: &shared.OutputTestRequest{
            Events: []shared.CriblEvent{
                shared.CriblEvent{
                    Raw: "sit",
                },
                shared.CriblEvent{
                    Raw: "recusandae",
                },
            },
        },
        ID: "f53a34a1-b8fe-4997-b1ad-c05d85ae2dfb",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OutputTestResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.PostSampleOutputRequest](../../models/operations/postsampleoutputrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PostSampleOutputResponse](../../models/operations/postsampleoutputresponse.md), error**

