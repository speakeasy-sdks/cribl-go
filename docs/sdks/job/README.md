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
	"cribl"
	"cribl/pkg/models/shared"
	"cribl/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Job.Cancel(ctx, operations.CancelJobRequest{
        ID: "c6891f82-ce11-4571-b230-5377dcfa89df",
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
	"cribl"
	"cribl/pkg/models/shared"
	"cribl/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Job.Delete(ctx, operations.DeleteJobRequest{
        ID: "975e3566-8609-42e9-83dd-c5f111dea102",
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
	"cribl"
	"cribl/pkg/models/shared"
	"cribl/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Job.Get(ctx, operations.GetJobRequest{
        ID: "6d541a4d-190f-4eb2-9780-bccc0dbbddb4",
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
	"cribl"
	"cribl/pkg/models/shared"
	"cribl/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Job.PauseJob(ctx, operations.PauseJobRequest{
        ID: "84708fb4-e391-4e6b-8158-c4c4e54599ea",
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
	"cribl"
	"cribl/pkg/models/shared"
	"cribl/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Job.Prevent(ctx, operations.PreventJobRequest{
        ID: "342260e9-b200-4ce7-8a1b-d8fb7a0a116c",
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
	"cribl"
	"cribl/pkg/models/shared"
	"cribl/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Job.Resume(ctx, operations.ResumeJobRequest{
        ID: "e723d409-7fa3-40e9-af72-5b29122030d8",
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

