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
    res, err := s.TaskErrors.Get(ctx, operations.GetTaskErrorsRequest{
        Group: "soluta",
        ID: "5744d08a-2267-4aae-a79e-3c71ad31becb",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskErrors != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetTaskErrorsRequest](../../models/operations/gettaskerrorsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetTaskErrorsResponse](../../models/operations/gettaskerrorsresponse.md), error**

