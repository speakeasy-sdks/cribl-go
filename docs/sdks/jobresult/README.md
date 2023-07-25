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
    res, err := s.JobResult.Get(ctx, operations.GetJobResultRequest{
        Group: "quam",
        ID: "9af4d357-24cd-4b0f-8d28-1187d56844ed",
        MaxFiles: cribl.Int64(921127),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.JobResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetJobResultRequest](../../models/operations/getjobresultrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetJobResultResponse](../../models/operations/getjobresultresponse.md), error**

