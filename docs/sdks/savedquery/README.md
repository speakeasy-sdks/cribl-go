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
        ID: "ddc1912e-bde6-44bf-8c54-69d4015dfa79",
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

