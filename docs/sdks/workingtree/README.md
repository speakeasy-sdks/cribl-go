# WorkingTree

### Available Operations

* [Get](#get) - get the the working tree status

## Get

get the the working tree status

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
    res, err := s.WorkingTree.Get(ctx, operations.GetWorkingTreeRequest{
        Group: cribl.String("vero"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GitStatusResults != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetWorkingTreeRequest](../../models/operations/getworkingtreerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetWorkingTreeResponse](../../models/operations/getworkingtreeresponse.md), error**

