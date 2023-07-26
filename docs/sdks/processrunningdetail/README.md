# ProcessRunningDetail

### Available Operations

* [Get](#get) - Get details of a process running on the edge host

## Get

Get details of a process running on the edge host

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
    res, err := s.ProcessRunningDetail.Get(ctx, operations.GetProcessRunningDetailRequest{
        Pid: "temporibus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Processes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.GetProcessRunningDetailRequest](../../models/operations/getprocessrunningdetailrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.GetProcessRunningDetailResponse](../../models/operations/getprocessrunningdetailresponse.md), error**

