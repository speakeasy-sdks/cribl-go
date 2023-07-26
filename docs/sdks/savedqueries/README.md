# SavedQueries

### Available Operations

* [Create](#create) - Create SavedQuery
* [Delete](#delete) - Delete SavedQuery
* [Get](#get) - Get a list of SavedQuery objects
* [Update](#update) - Update SavedQuery

## Create

Create SavedQuery

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
    res, err := s.SavedQueries.Create(ctx, shared.SavedQuery{
        Description: cribl.String("corporis"),
        Earliest: cribl.String("at"),
        ID: "03084fbb-a5cc-4eff-9cb0-1fe51e528a45",
        Latest: cribl.String("mollitia"),
        Name: "Morris Champlin",
        Query: "ullam",
        SampleRate: cribl.Int64(982059),
        Schedule: &shared.SavedQuerySchedule{
            CronSchedule: "corrupti",
            Enabled: false,
            KeepLastN: 725625,
            Tz: "placeat",
        },
        User: cribl.String("explicabo"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SavedQuery != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                              | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ctx`                                                  | [context.Context](https://pkg.go.dev/context#Context)  | :heavy_check_mark:                                     | The context to use for the request.                    |
| `request`                                              | [shared.SavedQuery](../../models/shared/savedquery.md) | :heavy_check_mark:                                     | The request object to use for the request.             |


### Response

**[*operations.CreateSavedQueriesResponse](../../models/operations/createsavedqueriesresponse.md), error**


## Delete

Delete SavedQuery

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
    res, err := s.SavedQueries.Delete(ctx, operations.DeleteSavedQueriesRequest{
        ID: "caba8da4-127d-4d59-bff4-711aa1bc74b8",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SavedQuery != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteSavedQueriesRequest](../../models/operations/deletesavedqueriesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.DeleteSavedQueriesResponse](../../models/operations/deletesavedqueriesresponse.md), error**


## Get

Get a list of SavedQuery objects

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
    res, err := s.SavedQueries.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.SavedQueries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetSavedQueriesResponse](../../models/operations/getsavedqueriesresponse.md), error**


## Update

Update SavedQuery

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
    res, err := s.SavedQueries.Update(ctx, operations.UpdateSavedQueriesRequest{
        SavedQuery: &shared.SavedQuery{
            Description: cribl.String("aliquid"),
            Earliest: cribl.String("placeat"),
            ID: "ecc74f77-b484-48bd-aa6f-0441d2c3b808",
            Latest: cribl.String("sit"),
            Name: "Alex Douglas",
            Query: "eveniet",
            SampleRate: cribl.Int64(56956),
            Schedule: &shared.SavedQuerySchedule{
                CronSchedule: "laboriosam",
                Enabled: false,
                KeepLastN: 25321,
                Tz: "labore",
            },
            User: cribl.String("ullam"),
        },
        ID: "9bebbad0-2f25-486b-8f15-2558daa95be6",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SavedQuery != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateSavedQueriesRequest](../../models/operations/updatesavedqueriesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateSavedQueriesResponse](../../models/operations/updatesavedqueriesresponse.md), error**

