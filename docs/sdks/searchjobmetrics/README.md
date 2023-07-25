# SearchJobMetrics

### Available Operations

* [Get](#get) - Get search job metrics

## Get

Get search job metrics

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
    res, err := s.SearchJobMetrics.Get(ctx, operations.GethSearchJobMetricsRequest{
        ID: "9e3c71ad-31be-4cb8-bd23-78ae3bfc23d9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GethSearchJobMetrics200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GethSearchJobMetricsRequest](../../models/operations/gethsearchjobmetricsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GethSearchJobMetricsResponse](../../models/operations/gethsearchjobmetricsresponse.md), error**

