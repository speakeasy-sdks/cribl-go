# EventBreakerRules

## Overview

Actions related to event breaker rules

### Available Operations

* [Delete](#delete) - Delete Event Breaker Ruleset
* [Get](#get) - Get Event Breaker Ruleset by ID
* [ListEventBreaker](#listeventbreaker) - Get a list of Event Breaker Ruleset objects
* [Post](#post) - Create Event Breaker Ruleset
* [Update](#update) - Update Event Breaker Ruleset

## Delete

Delete Event Breaker Ruleset

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
    id := "eius"

    ctx := context.Background()
    res, err := s.EventBreakerRules.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.EventBreakerRulesets != nil {
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

**[*operations.DeleteEventBreakerResponse](../../models/operations/deleteeventbreakerresponse.md), error**


## Get

Get Event Breaker Ruleset by ID

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
    id := "eos"

    ctx := context.Background()
    res, err := s.EventBreakerRules.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.EventBreakerRulesets != nil {
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

**[*operations.GetEventBreakerIDResponse](../../models/operations/geteventbreakeridresponse.md), error**


## ListEventBreaker

Get a list of Event Breaker Ruleset objects

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
    res, err := s.EventBreakerRules.ListEventBreaker(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.EventBreakerRulesets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListEventBreakerResponse](../../models/operations/listeventbreakerresponse.md), error**


## Post

Create Event Breaker Ruleset

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
    res, err := s.EventBreakerRules.Post(ctx, shared.EventBreakerRuleset{
        Description: cribl.String("voluptas"),
        ID: "1904e523-c7e0-4bc7-978e-4796f2a70c68",
        Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
        MinRawLength: cribl.Int64(510017),
        Rules: []shared.EventBreakerRulesetRules{
            shared.EventBreakerRulesetRules{
                Condition: "consequuntur",
                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                    DstField: cribl.String("deleniti"),
                    FieldFilterExpr: cribl.String("fugit"),
                    Fields: []string{
                        "fuga",
                    },
                    Keep: []string{
                        "mollitia",
                    },
                    Remove: []string{
                        "incidunt",
                    },
                },
                Disabled: cribl.Bool(false),
                Fields: []shared.EventBreakerRulesetRulesFields{
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Roy Hansen"),
                        Value: "sapiente",
                    },
                },
                MaxEventBytes: cribl.Int64(159870),
                Name: "Rose Turner",
                ParserEnabled: cribl.Bool(false),
                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                    Format: cribl.String("et"),
                    Length: cribl.Int64(456911),
                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                },
                TimestampAnchorRegex: "accusamus",
                TimestampEarliest: cribl.String("veritatis"),
                TimestampLatest: cribl.String("esse"),
                TimestampTimezone: cribl.String("quod"),
                Type: shared.EventBreakerRulesetRulesEventBreakerTypeTimestamp,
            },
        },
        Tags: cribl.String("vero"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EventBreakerRulesets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.EventBreakerRuleset](../../models/shared/eventbreakerruleset.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.PostEventBreakerResponse](../../models/operations/posteventbreakerresponse.md), error**


## Update

Update Event Breaker Ruleset

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
    id := "aliquid"
    eventBreakerRuleset := &shared.EventBreakerRuleset{
        Description: cribl.String("quasi"),
        ID: "e6b7b95b-c0ab-43c2-8c4f-3789fd871f99",
        Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
        MinRawLength: cribl.Int64(863023),
        Rules: []shared.EventBreakerRulesetRules{
            shared.EventBreakerRulesetRules{
                Condition: "possimus",
                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                    DstField: cribl.String("quia"),
                    FieldFilterExpr: cribl.String("eveniet"),
                    Fields: []string{
                        "asperiores",
                    },
                    Keep: []string{
                        "facere",
                    },
                    Remove: []string{
                        "veritatis",
                    },
                },
                Disabled: cribl.Bool(false),
                Fields: []shared.EventBreakerRulesetRulesFields{
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Ann Murphy"),
                        Value: "tenetur",
                    },
                },
                MaxEventBytes: cribl.Int64(62713),
                Name: "Chester King",
                ParserEnabled: cribl.Bool(false),
                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                    Format: cribl.String("illum"),
                    Length: cribl.Int64(742238),
                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                },
                TimestampAnchorRegex: "aliquam",
                TimestampEarliest: cribl.String("sapiente"),
                TimestampLatest: cribl.String("dicta"),
                TimestampTimezone: cribl.String("ullam"),
                Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
            },
        },
        Tags: cribl.String("ullam"),
    }

    ctx := context.Background()
    res, err := s.EventBreakerRules.Update(ctx, id, eventBreakerRuleset)
    if err != nil {
        log.Fatal(err)
    }

    if res.EventBreakerRulesets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `id`                                                                      | *string*                                                                  | :heavy_check_mark:                                                        | Unique ID                                                                 |
| `eventBreakerRuleset`                                                     | [*shared.EventBreakerRuleset](../../models/shared/eventbreakerruleset.md) | :heavy_minus_sign:                                                        | Event Breaker Ruleset object to be updated                                |


### Response

**[*operations.UpdateEventBreakerResponse](../../models/operations/updateeventbreakerresponse.md), error**

