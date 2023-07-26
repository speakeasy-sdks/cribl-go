# Collector

### Available Operations

* [Get](#get) - Get Collector by ID

## Get

Get Collector by ID

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
    res, err := s.Collector.Get(ctx, operations.GetCollectorRequest{
        ID: "350d8cdb-5a34-4181-8301-0421813d5208",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Collectors != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetCollectorRequest](../../models/operations/getcollectorrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetCollectorResponse](../../models/operations/getcollectorresponse.md), error**

