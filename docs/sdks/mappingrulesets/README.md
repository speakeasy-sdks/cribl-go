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
    id := "rem"

    ctx := context.Background()
    res, err := s.MappingRulesets.Delete(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


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
    id := "atque"
    mappingRuleset := &shared.MappingRuleset{
        Active: cribl.Bool(false),
        Conf: &shared.MappingRulesetConf{
            Functions: []map[string]interface{}{
                map[string]interface{}{
                    "pariatur": "sapiente",
                    "quo": "incidunt",
                    "quod": "minus",
                },
                map[string]interface{}{
                    "id": "excepturi",
                    "occaecati": "libero",
                    "quo": "esse",
                    "hic": "maxime",
                },
                map[string]interface{}{
                    "soluta": "fugit",
                },
                map[string]interface{}{
                    "eligendi": "recusandae",
                    "veritatis": "aut",
                    "laudantium": "iusto",
                    "dolor": "voluptates",
                },
            },
        },
        ID: "42b006d6-7887-48ba-8581-a58208c54fef",
    }

    ctx := context.Background()
    res, err := s.MappingRulesets.Update(ctx, id, mappingRuleset)
    if err != nil {
        log.Fatal(err)
    }

    if res.MappingRulesets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                       | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ctx`                                                           | [context.Context](https://pkg.go.dev/context#Context)           | :heavy_check_mark:                                              | The context to use for the request.                             |
| `id`                                                            | *string*                                                        | :heavy_check_mark:                                              | Unique ID                                                       |
| `mappingRuleset`                                                | [*shared.MappingRuleset](../../models/shared/mappingruleset.md) | :heavy_minus_sign:                                              | MappingRuleset object to be updated                             |


### Response

**[*operations.UpdateMappingRulesetsResponse](../../models/operations/updatemappingrulesetsresponse.md), error**

