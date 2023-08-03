# ExecuteDistributedUpgrade

### Available Operations

* [Post](#post) - Execute distributed group upgrade

## Post

Execute distributed group upgrade

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
    group := "fugiat"
    distributedUpgradeRequest := &shared.DistributedUpgradeRequest{
        PackageUrls: []shared.DistributedUpgradeRequestPackageUrls{
            shared.DistributedUpgradeRequestPackageUrls{
                PackageHashURL: cribl.String("officiis"),
                PackageURL: "ducimus",
            },
            shared.DistributedUpgradeRequestPackageUrls{
                PackageHashURL: cribl.String("dolor"),
                PackageURL: "dicta",
            },
            shared.DistributedUpgradeRequestPackageUrls{
                PackageHashURL: cribl.String("error"),
                PackageURL: "porro",
            },
        },
        UpgradeMode: shared.DistributedUpgradeRequestUpgradeModeRolling.ToPointer(),
        UpgradePercentage: cribl.Int64(491591),
        WorkerRetries: cribl.Int64(458970),
        WorkerRetryDelay: cribl.Int64(854115),
    }

    ctx := context.Background()
    res, err := s.ExecuteDistributedUpgrade.Post(ctx, group, distributedUpgradeRequest)
    if err != nil {
        log.Fatal(err)
    }

    if res.CriblPackage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `group`                                                                               | *string*                                                                              | :heavy_check_mark:                                                                    | Group to upgrade                                                                      |
| `distributedUpgradeRequest`                                                           | [*shared.DistributedUpgradeRequest](../../models/shared/distributedupgraderequest.md) | :heavy_minus_sign:                                                                    | distributedUpgrade object                                                             |


### Response

**[*operations.PostExecuteDistributedUpgradeResponse](../../models/operations/postexecutedistributedupgraderesponse.md), error**

