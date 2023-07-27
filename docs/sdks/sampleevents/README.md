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
    res, err := s.SampleEvents.Post(ctx, shared.PreviewDataParams{
        CPUProfile: cribl.Bool(false),
        Dropped: cribl.Bool(false),
        Events: []map[string]interface{}{
            map[string]interface{}{
                "voluptate": "cum",
                "esse": "alias",
                "consequuntur": "architecto",
            },
            map[string]interface{}{
                "nemo": "magni",
                "quae": "quaerat",
                "suscipit": "rerum",
            },
            map[string]interface{}{
                "aliquam": "repudiandae",
                "unde": "excepturi",
            },
            map[string]interface{}{
                "facilis": "doloremque",
                "officiis": "nisi",
                "reprehenderit": "necessitatibus",
                "alias": "provident",
            },
        },
        InputID: cribl.String("ut"),
        Level: cribl.Int64(941776),
        Memory: cribl.Int64(814159),
        Mode: shared.PreviewDataParamsModeRouteAndSend,
        PipelineID: "saepe",
        SampleID: "assumenda",
        SamplePipelineID: cribl.String("exercitationem"),
        Timeout: cribl.Int64(347919),
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

