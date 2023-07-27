# Job

### Available Operations

* [Cancel](#cancel) - Cancel a job by instance id
* [Delete](#delete) - Remove job from job inspector by instance id
* [Get](#get) - Get job info by instance id
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

    ctx := context.Background()
    res, err := s.Job.Cancel(ctx, operations.CancelJobRequest{
        ID: "d364ffd4-5590-46d1-a63d-48e935c2c9e8",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.JobCancel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.CancelJobRequest](../../models/operations/canceljobrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


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

    ctx := context.Background()
    res, err := s.Job.Delete(ctx, operations.DeleteJobRequest{
        ID: "1f30be3e-4320-42d7-a165-76506641870d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.JobDelete != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.DeleteJobRequest](../../models/operations/deletejobrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


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

    ctx := context.Background()
    res, err := s.Job.Get(ctx, operations.GetJobRequest{
        ID: "9d21f9ad-030c-44ec-811a-0836429068b8",
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [operations.GetJobRequest](../../models/operations/getjobrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.GetJobResponse](../../models/operations/getjobresponse.md), error**


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

    ctx := context.Background()
    res, err := s.Job.PauseJob(ctx, operations.PauseJobRequest{
        ID: "502a55e7-f73b-4c84-9e32-0a319f4badf9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.JobPause != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.PauseJobRequest](../../models/operations/pausejobrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


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

    ctx := context.Background()
    res, err := s.Job.Prevent(ctx, operations.PreventJobRequest{
        ID: "47c9a867-bc42-4426-a658-16ddca8ef51f",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.PreventJobRequest](../../models/operations/preventjobrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


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

    ctx := context.Background()
    res, err := s.Job.Resume(ctx, operations.ResumeJobRequest{
        ID: "cb4c593e-c12c-4daa-90ec-7afedbd80df4",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.JobResume != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ResumeJobRequest](../../models/operations/resumejobrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


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
    res, err := s.Job.RunJob(ctx, shared.SavedJob{})
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

