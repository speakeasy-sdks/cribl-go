# FunctionID

### Available Operations

* [Get](#get) - Get Function by ID

## Get

Get Function by ID

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
    res, err := s.FunctionID.Get(ctx, operations.GetFunctionIDRequest{
        ID: "75d2367f-e1a0-4cc8-9f79-f0a396d90c36",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Functions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetFunctionIDRequest](../../models/operations/getfunctionidrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetFunctionIDResponse](../../models/operations/getfunctionidresponse.md), error**

