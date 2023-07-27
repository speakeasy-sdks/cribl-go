# MappingRuleset

### Available Operations

* [Create](#create) - Create MappingRuleset
* [Delete](#delete) - Delete MappingRuleset
* [Get](#get) - Get MappingRuleset by ID
* [Update](#update) - Update MappingRuleset

## Create

Create MappingRuleset

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
    res, err := s.MappingRuleset.Create(ctx, shared.MappingRuleset{
        Active: cribl.Bool(false),
        Conf: &shared.MappingRulesetConf{
            Functions: []map[string]interface{}{
                map[string]interface{}{
                    "dolor": "rerum",
                    "sed": "accusamus",
                    "optio": "delectus",
                    "minus": "quo",
                },
                map[string]interface{}{
                    "asperiores": "voluptatum",
                    "iste": "corporis",
                    "accusantium": "illo",
                },
            },
        },
        ID: "0f5dd3d6-fa18-404e-94c8-2f168a363c88",
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

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [shared.MappingRuleset](../../models/shared/mappingruleset.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |


### Response

**[*operations.CreateMappingRulesetResponse](../../models/operations/createmappingrulesetresponse.md), error**


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
    id := "ducimus"

    ctx := context.Background()
    res, err := s.MappingRuleset.Delete(ctx, id)
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

**[*operations.DeleteMappingRulesetResponse](../../models/operations/deletemappingrulesetresponse.md), error**


## Get

Get MappingRuleset by ID

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
    id := "adipisci"

    ctx := context.Background()
    res, err := s.MappingRuleset.Get(ctx, id)
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

**[*operations.GetMappingRulesetResponse](../../models/operations/getmappingrulesetresponse.md), error**


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
    id := "recusandae"
    mappingRuleset := &shared.MappingRuleset{
        Active: cribl.Bool(false),
        Conf: &shared.MappingRulesetConf{
            Functions: []map[string]interface{}{
                map[string]interface{}{
                    "numquam": "sequi",
                    "voluptatum": "sit",
                    "rerum": "veritatis",
                },
                map[string]interface{}{
                    "autem": "quidem",
                    "totam": "porro",
                    "deserunt": "magni",
                    "nihil": "voluptas",
                },
            },
        },
        ID: "a60a04c4-95cc-4699-971b-51c1bdb1cf4b",
    }

    ctx := context.Background()
    res, err := s.MappingRuleset.Update(ctx, id, mappingRuleset)
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

**[*operations.UpdateMappingRulesetResponse](../../models/operations/updatemappingrulesetresponse.md), error**

