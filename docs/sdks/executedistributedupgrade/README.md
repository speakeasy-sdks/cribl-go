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

    ctx := context.Background()
    res, err := s.ExecuteDistributedUpgrade.Post(ctx, operations.PostExecuteDistributedUpgradeRequest{
        DistributedUpgradeRequest: &shared.DistributedUpgradeRequest{
            PackageUrls: []shared.DistributedUpgradeRequestPackageUrls{
                shared.DistributedUpgradeRequestPackageUrls{
                    PackageHashURL: cribl.String("nesciunt"),
                    PackageURL: "delectus",
                },
            },
            UpgradeMode: shared.DistributedUpgradeRequestUpgradeModeRegular.ToPointer(),
            UpgradePercentage: cribl.Int64(303421),
            WorkerRetries: cribl.Int64(644223),
            WorkerRetryDelay: cribl.Int64(266680),
        },
        Group: "sunt",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CriblPackage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.PostExecuteDistributedUpgradeRequest](../../models/operations/postexecutedistributedupgraderequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.PostExecuteDistributedUpgradeResponse](../../models/operations/postexecutedistributedupgraderesponse.md), error**

