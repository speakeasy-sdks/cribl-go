# TaskError

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
    res, err := s.TaskError.Get(ctx, operations.GetTaskErrorRequest{
        ID: "02ac1704-bf1c-4c9f-861a-ae5eb5f0c492",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetTaskErrorRequest](../../models/operations/gettaskerrorrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetTaskErrorResponse](../../models/operations/gettaskerrorresponse.md), error**

