# SampleEvents

### Available Operations

* [Post](#post) - Sends sample events through a pipeline and returns the results

## Post

Sends sample events through a pipeline and returns the results

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
    res, err := s.SampleEvents.Post(ctx, shared.PreviewDataParams{
        CPUProfile: cribl.Bool(false),
        Dropped: cribl.Bool(false),
        Events: []map[string]interface{}{
            map[string]interface{}{
                "quaerat": "nesciunt",
                "molestiae": "adipisci",
                "eveniet": "ipsa",
            },
        },
        InputID: cribl.String("laboriosam"),
        Level: cribl.Int64(25321),
        Memory: cribl.Int64(288525),
        Mode: shared.PreviewDataParamsModeRoute,
        PipelineID: "excepturi",
        SampleID: "soluta",
        SamplePipelineID: cribl.String("voluptates"),
        Timeout: cribl.Int64(700112),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LiveData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.PreviewDataParams](../../models/shared/previewdataparams.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.PostSampleEventsResponse](../../models/operations/postsampleeventsresponse.md), error**

