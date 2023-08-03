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
        CreateTime: cribl.Int64(34989),
        ID: "688f77c1-ffc7-41dc-a163-f2a3c80a97ff",
        Size: cribl.Int64(242013),
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
    id := "non"

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
    id := "optio"

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
    id := "illum"
    profilerItem := &shared.ProfilerItem{
        CreateTime: cribl.Int64(870183),
        ID: "f857a9e6-1876-4c6a-b21d-29dfc94d6fec",
        Size: cribl.Int64(818078),
        WorkerID: cribl.String("dignissimos"),
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

