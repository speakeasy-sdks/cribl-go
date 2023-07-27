# WorkingTree

### Available Operations

* [Get](#get) - get the the working tree status

## Get

get the the working tree status

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
    group := "molestias"

    ctx := context.Background()
    res, err := s.WorkingTree.Get(ctx, group)
    if err != nil {
        log.Fatal(err)
    }

    if res.GitStatusResults != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `group`                                               | **string*                                             | :heavy_minus_sign:                                    | Group ID                                              |


### Response

**[*operations.GetWorkingTreeResponse](../../models/operations/getworkingtreeresponse.md), error**

