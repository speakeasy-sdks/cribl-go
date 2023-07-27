# TaskErrors

### Available Operations

* [Get](#get) - Get Task errors for a job by id

## Get

Get Task errors for a job by id

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
    group := "autem"
    id := "vel"

    ctx := context.Background()
    res, err := s.TaskErrors.Get(ctx, group, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskErrors != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `group`                                               | *string*                                              | :heavy_check_mark:                                    | Group the job belongs to                              |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Instance id of the job to get results for             |


### Response

**[*operations.GetTaskErrorsResponse](../../models/operations/gettaskerrorsresponse.md), error**

