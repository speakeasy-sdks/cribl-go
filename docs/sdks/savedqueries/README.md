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
        Description: cribl.String("voluptatibus"),
        Earliest: cribl.String("numquam"),
        ID: "badf947c-9a86-47bc-8242-6665816ddca8",
        Latest: cribl.String("voluptates"),
        Name: "Maurice Bins",
        Query: "expedita",
        SampleRate: cribl.Int64(266370),
        Schedule: &shared.SavedQuerySchedule{
            CronSchedule: "cumque",
            Enabled: false,
            KeepLastN: 368599,
            Tz: "occaecati",
        },
        User: cribl.String("ipsum"),
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
    id := "accusamus"

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
    id := "quisquam"
    savedQuery := &shared.SavedQuery{
        Description: cribl.String("quasi"),
        Earliest: cribl.String("fugit"),
        ID: "cdaad0ec-7afe-4dbd-80df-448a47f9390c",
        Latest: cribl.String("minima"),
        Name: "Ms. Guy Macejkovic",
        Query: "ipsum",
        SampleRate: cribl.Int64(874283),
        Schedule: &shared.SavedQuerySchedule{
            CronSchedule: "fuga",
            Enabled: false,
            KeepLastN: 704732,
            Tz: "maiores",
        },
        User: cribl.String("error"),
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

