# InputStatusID

### Available Operations

* [Get](#get) - Get InputStatus by ID

## Get

Get InputStatus by ID

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
    res, err := s.InputStatusID.Get(ctx, operations.GetInputStatusIDRequest{
        ID: "4eded85a-9065-4e62-8bdf-c2032b6c8799",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.InputStatuses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetInputStatusIDRequest](../../models/operations/getinputstatusidrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetInputStatusIDResponse](../../models/operations/getinputstatusidresponse.md), error**

