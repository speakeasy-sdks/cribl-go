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
                    "nihil": "neque",
                    "aspernatur": "eaque",
                    "corporis": "cupiditate",
                },
            },
        },
        ID: "0ccc1096-4003-413b-be50-44f65fe72dc4",
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
        ID: "077d0cc3-f408-4efc-95ce-b4d6e1eae0f7",
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
        ID: "5aedf2ac-ab58-4b99-9c92-6ddb589461e7",
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
                        "inventore": "optio",
                    },
                    map[string]interface{}{
                        "recusandae": "commodi",
                        "possimus": "provident",
                        "veniam": "sit",
                    },
                },
            },
            ID: "2f0ea930-b69f-47ac-af72-f88500904911",
        },
        ID: "60820788-8ec6-4618-bbfe-9659eb40ec16",
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

