# Samples

## Overview

Actions related to samples

### Available Operations

* [CaptureLiveData](#capturelivedata) - Capture live incoming data
* [Delete](#delete) - Delete DataSample
* [Get](#get) - Get DataSample by ID
* [GetContent](#getcontent) - Get sample content by ID
* [ListDataSample](#listdatasample) - Get a list of DataSample objects
* [Post](#post) - Create DataSample
* [SendEvents](#sendevents) - Sends sample events through a pipeline and returns the results
* [Update](#update) - Update DataSample

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
    res, err := s.Samples.CaptureLiveData(ctx, shared.CaptureParams{
        Filter: "expedita",
        Level: 266370,
        MaxEvents: 768244,
        StepDuration: cribl.Int64(368599),
        WorkerID: cribl.String("occaecati"),
        WorkerThreshold: cribl.Int64(219279),
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


## Delete

Delete DataSample

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
    id := "accusamus"

    ctx := context.Background()
    res, err := s.Samples.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.DataSamples != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


### Response

**[*operations.DeleteDataSampleIDResponse](../../models/operations/deletedatasampleidresponse.md), error**


## Get

Get DataSample by ID

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
    id := "quisquam"

    ctx := context.Background()
    res, err := s.Samples.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.DataSamples != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


### Response

**[*operations.GetDataSampleIDResponse](../../models/operations/getdatasampleidresponse.md), error**


## GetContent

Get sample content by ID

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
    id := "quasi"

    ctx := context.Background()
    res, err := s.Samples.GetContent(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.SampleContents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Sample ID                                             |


### Response

**[*operations.GetSampleContentResponse](../../models/operations/getsamplecontentresponse.md), error**


## ListDataSample

Get a list of DataSample objects

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
    res, err := s.Samples.ListDataSample(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.DataSamples != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListDataSampleResponse](../../models/operations/listdatasampleresponse.md), error**


## Post

Create DataSample

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
    res, err := s.Samples.Post(ctx, map[string]interface{}{
        "quo": "temporibus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DataSamples != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [map[string]interface{}](../../models//.md)           | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostDataSampleResponse](../../models/operations/postdatasampleresponse.md), error**


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
    res, err := s.Samples.SendEvents(ctx, shared.PreviewDataParams{
        CPUProfile: cribl.Bool(false),
        Dropped: cribl.Bool(false),
        Events: []map[string]interface{}{
            map[string]interface{}{
                "quibusdam": "ipsa",
                "accusamus": "placeat",
                "quam": "similique",
            },
            map[string]interface{}{
                "saepe": "facere",
                "nobis": "at",
                "molestias": "aut",
                "temporibus": "tenetur",
            },
            map[string]interface{}{
                "numquam": "corrupti",
                "similique": "dolore",
            },
        },
        InputID: cribl.String("esse"),
        Level: cribl.Int64(968205),
        Memory: cribl.Int64(611792),
        Mode: shared.PreviewDataParamsModePipe,
        PipelineID: "occaecati",
        SampleID: "aut",
        SamplePipelineID: cribl.String("impedit"),
        Timeout: cribl.Int64(325855),
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


## Update

Update DataSample

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
    id := "quos"
    requestBody := map[string]interface{}{
        "quas": "voluptatem",
        "provident": "quas",
        "ipsum": "vero",
    }

    ctx := context.Background()
    res, err := s.Samples.Update(ctx, id, requestBody)
    if err != nil {
        log.Fatal(err)
    }

    if res.DataSamples != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |
| `requestBody`                                         | map[string]*interface{}*                              | :heavy_minus_sign:                                    | DataSample object to be updated                       |


### Response

**[*operations.UpdateDataSampleIDResponse](../../models/operations/updatedatasampleidresponse.md), error**

