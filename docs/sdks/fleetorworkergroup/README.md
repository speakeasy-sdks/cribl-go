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

    ctx := context.Background()
    res, err := s.FleetOrWorkerGroup.Deploy(ctx, operations.DeployFleetOrWorkerGroupRequest{
        DeployRequest: &shared.DeployRequest{
            Version: "est",
        },
        ID: "4fa87cf5-35a6-4fae-94eb-f60c321f023b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.DeployFleetOrWorkerGroupRequest](../../models/operations/deployfleetorworkergrouprequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.DeployFleetOrWorkerGroupResponse](../../models/operations/deployfleetorworkergroupresponse.md), error**

