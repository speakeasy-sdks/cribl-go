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
        Description: cribl.String("error"),
        Earliest: cribl.String("recusandae"),
        ID: "f3ffdd9f-7f07-49af-8d35-724cdb0f4d28",
        Latest: cribl.String("quasi"),
        Name: "Irma Kub",
        Query: "iure",
        SampleRate: cribl.Int64(512408),
        Schedule: &shared.SavedQuerySchedule{
            CronSchedule: "modi",
            Enabled: false,
            KeepLastN: 301701,
            Tz: "accusamus",
        },
        User: cribl.String("nulla"),
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
    id := "repudiandae"

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
    id := "quibusdam"

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
    id := "praesentium"
    savedQuery := &shared.SavedQuery{
        Description: cribl.String("enim"),
        Earliest: cribl.String("animi"),
        ID: "9065e628-bdfc-4203-ab6c-879923b7e135",
        Latest: cribl.String("atque"),
        Name: "Nettie Kilback",
        Query: "et",
        SampleRate: cribl.Int64(181267),
        Schedule: &shared.SavedQuerySchedule{
            CronSchedule: "impedit",
            Enabled: false,
            KeepLastN: 401388,
            Tz: "praesentium",
        },
        User: cribl.String("natus"),
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

