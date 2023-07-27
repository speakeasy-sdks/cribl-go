# JobResults

### Available Operations

* [Get](#get) - Get results for a discover job by instance id

## Get

Get results for a discover job by instance id

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
    res, err := s.JobResults.Get(ctx, operations.GetJobResultsRequest{
        ID: "d85a9065-e628-4bdf-8203-2b6c879923b7",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetJobResults200ApplicationXNdjsonBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetJobResultsRequest](../../models/operations/getjobresultsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetJobResultsResponse](../../models/operations/getjobresultsresponse.md), error**

