# Roles

### Available Operations

* [Create](#create) - Create Role
* [Delete](#delete) - Delete Role
* [Get](#get) - Get Role by ID
* [ListRoles](#listroles) - Get a list of Role objects
* [Update](#update) - Update Role

## Create

Create Role

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
    res, err := s.Roles.Create(ctx, shared.Role{
        ID: "1263d48e-935c-42c9-a81f-30be3e43202d",
        Policy: []string{
            "magni",
            "beatae",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Roles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Role](../../models/shared/role.md)            | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateRoleResponse](../../models/operations/createroleresponse.md), error**


## Delete

Delete Role

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
    id := "aliquid"

    ctx := context.Background()
    res, err := s.Roles.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Roles != nil {
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

**[*operations.DeleteRoleResponse](../../models/operations/deleteroleresponse.md), error**


## Get

Get Role by ID

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
    id := "ad"

    ctx := context.Background()
    res, err := s.Roles.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Roles != nil {
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

**[*operations.GetRoleResponse](../../models/operations/getroleresponse.md), error**


## ListRoles

Get a list of Role objects

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
    res, err := s.Roles.ListRoles(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Roles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListRolesResponse](../../models/operations/listrolesresponse.md), error**


## Update

Update Role

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
    id := "voluptate"
    role := &shared.Role{
        ID: "65066418-70d9-4d21-b9ad-030c4ecc11a0",
        Policy: []string{
            "ratione",
            "iure",
            "tempora",
        },
    }

    ctx := context.Background()
    res, err := s.Roles.Update(ctx, id, role)
    if err != nil {
        log.Fatal(err)
    }

    if res.Roles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |
| `role`                                                | [*shared.Role](../../models/shared/role.md)           | :heavy_minus_sign:                                    | Role object to be updated                             |


### Response

**[*operations.UpdateRoleResponse](../../models/operations/updateroleresponse.md), error**

