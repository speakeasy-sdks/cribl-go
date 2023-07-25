# RouteListID

### Available Operations

* [Get](#get) - List all routes by id

## Get

List all routes by id

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
    res, err := s.RouteListID.Get(ctx, operations.GetRouteListIDRequest{
        ID: "61eca16e-f894-451b-976e-eeb518c4da1f",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetRouteListIDRequest](../../models/operations/getroutelistidrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetRouteListIDResponse](../../models/operations/getroutelistidresponse.md), error**

