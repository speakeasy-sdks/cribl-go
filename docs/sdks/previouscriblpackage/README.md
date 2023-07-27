# PreviousCriblPackage

### Available Operations

* [Get](#get) - Get the previously downloaded Cribl package

## Get

Get the previously downloaded Cribl package

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
    file := "suscipit"

    ctx := context.Background()
    res, err := s.PreviousCriblPackage.Get(ctx, file)
    if err != nil {
        log.Fatal(err)
    }

    if res.CriblPackage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `file`                                                | *string*                                              | :heavy_check_mark:                                    | Name of the file to be downloaded                     |


### Response

**[*operations.GetPreviousCriblPackageResponse](../../models/operations/getpreviouscriblpackageresponse.md), error**

