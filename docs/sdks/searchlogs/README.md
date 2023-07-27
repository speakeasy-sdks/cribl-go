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
    id := "esse"

    ctx := context.Background()
    res, err := s.SearchLogs.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchLogs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | search job id                                         |


### Response

**[*operations.GetSearchLogsResponse](../../models/operations/getsearchlogsresponse.md), error**

