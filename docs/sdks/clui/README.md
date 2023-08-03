# Clui

### Available Operations

* [ListCluisResults](#listcluisresults) - Get CLUI search results

## ListCluisResults

Get CLUI search results

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
    query := "minima"
    context := "beatae"

    ctx := context.Background()
    res, err := s.Clui.ListCluisResults(ctx, query, context)
    if err != nil {
        log.Fatal(err)
    }

    if res.CluiItems != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `query`                                               | *string*                                              | :heavy_check_mark:                                    | Search query                                          |
| `context`                                             | **string*                                             | :heavy_minus_sign:                                    | Search query context, either "stream" or "edge"       |


### Response

**[*operations.ListCluisResultsResponse](../../models/operations/listcluisresultsresponse.md), error**

