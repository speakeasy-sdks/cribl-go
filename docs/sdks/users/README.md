# Users

## Overview

Actions related to users

### Available Operations

* [CreateUser](#createuser) - Create User
* [Delete](#delete) - Delete User
* [Get](#get) - Get User by ID
* [ListUsers](#listusers) - Get a list of User objects
* [UpdateInfo](#updateinfo) - Update User except for their roles
* [UpdateProperties](#updateproperties) - Update User properties – admin use only

## CreateUser

Create User

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
    res, err := s.Users.CreateUser(ctx, shared.User{
        CurrentPassword: cribl.String("itaque"),
        Disabled: false,
        Email: "Braeden.Farrell54@gmail.com",
        First: "non",
        ID: "7ec59e1f-67f3-4c4c-8e4b-6d7696ff3c57",
        Last: "eius",
        Password: cribl.String("dignissimos"),
        Roles: []string{
            "perferendis",
            "architecto",
        },
        Username: "Daphney_Hessel",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Users != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.User](../../models/shared/user.md)            | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateUserResponse](../../models/operations/createuserresponse.md), error**


## Delete

Delete User

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
    id := "dolore"

    ctx := context.Background()
    res, err := s.Users.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Users != nil {
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

**[*operations.DeleteUserIDResponse](../../models/operations/deleteuseridresponse.md), error**


## Get

Get User by ID

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
    id := "magnam"

    ctx := context.Background()
    res, err := s.Users.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Users != nil {
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

**[*operations.GetUserIDResponse](../../models/operations/getuseridresponse.md), error**


## ListUsers

Get a list of User objects

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
    res, err := s.Users.ListUsers(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Users != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListUsersResponse](../../models/operations/listusersresponse.md), error**


## UpdateInfo

Update User except for their roles

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
    id := "maiores"
    user := &shared.User{
        CurrentPassword: cribl.String("ipsam"),
        Disabled: false,
        Email: "Tyrel.Lakin51@gmail.com",
        First: "tempora",
        ID: "c3197e19-3a24-4546-bf94-874c2d5cc497",
        Last: "qui",
        Password: cribl.String("qui"),
        Roles: []string{
            "ratione",
        },
        Username: "Summer_Kertzmann",
    }

    ctx := context.Background()
    res, err := s.Users.UpdateInfo(ctx, id, user)
    if err != nil {
        log.Fatal(err)
    }

    if res.Users != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |
| `user`                                                | [*shared.User](../../models/shared/user.md)           | :heavy_minus_sign:                                    | User object                                           |


### Response

**[*operations.UpdateUserObjectResponse](../../models/operations/updateuserobjectresponse.md), error**


## UpdateProperties

Update User properties – admin use only

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
    id := "fugiat"
    userProfile := &shared.UserProfile{
        Disabled: false,
        Email: "Wilbert_Torphy@yahoo.com",
        First: "aut",
        ID: "0b979ef2-0387-4320-990c-cc1096400313",
        Last: "rerum",
        Password: cribl.String("ipsum"),
        Roles: []string{
            "quis",
            "eaque",
            "incidunt",
            "ut",
        },
        Username: "Vilma_Jacobi",
    }

    ctx := context.Background()
    res, err := s.Users.UpdateProperties(ctx, id, userProfile)
    if err != nil {
        log.Fatal(err)
    }

    if res.UserProfiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |
| `id`                                                      | *string*                                                  | :heavy_check_mark:                                        | Unique ID                                                 |
| `userProfile`                                             | [*shared.UserProfile](../../models/shared/userprofile.md) | :heavy_minus_sign:                                        | UserProfile object                                        |


### Response

**[*operations.UpdateUserPropertiesResponse](../../models/operations/updateuserpropertiesresponse.md), error**

