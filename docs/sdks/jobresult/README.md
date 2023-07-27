# JobResult

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
    group := "iste"
    id := "dicta"
    maxFiles := 552439

    ctx := context.Background()
    res, err := s.JobResult.Get(ctx, group, id, maxFiles)
    if err != nil {
        log.Fatal(err)
    }

    if res.JobResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `group`                                               | *string*                                              | :heavy_check_mark:                                    | Group the job belongs to                              |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Instance id of the job to get results for             |
| `maxFiles`                                            | **int64*                                              | :heavy_minus_sign:                                    | Maximum files to get job results                      |


### Response

**[*operations.GetJobResultResponse](../../models/operations/getjobresultresponse.md), error**

