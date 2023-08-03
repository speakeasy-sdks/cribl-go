# FleetOrWorkerGroup

### Available Operations

* [Deploy](#deploy) - Deploy commits for a Fleet or Worker Group

## Deploy

Deploy commits for a Fleet or Worker Group

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
    id := "voluptatem"
    deployRequest := &shared.DeployRequest{
        Version: "ad",
    }

    ctx := context.Background()
    res, err := s.FleetOrWorkerGroup.Deploy(ctx, id, deployRequest)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                     | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ctx`                                                         | [context.Context](https://pkg.go.dev/context#Context)         | :heavy_check_mark:                                            | The context to use for the request.                           |
| `id`                                                          | *string*                                                      | :heavy_check_mark:                                            | Unique ID                                                     |
| `deployRequest`                                               | [*shared.DeployRequest](../../models/shared/deployrequest.md) | :heavy_minus_sign:                                            | DeployRequest object                                          |


### Response

**[*operations.DeployFleetOrWorkerGroupResponse](../../models/operations/deployfleetorworkergroupresponse.md), error**

