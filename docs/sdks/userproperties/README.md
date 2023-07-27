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
    id := "natus"
    userProfile := &shared.UserProfile{
        Disabled: false,
        Email: "Shakira.Huels@gmail.com",
        First: "vitae",
        ID: "adb55f9e-5d75-41c9-be8f-7502bfdc3450",
        Last: "blanditiis",
        Password: cribl.String("modi"),
        Roles: []string{
            "a",
        },
        Username: "Assunta.Krajcik",
    }

    ctx := context.Background()
    res, err := s.UserProperties.Update(ctx, id, userProfile)
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

