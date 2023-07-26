# UserProperties

### Available Operations

* [Update](#update) - Update User properties – admin use only

## Update

Update User properties – admin use only

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
    res, err := s.UserProperties.Update(ctx, operations.UpdateUserPropertiesRequest{
        UserProfile: &shared.UserProfile{
            Disabled: false,
            Email: "Thora.Fisher@yahoo.com",
            First: "facere",
            ID: "b1344ba9-f78a-45c0-ad7a-ab62e97261fb",
            Last: "alias",
            Password: cribl.String("quod"),
            Roles: []string{
                "corrupti",
                "temporibus",
            },
            Username: "Brooklyn.Klein10",
        },
        ID: "996b5b4b-50ee-4f71-ab7a-7ab0344b1710",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserProfiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateUserPropertiesRequest](../../models/operations/updateuserpropertiesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.UpdateUserPropertiesResponse](../../models/operations/updateuserpropertiesresponse.md), error**

