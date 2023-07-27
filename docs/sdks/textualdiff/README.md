# TextualDiff

### Available Operations

* [Get](#get) - get the textual diff for given commit

## Get

get the textual diff for given commit

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
    res, err := s.TextualDiff.Get(ctx, operations.GetTextualDiffRequest{
        Commit: cribl.String("deleniti"),
        Group: cribl.String("mollitia"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TextualDiff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetTextualDiffRequest](../../models/operations/gettextualdiffrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetTextualDiffResponse](../../models/operations/gettextualdiffresponse.md), error**

