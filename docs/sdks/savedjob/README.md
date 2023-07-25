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
        ID: "354aa432-b47e-4176-bc52-08c23e9802d8",
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
        ID: "2f0d45eb-4a8b-4674-ae5c-fc18edc7f787",
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
        ID: "e32e04b3-d3ed-40c5-a70e-f42bd3c9f1cc",
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

