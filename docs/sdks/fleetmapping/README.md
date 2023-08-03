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
    res, err := s.FleetMapping.Create(ctx, shared.MappingRuleset{
        Active: cribl.Bool(false),
        Conf: &shared.MappingRulesetConf{
            Functions: []map[string]interface{}{
                map[string]interface{}{
                    "dignissimos": "libero",
                    "illo": "ab",
                },
                map[string]interface{}{
                    "accusamus": "saepe",
                    "tempore": "veniam",
                },
                map[string]interface{}{
                    "reiciendis": "earum",
                },
                map[string]interface{}{
                    "praesentium": "nemo",
                    "repellat": "quisquam",
                },
            },
        },
        ID: "37814d4c-98e0-4c2b-b89e-b75dad636c60",
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

