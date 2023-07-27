# DistributedDeployment

### Available Operations

* [Get](#get) - Get summary of Distributed deployment

## Get

Get summary of Distributed deployment

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
    res, err := s.DistributedDeployment.Get(ctx, operations.GetDistributedDeploymentRequest{
        Mode: cribl.String("impedit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DistributedSummaries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetDistributedDeploymentRequest](../../models/operations/getdistributeddeploymentrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetDistributedDeploymentResponse](../../models/operations/getdistributeddeploymentresponse.md), error**

