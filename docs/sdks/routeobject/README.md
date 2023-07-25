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
    res, err := s.RouteObject.Update(ctx, operations.UpdateRouteObjectRequest{
        Routes: &shared.Routes{
            Comments: []map[string]interface{}{
                map[string]interface{}{
                    "dolor": "nemo",
                    "quis": "quasi",
                    "odit": "delectus",
                    "doloremque": "laboriosam",
                },
                map[string]interface{}{
                    "dolore": "itaque",
                    "enim": "nam",
                    "dignissimos": "consequuntur",
                    "sapiente": "alias",
                },
                map[string]interface{}{
                    "nemo": "dolore",
                    "corrupti": "exercitationem",
                    "commodi": "laudantium",
                    "est": "consequatur",
                },
            },
            Groups: map[string]shared.RoutesGroups{
                "dolores": shared.RoutesGroups{
                    Description: cribl.String("labore"),
                    Disabled: cribl.Bool(false),
                    Name: "Michael Barrows PhD",
                },
                "autem": shared.RoutesGroups{
                    Description: cribl.String("voluptates"),
                    Disabled: cribl.Bool(false),
                    Name: "Terrance Glover",
                },
            },
            ID: cribl.String("645d0308-4fbb-4a5c-8eff-5cb01fe51e52"),
            Routes: []map[string]interface{}{
                map[string]interface{}{
                    "eius": "ad",
                    "mollitia": "minus",
                    "quos": "explicabo",
                },
                map[string]interface{}{
                    "praesentium": "ullam",
                    "maiores": "corrupti",
                    "libero": "placeat",
                },
                map[string]interface{}{
                    "placeat": "animi",
                },
            },
        },
        ID: "ba8da412-7dd5-497f-b471-1aa1bc74b86c",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Routes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateRouteObjectRequest](../../models/operations/updaterouteobjectrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateRouteObjectResponse](../../models/operations/updaterouteobjectresponse.md), error**

