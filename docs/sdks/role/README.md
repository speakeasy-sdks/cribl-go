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
    res, err := s.Role.Create(ctx, shared.Role{
        ID: "0d03f3de-ba29-47be-be90-bc40df868fd5",
        Policy: []string{
            "dolore",
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
    res, err := s.Role.Delete(ctx, operations.DeleteRoleRequest{
        ID: "05cb331d-492f-44f1-a7fb-0e0bf1f82179",
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
    res, err := s.Role.Get(ctx, operations.GetRoleRequest{
        ID: "78d0acca-77ae-4b7b-b021-a52046b64e99",
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
    res, err := s.Role.Update(ctx, operations.UpdateRoleRequest{
        Role: &shared.Role{
            ID: "fb0e67e0-94fd-4fed-9540-ef53a34a1b8f",
            Policy: []string{
                "perspiciatis",
                "natus",
                "reprehenderit",
                "dolorem",
            },
        },
        ID: "1adc05d8-5ae2-4dfb-b0fb-3874290d3365",
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

