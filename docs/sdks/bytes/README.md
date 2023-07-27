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

    ctx := context.Background()
    res, err := s.Bytes.Get(ctx, operations.GetBytesRequest{
        BytesRequested: cribl.Int64(715208),
        Path: "voluptatum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SampleFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetBytesRequest](../../models/operations/getbytesrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetBytesResponse](../../models/operations/getbytesresponse.md), error**

