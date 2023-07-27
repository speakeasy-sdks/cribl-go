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

    ctx := context.Background()
    res, err := s.Groups.Get(ctx, operations.GetGroupsRequest{
        Fields: cribl.String("quidem"),
        Product: cribl.String("non"),
    })
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
| `request`                                                                  | [operations.GetGroupsRequest](../../models/operations/getgroupsrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetGroupsResponse](../../models/operations/getgroupsresponse.md), error**

