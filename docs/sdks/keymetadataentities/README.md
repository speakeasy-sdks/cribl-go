# KeyMetadataEntities

### Available Operations

* [Get](#get) - Get a list of KeyMetadataEntity objects

## Get

Get a list of KeyMetadataEntity objects

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
    res, err := s.KeyMetadataEntities.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.KeyMetadataEntities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetKeyMetadataEntitiesResponse](../../models/operations/getkeymetadataentitiesresponse.md), error**

