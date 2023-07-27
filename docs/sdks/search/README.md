# Search

### Available Operations

* [DispatchSearch](#dispatchsearch) - Dispatch search *id* to worker nodes filtered by worker node filter using RPC

## DispatchSearch

Dispatch search *id* to worker nodes filtered by worker node filter using RPC

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
    res, err := s.Search.DispatchSearch(ctx, operations.DispatchSearchRequest{
        ID: "2bf7d67c-a84a-4d99-b41d-61243531870c",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DispatchSearchRequest](../../models/operations/dispatchsearchrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DispatchSearchResponse](../../models/operations/dispatchsearchresponse.md), error**

