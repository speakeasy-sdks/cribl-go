# EdgeListing

### Available Operations

* [Get](#get) - Get a directory listing of the given path

## Get

Get a directory listing of the given path

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
    res, err := s.EdgeListing.Get(ctx, operations.GetEdgeListingRequest{
        Path: "quasi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FilesystemEntries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetEdgeListingRequest](../../models/operations/getedgelistingrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetEdgeListingResponse](../../models/operations/getedgelistingresponse.md), error**

