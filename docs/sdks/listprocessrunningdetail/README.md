# ListProcessRunningDetail

### Available Operations

* [Get](#get) - Get a detailed list of processes running on the edge host

## Get

Get a detailed list of processes running on the edge host

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ListProcessRunningDetail.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Processes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetListProcessRunningDetailResponse](../../models/operations/getlistprocessrunningdetailresponse.md), error**

