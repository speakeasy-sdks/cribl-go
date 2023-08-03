# RouteObject

### Available Operations

* [Update](#update) - Add, delete or update the routes with the required content.

## Update

Add, delete or update the routes with the required content.

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
    id := "excepturi"
    routes := &shared.Routes{
        Comments: []map[string]interface{}{
            map[string]interface{}{
                "dicta": "sapiente",
                "ipsum": "consequatur",
                "soluta": "necessitatibus",
            },
            map[string]interface{}{
                "recusandae": "labore",
            },
            map[string]interface{}{
                "magni": "aperiam",
            },
            map[string]interface{}{
                "illum": "iusto",
            },
        },
        Groups: map[string]shared.RoutesGroups{
            "beatae": shared.RoutesGroups{
                Description: cribl.String("aliquid"),
                Disabled: cribl.Bool(false),
                Name: "Colleen Kautzer III",
            },
        },
        ID: cribl.String("641870d9-d21f-49ad-830c-4ecc11a08364"),
        Routes: []map[string]interface{}{
            map[string]interface{}{
                "voluptatem": "suscipit",
                "laudantium": "facilis",
                "laudantium": "ullam",
            },
        },
    }

    ctx := context.Background()
    res, err := s.RouteObject.Update(ctx, id, routes)
    if err != nil {
        log.Fatal(err)
    }

    if res.Routes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `id`                                                                       | *string*                                                                   | :heavy_check_mark:                                                         | There is only one route entity and it should be accessed with id: default. |
| `routes`                                                                   | [*shared.Routes](../../models/shared/routes.md)                            | :heavy_minus_sign:                                                         | Routes object                                                              |


### Response

**[*operations.UpdateRouteObjectResponse](../../models/operations/updaterouteobjectresponse.md), error**

