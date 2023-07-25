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
    res, err := s.Profiler.Create(ctx, shared.ProfilerItem{
        CreateTime: cribl.Int64(234291),
        ID: "7bb8e0cc-8851-487e-8de0-4af28c5dddb4",
        Size: cribl.Int64(413962),
        WorkerID: cribl.String("mollitia"),
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
    res, err := s.Profiler.Delete(ctx, operations.DeleteProfilerRequest{
        ID: "a1cfd6d8-28da-4013-9911-29646645c1d8",
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
    res, err := s.Profiler.Get(ctx, operations.GetProfilerRequest{
        ID: "1f29042f-569b-47af-b0ea-2216cbe071bc",
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
    res, err := s.Profiler.Update(ctx, operations.UpdateProfilerRequest{
        ProfilerItem: &shared.ProfilerItem{
            CreateTime: cribl.Int64(69788),
            ID: "63e279a3-b084-4da9-9257-d04f40847a74",
            Size: cribl.Int64(155785),
            WorkerID: cribl.String("repellendus"),
        },
        ID: "84496cbd-eecf-46b9-9bc6-3562ebfdf55c",
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

