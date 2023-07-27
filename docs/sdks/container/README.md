# Container

### Available Operations

* [Get](#get) - Get details for a single container on the edge host. Add stream=true to get a stream instead.

## Get

Get details for a single container on the edge host. Add stream=true to get a stream instead.

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
    id := "culpa"

    ctx := context.Background()
    res, err := s.Container.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Containers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


### Response

**[*operations.GetContainerResponse](../../models/operations/getcontainerresponse.md), error**

