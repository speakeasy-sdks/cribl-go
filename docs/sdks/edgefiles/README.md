# EdgeFiles

### Available Operations

* [ListEdgeHostFiles](#listedgehostfiles) - Get details about a file on the edge host.

## ListEdgeHostFiles

Get details about a file on the edge host.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.EdgeFiles.ListEdgeHostFiles(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.EdgeHostFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListEdgeHostFilesResponse](../../models/operations/listedgehostfilesresponse.md), error**

