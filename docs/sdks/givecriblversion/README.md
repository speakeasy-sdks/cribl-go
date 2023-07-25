# GiveCriblVersion

### Available Operations

* [Post](#post) - Upgrade Cribl to a given version

## Post

Upgrade Cribl to a given version

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
    res, err := s.GiveCriblVersion.Post(ctx, operations.PostGiveCriblVersionRequest{
        Version: "labore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Cribl != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PostGiveCriblVersionRequest](../../models/operations/postgivecriblversionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.PostGiveCriblVersionResponse](../../models/operations/postgivecriblversionresponse.md), error**

