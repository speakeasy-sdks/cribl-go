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
    res, err := s.MappingRuleset.Create(ctx, shared.MappingRuleset{
        Active: cribl.Bool(false),
        Conf: &shared.MappingRulesetConf{
            Functions: []map[string]interface{}{
                map[string]interface{}{
                    "consequatur": "delectus",
                    "incidunt": "dolore",
                },
            },
        },
        ID: "5accf667-aaf9-4bba-9185-fe431d6bf5c8",
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
    res, err := s.MappingRuleset.Delete(ctx, operations.DeleteMappingRulesetRequest{
        ID: "38fbb8c2-0cb6-47fc-8b42-5e99e6234c9f",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteMappingRulesetRequest](../../models/operations/deletemappingrulesetrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


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
    res, err := s.MappingRuleset.Get(ctx, operations.GetMappingRulesetRequest{
        ID: "7b79dfeb-77a5-4c38-94ba-f91e506ef890",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetMappingRulesetRequest](../../models/operations/getmappingrulesetrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


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
    res, err := s.MappingRuleset.Update(ctx, operations.UpdateMappingRulesetRequest{
        MappingRuleset: &shared.MappingRuleset{
            Active: cribl.Bool(false),
            Conf: &shared.MappingRulesetConf{
                Functions: []map[string]interface{}{
                    map[string]interface{}{
                        "tempora": "nam",
                        "numquam": "odio",
                    },
                    map[string]interface{}{
                        "maiores": "veritatis",
                        "autem": "earum",
                    },
                    map[string]interface{}{
                        "ex": "possimus",
                        "nesciunt": "corrupti",
                    },
                },
            },
            ID: "5a3c4ac6-31b9-49e2-aced-8f9fdb9410f6",
        },
        ID: "3bbf8178-37b0-41af-9d78-8624189eb448",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateMappingRulesetRequest](../../models/operations/updatemappingrulesetrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.UpdateMappingRulesetResponse](../../models/operations/updatemappingrulesetresponse.md), error**

