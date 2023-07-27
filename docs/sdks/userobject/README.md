# UserObject

### Available Operations

* [Get](#get) - Get a list of User objects
* [Update](#update) - Update User except for their roles

## Get

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
    res, err := s.UserObject.Get(ctx)
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

**[*operations.GetUserObjectResponse](../../models/operations/getuserobjectresponse.md), error**


## Update

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
    id := "enim"
    user := &shared.User{
        CurrentPassword: cribl.String("nihil"),
        Disabled: false,
        Email: "Una_Turcotte27@gmail.com",
        First: "quam",
        ID: "353f63c8-2093-479a-a69c-d5fbcf79da18",
        Last: "fuga",
        Password: cribl.String("odio"),
        Roles: []string{
            "magni",
            "eos",
            "harum",
        },
        Username: "Watson_Moen",
    }

    ctx := context.Background()
    res, err := s.UserObject.Update(ctx, id, user)
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

