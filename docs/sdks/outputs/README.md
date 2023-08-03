# Outputs

## Overview

Actions related to outputs

### Available Operations

* [Create](#create) - Create Output
* [Delete](#delete) - Delete Output
* [DeletePQ](#deletepq) - Delete destination persistent queue
* [Get](#get) - Get Output by ID
* [GetLatestPQ](#getlatestpq) - Get status of latest clear PQ job for an output
* [GetSamples](#getsamples) - Get samples data for the specified output. Used to get sample data for the test action.
* [GetStatus](#getstatus) - Get OutputStatus by ID
* [ListOutputObjects](#listoutputobjects) - Get a list of Output objects
* [ListOutputStatus](#listoutputstatus) - Get a list of OutputStatus objects
* [Post](#post) - Send sample data to an output to validate configuration or test connectivity
* [Update](#update) - Update Output

## Create

Create Output

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
    res, err := s.Outputs.Create(ctx, shared.Output{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Outputs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Output](../../models/shared/output.md)        | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateOutputObjectResponse](../../models/operations/createoutputobjectresponse.md), error**


## Delete

Delete Output

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
    id := "enim"

    ctx := context.Background()
    res, err := s.Outputs.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Outputs != nil {
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

**[*operations.DeleteOutputIDResponse](../../models/operations/deleteoutputidresponse.md), error**


## DeletePQ

Delete destination persistent queue

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
    id := "rem"

    ctx := context.Background()
    res, err := s.Outputs.DeletePQ(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteDestinationQueue200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Output Id                                             |


### Response

**[*operations.DeleteDestinationQueueResponse](../../models/operations/deletedestinationqueueresponse.md), error**


## Get

Get Output by ID

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
    id := "magni"

    ctx := context.Background()
    res, err := s.Outputs.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Outputs != nil {
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

**[*operations.GetOutputIDResponse](../../models/operations/getoutputidresponse.md), error**


## GetLatestPQ

Get status of latest clear PQ job for an output

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
    id := "quae"

    ctx := context.Background()
    res, err := s.Outputs.GetLatestPQ(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetLatestPQ200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Output Id                                             |


### Response

**[*operations.GetLatestPQResponse](../../models/operations/getlatestpqresponse.md), error**


## GetSamples

Get samples data for the specified output. Used to get sample data for the test action.

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
    id := "quas"

    ctx := context.Background()
    res, err := s.Outputs.GetSamples(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSpecifiedOutput200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Output Id                                             |


### Response

**[*operations.GetSpecifiedOutputResponse](../../models/operations/getspecifiedoutputresponse.md), error**


## GetStatus

Get OutputStatus by ID

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
    id := "placeat"

    ctx := context.Background()
    res, err := s.Outputs.GetStatus(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.OutputStatuses != nil {
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

**[*operations.GetOutputStatusIDResponse](../../models/operations/getoutputstatusidresponse.md), error**


## ListOutputObjects

Get a list of Output objects

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
    res, err := s.Outputs.ListOutputObjects(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Outputs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListOutputObjectsResponse](../../models/operations/listoutputobjectsresponse.md), error**


## ListOutputStatus

Get a list of OutputStatus objects

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
    res, err := s.Outputs.ListOutputStatus(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.OutputStatuses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListOutputStatusResponse](../../models/operations/listoutputstatusresponse.md), error**


## Post

Send sample data to an output to validate configuration or test connectivity

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
    id := "enim"
    outputTestRequest := &shared.OutputTestRequest{
        Events: []shared.CriblEvent{
            shared.CriblEvent{
                Raw: "sapiente",
            },
            shared.CriblEvent{
                Raw: "saepe",
            },
        },
    }

    ctx := context.Background()
    res, err := s.Outputs.Post(ctx, id, outputTestRequest)
    if err != nil {
        log.Fatal(err)
    }

    if res.OutputTestResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |
| `id`                                                                  | *string*                                                              | :heavy_check_mark:                                                    | Output Id                                                             |
| `outputTestRequest`                                                   | [*shared.OutputTestRequest](../../models/shared/outputtestrequest.md) | :heavy_minus_sign:                                                    | OutputTestRequest object                                              |


### Response

**[*operations.PostSampleOutputResponse](../../models/operations/postsampleoutputresponse.md), error**


## Update

Update Output

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
    id := "delectus"
    output := shared.Output{}

    ctx := context.Background()
    res, err := s.Outputs.Update(ctx, id, output)
    if err != nil {
        log.Fatal(err)
    }

    if res.Outputs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |
| `output`                                              | [*shared.Output](../../models/shared/output.md)       | :heavy_minus_sign:                                    | Output object to be updated                           |


### Response

**[*operations.UpdateOutputIDResponse](../../models/operations/updateoutputidresponse.md), error**

