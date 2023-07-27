# WorkerEdgeNodesCount

### Available Operations

* [Get](#get) - get worker and edge nodes count

## Get

get worker and edge nodes count

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
    filterExp := "voluptate"

    ctx := context.Background()
    res, err := s.WorkerEdgeNodesCount.Get(ctx, filterExp)
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkerEdgeNodes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `filterExp`                                           | **string*                                             | :heavy_minus_sign:                                    | Filter expression evaluated against nodes             |


### Response

**[*operations.GetWorkerEdgeNodesCountResponse](../../models/operations/getworkeredgenodescountresponse.md), error**

