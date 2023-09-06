# Routes

## Overview

Actions related to routes

### Available Operations

* [Get](#get) - List all routes by id
* [ListRouteLists](#listroutelists) - List all routes
* [Update](#update) - Add, delete or update the routes with the required content.

## Get

List all routes by id

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
    id := "voluptatibus"

    ctx := context.Background()
    res, err := s.Routes.Get(ctx, id)
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


### Response

**[*operations.GetRouteListIDResponse](../../models/operations/getroutelistidresponse.md), error**


## ListRouteLists

List all routes

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
    res, err := s.Routes.ListRouteLists(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Routes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListRouteListsResponse](../../models/operations/listroutelistsresponse.md), error**


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
    id := "molestias"
    routes := &shared.Routes{
        Comments: []map[string]interface{}{
            map[string]interface{}{
                "officia": "libero",
            },
        },
        Groups: map[string]shared.RoutesGroups{
            "totam": shared.RoutesGroups{
                Description: cribl.String("sequi"),
                Disabled: cribl.Bool(false),
                Name: "Gertrude Ryan",
            },
        },
        ID: cribl.String("3ffda9e0-6bee-4482-9c1f-c0e115c80bff"),
        Routes: []map[string]interface{}{
            map[string]interface{}{
                "iste": "dicta",
            },
        },
    }

    ctx := context.Background()
    res, err := s.Routes.Update(ctx, id, routes)
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

