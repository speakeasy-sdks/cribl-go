# OutputStatusID

### Available Operations

* [Get](#get) - Get OutputStatus by ID

## Get

Get OutputStatus by ID

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
    res, err := s.OutputStatusID.Get(ctx, operations.GetOutputStatusIDRequest{
        ID: "c9b2ad32-dafe-481a-88f4-444573fecd47",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OutputStatuses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetOutputStatusIDRequest](../../models/operations/getoutputstatusidrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetOutputStatusIDResponse](../../models/operations/getoutputstatusidresponse.md), error**

