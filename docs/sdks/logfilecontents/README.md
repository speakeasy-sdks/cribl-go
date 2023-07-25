# LogFileContents

### Available Operations

* [Get](#get) - Get contents of the log file

## Get

Get contents of the log file

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
    res, err := s.LogFileContents.Get(ctx, operations.GetLogFileContentsRequest{
        EndOffset: cribl.Int64(120432),
        Et: cribl.Int64(926266),
        Filter: cribl.String("sapiente"),
        GroupID: "id",
        ID: "2198258f-d0a9-4eba-87f7-d3ef049640d6",
        Limit: cribl.Int64(632850),
        Lt: cribl.Int64(116705),
        Offset: cribl.Int64(538877),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LogFileContents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetLogFileContentsRequest](../../models/operations/getlogfilecontentsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetLogFileContentsResponse](../../models/operations/getlogfilecontentsresponse.md), error**

