# SavedQueries

### Available Operations

* [Create](#create) - Create SavedQuery
* [Delete](#delete) - Delete SavedQuery
* [Get](#get) - Get SavedQuery by ID
* [ListSavedQueries](#listsavedqueries) - Get a list of SavedQuery objects
* [Update](#update) - Update SavedQuery

## Create

Create SavedQuery

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
    res, err := s.SavedQueries.Create(ctx, shared.SavedQuery{
        Description: cribl.String("voluptas"),
        Earliest: cribl.String("consequuntur"),
        ID: "a7b408f0-5e3d-448f-9af3-13a1f5fd9425",
        Latest: cribl.String("cupiditate"),
        Name: "John Predovic",
        Query: "sapiente",
        SampleRate: cribl.Int64(164319),
        Schedule: &shared.SavedQuerySchedule{
            CronSchedule: "veniam",
            Enabled: false,
            KeepLastN: 893773,
            Tz: "officia",
        },
        User: cribl.String("sint"),
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
    id := "ut"

    ctx := context.Background()
    res, err := s.SavedQueries.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.SavedQuery != nil {
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

**[*operations.DeleteSavedQueriesResponse](../../models/operations/deletesavedqueriesresponse.md), error**


## Get

Get SavedQuery by ID

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
    id := "numquam"

    ctx := context.Background()
    res, err := s.SavedQueries.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.SavedQuery != nil {
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

**[*operations.GetSavedQueryResponse](../../models/operations/getsavedqueryresponse.md), error**


## ListSavedQueries

Get a list of SavedQuery objects

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
    res, err := s.SavedQueries.ListSavedQueries(ctx)
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

**[*operations.ListSavedQueriesResponse](../../models/operations/listsavedqueriesresponse.md), error**


## Update

Update SavedQuery

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
    id := "tenetur"
    savedQuery := &shared.SavedQuery{
        Description: cribl.String("adipisci"),
        Earliest: cribl.String("libero"),
        ID: "756c11f6-c37a-4512-a243-835bbc05a23a",
        Latest: cribl.String("modi"),
        Name: "Traci Waters",
        Query: "enim",
        SampleRate: cribl.Int64(987759),
        Schedule: &shared.SavedQuerySchedule{
            CronSchedule: "assumenda",
            Enabled: false,
            KeepLastN: 887363,
            Tz: "architecto",
        },
        User: cribl.String("alias"),
    }

    ctx := context.Background()
    res, err := s.SavedQueries.Update(ctx, id, savedQuery)
    if err != nil {
        log.Fatal(err)
    }

    if res.SavedQuery != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                               | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ctx`                                                   | [context.Context](https://pkg.go.dev/context#Context)   | :heavy_check_mark:                                      | The context to use for the request.                     |
| `id`                                                    | *string*                                                | :heavy_check_mark:                                      | Unique ID                                               |
| `savedQuery`                                            | [*shared.SavedQuery](../../models/shared/savedquery.md) | :heavy_minus_sign:                                      | SavedQuery object to be updated                         |


### Response

**[*operations.UpdateSavedQueriesResponse](../../models/operations/updatesavedqueriesresponse.md), error**

