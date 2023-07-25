# UserID

### Available Operations

* [Delete](#delete) - Delete User
* [Get](#get) - Get User by ID

## Delete

Delete User

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
    res, err := s.UserID.Delete(ctx, operations.DeleteUserIDRequest{
        ID: "04fa4470-7bf3-475b-8428-2821fdb2f69e",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteUserIDRequest](../../models/operations/deleteuseridrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


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
    res, err := s.UserID.Get(ctx, operations.GetUserIDRequest{
        ID: "59267c71-cc8d-43cd-8258-d0358a82c808",
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetUserIDRequest](../../models/operations/getuseridrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetUserIDResponse](../../models/operations/getuseridresponse.md), error**

