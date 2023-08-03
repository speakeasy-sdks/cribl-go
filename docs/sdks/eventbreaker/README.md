# EventBreaker

### Available Operations

* [Delete](#delete) - Delete Event Breaker Ruleset
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
    id := "voluptatibus"

    ctx := context.Background()
    res, err := s.EventBreaker.Delete(ctx, id)
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
    res, err := s.EventBreaker.Post(ctx, shared.EventBreakerRuleset{
        Description: cribl.String("quia"),
        ID: "c4310661-e963-449e-9cf9-e06e3a437000",
        Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
        MinRawLength: cribl.Int64(639705),
        Rules: []shared.EventBreakerRulesetRules{
            shared.EventBreakerRulesetRules{
                Condition: "ea",
                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                    DstField: cribl.String("quidem"),
                    FieldFilterExpr: cribl.String("voluptas"),
                    Fields: []string{
                        "placeat",
                        "perspiciatis",
                        "expedita",
                    },
                    Keep: []string{
                        "a",
                        "voluptate",
                        "ullam",
                    },
                    Remove: []string{
                        "necessitatibus",
                        "animi",
                        "impedit",
                    },
                },
                Disabled: cribl.Bool(false),
                Fields: []shared.EventBreakerRulesetRulesFields{
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Jodi Mueller"),
                        Value: "veritatis",
                    },
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Mrs. Glenn Bruen"),
                        Value: "qui",
                    },
                },
                MaxEventBytes: cribl.Int64(611328),
                Name: "Vivian Rodriguez",
                ParserEnabled: cribl.Bool(false),
                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                    Format: cribl.String("dolorum"),
                    Length: cribl.Int64(487676),
                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                },
                TimestampAnchorRegex: "alias",
                TimestampEarliest: cribl.String("magni"),
                TimestampLatest: cribl.String("vel"),
                TimestampTimezone: cribl.String("quae"),
                Type: shared.EventBreakerRulesetRulesEventBreakerTypeRegex,
            },
            shared.EventBreakerRulesetRules{
                Condition: "modi",
                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                    DstField: cribl.String("neque"),
                    FieldFilterExpr: cribl.String("exercitationem"),
                    Fields: []string{
                        "et",
                        "ipsum",
                        "unde",
                        "nulla",
                    },
                    Keep: []string{
                        "maxime",
                        "quia",
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
            shared.EventBreakerRulesetRules{
                Condition: "unde",
                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                    DstField: cribl.String("molestiae"),
                    FieldFilterExpr: cribl.String("delectus"),
                    Fields: []string{
                        "fugit",
                        "numquam",
                        "numquam",
                    },
                    Keep: []string{
                        "at",
                    },
                    Remove: []string{
                        "dignissimos",
                        "optio",
                        "necessitatibus",
                    },
                },
                Disabled: cribl.Bool(false),
                Fields: []shared.EventBreakerRulesetRulesFields{
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Kristy Lemke"),
                        Value: "placeat",
                    },
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Gladys King"),
                        Value: "modi",
                    },
                },
                MaxEventBytes: cribl.Int64(357347),
                Name: "Tasha Wolff DDS",
                ParserEnabled: cribl.Bool(false),
                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                    Format: cribl.String("ratione"),
                    Length: cribl.Int64(289913),
                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                },
                TimestampAnchorRegex: "occaecati",
                TimestampEarliest: cribl.String("voluptas"),
                TimestampLatest: cribl.String("quo"),
                TimestampTimezone: cribl.String("velit"),
                Type: shared.EventBreakerRulesetRulesEventBreakerTypeTimestamp,
            },
        },
        Tags: cribl.String("fuga"),
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
    id := "nostrum"
    eventBreakerRuleset := &shared.EventBreakerRuleset{
        Description: cribl.String("est"),
        ID: "cfbe2fd5-7075-4779-a917-7deac646ecb5",
        Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
        MinRawLength: cribl.Int64(463344),
        Rules: []shared.EventBreakerRulesetRules{
            shared.EventBreakerRulesetRules{
                Condition: "modi",
                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                    DstField: cribl.String("ipsa"),
                    FieldFilterExpr: cribl.String("sint"),
                    Fields: []string{
                        "sequi",
                        "repudiandae",
                        "cum",
                        "dicta",
                    },
                    Keep: []string{
                        "veniam",
                        "animi",
                        "dolores",
                        "nam",
                    },
                    Remove: []string{
                        "consequuntur",
                    },
                },
                Disabled: cribl.Bool(false),
                Fields: []shared.EventBreakerRulesetRulesFields{
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Larry Kuphal Sr."),
                        Value: "laboriosam",
                    },
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Pete Mann"),
                        Value: "aliquam",
                    },
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Nettie Rosenbaum"),
                        Value: "hic",
                    },
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Willard Larson"),
                        Value: "eaque",
                    },
                },
                MaxEventBytes: cribl.Int64(901163),
                Name: "Billie Morar",
                ParserEnabled: cribl.Bool(false),
                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                    Format: cribl.String("libero"),
                    Length: cribl.Int64(244032),
                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                },
                TimestampAnchorRegex: "delectus",
                TimestampEarliest: cribl.String("impedit"),
                TimestampLatest: cribl.String("cum"),
                TimestampTimezone: cribl.String("ipsum"),
                Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSON,
            },
        },
        Tags: cribl.String("saepe"),
    }

    ctx := context.Background()
    res, err := s.EventBreaker.Update(ctx, id, eventBreakerRuleset)
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

