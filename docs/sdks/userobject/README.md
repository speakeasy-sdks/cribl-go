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

    ctx := context.Background()
    res, err := s.UserObject.Update(ctx, operations.UpdateUserObjectRequest{
        User: &shared.User{
            CurrentPassword: cribl.String("ratione"),
            Disabled: false,
            Email: "Eric.Jacobi57@hotmail.com",
            First: "architecto",
            ID: "80ff60eb-9a66-458e-a9a4-b843d382dbec",
            Last: "ducimus",
            Password: cribl.String("ad"),
            Roles: []string{
                "ea",
                "corrupti",
                "minus",
                "autem",
            },
            Username: "Aglae58",
        },
        ID: "468ce304-d884-49bf-8214-c337f96bb0c6",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateUserObjectRequest](../../models/operations/updateuserobjectrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateUserObjectResponse](../../models/operations/updateuserobjectresponse.md), error**

