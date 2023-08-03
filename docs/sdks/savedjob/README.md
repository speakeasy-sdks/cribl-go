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
    id := "amet"

    ctx := context.Background()
    res, err := s.SavedJob.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.SavedJobs != nil {
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

**[*operations.DeleteSavedJobResponse](../../models/operations/deletesavedjobresponse.md), error**


## Get

Get SavedJob by ID

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
    id := "veritatis"

    ctx := context.Background()
    res, err := s.SavedJob.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.SavedJobs != nil {
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

**[*operations.GetSavedJobResponse](../../models/operations/getsavedjobresponse.md), error**


## Update

Update SavedJob

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
    id := "error"
    savedJob := shared.SavedJob{}

    ctx := context.Background()
    res, err := s.SavedJob.Update(ctx, id, savedJob)
    if err != nil {
        log.Fatal(err)
    }

    if res.SavedJobs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |
| `savedJob`                                            | [*shared.SavedJob](../../models/shared/savedjob.md)   | :heavy_minus_sign:                                    | SavedJob object to be updated                         |


### Response

**[*operations.UpdateSavedJobResponse](../../models/operations/updatesavedjobresponse.md), error**

