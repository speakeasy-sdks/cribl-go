# DestinationQueue

### Available Operations

* [Delete](#delete) - Delete destination persistent queue

## Delete

Delete destination persistent queue

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
    id := "magnam"

    ctx := context.Background()
    res, err := s.DestinationQueue.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteDestinationQueue200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Output Id                                             |


### Response

**[*operations.DeleteDestinationQueueResponse](../../models/operations/deletedestinationqueueresponse.md), error**

