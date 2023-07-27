# SearchTimeline

### Available Operations

* [Get](#get) - Get search timeline

## Get

Get search timeline

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

    ctx := context.Background()
    res, err := s.SearchTimeline.Get(ctx, operations.GetSearchTimelineRequest{
        ID: "97e15229-7510-4da8-8312-292cc61c2a70",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchTimeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetSearchTimelineRequest](../../models/operations/getsearchtimelinerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetSearchTimelineResponse](../../models/operations/getsearchtimelineresponse.md), error**

