# Distributed

### Available Operations

* [Get](#get) - Get summary of Distributed deployment
* [ListDiagBundles](#listdiagbundles) - Get list of existing diag bundles

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
    mode := "pariatur"

    ctx := context.Background()
    res, err := s.Distributed.Get(ctx, mode)
    if err != nil {
        log.Fatal(err)
    }

    if res.DistributedSummaries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |
| `mode`                                                                              | **string*                                                                           | :heavy_minus_sign:                                                                  | Filter for worker/group type, either "worker" for Stream or "managed-edge" for Edge |


### Response

**[*operations.GetDistributedDeploymentResponse](../../models/operations/getdistributeddeploymentresponse.md), error**


## ListDiagBundles

Get list of existing diag bundles

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
    res, err := s.Distributed.ListDiagBundles(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListDiagBundles200ApplicationTarPlusGzipBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListDiagBundlesResponse](../../models/operations/listdiagbundlesresponse.md), error**

