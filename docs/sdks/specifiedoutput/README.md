# SpecifiedOutput

### Available Operations

* [Get](#get) - Get samples data for the specified output. Used to get sample data for the test action.

## Get

Get samples data for the specified output. Used to get sample data for the test action.

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
    id := "blanditiis"

    ctx := context.Background()
    res, err := s.SpecifiedOutput.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSpecifiedOutput200ApplicationJSONObject != nil {
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

**[*operations.GetSpecifiedOutputResponse](../../models/operations/getspecifiedoutputresponse.md), error**

