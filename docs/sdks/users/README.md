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
        CurrentPassword: cribl.String("voluptates"),
        Disabled: false,
        Email: "Billy39@yahoo.com",
        First: "necessitatibus",
        ID: "2813fa4a-41c4-480d-bf21-32af03102d51",
        Last: "magnam",
        Password: cribl.String("delectus"),
        Roles: []string{
            "numquam",
        },
        Username: "Noelia_Rolfson",
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
    id := "quae"

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
    id := "deleniti"

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
    id := "expedita"
    user := &shared.User{
        CurrentPassword: cribl.String("hic"),
        Disabled: false,
        Email: "Gracie62@yahoo.com",
        First: "ea",
        ID: "a4f77a87-ee3e-44be-b52c-65b34418e3bb",
        Last: "iste",
        Password: cribl.String("illo"),
        Roles: []string{
            "minus",
        },
        Username: "Kay.Smith",
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
    id := "enim"
    userProfile := &shared.UserProfile{
        Disabled: false,
        Email: "Alvah.Von@hotmail.com",
        First: "quae",
        ID: "9d8f84f1-44f3-4e07-adcc-4aa5f3cabd90",
        Last: "ipsam",
        Password: cribl.String("officia"),
        Roles: []string{
            "cupiditate",
        },
        Username: "Imelda.Crona32",
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

