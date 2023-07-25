# FleetMapping

### Available Operations

* [Create](#create) - Create MappingRuleset

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
    res, err := s.FleetMapping.Create(ctx, shared.MappingRuleset{
        Active: cribl.Bool(false),
        Conf: &shared.MappingRulesetConf{
            Functions: []map[string]interface{}{
                map[string]interface{}{
                    "saepe": "consequatur",
                },
                map[string]interface{}{
                    "debitis": "facere",
                    "quisquam": "cumque",
                },
                map[string]interface{}{
                    "dolorum": "deserunt",
                    "ad": "reiciendis",
                },
                map[string]interface{}{
                    "porro": "laborum",
                },
            },
        },
        ID: "bd905a97-2e05-4672-8227-b2d309470bf7",
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

