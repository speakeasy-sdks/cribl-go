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
    res, err := s.UserObject.Update(ctx, operations.UpdateUserObjectRequest{
        User: &shared.User{
            CurrentPassword: cribl.String("repellat"),
            Disabled: false,
            Email: "Brooke.Konopelski@hotmail.com",
            First: "laborum",
            ID: "2047c044-9e14-43f9-a19b-b7d40d5a11fa",
            Last: "modi",
            Password: cribl.String("nesciunt"),
            Roles: []string{
                "repudiandae",
                "commodi",
            },
            Username: "Candida.Heidenreich",
        },
        ID: "33f95c9d-2373-497c-b85b-5db4f500183f",
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

