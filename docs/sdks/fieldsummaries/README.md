# FieldSummaries

### Available Operations

* [Get](#get) - List field summaries

## Get

List field summaries

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
    res, err := s.FieldSummaries.Get(ctx, operations.GetFieldSummariesRequest{
        ID: "418e3bb9-1c8d-4975-a0e8-419d8f84f144",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FieldSummaries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetFieldSummariesRequest](../../models/operations/getfieldsummariesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetFieldSummariesResponse](../../models/operations/getfieldsummariesresponse.md), error**

