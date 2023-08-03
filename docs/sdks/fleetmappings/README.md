# FleetMappings

### Available Operations

* [Create](#create) - Create MappingRuleset
* [Delete](#delete) - Delete MappingRuleset
* [Get](#get) - Get MappingRuleset by ID
* [ListFleetMappings](#listfleetmappings) - Get a list of MappingRuleset objects
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
    res, err := s.FleetMappings.Create(ctx, shared.MappingRuleset{
        Active: cribl.Bool(false),
        Conf: &shared.MappingRulesetConf{
            Functions: []map[string]interface{}{
                map[string]interface{}{
                    "corporis": "qui",
                    "expedita": "voluptatum",
                    "cupiditate": "minima",
                    "placeat": "enim",
                },
                map[string]interface{}{
                    "in": "minus",
                },
                map[string]interface{}{
                    "modi": "corporis",
                    "magnam": "voluptates",
                },
                map[string]interface{}{
                    "tempore": "aperiam",
                    "libero": "ratione",
                    "labore": "totam",
                    "occaecati": "voluptas",
                },
            },
        },
        ID: "c3ca5acf-be2f-4d57-8757-7929177deac6",
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

**[*operations.CreateFleetMappingResponse](../../models/operations/createfleetmappingresponse.md), error**


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
    id := "quaerat"

    ctx := context.Background()
    res, err := s.FleetMappings.Delete(ctx, id)
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
    id := "commodi"

    ctx := context.Background()
    res, err := s.FleetMappings.Get(ctx, id)
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


## ListFleetMappings

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
    res, err := s.FleetMappings.ListFleetMappings(ctx)
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

**[*operations.ListFleetMappingsResponse](../../models/operations/listfleetmappingsresponse.md), error**


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
    id := "officiis"
    mappingRuleset := &shared.MappingRuleset{
        Active: cribl.Bool(false),
        Conf: &shared.MappingRulesetConf{
            Functions: []map[string]interface{}{
                map[string]interface{}{
                    "exercitationem": "quam",
                    "dolorem": "modi",
                    "ipsa": "sint",
                },
                map[string]interface{}{
                    "sequi": "repudiandae",
                    "cum": "dicta",
                    "earum": "veniam",
                    "animi": "dolores",
                },
                map[string]interface{}{
                    "dicta": "consequuntur",
                    "necessitatibus": "nobis",
                    "ipsa": "ducimus",
                },
                map[string]interface{}{
                    "veritatis": "quasi",
                    "laboriosam": "pariatur",
                    "libero": "excepturi",
                    "occaecati": "nemo",
                },
            },
        },
        ID: "45fc95fa-8897-40e1-89db-b30fcb33ea05",
    }

    ctx := context.Background()
    res, err := s.FleetMappings.Update(ctx, id, mappingRuleset)
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

