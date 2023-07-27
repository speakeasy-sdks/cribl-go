# Bytes

### Available Operations

* [Get](#get) - Get some number of bytes from the file at the given path

## Get

Get some number of bytes from the file at the given path

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
    path := "perspiciatis"
    bytesRequested := 446394

    ctx := context.Background()
    res, err := s.Bytes.Get(ctx, path, bytesRequested)
    if err != nil {
        log.Fatal(err)
    }

    if res.SampleFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `path`                                                                             | *string*                                                                           | :heavy_check_mark:                                                                 | The path to the file to sample                                                     |
| `bytesRequested`                                                                   | **int64*                                                                           | :heavy_minus_sign:                                                                 | The number of bytes to return;   this value could be constrained by system limits. |


### Response

**[*operations.GetBytesResponse](../../models/operations/getbytesresponse.md), error**

