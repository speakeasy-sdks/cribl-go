# HostMetadataStructure

### Available Operations

* [Get](#get) - Get the host's metadata structure

## Get

Get the host's metadata structure

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
    res, err := s.HostMetadataStructure.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.EdgeMetadatas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetHostMetadataStructureResponse](../../models/operations/gethostmetadatastructureresponse.md), error**

