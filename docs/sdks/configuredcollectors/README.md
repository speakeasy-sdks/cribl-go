# ConfiguredCollectors

### Available Operations

* [Get](#get) - Get list of configured collectors

## Get

Get list of configured collectors

### Example Usage

```go
package main

import(
	"context"
	"log"
	"cribl"
	"cribl/pkg/models/shared"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ConfiguredCollectors.Get(ctx)
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

**[*operations.GetConfiguredCollectorsResponse](../../models/operations/getconfiguredcollectorsresponse.md), error**
