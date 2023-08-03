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
    res, err := s.User.CreateUser(ctx, shared.User{
        CurrentPassword: cribl.String("nihil"),
        Disabled: false,
        Email: "Maymie.Auer19@yahoo.com",
        First: "eos",
        ID: "a4da37cb-aaf4-4452-8484-2c9b2ad32daf",
        Last: "necessitatibus",
        Password: cribl.String("voluptatum"),
        Roles: []string{
            "mollitia",
        },
        Username: "Justina.Kutch27",
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

