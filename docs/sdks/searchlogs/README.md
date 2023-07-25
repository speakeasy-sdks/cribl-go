# SearchLogs

### Available Operations

* [Get](#get) - Get search logs

## Get

Get search logs

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
    res, err := s.SearchLogs.Get(ctx, operations.GetSearchLogsRequest{
        ID: "450a986a-495b-4ac7-87f0-6b28ecc86492",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchLogs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetSearchLogsRequest](../../models/operations/getsearchlogsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetSearchLogsResponse](../../models/operations/getsearchlogsresponse.md), error**

