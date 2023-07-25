# ListCriblVersion

### Available Operations

* [Get](#get) - Get a list of Cribl versions available for upgrade

## Get

Get a list of Cribl versions available for upgrade

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
    res, err := s.ListCriblVersion.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpgradeResults != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetListCriblVersionResponse](../../models/operations/getlistcriblversionresponse.md), error**

