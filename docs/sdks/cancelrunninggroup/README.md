# CancelRunningGroup

### Available Operations

* [Post](#post) - Cancel a running group upgrade

## Post

Cancel a running group upgrade

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
    group := "adipisci"

    ctx := context.Background()
    res, err := s.CancelRunningGroup.Post(ctx, group)
    if err != nil {
        log.Fatal(err)
    }

    if res.CancelRunningGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `group`                                               | *string*                                              | :heavy_check_mark:                                    | id of the group                                       |


### Response

**[*operations.PostCancelRunningGroupResponse](../../models/operations/postcancelrunninggroupresponse.md), error**

