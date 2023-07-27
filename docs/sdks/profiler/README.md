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
        CreateTime: cribl.Int64(496578),
        ID: "00878756-143f-45a6-898b-55554080d40b",
        Size: cribl.Int64(751298),
        WorkerID: cribl.String("similique"),
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
    id := "porro"

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
    id := "impedit"

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
    id := "nisi"
    profilerItem := &shared.ProfilerItem{
        CreateTime: cribl.Int64(768920),
        ID: "bd6b5f3e-c909-4304-b926-bad2553819b4",
        Size: cribl.Int64(452653),
        WorkerID: cribl.String("eius"),
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

