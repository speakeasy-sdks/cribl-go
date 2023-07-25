# MasterNodePackage

### Available Operations

* [Post](#post) - Upgrade master node with the provided package

## Post

Upgrade master node with the provided package

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
    res, err := s.MasterNodePackage.Post(ctx, shared.UpgradeMasterRequest{
        Packages: []shared.UpgradeMasterRequestPackages{
            shared.UpgradeMasterRequestPackages{
                PackageHashURL: cribl.String("aut"),
                PackageURL: "magni",
            },
            shared.UpgradeMasterRequestPackages{
                PackageHashURL: cribl.String("rerum"),
                PackageURL: "voluptatibus",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Cribl != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.UpgradeMasterRequest](../../models/shared/upgrademasterrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.PostMasterNodePackageResponse](../../models/operations/postmasternodepackageresponse.md), error**

