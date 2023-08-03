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
    version := "amet"

    ctx := context.Background()
    res, err := s.GiveCriblVersion.Post(ctx, version)
    if err != nil {
        log.Fatal(err)
    }

    if res.Cribl != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `version`                                             | *string*                                              | :heavy_check_mark:                                    | Version number                                        |


### Response

**[*operations.PostGiveCriblVersionResponse](../../models/operations/postgivecriblversionresponse.md), error**

