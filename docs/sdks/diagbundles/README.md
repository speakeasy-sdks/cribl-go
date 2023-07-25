# DiagBundles

### Available Operations

* [Get](#get) - Get list of existing diag bundles

## Get

Get list of existing diag bundles

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
    res, err := s.DiagBundles.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDiagBundles200ApplicationTarPlusGzipBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetDiagBundlesResponse](../../models/operations/getdiagbundlesresponse.md), error**

