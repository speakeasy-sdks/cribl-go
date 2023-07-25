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
            Email: "Myah.Swift45@hotmail.com",
            First: "aliquid",
            ID: "b7206dab-7500-452a-9647-edc439ed8c43",
            Last: "explicabo",
            Password: cribl.String("eaque"),
            Roles: []string{
                "incidunt",
                "quae",
                "eos",
                "eius",
            },
            Username: "Alena27",
        },
        ID: "87ac693b-94c3-4b9d-a488-d795aa42fc40",
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

