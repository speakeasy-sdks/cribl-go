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
    allow := "autem"
    depth := 779409
    mode := "nesciunt"
    path := "illum"

    ctx := context.Background()
    res, err := s.LogFileList.Get(ctx, allow, depth, mode, path)
    if err != nil {
        log.Fatal(err)
    }

    if res.EdgeFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |
| `allow`                                                                 | **string*                                                               | :heavy_minus_sign:                                                      | query array[string] optional List of allowed-filename wildcard patterns |
| `depth`                                                                 | **int64*                                                                | :heavy_minus_sign:                                                      | Search depth for "manual" mode                                          |
| `mode`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | Discovery Mode (default is "auto")                                      |
| `path`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | Search directory for "manual" mode                                      |


### Response

**[*operations.GetLogFileListResponse](../../models/operations/getlogfilelistresponse.md), error**

