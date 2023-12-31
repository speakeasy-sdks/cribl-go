# EdgeLs

### Available Operations

* [GetDirectoryListing](#getdirectorylisting) - Get a directory listing of the given path

## GetDirectoryListing

Get a directory listing of the given path

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
    path := "nemo"

    ctx := context.Background()
    res, err := s.EdgeLs.GetDirectoryListing(ctx, path)
    if err != nil {
        log.Fatal(err)
    }

    if res.FilesystemEntries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `path`                                                | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


### Response

**[*operations.GetEdgeListingResponse](../../models/operations/getedgelistingresponse.md), error**

