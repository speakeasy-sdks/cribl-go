# LiveData

### Available Operations

* [Post](#post) - Capture live incoming data

## Post

Capture live incoming data

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
    res, err := s.LiveData.Post(ctx, shared.CaptureParams{
        Filter: "fuga",
        Level: 326903,
        MaxEvents: 102446,
        StepDuration: cribl.Int64(181476),
        WorkerID: cribl.String("aliquid"),
        WorkerThreshold: cribl.Int64(166289),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LiveData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.CaptureParams](../../models/shared/captureparams.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |


### Response

**[*operations.PostLiveDataResponse](../../models/operations/postlivedataresponse.md), error**

