# SavedQuery

### Available Operations

* [Get](#get) - Get SavedQuery by ID

## Get

Get SavedQuery by ID

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
    res, err := s.SavedQuery.Get(ctx, operations.GetSavedQueryRequest{
        ID: "cd02756c-354a-4a43-ab47-e1763c5208c2",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SavedQuery != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetSavedQueryRequest](../../models/operations/getsavedqueryrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetSavedQueryResponse](../../models/operations/getsavedqueryresponse.md), error**

