# ChangedFiles

### Available Operations

* [Get](#get) - get the files changed

## Get

get the files changed

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
    id := "ullam"
    group := "in"

    ctx := context.Background()
    res, err := s.ChangedFiles.Get(ctx, id, group)
    if err != nil {
        log.Fatal(err)
    }

    if res.ChangedFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | **string*                                             | :heavy_minus_sign:                                    | Commit ID                                             |
| `group`                                               | **string*                                             | :heavy_minus_sign:                                    | Group ID                                              |


### Response

**[*operations.GetChangedFilesResponse](../../models/operations/getchangedfilesresponse.md), error**

