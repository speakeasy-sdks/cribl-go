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
    id := "aspernatur"

    ctx := context.Background()
    res, err := s.Feature.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.FeaturesEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Feature id                                            |


### Response

**[*operations.GetFeatureResponse](../../models/operations/getfeatureresponse.md), error**

