# User

### Available Operations

* [CreateUser](#createuser) - Create User

## CreateUser

Create User

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
    res, err := s.User.CreateUser(ctx, shared.User{
        CurrentPassword: cribl.String("tenetur"),
        Disabled: false,
        Email: "Blake91@gmail.com",
        First: "aliquam",
        ID: "e080aa10-4186-4ec7-99e0-2f3702c5c8e2",
        Last: "fugiat",
        Password: cribl.String("sequi"),
        Roles: []string{
            "voluptates",
        },
        Username: "Lorenzo12",
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

