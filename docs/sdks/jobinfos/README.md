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
    res, err := s.JobInfos.Get(ctx, operations.GetJobInfosRequest{
        CollectorID: cribl.String("numquam"),
        ID: cribl.String("8a47f939-0c58-4880-983d-abf9ef3ffdd9"),
        Limit: cribl.Int64(966652),
        Offset: cribl.Int64(487765),
        RunType: cribl.String("voluptatibus"),
        State: cribl.String("aut"),
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

