# QuerySnippet

### Available Operations

* [Apply](#apply) - Applies a query snippet on a set of input events for preview

## Apply

Applies a query snippet on a set of input events for preview

### Example Usage

```go
package main

import(
	"context"
	"log"
	"cribl"
	"cribl/pkg/models/shared"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.QuerySnippet.Apply(ctx, shared.PreviewRequestBody{
        Events: []interface{}{
            "iste",
        },
        Options: &shared.PreviewOptions{},
        Query: "labore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PreviewResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.PreviewRequestBody](../../models/shared/previewrequestbody.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.ApplyQuerySnippetResponse](../../models/operations/applyquerysnippetresponse.md), error**

