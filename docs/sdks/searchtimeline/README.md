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
    res, err := s.SearchTimeline.Get(ctx, operations.GetSearchTimelineRequest{
        ID: "386f62c9-69c4-4cc6-b788-90a3fd3c81da",
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

