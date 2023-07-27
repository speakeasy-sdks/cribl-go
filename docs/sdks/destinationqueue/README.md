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

    ctx := context.Background()
    res, err := s.DestinationQueue.Delete(ctx, operations.DeleteDestinationQueueRequest{
        ID: "aea4c506-a8aa-494c-8264-4cf5e9d9a457",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteDestinationQueue200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeleteDestinationQueueRequest](../../models/operations/deletedestinationqueuerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.DeleteDestinationQueueResponse](../../models/operations/deletedestinationqueueresponse.md), error**

