# Jobs

### Available Operations

* [Cancel](#cancel) - Cancel a job by instance id
* [Delete](#delete) - Remove job from job inspector by instance id
* [Get](#get) - Get job info by instance id
* [GetError](#geterror) - Get Task errors for a job by id
* [GetResult](#getresult) - Get results for a discover job by instance id
* [ListJobInfos](#listjobinfos) - Get info on jobs
* [ListJobResults](#listjobresults) - Get results for a discover job by instance id
* [ListTaskErrors](#listtaskerrors) - Get Task errors for a job by id
* [PauseJob](#pausejob) - Pause a job by instance id
* [Prevent](#prevent) - prevent job from being deleted automatically
* [Resume](#resume) - Resume a job by instance id
* [RunJob](#runjob) - Run or schedule a job

## Cancel

Cancel a job by instance id

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
    id := "dolores"

    ctx := context.Background()
    res, err := s.Jobs.Cancel(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.JobCancel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Job instance id                                       |


### Response

**[*operations.CancelJobResponse](../../models/operations/canceljobresponse.md), error**


## Delete

Remove job from job inspector by instance id

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
    res, err := s.Jobs.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.JobDelete != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Job instance id                                       |


### Response

**[*operations.DeleteJobResponse](../../models/operations/deletejobresponse.md), error**


## Get

Get job info by instance id

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
    id := "laboriosam"

    ctx := context.Background()
    res, err := s.Jobs.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.JobInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Job instance id                                       |


### Response

**[*operations.GetJobResponse](../../models/operations/getjobresponse.md), error**


## GetError

Get Task errors for a job by id

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
    id := "velit"

    ctx := context.Background()
    res, err := s.Jobs.GetError(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskErrors != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Instance id of the job to get results for             |


### Response

**[*operations.GetTaskErrorResponse](../../models/operations/gettaskerrorresponse.md), error**


## GetResult

Get results for a discover job by instance id

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
    group := "a"
    id := "molestias"
    maxFiles := 300029

    ctx := context.Background()
    res, err := s.Jobs.GetResult(ctx, group, id, maxFiles)
    if err != nil {
        log.Fatal(err)
    }

    if res.JobResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `group`                                               | *string*                                              | :heavy_check_mark:                                    | Group the job belongs to                              |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Instance id of the job to get results for             |
| `maxFiles`                                            | **int64*                                              | :heavy_minus_sign:                                    | Maximum files to get job results                      |


### Response

**[*operations.GetJobResultResponse](../../models/operations/getjobresultresponse.md), error**


## ListJobInfos

Get info on jobs

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
    res, err := s.Jobs.ListJobInfos(ctx, operations.ListJobInfosRequest{
        CollectorID: cribl.String("saepe"),
        ID: cribl.String("29e973e9-22a5-47a1-9be3-e060807e2b6e"),
        Limit: cribl.Int64(187552),
        Offset: cribl.Int64(672582),
        RunType: cribl.String("distinctio"),
        State: cribl.String("voluptatum"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.JobInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListJobInfosRequest](../../models/operations/listjobinfosrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListJobInfosResponse](../../models/operations/listjobinfosresponse.md), error**


## ListJobResults

Get results for a discover job by instance id

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
    res, err := s.Jobs.ListJobResults(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListJobResults200ApplicationXNdjsonBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Job instance id                                       |


### Response

**[*operations.ListJobResultsResponse](../../models/operations/listjobresultsresponse.md), error**


## ListTaskErrors

Get Task errors for a job by id

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
    group := "aliquam"
    id := "ad"

    ctx := context.Background()
    res, err := s.Jobs.ListTaskErrors(ctx, group, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskErrors != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `group`                                               | *string*                                              | :heavy_check_mark:                                    | Group the job belongs to                              |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Instance id of the job to get results for             |


### Response

**[*operations.ListTaskErrorsResponse](../../models/operations/listtaskerrorsresponse.md), error**


## PauseJob

Pause a job by instance id

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
    id := "repellat"

    ctx := context.Background()
    res, err := s.Jobs.PauseJob(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.JobPause != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Job instance id                                       |


### Response

**[*operations.PauseJobResponse](../../models/operations/pausejobresponse.md), error**


## Prevent

prevent job from being deleted automatically

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
    id := "alias"

    ctx := context.Background()
    res, err := s.Jobs.Prevent(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.JobInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Job Instance id                                       |


### Response

**[*operations.PreventJobResponse](../../models/operations/preventjobresponse.md), error**


## Resume

Resume a job by instance id

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
    id := "corporis"

    ctx := context.Background()
    res, err := s.Jobs.Resume(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.JobResume != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Job instance id                                       |


### Response

**[*operations.ResumeJobResponse](../../models/operations/resumejobresponse.md), error**


## RunJob

Run or schedule a job

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
    res, err := s.Jobs.RunJob(ctx, shared.SavedJob{})
    if err != nil {
        log.Fatal(err)
    }

    if res.JobRun != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.SavedJob](../../models/shared/savedjob.md)    | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.RunJobResponse](../../models/operations/runjobresponse.md), error**

