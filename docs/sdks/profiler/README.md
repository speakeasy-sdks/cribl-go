# Profiler

### Available Operations

* [Create](#create) - Create ProfilerItem
* [Delete](#delete) - Delete ProfilerItem
* [Get](#get) - Get ProfilerItem by ID
* [ListProfilers](#listprofilers) - Get a list of ProfilerItem objects
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
        CreateTime: cribl.Int64(334018),
        ID: "2ff785fc-3781-44d4-898e-0c2bb89eb75d",
        Size: cribl.Int64(629377),
        WorkerID: cribl.String("repellendus"),
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
    id := "iure"

    ctx := context.Background()
    res, err := s.Profiler.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.ProfilerItem != nil {
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
    id := "dolorem"

    ctx := context.Background()
    res, err := s.Profiler.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.ProfilerItem != nil {
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

**[*operations.GetProfilerResponse](../../models/operations/getprofilerresponse.md), error**


## ListProfilers

Get a list of ProfilerItem objects

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
    res, err := s.Profiler.ListProfilers(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ProfilerItems != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListProfilersResponse](../../models/operations/listprofilersresponse.md), error**


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
    id := "commodi"
    profilerItem := &shared.ProfilerItem{
        CreateTime: cribl.Int64(771226),
        ID: "600503d8-bb31-4180-b739-ae9e057eb809",
        Size: cribl.Int64(879522),
        WorkerID: cribl.String("eos"),
    }

    ctx := context.Background()
    res, err := s.Profiler.Update(ctx, id, profilerItem)
    if err != nil {
        log.Fatal(err)
    }

    if res.ProfilerItem != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |
| `id`                                                        | *string*                                                    | :heavy_check_mark:                                          | Unique ID                                                   |
| `profilerItem`                                              | [*shared.ProfilerItem](../../models/shared/profileritem.md) | :heavy_minus_sign:                                          | ProfilerItem object to be updated                           |


### Response

**[*operations.UpdateProfilerResponse](../../models/operations/updateprofilerresponse.md), error**

