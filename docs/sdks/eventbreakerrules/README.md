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
    id := "ratione"

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
    id := "ullam"

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
        Description: cribl.String("perferendis"),
        ID: "d8cdb5a3-4181-4430-9042-1813d5208ece",
        Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
        MinRawLength: cribl.Int64(456410),
        Rules: []shared.EventBreakerRulesetRules{
            shared.EventBreakerRulesetRules{
                Condition: "sed",
                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                    DstField: cribl.String("veniam"),
                    FieldFilterExpr: cribl.String("nesciunt"),
                    Fields: []string{
                        "eum",
                        "vel",
                        "voluptatum",
                    },
                    Keep: []string{
                        "exercitationem",
                        "ab",
                    },
                    Remove: []string{
                        "autem",
                        "nobis",
                        "laboriosam",
                        "recusandae",
                    },
                },
                Disabled: cribl.Bool(false),
                Fields: []shared.EventBreakerRulesetRulesFields{
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Mrs. June Tremblay"),
                        Value: "vero",
                    },
                },
                MaxEventBytes: cribl.Int64(667593),
                Name: "Johnny Yundt",
                ParserEnabled: cribl.Bool(false),
                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                    Format: cribl.String("occaecati"),
                    Length: cribl.Int64(364544),
                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                },
                TimestampAnchorRegex: "blanditiis",
                TimestampEarliest: cribl.String("officia"),
                TimestampLatest: cribl.String("voluptas"),
                TimestampTimezone: cribl.String("numquam"),
                Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
            },
            shared.EventBreakerRulesetRules{
                Condition: "quos",
                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                    DstField: cribl.String("eius"),
                    FieldFilterExpr: cribl.String("aspernatur"),
                    Fields: []string{
                        "nesciunt",
                        "fuga",
                    },
                    Keep: []string{
                        "incidunt",
                        "quasi",
                        "rem",
                    },
                    Remove: []string{
                        "dicta",
                        "nisi",
                        "consequuntur",
                        "consectetur",
                    },
                },
                Disabled: cribl.Bool(false),
                Fields: []shared.EventBreakerRulesetRulesFields{
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Miss Dominick Rogahn"),
                        Value: "occaecati",
                    },
                },
                MaxEventBytes: cribl.Int64(612867),
                Name: "Stephanie Pfannerstill",
                ParserEnabled: cribl.Bool(false),
                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                    Format: cribl.String("distinctio"),
                    Length: cribl.Int64(608593),
                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                },
                TimestampAnchorRegex: "minima",
                TimestampEarliest: cribl.String("praesentium"),
                TimestampLatest: cribl.String("maxime"),
                TimestampTimezone: cribl.String("magnam"),
                Type: shared.EventBreakerRulesetRulesEventBreakerTypeCsv,
            },
            shared.EventBreakerRulesetRules{
                Condition: "quos",
                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                    DstField: cribl.String("commodi"),
                    FieldFilterExpr: cribl.String("itaque"),
                    Fields: []string{
                        "totam",
                        "earum",
                    },
                    Keep: []string{
                        "nam",
                        "vero",
                    },
                    Remove: []string{
                        "ipsam",
                    },
                },
                Disabled: cribl.Bool(false),
                Fields: []shared.EventBreakerRulesetRulesFields{
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Frances Franey"),
                        Value: "sint",
                    },
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Gerard Koch"),
                        Value: "est",
                    },
                },
                MaxEventBytes: cribl.Int64(336102),
                Name: "Tomas Ryan",
                ParserEnabled: cribl.Bool(false),
                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                    Format: cribl.String("asperiores"),
                    Length: cribl.Int64(404306),
                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                },
                TimestampAnchorRegex: "debitis",
                TimestampEarliest: cribl.String("delectus"),
                TimestampLatest: cribl.String("quae"),
                TimestampTimezone: cribl.String("minus"),
                Type: shared.EventBreakerRulesetRulesEventBreakerTypeTimestamp,
            },
            shared.EventBreakerRulesetRules{
                Condition: "laborum",
                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                    DstField: cribl.String("consectetur"),
                    FieldFilterExpr: cribl.String("velit"),
                    Fields: []string{
                        "ipsum",
                        "impedit",
                        "magni",
                    },
                    Keep: []string{
                        "repudiandae",
                        "nam",
                        "dolore",
                    },
                    Remove: []string{
                        "voluptate",
                        "sequi",
                    },
                },
                Disabled: cribl.Bool(false),
                Fields: []shared.EventBreakerRulesetRulesFields{
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Angelica Leuschke"),
                        Value: "odit",
                    },
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Cecil Gutkowski DDS"),
                        Value: "libero",
                    },
                },
                MaxEventBytes: cribl.Int64(102413),
                Name: "Howard Sauer",
                ParserEnabled: cribl.Bool(false),
                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                    Format: cribl.String("illo"),
                    Length: cribl.Int64(36561),
                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                },
                TimestampAnchorRegex: "ea",
                TimestampEarliest: cribl.String("beatae"),
                TimestampLatest: cribl.String("vero"),
                TimestampTimezone: cribl.String("excepturi"),
                Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
            },
        },
        Tags: cribl.String("velit"),
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
    id := "ut"
    eventBreakerRuleset := &shared.EventBreakerRuleset{
        Description: cribl.String("perspiciatis"),
        ID: "e1cf9e06-e3a4-4370-80ae-6b6bc9b8f759",
        Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
        MinRawLength: cribl.Int64(897543),
        Rules: []shared.EventBreakerRulesetRules{
            shared.EventBreakerRulesetRules{
                Condition: "impedit",
                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                    DstField: cribl.String("ipsam"),
                    FieldFilterExpr: cribl.String("corporis"),
                    Fields: []string{
                        "error",
                        "esse",
                        "labore",
                    },
                    Keep: []string{
                        "vero",
                    },
                    Remove: []string{
                        "vitae",
                    },
                },
                Disabled: cribl.Bool(false),
                Fields: []shared.EventBreakerRulesetRulesFields{
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Lauren Deckow"),
                        Value: "nemo",
                    },
                },
                MaxEventBytes: cribl.Int64(745233),
                Name: "Wallace Pagac",
                ParserEnabled: cribl.Bool(false),
                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                    Format: cribl.String("alias"),
                    Length: cribl.Int64(168042),
                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                },
                TimestampAnchorRegex: "quae",
                TimestampEarliest: cribl.String("quae"),
                TimestampLatest: cribl.String("modi"),
                TimestampTimezone: cribl.String("neque"),
                Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
            },
            shared.EventBreakerRulesetRules{
                Condition: "itaque",
                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                    DstField: cribl.String("et"),
                    FieldFilterExpr: cribl.String("ipsum"),
                    Fields: []string{
                        "nulla",
                        "distinctio",
                        "maxime",
                    },
                    Keep: []string{
                        "quia",
                    },
                    Remove: []string{
                        "omnis",
                        "libero",
                    },
                },
                Disabled: cribl.Bool(false),
                Fields: []shared.EventBreakerRulesetRulesFields{
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Wm Steuber"),
                        Value: "placeat",
                    },
                },
                MaxEventBytes: cribl.Int64(25756),
                Name: "Mr. Angela Volkman",
                ParserEnabled: cribl.Bool(false),
                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                    Format: cribl.String("dolore"),
                    Length: cribl.Int64(755106),
                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                },
                TimestampAnchorRegex: "voluptatem",
                TimestampEarliest: cribl.String("autem"),
                TimestampLatest: cribl.String("esse"),
                TimestampTimezone: cribl.String("dolores"),
                Type: shared.EventBreakerRulesetRulesEventBreakerTypeTimestamp,
            },
            shared.EventBreakerRulesetRules{
                Condition: "beatae",
                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                    DstField: cribl.String("est"),
                    FieldFilterExpr: cribl.String("facere"),
                    Fields: []string{
                        "molestiae",
                        "provident",
                        "accusamus",
                    },
                    Keep: []string{
                        "tempore",
                        "sint",
                        "ea",
                        "autem",
                    },
                    Remove: []string{
                        "rerum",
                        "laudantium",
                    },
                },
                Disabled: cribl.Bool(false),
                Fields: []shared.EventBreakerRulesetRulesFields{
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Boyd Rippin Sr."),
                        Value: "quidem",
                    },
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Phil Barton"),
                        Value: "eos",
                    },
                },
                MaxEventBytes: cribl.Int64(844854),
                Name: "Mrs. Mabel Connelly",
                ParserEnabled: cribl.Bool(false),
                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                    Format: cribl.String("earum"),
                    Length: cribl.Int64(239337),
                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                },
                TimestampAnchorRegex: "similique",
                TimestampEarliest: cribl.String("ut"),
                TimestampLatest: cribl.String("quidem"),
                TimestampTimezone: cribl.String("quis"),
                Type: shared.EventBreakerRulesetRulesEventBreakerTypeRegex,
            },
        },
        Tags: cribl.String("unde"),
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

