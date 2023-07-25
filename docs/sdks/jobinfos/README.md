# JobInfos

### Available Operations

* [Get](#get) - Get info on jobs

## Get

Get info on jobs

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
    res, err := s.JobInfos.Get(ctx, operations.GetJobInfosRequest{
        CollectorID: cribl.String("velit"),
        ID: cribl.String("f5aeb779-9d22-4e8c-9f84-93825fdc42c8"),
        Limit: cribl.Int64(448482),
        Offset: cribl.Int64(413768),
        RunType: cribl.String("maxime"),
        State: cribl.String("sed"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.JobInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetJobInfosRequest](../../models/operations/getjobinfosrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetJobInfosResponse](../../models/operations/getjobinfosresponse.md), error**

