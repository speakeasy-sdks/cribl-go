# LogFileList

### Available Operations

* [Get](#get) - list log files

## Get

list log files

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
    res, err := s.LogFileList.Get(ctx, operations.GetLogFileListRequest{
        Allow: cribl.String("voluptatem"),
        Depth: cribl.Int64(396184),
        Mode: cribl.String("laudantium"),
        Path: cribl.String("unde"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EdgeFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetLogFileListRequest](../../models/operations/getlogfilelistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetLogFileListResponse](../../models/operations/getlogfilelistresponse.md), error**

