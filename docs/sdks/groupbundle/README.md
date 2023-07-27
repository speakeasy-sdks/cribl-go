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

    ctx := context.Background()
    res, err := s.GroupBundle.Get(ctx, operations.GetGroupBundleRequest{
        DeployRequest: &shared.DeployRequest{
            Version: "eum",
        },
        ID: "6a6d2d00-0355-4338-8ec0-86fa21e9152c",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GroupBundle != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetGroupBundleRequest](../../models/operations/getgroupbundlerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetGroupBundleResponse](../../models/operations/getgroupbundleresponse.md), error**

