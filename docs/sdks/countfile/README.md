# CountFile

### Available Operations

* [Get](#get) - get the count of files of changed

## Get

get the count of files of changed

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
    group := "voluptatem"

    ctx := context.Background()
    res, err := s.CountFile.Get(ctx, group)
    if err != nil {
        log.Fatal(err)
    }

    if res.CountFile != nil {
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

**[*operations.GetCountFileResponse](../../models/operations/getcountfileresponse.md), error**

