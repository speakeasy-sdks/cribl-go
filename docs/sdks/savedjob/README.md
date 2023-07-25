# SavedJob

### Available Operations

* [Delete](#delete) - Delete SavedJob
* [Get](#get) - Get SavedJob by ID
* [Update](#update) - Update SavedJob

## Delete

Delete SavedJob

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
    res, err := s.SavedJob.Delete(ctx, operations.DeleteSavedJobRequest{
        ID: "70fb3874-290d-4336-961e-ca16ef89451b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SavedJobs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteSavedJobRequest](../../models/operations/deletesavedjobrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DeleteSavedJobResponse](../../models/operations/deletesavedjobresponse.md), error**


## Get

Get SavedJob by ID

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
    res, err := s.SavedJob.Get(ctx, operations.GetSavedJobRequest{
        ID: "d76eeeb5-18c4-4da1-bad3-5512f06d4e5b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SavedJobs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetSavedJobRequest](../../models/operations/getsavedjobrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetSavedJobResponse](../../models/operations/getsavedjobresponse.md), error**


## Update

Update SavedJob

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
    res, err := s.SavedJob.Update(ctx, operations.UpdateSavedJobRequest{
        SavedJob: &shared.SavedJob{},
        ID: "72f0f548-568a-4042-8e00-a1d6eb943464",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SavedJobs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateSavedJobRequest](../../models/operations/updatesavedjobrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateSavedJobResponse](../../models/operations/updatesavedjobresponse.md), error**

