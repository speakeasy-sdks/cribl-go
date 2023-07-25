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
    res, err := s.StageDistributedPackage.Post(ctx, operations.PostStageDistributedPackageRequest{
        Group: "ullam",
        UpgradePercentage: cribl.String("dicta"),
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PostStageDistributedPackageRequest](../../models/operations/poststagedistributedpackagerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.PostStageDistributedPackageResponse](../../models/operations/poststagedistributedpackageresponse.md), error**

