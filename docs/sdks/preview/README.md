# Preview

## Overview

Actions related to data preview

### Available Operations

* [CaptureLiveData](#capturelivedata) - Capture live incoming data
* [SendEvents](#sendevents) - Sends sample events through a pipeline and returns the results

## CaptureLiveData

Capture live incoming data

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
    res, err := s.Preview.CaptureLiveData(ctx, shared.CaptureParams{
        Filter: "dignissimos",
        Level: 458970,
        MaxEvents: 854115,
        StepDuration: cribl.Int64(322333),
        WorkerID: cribl.String("aspernatur"),
        WorkerThreshold: cribl.Int64(316501),
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

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.CaptureParams](../../models/shared/captureparams.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |


### Response

**[*operations.PostLiveDataResponse](../../models/operations/postlivedataresponse.md), error**


## SendEvents

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
    res, err := s.Preview.SendEvents(ctx, shared.PreviewDataParams{
        CPUProfile: cribl.Bool(false),
        Dropped: cribl.Bool(false),
        Events: []map[string]interface{}{
            map[string]interface{}{
                "delectus": "iusto",
            },
        },
        InputID: cribl.String("dignissimos"),
        Level: cribl.Int64(729828),
        Memory: cribl.Int64(72350),
        Mode: shared.PreviewDataParamsModePipe,
        PipelineID: "incidunt",
        SampleID: "accusamus",
        SamplePipelineID: cribl.String("saepe"),
        Timeout: cribl.Int64(734814),
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

