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
    res, err := s.SampleContent.Get(ctx, operations.GetSampleContentRequest{
        ID: "f4f127fb-0e0b-4f1f-8217-978d0acca77a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SampleContents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetSampleContentRequest](../../models/operations/getsamplecontentrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetSampleContentResponse](../../models/operations/getsamplecontentresponse.md), error**

