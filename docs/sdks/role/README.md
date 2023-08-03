# Role

### Available Operations

* [Create](#create) - Create Role
* [Delete](#delete) - Delete Role
* [Get](#get) - Get Role by ID
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
    res, err := s.Role.Create(ctx, shared.Role{
        ID: "355338ce-c086-4fa2-9e91-52cb3119167b",
        Policy: []string{
            "necessitatibus",
            "ipsum",
            "impedit",
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
    id := "quos"

    ctx := context.Background()
    res, err := s.Role.Delete(ctx, id)
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
    id := "illum"

    ctx := context.Background()
    res, err := s.Role.Get(ctx, id)
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
    id := "distinctio"
    role := &shared.Role{
        ID: "03408d6d-364f-4fd4-9590-6d1263d48e93",
        Policy: []string{
            "quod",
            "dolores",
        },
    }

    ctx := context.Background()
    res, err := s.Role.Update(ctx, id, role)
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

