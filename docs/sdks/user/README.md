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
        CurrentPassword: cribl.String("nulla"),
        Disabled: false,
        Email: "Devon65@yahoo.com",
        First: "unde",
        ID: "86a495ba-c707-4f06-b28e-cc86492386f6",
        Last: "qui",
        Password: cribl.String("eligendi"),
        Roles: []string{
            "eum",
            "sint",
            "eligendi",
        },
        Username: "Douglas.Schamberger69",
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

