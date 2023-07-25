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
    res, err := s.PreviousCriblPackage.Get(ctx, operations.GetPreviousCriblPackageRequest{
        File: "harum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CriblPackage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetPreviousCriblPackageRequest](../../models/operations/getpreviouscriblpackagerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetPreviousCriblPackageResponse](../../models/operations/getpreviouscriblpackageresponse.md), error**

