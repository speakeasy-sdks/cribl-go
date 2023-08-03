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
        CollectorID: cribl.String("aliquid"),
        ID: cribl.String("6c723ffd-a9e0-46be-a482-5c1fc0e115c8"),
        Limit: cribl.Int64(17060),
        Offset: cribl.Int64(704271),
        RunType: cribl.String("reiciendis"),
        State: cribl.String("a"),
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

