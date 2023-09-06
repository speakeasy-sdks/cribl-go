# Workers

## Overview

Actions related to workers

### Available Operations

* [Get](#get) - get worker and edge nodes count
* [ListWorkerEdgeNodes](#listworkeredgenodes) - get worker and edge nodes
* [Restart](#restart) - restarts worker nodes

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
    filterExp := "nesciunt"

    ctx := context.Background()
    res, err := s.Workers.Get(ctx, filterExp)
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


## ListWorkerEdgeNodes

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
    filterExp := "ipsa"
    limit := 575139
    offset := 291389
    sortExp := "esse"

    ctx := context.Background()
    res, err := s.Workers.ListWorkerEdgeNodes(ctx, filterExp, limit, offset, sortExp)
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

**[*operations.ListWorkerEdgeNodesResponse](../../models/operations/listworkeredgenodesresponse.md), error**


## Restart

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
    res, err := s.Workers.Restart(ctx)
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

