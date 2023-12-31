# EdgeEvents

### Available Operations

* [ListConfiguredCollectors](#listconfiguredcollectors) - Get list of configured collectors

## ListConfiguredCollectors

Get list of configured collectors

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
    res, err := s.EdgeEvents.ListConfiguredCollectors(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfiguredCollectors != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListConfiguredCollectorsResponse](../../models/operations/listconfiguredcollectorsresponse.md), error**

