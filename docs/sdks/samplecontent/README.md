# SampleContent

### Available Operations

* [Get](#get) - Get sample content by ID

## Get

Get sample content by ID

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
    id := "aut"

    ctx := context.Background()
    res, err := s.SampleContent.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.SampleContents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Sample ID                                             |


### Response

**[*operations.GetSampleContentResponse](../../models/operations/getsamplecontentresponse.md), error**

