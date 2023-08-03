# LogFileContent

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
    res, err := s.LogFileContent.Get(ctx, operations.GetLogFileContentRequest{
        EndOffset: cribl.Int64(276650),
        Et: cribl.Int64(240490),
        Filter: cribl.String("praesentium"),
        ID: "35bbc05a-23a4-45ce-bc5f-de10a0ce2169",
        Limit: cribl.Int64(892708),
        Lt: cribl.Int64(354821),
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetLogFileContentRequest](../../models/operations/getlogfilecontentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetLogFileContentResponse](../../models/operations/getlogfilecontentresponse.md), error**

