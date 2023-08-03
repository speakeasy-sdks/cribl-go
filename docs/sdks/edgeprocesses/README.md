# EdgeProcesses

### Available Operations

* [GetRunDetails](#getrundetails) - Get details of a process running on the edge host
* [ListProcessRunningDetail](#listprocessrunningdetail) - Get a detailed list of processes running on the edge host

## GetRunDetails

Get details of a process running on the edge host

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
    pid := "asperiores"

    ctx := context.Background()
    res, err := s.EdgeProcesses.GetRunDetails(ctx, pid)
    if err != nil {
        log.Fatal(err)
    }

    if res.Processes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `pid`                                                 | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


### Response

**[*operations.GetProcessRunningDetailResponse](../../models/operations/getprocessrunningdetailresponse.md), error**


## ListProcessRunningDetail

Get a detailed list of processes running on the edge host

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.EdgeProcesses.ListProcessRunningDetail(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Processes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListProcessRunningDetailResponse](../../models/operations/listprocessrunningdetailresponse.md), error**

