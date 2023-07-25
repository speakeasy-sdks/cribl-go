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
    res, err := s.MappingRulesets.Delete(ctx, operations.DeleteMappingRulesetsRequest{
        ID: "c9b2ad32-dafe-481a-88f4-444573fecd47",
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
    res, err := s.MappingRulesets.Update(ctx, operations.UpdateMappingRulesetsRequest{
        MappingRuleset: &shared.MappingRuleset{
            Active: cribl.Bool(false),
            Conf: &shared.MappingRulesetConf{
                Functions: []map[string]interface{}{
                    map[string]interface{}{
                        "nesciunt": "earum",
                        "eum": "dolor",
                    },
                },
            },
            ID: "c8209379-aa69-4cd5-bbcf-79da18a7822b",
        },
        ID: "f95894e6-861a-4db5-9f9e-5d751c9fe8f7",
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

