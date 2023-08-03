# Groups

### Available Operations

* [Get](#get) - Get a list of ConfigGroup objects

## Get

Get a list of ConfigGroup objects

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
    fields := "saepe"
    product := "delectus"

    ctx := context.Background()
    res, err := s.Groups.Get(ctx, fields, product)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfigGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `fields`                                                                   | **string*                                                                  | :heavy_minus_sign:                                                         | additional fields to add to results: git.commit, git.localChanges, git.log |
| `product`                                                                  | **string*                                                                  | :heavy_minus_sign:                                                         | filter to specific product: "stream" or "edge"                             |


### Response

**[*operations.GetGroupsResponse](../../models/operations/getgroupsresponse.md), error**

