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
        EndOffset: cribl.Int64(981142),
        Et: cribl.Int64(35219),
        Filter: cribl.String("sed"),
        GroupID: "eos",
        ID: "25194db5-5410-4adc-a69a-f90a26c7cdc9",
        Limit: cribl.Int64(536999),
        Lt: cribl.Int64(94288),
        Offset: cribl.Int64(978857),
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

