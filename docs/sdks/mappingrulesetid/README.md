# MappingRulesetID

### Available Operations

* [Get](#get) - Get MappingRuleset by ID

## Get

Get MappingRuleset by ID

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
    res, err := s.MappingRulesetID.Get(ctx, operations.GetMappingRulesetIDRequest{
        ID: "73f5033f-19db-4f12-9ce4-152eab9cd7e5",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MappingRulesets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetMappingRulesetIDRequest](../../models/operations/getmappingrulesetidrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetMappingRulesetIDResponse](../../models/operations/getmappingrulesetidresponse.md), error**

