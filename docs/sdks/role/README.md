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
        ID: "85b26059-1d74-45e3-8205-9c9c3f567e0e",
        Policy: []string{
            "corporis",
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

    ctx := context.Background()
    res, err := s.Role.Delete(ctx, operations.DeleteRoleRequest{
        ID: "2765b1d6-2fcd-4ace-9f01-216ce2239e8f",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.DeleteRoleRequest](../../models/operations/deleterolerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


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

    ctx := context.Background()
    res, err := s.Role.Get(ctx, operations.GetRoleRequest{
        ID: "25cd0d19-d959-4f43-9e39-266cbd95f7aa",
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.GetRoleRequest](../../models/operations/getrolerequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


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

    ctx := context.Background()
    res, err := s.Role.Update(ctx, operations.UpdateRoleRequest{
        Role: &shared.Role{
            ID: "2b241136-95d1-4e66-98fc-c4596217c297",
            Policy: []string{
                "aliquid",
                "voluptate",
            },
        },
        ID: "63342540-38bf-4b59-b1e9-8190557389ce",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.UpdateRoleRequest](../../models/operations/updaterolerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.UpdateRoleResponse](../../models/operations/updateroleresponse.md), error**

