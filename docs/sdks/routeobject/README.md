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

    ctx := context.Background()
    res, err := s.RouteObject.Update(ctx, operations.UpdateRouteObjectRequest{
        Routes: &shared.Routes{
            Comments: []map[string]interface{}{
                map[string]interface{}{
                    "iure": "voluptatibus",
                    "id": "qui",
                    "explicabo": "accusantium",
                },
                map[string]interface{}{
                    "nesciunt": "commodi",
                    "molestias": "atque",
                },
            },
            Groups: map[string]shared.RoutesGroups{
                "totam": shared.RoutesGroups{
                    Description: cribl.String("ipsam"),
                    Disabled: cribl.Bool(false),
                    Name: "Lillian Ruecker Jr.",
                },
            },
            ID: cribl.String("006bef49-21ec-4205-bb74-9366ac8ee0f2"),
            Routes: []map[string]interface{}{
                map[string]interface{}{
                    "veritatis": "provident",
                    "veniam": "quos",
                    "totam": "facere",
                    "eius": "doloremque",
                },
                map[string]interface{}{
                    "aut": "sequi",
                    "reiciendis": "neque",
                    "assumenda": "saepe",
                    "nobis": "est",
                },
                map[string]interface{}{
                    "natus": "molestiae",
                },
            },
        },
        ID: "be3e90bc-40df-4868-bd52-405cb331d492",
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

