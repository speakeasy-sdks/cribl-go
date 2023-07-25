# Feature

### Available Operations

* [Get](#get) - Get feature by Id

## Get

Get feature by id (i.e. 'type/name`)

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
    res, err := s.Feature.Get(ctx, operations.GetFeatureRequest{
        ID: "f9621a6a-4f77-4a87-ae3e-4be752c65b34",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FeaturesEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetFeatureRequest](../../models/operations/getfeaturerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetFeatureResponse](../../models/operations/getfeatureresponse.md), error**

