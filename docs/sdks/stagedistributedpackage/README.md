# StageDistributedPackage

### Available Operations

* [Post](#post) - Stage distributed group upgrade

## Post

Stage distributed group upgrade

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
    group := "laudantium"
    upgradePercentage := "voluptates"

    ctx := context.Background()
    res, err := s.StageDistributedPackage.Post(ctx, group, upgradePercentage)
    if err != nil {
        log.Fatal(err)
    }

    if res.CriblPackage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `group`                                                        | *string*                                                       | :heavy_check_mark:                                             | Group to upgrade                                               |
| `upgradePercentage`                                            | **string*                                                      | :heavy_minus_sign:                                             | body number percentage of nodes on the worker group to upgrade |


### Response

**[*operations.PostStageDistributedPackageResponse](../../models/operations/poststagedistributedpackageresponse.md), error**

