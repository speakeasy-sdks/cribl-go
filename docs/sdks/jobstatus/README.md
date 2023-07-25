# JobStatus

### Available Operations

* [Get](#get) - Get job status

## Get

Get job status

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
    res, err := s.JobStatus.Get(ctx, operations.GetJobStatusRequest{
        ID: "e13584f7-ae12-4c68-91f8-2ce115717230",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchJobStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetJobStatusRequest](../../models/operations/getjobstatusrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetJobStatusResponse](../../models/operations/getjobstatusresponse.md), error**

