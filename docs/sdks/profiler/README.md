# Profiler

### Available Operations

* [Create](#create) - Create ProfilerItem
* [Delete](#delete) - Delete ProfilerItem
* [Get](#get) - Get ProfilerItem by ID
* [Update](#update) - Update ProfilerItem

## Create

Create ProfilerItem

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
    res, err := s.Profiler.Create(ctx, shared.ProfilerItem{
        CreateTime: cribl.Int64(447678),
        ID: "1cffbd0e-b74b-4842-9953-b44bd3c43159",
        Size: cribl.Int64(825303),
        WorkerID: cribl.String("adipisci"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ProfilerItem != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [shared.ProfilerItem](../../models/shared/profileritem.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |


### Response

**[*operations.CreateProfilerResponse](../../models/operations/createprofilerresponse.md), error**


## Delete

Delete ProfilerItem

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
    res, err := s.Profiler.Delete(ctx, operations.DeleteProfilerRequest{
        ID: "3e5953c0-0113-4986-baa4-1e6c31cc2f1f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ProfilerItem != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteProfilerRequest](../../models/operations/deleteprofilerrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DeleteProfilerResponse](../../models/operations/deleteprofilerresponse.md), error**


## Get

Get ProfilerItem by ID

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
    res, err := s.Profiler.Get(ctx, operations.GetProfilerRequest{
        ID: "cb51c9a4-1ffb-4e9c-bd79-5ee65e076cc7",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ProfilerItem != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetProfilerRequest](../../models/operations/getprofilerrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetProfilerResponse](../../models/operations/getprofilerresponse.md), error**


## Update

Update ProfilerItem

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
    res, err := s.Profiler.Update(ctx, operations.UpdateProfilerRequest{
        ProfilerItem: &shared.ProfilerItem{
            CreateTime: cribl.Int64(648469),
            ID: "bf616ea5-c716-4419-b4b9-0f2e09d19d2f",
            Size: cribl.Int64(768772),
            WorkerID: cribl.String("explicabo"),
        },
        ID: "f9e2e105-944b-4935-9237-a72f90849d6a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ProfilerItem != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateProfilerRequest](../../models/operations/updateprofilerrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateProfilerResponse](../../models/operations/updateprofilerresponse.md), error**

