# LatestPQ

### Available Operations

* [Get](#get) - Get status of latest clear PQ job for an output

## Get

Get status of latest clear PQ job for an output

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
    id := "distinctio"

    ctx := context.Background()
    res, err := s.LatestPQ.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetLatestPQ200ApplicationJSONObject != nil {
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

**[*operations.GetLatestPQResponse](../../models/operations/getlatestpqresponse.md), error**

