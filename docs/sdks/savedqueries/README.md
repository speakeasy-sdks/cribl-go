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
        Description: cribl.String("veniam"),
        Earliest: cribl.String("aperiam"),
        ID: "3f6c39bc-d0a6-4290-b957-f385189ad7ef",
        Latest: cribl.String("voluptatum"),
        Name: "Melinda Okuneva",
        Query: "quae",
        SampleRate: cribl.Int64(236564),
        Schedule: &shared.SavedQuerySchedule{
            CronSchedule: "hic",
            Enabled: false,
            KeepLastN: 213707,
            Tz: "adipisci",
        },
        User: cribl.String("optio"),
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
        ID: "a79fb9de-4032-4ba2-afd3-68ba9216bcb4",
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
            Description: cribl.String("dicta"),
            Earliest: cribl.String("nostrum"),
            ID: "835c7364-1723-4133-adc0-46bc5163bbca",
            Latest: cribl.String("eius"),
            Name: "Phillip D'Amore",
            Query: "non",
            SampleRate: cribl.Int64(141414),
            Schedule: &shared.SavedQuerySchedule{
                CronSchedule: "maxime",
                Enabled: false,
                KeepLastN: 134798,
                Tz: "magni",
            },
            User: cribl.String("minus"),
        },
        ID: "55350495-c5db-4b3c-97c1-e4981e8aa257",
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

