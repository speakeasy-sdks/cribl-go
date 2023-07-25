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
    res, err := s.WorkerEdgeNodesCount.Get(ctx, operations.GetWorkerEdgeNodesCountRequest{
        FilterExp: cribl.String("tenetur"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkerEdgeNodes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetWorkerEdgeNodesCountRequest](../../models/operations/getworkeredgenodescountrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetWorkerEdgeNodesCountResponse](../../models/operations/getworkeredgenodescountresponse.md), error**

