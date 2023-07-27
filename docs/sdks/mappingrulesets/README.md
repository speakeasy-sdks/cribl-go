# MappingRulesets

### Available Operations

* [Delete](#delete) - Delete MappingRuleset
* [Get](#get) - Get a list of MappingRuleset objects
* [Update](#update) - Update MappingRuleset

## Delete

Delete MappingRuleset

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
    res, err := s.MappingRulesets.Delete(ctx, operations.DeleteMappingRulesetsRequest{
        ID: "224a6a0e-123b-4784-bec5-9e1f67f3c4cc",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.DeleteMappingRulesetsRequest](../../models/operations/deletemappingrulesetsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.DeleteMappingRulesetsResponse](../../models/operations/deletemappingrulesetsresponse.md), error**


## Get

Get a list of MappingRuleset objects

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.MappingRulesets.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.MappingRulesets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetMappingRulesetsResponse](../../models/operations/getmappingrulesetsresponse.md), error**


## Update

Update MappingRuleset

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
    res, err := s.MappingRulesets.Update(ctx, operations.UpdateMappingRulesetsRequest{
        MappingRuleset: &shared.MappingRuleset{
            Active: cribl.Bool(false),
            Conf: &shared.MappingRulesetConf{
                Functions: []map[string]interface{}{
                    map[string]interface{}{
                        "libero": "suscipit",
                        "illum": "iusto",
                    },
                    map[string]interface{}{
                        "sint": "aliquid",
                        "repellat": "sapiente",
                    },
                    map[string]interface{}{
                        "eligendi": "ullam",
                    },
                    map[string]interface{}{
                        "eius": "dignissimos",
                        "corporis": "perferendis",
                    },
                },
            },
            ID: "1357e44f-51f8-4b08-8c31-97e193a24546",
        },
        ID: "7f94874c-2d5c-4c49-b223-3e66bd8fe5d0",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateMappingRulesetsRequest](../../models/operations/updatemappingrulesetsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.UpdateMappingRulesetsResponse](../../models/operations/updatemappingrulesetsresponse.md), error**

