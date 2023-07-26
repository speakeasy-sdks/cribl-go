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
    res, err := s.EventBreaker.Delete(ctx, operations.DeleteEventBreakerRequest{
        ID: "ac600dec-001a-4c80-ae2e-c09ff8f0f816",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteEventBreakerRequest](../../models/operations/deleteeventbreakerrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


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
    res, err := s.EventBreaker.Post(ctx, shared.EventBreakerRuleset{
        Description: cribl.String("earum"),
        ID: "f3477c13-e902-4c14-925b-0960a668151a",
        Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
        MinRawLength: cribl.Int64(278116),
        Rules: []shared.EventBreakerRulesetRules{
            shared.EventBreakerRulesetRules{
                Condition: "magni",
                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                    DstField: cribl.String("deserunt"),
                    FieldFilterExpr: cribl.String("delectus"),
                    Fields: []string{
                        "sed",
                        "nesciunt",
                        "maxime",
                    },
                    Keep: []string{
                        "cupiditate",
                        "aliquam",
                    },
                    Remove: []string{
                        "maiores",
                        "laudantium",
                        "velit",
                    },
                },
                Disabled: cribl.Bool(false),
                Fields: []shared.EventBreakerRulesetRulesFields{
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Renee Beer"),
                        Value: "quas",
                    },
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Elsie Yundt"),
                        Value: "perspiciatis",
                    },
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Mildred Schinner"),
                        Value: "porro",
                    },
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Abraham Gleason"),
                        Value: "eius",
                    },
                },
                MaxEventBytes: cribl.Int64(194058),
                Name: "Marlon Koelpin",
                ParserEnabled: cribl.Bool(false),
                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                    Format: cribl.String("repellat"),
                    Length: cribl.Int64(955047),
                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                },
                TimestampAnchorRegex: "maiores",
                TimestampEarliest: cribl.String("itaque"),
                TimestampLatest: cribl.String("nulla"),
                TimestampTimezone: cribl.String("deserunt"),
                Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
            },
            shared.EventBreakerRulesetRules{
                Condition: "velit",
                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                    DstField: cribl.String("officiis"),
                    FieldFilterExpr: cribl.String("enim"),
                    Fields: []string{
                        "saepe",
                        "eum",
                        "repudiandae",
                    },
                    Keep: []string{
                        "officia",
                    },
                    Remove: []string{
                        "quasi",
                        "blanditiis",
                        "eius",
                        "quisquam",
                    },
                },
                Disabled: cribl.Bool(false),
                Fields: []shared.EventBreakerRulesetRulesFields{
                    shared.EventBreakerRulesetRulesFields{
                        Name: cribl.String("Jeremiah Schimmel"),
                        Value: "reprehenderit",
                    },
                },
                MaxEventBytes: cribl.Int64(800799),
                Name: "Byron Farrell",
                ParserEnabled: cribl.Bool(false),
                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                    Format: cribl.String("laborum"),
                    Length: cribl.Int64(266946),
                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                },
                TimestampAnchorRegex: "necessitatibus",
                TimestampEarliest: cribl.String("architecto"),
                TimestampLatest: cribl.String("molestias"),
                TimestampTimezone: cribl.String("dolore"),
                Type: shared.EventBreakerRulesetRulesEventBreakerTypeRegex,
            },
        },
        Tags: cribl.String("maiores"),
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
    res, err := s.EventBreaker.Update(ctx, operations.UpdateEventBreakerRequest{
        EventBreakerRuleset: &shared.EventBreakerRuleset{
            Description: cribl.String("neque"),
            ID: "2e550557-56f5-4d56-90bd-0af2dfe13db4",
            Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
            MinRawLength: cribl.Int64(976762),
            Rules: []shared.EventBreakerRulesetRules{
                shared.EventBreakerRulesetRules{
                    Condition: "explicabo",
                    Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                        DstField: cribl.String("minus"),
                        FieldFilterExpr: cribl.String("soluta"),
                        Fields: []string{
                            "velit",
                            "earum",
                            "praesentium",
                        },
                        Keep: []string{
                            "non",
                            "quasi",
                            "mollitia",
                        },
                        Remove: []string{
                            "harum",
                            "cumque",
                            "doloremque",
                            "expedita",
                        },
                    },
                    Disabled: cribl.Bool(false),
                    Fields: []shared.EventBreakerRulesetRulesFields{
                        shared.EventBreakerRulesetRulesFields{
                            Name: cribl.String("Sandy Hyatt"),
                            Value: "tempora",
                        },
                        shared.EventBreakerRulesetRulesFields{
                            Name: cribl.String("Rodney Prohaska"),
                            Value: "optio",
                        },
                        shared.EventBreakerRulesetRulesFields{
                            Name: cribl.String("Kim Ryan"),
                            Value: "voluptatum",
                        },
                    },
                    MaxEventBytes: cribl.Int64(614770),
                    Name: "Dr. Ruth Blanda",
                    ParserEnabled: cribl.Bool(false),
                    Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                        Format: cribl.String("at"),
                        Length: cribl.Int64(823472),
                        Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                    },
                    TimestampAnchorRegex: "pariatur",
                    TimestampEarliest: cribl.String("vel"),
                    TimestampLatest: cribl.String("sapiente"),
                    TimestampTimezone: cribl.String("mollitia"),
                    Type: shared.EventBreakerRulesetRulesEventBreakerTypeRegex,
                },
                shared.EventBreakerRulesetRules{
                    Condition: "quos",
                    Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                        DstField: cribl.String("aperiam"),
                        FieldFilterExpr: cribl.String("non"),
                        Fields: []string{
                            "ad",
                            "aliquam",
                            "quisquam",
                            "quas",
                        },
                        Keep: []string{
                            "maiores",
                        },
                        Remove: []string{
                            "aliquid",
                        },
                    },
                    Disabled: cribl.Bool(false),
                    Fields: []shared.EventBreakerRulesetRulesFields{
                        shared.EventBreakerRulesetRulesFields{
                            Name: cribl.String("Rodney Jacobs"),
                            Value: "rem",
                        },
                        shared.EventBreakerRulesetRulesFields{
                            Name: cribl.String("Allan Ferry"),
                            Value: "blanditiis",
                        },
                        shared.EventBreakerRulesetRulesFields{
                            Name: cribl.String("Miss Emily Lemke DVM"),
                            Value: "autem",
                        },
                    },
                    MaxEventBytes: cribl.Int64(694088),
                    Name: "Lowell Oberbrunner",
                    ParserEnabled: cribl.Bool(false),
                    Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                        Format: cribl.String("voluptas"),
                        Length: cribl.Int64(658199),
                        Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                    },
                    TimestampAnchorRegex: "alias",
                    TimestampEarliest: cribl.String("fuga"),
                    TimestampLatest: cribl.String("aut"),
                    TimestampTimezone: cribl.String("dolore"),
                    Type: shared.EventBreakerRulesetRulesEventBreakerTypeTimestamp,
                },
            },
            Tags: cribl.String("aliquam"),
        },
        ID: "95cc6991-71b5-41c1-bdb1-cf4b888ebdfc",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateEventBreakerRequest](../../models/operations/updateeventbreakerrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateEventBreakerResponse](../../models/operations/updateeventbreakerresponse.md), error**

