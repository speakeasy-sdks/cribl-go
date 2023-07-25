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
	"cribl"
	"cribl/pkg/models/shared"
	"cribl/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.WorkerEdgeNodes.Get(ctx, operations.GetWorkerEdgeNodesRequest{
        FilterExp: cribl.String("corporis"),
        Limit: cribl.Int64(388180),
        Offset: cribl.Int64(385846),
        SortExp: cribl.String("omnis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MasterWorkerEntries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetWorkerEdgeNodesRequest](../../models/operations/getworkeredgenodesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


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
	"cribl"
	"cribl/pkg/models/shared"
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

