# WorkerEdgeNodes

### Available Operations

* [Get](#get) - get worker and edge nodes
* [Restarts](#restarts) - restarts worker nodes

## Get

get worker and edge nodes

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
    filterExp := "eius"
    limit := 342393
    offset := 406037
    sortExp := "amet"

    ctx := context.Background()
    res, err := s.WorkerEdgeNodes.Get(ctx, filterExp, limit, offset, sortExp)
    if err != nil {
        log.Fatal(err)
    }

    if res.MasterWorkerEntries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `filterExp`                                           | **string*                                             | :heavy_minus_sign:                                    | Filter expression evaluated against nodes             |
| `limit`                                               | **int64*                                              | :heavy_minus_sign:                                    | Maximum number of nodes to return                     |
| `offset`                                              | **int64*                                              | :heavy_minus_sign:                                    | Pagination offset                                     |
| `sortExp`                                             | **string*                                             | :heavy_minus_sign:                                    | Sorting expression evaluated against nodes            |


### Response

**[*operations.GetWorkerEdgeNodesResponse](../../models/operations/getworkeredgenodesresponse.md), error**


## Restarts

restarts worker nodes

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
    res, err := s.WorkerEdgeNodes.Restarts(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.RestartResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.RestartsWorkerEdgeNodesResponse](../../models/operations/restartsworkeredgenodesresponse.md), error**

