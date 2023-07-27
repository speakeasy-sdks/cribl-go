# SavedJobs

### Available Operations

* [Create](#create) - Create SavedJob
* [Get](#get) - Get a list of SavedJob objects

## Create

Create SavedJob

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
    res, err := s.SavedJobs.Create(ctx, shared.SavedJob{})
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
| `request`                                             | [shared.SavedJob](../../models/shared/savedjob.md)    | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateSavedJobsResponse](../../models/operations/createsavedjobsresponse.md), error**


## Get

Get a list of SavedJob objects

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
    res, err := s.SavedJobs.Get(ctx)
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


### Response

**[*operations.GetSavedJobsResponse](../../models/operations/getsavedjobsresponse.md), error**

