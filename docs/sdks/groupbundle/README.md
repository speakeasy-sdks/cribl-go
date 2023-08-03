# GroupBundle

### Available Operations

* [Get](#get) - Get effective bundle version for given Group

## Get

Get effective bundle version for given Group

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
    id := "quae"
    deployRequest := &shared.DeployRequest{
        Version: "eaque",
    }

    ctx := context.Background()
    res, err := s.GroupBundle.Get(ctx, id, deployRequest)
    if err != nil {
        log.Fatal(err)
    }

    if res.GroupBundle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                     | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ctx`                                                         | [context.Context](https://pkg.go.dev/context#Context)         | :heavy_check_mark:                                            | The context to use for the request.                           |
| `id`                                                          | *string*                                                      | :heavy_check_mark:                                            | Group ID                                                      |
| `deployRequest`                                               | [*shared.DeployRequest](../../models/shared/deployrequest.md) | :heavy_minus_sign:                                            | DeployRequest object                                          |


### Response

**[*operations.GetGroupBundleResponse](../../models/operations/getgroupbundleresponse.md), error**

