# Cluis

### Available Operations

* [Get](#get) - Get CLUI search results

## Get

Get CLUI search results

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
    res, err := s.Cluis.Get(ctx, operations.GetCluisRequest{
        Context: cribl.String("nemo"),
        Query: "asperiores",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CluiItems != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetCluisRequest](../../models/operations/getcluisrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetCluisResponse](../../models/operations/getcluisresponse.md), error**

