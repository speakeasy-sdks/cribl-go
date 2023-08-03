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

    ctx := context.Background()
    res, err := s.LogFileContents.Get(ctx, operations.GetLogFileContentsRequest{
        EndOffset: cribl.Int64(103578),
        Et: cribl.Int64(33771),
        Filter: cribl.String("perferendis"),
        GroupID: "veritatis",
        ID: "9c6dc5e3-4762-4799-bfbb-e6949fb2bb4e",
        Limit: cribl.Int64(797254),
        Lt: cribl.Int64(664666),
        Offset: cribl.Int64(904440),
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

