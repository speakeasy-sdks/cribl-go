# SearchJob

### Available Operations

* [Create](#create) - Create SearchJob
* [Delete](#delete) - Delete SearchJob
* [Get](#get) - Get SearchJob by ID
* [Update](#update) - Update SearchJob

## Create

Create SearchJob

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
    res, err := s.SearchJob.Create(ctx, shared.SearchJob{
        ChartConfig: &shared.ChartConfig{
            Axis: &shared.ChartConfigAxis{
                XAxis: cribl.String("delectus"),
                YAxis: []string{
                    "totam",
                    "quidem",
                },
            },
            Legend: &shared.Legend{},
            Series: []shared.ChartSeries{
                shared.ChartSeries{
                    Color: cribl.String("sequi"),
                    Data: []map[string]interface{}{
                        map[string]interface{}{
                            "modi": "dolores",
                            "beatae": "rerum",
                            "repellendus": "ut",
                            "nesciunt": "facere",
                        },
                        map[string]interface{}{
                            "delectus": "ipsa",
                        },
                        map[string]interface{}{
                            "libero": "eaque",
                            "animi": "doloremque",
                            "consequatur": "perferendis",
                            "dolor": "earum",
                        },
                    },
                    DataExpression: "facilis",
                    DataFilter: map[string]interface{}{
                        "sed": "fugiat",
                    },
                    Name: "Garry Effertz",
                    Type: &shared.ChartType{},
                },
            },
            Settings: &shared.Settings{
                Color: cribl.String("aperiam"),
                ColorPalette: 867423,
                Type: "excepturi",
            },
            SingleValue: shared.SingleValue{
                Color: cribl.String("aliquam"),
                Decimals: cribl.Int64(973256),
                Label: cribl.String("fuga"),
                Prefix: cribl.String("deserunt"),
                Suffix: cribl.String("iure"),
                Type: cribl.String("labore"),
            },
            XAxis: &shared.Axis{
                AxisLabel: &shared.AxisLabel{
                    Formatter: &shared.AxisLabelFormatter{},
                    FormatterData: []int64{
                        767608,
                    },
                },
                Type: cribl.String("corporis"),
            },
        },
        CompatibilityChecks: &shared.SearchJobCompatibilityChecks{
            Datatypes: cribl.Bool(false),
        },
        CorrelationID: cribl.String("odio"),
        CPUMetrics: &shared.CPUTimeMetric{
            TotalCPUSeconds: 841280,
            TotalExecCPUSeconds: 116932,
        },
        DatatypeOverrides: &shared.DatatypeOverrides{
            BreakerRulesets: []shared.EventBreakerRuleset{
                shared.EventBreakerRuleset{
                    Description: cribl.String("vero"),
                    ID: "dc2050d3-8dc3-4ce1-8547-2f9ee69166a8",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(734741),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "adipisci",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("dolore"),
                                FieldFilterExpr: cribl.String("tempora"),
                                Fields: []string{
                                    "debitis",
                                    "similique",
                                },
                                Keep: []string{
                                    "blanditiis",
                                    "libero",
                                    "sequi",
                                    "laborum",
                                },
                                Remove: []string{
                                    "totam",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Angelica Huel DVM"),
                                    Value: "officiis",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Lisa Huel IV"),
                                    Value: "repellendus",
                                },
                            },
                            MaxEventBytes: cribl.Int64(184285),
                            Name: "Evan Russel",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("officia"),
                                Length: cribl.Int64(920194),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "voluptatem",
                            TimestampEarliest: cribl.String("quod"),
                            TimestampLatest: cribl.String("vitae"),
                            TimestampTimezone: cribl.String("vel"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "laboriosam",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("veritatis"),
                                FieldFilterExpr: cribl.String("similique"),
                                Fields: []string{
                                    "temporibus",
                                },
                                Keep: []string{
                                    "veritatis",
                                    "ipsum",
                                    "iure",
                                },
                                Remove: []string{
                                    "molestiae",
                                    "itaque",
                                    "voluptatum",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Mrs. Ethel Deckow"),
                                    Value: "velit",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Leslie Hirthe"),
                                    Value: "veniam",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Krystal Quitzon"),
                                    Value: "aliquam",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Dean Mayert MD"),
                                    Value: "veniam",
                                },
                            },
                            MaxEventBytes: cribl.Int64(409677),
                            Name: "Terence Pfeffer",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("pariatur"),
                                Length: cribl.Int64(672505),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                            },
                            TimestampAnchorRegex: "omnis",
                            TimestampEarliest: cribl.String("maxime"),
                            TimestampLatest: cribl.String("officia"),
                            TimestampTimezone: cribl.String("iusto"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeHeader,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "ab",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("deleniti"),
                                FieldFilterExpr: cribl.String("et"),
                                Fields: []string{
                                    "sint",
                                    "ipsam",
                                    "laboriosam",
                                    "molestiae",
                                },
                                Keep: []string{
                                    "ex",
                                },
                                Remove: []string{
                                    "dolorem",
                                    "minus",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Cynthia Reichert"),
                                    Value: "nisi",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Frances Johnson"),
                                    Value: "non",
                                },
                            },
                            MaxEventBytes: cribl.Int64(382928),
                            Name: "Mr. Margie Heller",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("tempore"),
                                Length: cribl.Int64(144579),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "magni",
                            TimestampEarliest: cribl.String("sunt"),
                            TimestampLatest: cribl.String("quidem"),
                            TimestampTimezone: cribl.String("perspiciatis"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeCsv,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "eos",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("saepe"),
                                FieldFilterExpr: cribl.String("ipsa"),
                                Fields: []string{
                                    "consequuntur",
                                    "aliquam",
                                },
                                Keep: []string{
                                    "reprehenderit",
                                    "quidem",
                                },
                                Remove: []string{
                                    "officia",
                                    "modi",
                                    "alias",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Paul Harvey"),
                                    Value: "libero",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Cristina Huels DVM"),
                                    Value: "facere",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Sean Daugherty"),
                                    Value: "incidunt",
                                },
                            },
                            MaxEventBytes: cribl.Int64(866777),
                            Name: "Deborah Waters",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("accusantium"),
                                Length: cribl.Int64(890631),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "aliquam",
                            TimestampEarliest: cribl.String("quasi"),
                            TimestampLatest: cribl.String("repellendus"),
                            TimestampTimezone: cribl.String("architecto"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeCsv,
                        },
                    },
                    Tags: cribl.String("labore"),
                },
                shared.EventBreakerRuleset{
                    Description: cribl.String("nisi"),
                    ID: "5e85156f-ff73-4fdf-94fd-d5ea9543398d",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(671516),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "soluta",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("eius"),
                                FieldFilterExpr: cribl.String("dolores"),
                                Fields: []string{
                                    "deleniti",
                                    "assumenda",
                                    "iure",
                                },
                                Keep: []string{
                                    "neque",
                                },
                                Remove: []string{
                                    "quos",
                                    "saepe",
                                    "incidunt",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("David Dooley"),
                                    Value: "culpa",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Tabitha Monahan IV"),
                                    Value: "similique",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Rhonda Gislason"),
                                    Value: "nisi",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Guadalupe Ankunding"),
                                    Value: "nulla",
                                },
                            },
                            MaxEventBytes: cribl.Int64(676138),
                            Name: "Sammy Franecki",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("quibusdam"),
                                Length: cribl.Int64(40155),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                            },
                            TimestampAnchorRegex: "nobis",
                            TimestampEarliest: cribl.String("laboriosam"),
                            TimestampLatest: cribl.String("esse"),
                            TimestampTimezone: cribl.String("et"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeCsv,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "optio",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("quam"),
                                FieldFilterExpr: cribl.String("delectus"),
                                Fields: []string{
                                    "vero",
                                },
                                Keep: []string{
                                    "culpa",
                                },
                                Remove: []string{
                                    "quasi",
                                    "veniam",
                                    "provident",
                                    "consequuntur",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Eduardo Bayer DDS"),
                                    Value: "magnam",
                                },
                            },
                            MaxEventBytes: cribl.Int64(615698),
                            Name: "Katherine Ziemann",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("quibusdam"),
                                Length: cribl.Int64(519586),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "maxime",
                            TimestampEarliest: cribl.String("totam"),
                            TimestampLatest: cribl.String("id"),
                            TimestampTimezone: cribl.String("neque"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSON,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "vel",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("ipsum"),
                                FieldFilterExpr: cribl.String("occaecati"),
                                Fields: []string{
                                    "similique",
                                    "quis",
                                    "facilis",
                                    "in",
                                },
                                Keep: []string{
                                    "nisi",
                                    "natus",
                                    "eaque",
                                },
                                Remove: []string{
                                    "facilis",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Joshua Padberg"),
                                    Value: "maiores",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Emma Dickinson"),
                                    Value: "tempora",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Miss Willard Welch"),
                                    Value: "rem",
                                },
                            },
                            MaxEventBytes: cribl.Int64(796754),
                            Name: "Faye Braun",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("dolor"),
                                Length: cribl.Int64(155108),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                            },
                            TimestampAnchorRegex: "occaecati",
                            TimestampEarliest: cribl.String("doloribus"),
                            TimestampLatest: cribl.String("libero"),
                            TimestampTimezone: cribl.String("culpa"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeCsv,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "molestias",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("magnam"),
                                FieldFilterExpr: cribl.String("voluptate"),
                                Fields: []string{
                                    "officia",
                                    "qui",
                                },
                                Keep: []string{
                                    "vero",
                                    "quas",
                                    "possimus",
                                },
                                Remove: []string{
                                    "quo",
                                    "ullam",
                                    "ipsa",
                                    "placeat",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Mr. Jeffery Hegmann"),
                                    Value: "iusto",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Allison Kutch"),
                                    Value: "provident",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Linda Kulas MD"),
                                    Value: "alias",
                                },
                            },
                            MaxEventBytes: cribl.Int64(25872),
                            Name: "Sergio Johns",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("eligendi"),
                                Length: cribl.Int64(496307),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                            },
                            TimestampAnchorRegex: "eum",
                            TimestampEarliest: cribl.String("possimus"),
                            TimestampLatest: cribl.String("dolore"),
                            TimestampTimezone: cribl.String("neque"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeRegex,
                        },
                    },
                    Tags: cribl.String("omnis"),
                },
                shared.EventBreakerRuleset{
                    Description: cribl.String("quaerat"),
                    ID: "398c783c-9239-48ed-bd3a-b7ca3c5ca864",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(624011),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "ducimus",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("perferendis"),
                                FieldFilterExpr: cribl.String("impedit"),
                                Fields: []string{
                                    "quibusdam",
                                    "veniam",
                                    "pariatur",
                                    "commodi",
                                },
                                Keep: []string{
                                    "corrupti",
                                    "iste",
                                    "distinctio",
                                },
                                Remove: []string{
                                    "consequuntur",
                                    "voluptatem",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Ms. Erin Bergnaum"),
                                    Value: "illum",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Katrina Tillman"),
                                    Value: "ratione",
                                },
                            },
                            MaxEventBytes: cribl.Int64(817807),
                            Name: "Erika Cruickshank",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("inventore"),
                                Length: cribl.Int64(280313),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                            },
                            TimestampAnchorRegex: "molestias",
                            TimestampEarliest: cribl.String("fuga"),
                            TimestampLatest: cribl.String("odit"),
                            TimestampTimezone: cribl.String("minus"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeRegex,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "sint",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("exercitationem"),
                                FieldFilterExpr: cribl.String("magnam"),
                                Fields: []string{
                                    "quaerat",
                                    "minima",
                                },
                                Keep: []string{
                                    "unde",
                                    "ullam",
                                    "enim",
                                    "facere",
                                },
                                Remove: []string{
                                    "cumque",
                                    "et",
                                    "praesentium",
                                    "minima",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Mr. Oscar Mitchell"),
                                    Value: "iure",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Calvin Dietrich"),
                                    Value: "explicabo",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Cameron O'Kon"),
                                    Value: "incidunt",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Dominic Nienow"),
                                    Value: "sed",
                                },
                            },
                            MaxEventBytes: cribl.Int64(229336),
                            Name: "Silvia Stoltenberg",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("dolor"),
                                Length: cribl.Int64(516716),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                            },
                            TimestampAnchorRegex: "vitae",
                            TimestampEarliest: cribl.String("laborum"),
                            TimestampLatest: cribl.String("beatae"),
                            TimestampTimezone: cribl.String("vitae"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "non",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("laudantium"),
                                FieldFilterExpr: cribl.String("consequuntur"),
                                Fields: []string{
                                    "nulla",
                                    "ducimus",
                                    "eveniet",
                                },
                                Keep: []string{
                                    "enim",
                                    "voluptas",
                                    "veniam",
                                    "voluptatem",
                                },
                                Remove: []string{
                                    "vel",
                                    "aspernatur",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Leon Lesch"),
                                    Value: "nulla",
                                },
                            },
                            MaxEventBytes: cribl.Int64(481914),
                            Name: "Sherri Hirthe",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("modi"),
                                Length: cribl.Int64(763869),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                            },
                            TimestampAnchorRegex: "accusantium",
                            TimestampEarliest: cribl.String("fuga"),
                            TimestampLatest: cribl.String("doloremque"),
                            TimestampTimezone: cribl.String("reprehenderit"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeRegex,
                        },
                    },
                    Tags: cribl.String("veritatis"),
                },
                shared.EventBreakerRuleset{
                    Description: cribl.String("similique"),
                    ID: "961d24a7-dbb8-4f53-ad89-2cf7812cb512",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(771645),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "esse",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("voluptatum"),
                                FieldFilterExpr: cribl.String("aspernatur"),
                                Fields: []string{
                                    "accusantium",
                                    "nobis",
                                },
                                Keep: []string{
                                    "corporis",
                                    "tempora",
                                    "voluptatum",
                                    "a",
                                },
                                Remove: []string{
                                    "blanditiis",
                                    "hic",
                                    "blanditiis",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Miss Maryann Wiegand"),
                                    Value: "totam",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Willie Wunsch III"),
                                    Value: "nulla",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Meredith Lynch PhD"),
                                    Value: "doloremque",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Ms. Mabel Bernhard DVM"),
                                    Value: "voluptas",
                                },
                            },
                            MaxEventBytes: cribl.Int64(470523),
                            Name: "Loretta Howe",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("doloribus"),
                                Length: cribl.Int64(187770),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "ex",
                            TimestampEarliest: cribl.String("quos"),
                            TimestampLatest: cribl.String("dicta"),
                            TimestampTimezone: cribl.String("minus"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "molestiae",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("iure"),
                                FieldFilterExpr: cribl.String("blanditiis"),
                                Fields: []string{
                                    "impedit",
                                    "itaque",
                                    "molestiae",
                                    "quaerat",
                                },
                                Keep: []string{
                                    "dolore",
                                },
                                Remove: []string{
                                    "excepturi",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Heather Heller IV"),
                                    Value: "voluptas",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Jean Goldner"),
                                    Value: "officia",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Ora Homenick"),
                                    Value: "neque",
                                },
                            },
                            MaxEventBytes: cribl.Int64(648985),
                            Name: "Melvin Stracke",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("quibusdam"),
                                Length: cribl.Int64(841408),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "dolorem",
                            TimestampEarliest: cribl.String("velit"),
                            TimestampLatest: cribl.String("vero"),
                            TimestampTimezone: cribl.String("placeat"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeCsv,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "vel",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("non"),
                                FieldFilterExpr: cribl.String("incidunt"),
                                Fields: []string{
                                    "ipsum",
                                    "itaque",
                                    "non",
                                },
                                Keep: []string{
                                    "esse",
                                    "id",
                                    "natus",
                                },
                                Remove: []string{
                                    "saepe",
                                    "modi",
                                    "assumenda",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Pearl Trantow"),
                                    Value: "libero",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Arturo Harris"),
                                    Value: "aliquam",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Mrs. Tracy Walker"),
                                    Value: "numquam",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Mr. Todd Feil"),
                                    Value: "perferendis",
                                },
                            },
                            MaxEventBytes: cribl.Int64(439135),
                            Name: "Bryant Dickens",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("quo"),
                                Length: cribl.Int64(14112),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "sed",
                            TimestampEarliest: cribl.String("harum"),
                            TimestampLatest: cribl.String("vero"),
                            TimestampTimezone: cribl.String("nihil"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSON,
                        },
                    },
                    Tags: cribl.String("incidunt"),
                },
            },
            DisableBreakers: false,
        },
        EarliestEpoch: cribl.Int64(953959),
        ErrorStateConfig: &shared.SearchJobErrorStateConfig{
            Coordinated: false,
            ErrorMessages: []string{
                "consequuntur",
            },
        },
        Group: "numquam",
        ID: "49d86f4b-b20f-4e5d-911c-bfe749caf45a",
        LatestEpoch: cribl.Int64(151039),
        NumEventsAfter: cribl.Int64(474867),
        NumEventsBefore: cribl.Int64(964857),
        Query: "ex",
        SampleRate: cribl.Int64(624255),
        SearchConfig: shared.SearchConfig{
            Datasets: []string{
                "consequuntur",
                "quod",
                "occaecati",
                "earum",
            },
            HasSendOperator: false,
            OrderedFieldNames: []string{
                "temporibus",
                "quae",
            },
            SearchTerms: []shared.SearchTerm{
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "debitis",
                },
            },
            SortFields: []shared.SortByField{
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionDescending,
                    FieldName: "quidem",
                    NullPosition: shared.SortByFieldNullPositionNullsFirst,
                },
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionDescending,
                    FieldName: "repellendus",
                    NullPosition: shared.SortByFieldNullPositionNullsFirst,
                },
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionDescending,
                    FieldName: "vel",
                    NullPosition: shared.SortByFieldNullPositionNullsLast,
                },
            },
            SuppressEventMarking: false,
            UseFormattedVisualization: false,
        },
        SearchParameterDeclarations: []shared.SearchParameter{
            shared.SearchParameter{
                DefaultValue: cribl.String("adipisci"),
                Name: "Ruth MacGyver",
                Type: shared.SearchParameterTypeBoolean,
            },
        },
        SearchParameterValues: cribl.String("nesciunt"),
        Status: shared.SearchJobStatusRunning,
        TableConfig: &shared.TableViewSettings{
            ColumnFilterSettings: &shared.ColumnFilterSettings{
                Contains: "in",
            },
            ColumnFormatSettings: &shared.ColumnFormatSettings{
                Palette: "modi",
                Precision: "nihil",
                Prefix: "velit",
                Suffix: "aut",
            },
            ColumnOrderSettings: &shared.ColumnOrderSettings{
                Order: "rem",
            },
            ColumnSortSettings: &shared.ColumnSortSettings{
                Sort: "odit",
            },
            RowNumberColumnWidth: cribl.Int64(727641),
            ShowColumnTotals: false,
            ShowColumnTotalsPinned: false,
            ShowRowNumbers: false,
            ShowRowTotals: false,
            ShowRowTotalsPinned: false,
            ViewMode: shared.TableViewSettingsViewModeTable,
        },
        TargetEventTime: cribl.Int64(278751),
        TimeCompleted: cribl.Int64(992821),
        TimeCreated: 181311,
        TimeStarted: cribl.Int64(650392),
        Type: shared.SearchJobTypeScheduled.ToPointer(),
        User: "dicta",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.SearchJob](../../models/shared/searchjob.md)  | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateSearchJobResponse](../../models/operations/createsearchjobresponse.md), error**


## Delete

Delete SearchJob

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
    res, err := s.SearchJob.Delete(ctx, operations.DeleteSearchJobRequest{
        ID: "fd5671e9-c326-4350-a467-143789ce0e99",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteSearchJobRequest](../../models/operations/deletesearchjobrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.DeleteSearchJobResponse](../../models/operations/deletesearchjobresponse.md), error**


## Get

Get SearchJob by ID

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
    res, err := s.SearchJob.Get(ctx, operations.GetSearchJobRequest{
        ID: "1594d93a-74c0-4252-be3b-4b4db8b778eb",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetSearchJobRequest](../../models/operations/getsearchjobrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetSearchJobResponse](../../models/operations/getsearchjobresponse.md), error**


## Update

Update SearchJob

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
    res, err := s.SearchJob.Update(ctx, operations.UpdateSearchJobRequest{
        SearchJob: &shared.SearchJob{
            ChartConfig: &shared.ChartConfig{
                Axis: &shared.ChartConfigAxis{
                    XAxis: cribl.String("nam"),
                    YAxis: []string{
                        "officiis",
                        "quasi",
                    },
                },
                Legend: &shared.Legend{},
                Series: []shared.ChartSeries{
                    shared.ChartSeries{
                        Color: cribl.String("aspernatur"),
                        Data: []map[string]interface{}{
                            map[string]interface{}{
                                "ad": "eaque",
                                "aspernatur": "expedita",
                                "laborum": "hic",
                                "rerum": "explicabo",
                            },
                            map[string]interface{}{
                                "nam": "placeat",
                                "aliquam": "aliquid",
                                "adipisci": "ipsam",
                                "fugiat": "enim",
                            },
                            map[string]interface{}{
                                "eum": "exercitationem",
                                "at": "culpa",
                                "alias": "eos",
                                "quos": "quisquam",
                            },
                            map[string]interface{}{
                                "accusamus": "sint",
                            },
                        },
                        DataExpression: "enim",
                        DataFilter: map[string]interface{}{
                            "mollitia": "veritatis",
                        },
                        Name: "Philip Armstrong",
                        Type: &shared.ChartType{},
                    },
                    shared.ChartSeries{
                        Color: cribl.String("deserunt"),
                        Data: []map[string]interface{}{
                            map[string]interface{}{
                                "commodi": "magnam",
                                "rem": "occaecati",
                            },
                            map[string]interface{}{
                                "voluptate": "nam",
                                "quam": "blanditiis",
                                "laboriosam": "odio",
                                "adipisci": "necessitatibus",
                            },
                            map[string]interface{}{
                                "consectetur": "fuga",
                            },
                        },
                        DataExpression: "quasi",
                        DataFilter: map[string]interface{}{
                            "dolorum": "vel",
                        },
                        Name: "Terrance Mante",
                        Type: &shared.ChartType{},
                    },
                    shared.ChartSeries{
                        Color: cribl.String("perspiciatis"),
                        Data: []map[string]interface{}{
                            map[string]interface{}{
                                "natus": "numquam",
                                "tempora": "corrupti",
                            },
                            map[string]interface{}{
                                "asperiores": "veniam",
                                "cumque": "praesentium",
                            },
                        },
                        DataExpression: "ut",
                        DataFilter: map[string]interface{}{
                            "blanditiis": "nesciunt",
                        },
                        Name: "Hannah Leffler",
                        Type: &shared.ChartType{},
                    },
                    shared.ChartSeries{
                        Color: cribl.String("consectetur"),
                        Data: []map[string]interface{}{
                            map[string]interface{}{
                                "a": "ex",
                                "dolore": "dicta",
                                "minima": "facilis",
                                "sit": "incidunt",
                            },
                            map[string]interface{}{
                                "molestias": "hic",
                                "error": "repellendus",
                            },
                            map[string]interface{}{
                                "dicta": "ratione",
                                "delectus": "ut",
                                "officiis": "itaque",
                                "nulla": "distinctio",
                            },
                            map[string]interface{}{
                                "in": "deleniti",
                                "tempore": "reiciendis",
                                "commodi": "sit",
                                "ea": "molestias",
                            },
                        },
                        DataExpression: "quia",
                        DataFilter: map[string]interface{}{
                            "rem": "molestias",
                            "eius": "necessitatibus",
                        },
                        Name: "Harvey Jacobi",
                        Type: &shared.ChartType{},
                    },
                },
                Settings: &shared.Settings{
                    Color: cribl.String("corporis"),
                    ColorPalette: 771363,
                    Type: "in",
                },
                SingleValue: shared.SingleValue{
                    Color: cribl.String("fugit"),
                    Decimals: cribl.Int64(454934),
                    Label: cribl.String("provident"),
                    Prefix: cribl.String("quis"),
                    Suffix: cribl.String("expedita"),
                    Type: cribl.String("quam"),
                },
                XAxis: &shared.Axis{
                    AxisLabel: &shared.AxisLabel{
                        Formatter: &shared.AxisLabelFormatter{},
                        FormatterData: []int64{
                            346519,
                            113947,
                            300584,
                        },
                    },
                    Type: cribl.String("blanditiis"),
                },
            },
            CompatibilityChecks: &shared.SearchJobCompatibilityChecks{
                Datatypes: cribl.Bool(false),
            },
            CorrelationID: cribl.String("nulla"),
            CPUMetrics: &shared.CPUTimeMetric{
                TotalCPUSeconds: 432899,
                TotalExecCPUSeconds: 845164,
            },
            DatatypeOverrides: &shared.DatatypeOverrides{
                BreakerRulesets: []shared.EventBreakerRuleset{
                    shared.EventBreakerRuleset{
                        Description: cribl.String("ut"),
                        ID: "9e5635b3-3bc0-4f97-8c42-fc9f4844225e",
                        Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                        MinRawLength: cribl.Int64(450324),
                        Rules: []shared.EventBreakerRulesetRules{
                            shared.EventBreakerRulesetRules{
                                Condition: "expedita",
                                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                    DstField: cribl.String("ducimus"),
                                    FieldFilterExpr: cribl.String("excepturi"),
                                    Fields: []string{
                                        "doloremque",
                                        "iure",
                                    },
                                    Keep: []string{
                                        "cumque",
                                        "eaque",
                                    },
                                    Remove: []string{
                                        "earum",
                                        "culpa",
                                        "vel",
                                        "sapiente",
                                    },
                                },
                                Disabled: cribl.Bool(false),
                                Fields: []shared.EventBreakerRulesetRulesFields{
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Miss Kelli Maggio MD"),
                                        Value: "deleniti",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Austin Hartmann"),
                                        Value: "vitae",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Dana Gottlieb"),
                                        Value: "adipisci",
                                    },
                                },
                                MaxEventBytes: cribl.Int64(591998),
                                Name: "Theodore Wiza",
                                ParserEnabled: cribl.Bool(false),
                                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                    Format: cribl.String("voluptate"),
                                    Length: cribl.Int64(130289),
                                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                                },
                                TimestampAnchorRegex: "perferendis",
                                TimestampEarliest: cribl.String("possimus"),
                                TimestampLatest: cribl.String("dicta"),
                                TimestampTimezone: cribl.String("delectus"),
                                Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                            },
                            shared.EventBreakerRulesetRules{
                                Condition: "exercitationem",
                                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                    DstField: cribl.String("quis"),
                                    FieldFilterExpr: cribl.String("quos"),
                                    Fields: []string{
                                        "cupiditate",
                                        "excepturi",
                                        "quod",
                                        "iure",
                                    },
                                    Keep: []string{
                                        "odit",
                                    },
                                    Remove: []string{
                                        "fugit",
                                        "nam",
                                        "optio",
                                        "accusantium",
                                    },
                                },
                                Disabled: cribl.Bool(false),
                                Fields: []shared.EventBreakerRulesetRulesFields{
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Jesus Barrows"),
                                        Value: "possimus",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Simon Nicolas"),
                                        Value: "doloremque",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Rachel Sporer"),
                                        Value: "placeat",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Carlton Schowalter"),
                                        Value: "modi",
                                    },
                                },
                                MaxEventBytes: cribl.Int64(795501),
                                Name: "Oliver Breitenberg",
                                ParserEnabled: cribl.Bool(false),
                                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                    Format: cribl.String("eveniet"),
                                    Length: cribl.Int64(573938),
                                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                                },
                                TimestampAnchorRegex: "illo",
                                TimestampEarliest: cribl.String("minima"),
                                TimestampLatest: cribl.String("nulla"),
                                TimestampTimezone: cribl.String("reiciendis"),
                                Type: shared.EventBreakerRulesetRulesEventBreakerTypeHeader,
                            },
                        },
                        Tags: cribl.String("ipsa"),
                    },
                    shared.EventBreakerRuleset{
                        Description: cribl.String("amet"),
                        ID: "907f3783-1983-4d42-a54a-85466597c502",
                        Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                        MinRawLength: cribl.Int64(213871),
                        Rules: []shared.EventBreakerRulesetRules{
                            shared.EventBreakerRulesetRules{
                                Condition: "maxime",
                                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                    DstField: cribl.String("inventore"),
                                    FieldFilterExpr: cribl.String("modi"),
                                    Fields: []string{
                                        "vitae",
                                        "quibusdam",
                                    },
                                    Keep: []string{
                                        "et",
                                        "dolorum",
                                    },
                                    Remove: []string{
                                        "mollitia",
                                        "autem",
                                        "illum",
                                    },
                                },
                                Disabled: cribl.Bool(false),
                                Fields: []shared.EventBreakerRulesetRulesFields{
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Corey Pacocha"),
                                        Value: "aliquid",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Penny Kihn"),
                                        Value: "reiciendis",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Brandon Rogahn"),
                                        Value: "aspernatur",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("George Ankunding"),
                                        Value: "maiores",
                                    },
                                },
                                MaxEventBytes: cribl.Int64(422823),
                                Name: "Mr. Elbert Bins II",
                                ParserEnabled: cribl.Bool(false),
                                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                    Format: cribl.String("ducimus"),
                                    Length: cribl.Int64(380884),
                                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                                },
                                TimestampAnchorRegex: "alias",
                                TimestampEarliest: cribl.String("libero"),
                                TimestampLatest: cribl.String("quibusdam"),
                                TimestampTimezone: cribl.String("fuga"),
                                Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                            },
                        },
                        Tags: cribl.String("similique"),
                    },
                },
                DisableBreakers: false,
            },
            EarliestEpoch: cribl.Int64(995708),
            ErrorStateConfig: &shared.SearchJobErrorStateConfig{
                Coordinated: false,
                ErrorMessages: []string{
                    "accusamus",
                    "illum",
                    "blanditiis",
                    "tempora",
                },
            },
            Group: "similique",
            ID: "35a41238-e1a7-435a-826a-e33bef971a8f",
            LatestEpoch: cribl.Int64(267918),
            NumEventsAfter: cribl.Int64(401466),
            NumEventsBefore: cribl.Int64(749037),
            Query: "placeat",
            SampleRate: cribl.Int64(651310),
            SearchConfig: shared.SearchConfig{
                Datasets: []string{
                    "quae",
                },
                HasSendOperator: false,
                OrderedFieldNames: []string{
                    "commodi",
                },
                SearchTerms: []shared.SearchTerm{
                    shared.SearchTerm{
                        IsCaseSensitive: false,
                        Term: "officiis",
                    },
                    shared.SearchTerm{
                        IsCaseSensitive: false,
                        Term: "omnis",
                    },
                    shared.SearchTerm{
                        IsCaseSensitive: false,
                        Term: "ea",
                    },
                    shared.SearchTerm{
                        IsCaseSensitive: false,
                        Term: "ipsam",
                    },
                },
                SortFields: []shared.SortByField{
                    shared.SortByField{
                        Direction: shared.SortByFieldDirectionAscending,
                        FieldName: "vitae",
                        NullPosition: shared.SortByFieldNullPositionNullsFirst,
                    },
                    shared.SortByField{
                        Direction: shared.SortByFieldDirectionDescending,
                        FieldName: "voluptatem",
                        NullPosition: shared.SortByFieldNullPositionNullsLast,
                    },
                    shared.SortByField{
                        Direction: shared.SortByFieldDirectionDescending,
                        FieldName: "tenetur",
                        NullPosition: shared.SortByFieldNullPositionNullsLast,
                    },
                },
                SuppressEventMarking: false,
                UseFormattedVisualization: false,
            },
            SearchParameterDeclarations: []shared.SearchParameter{
                shared.SearchParameter{
                    DefaultValue: cribl.String("necessitatibus"),
                    Name: "Jackie Witting",
                    Type: shared.SearchParameterTypeNumber,
                },
                shared.SearchParameter{
                    DefaultValue: cribl.String("occaecati"),
                    Name: "Miss Maurice Hauck",
                    Type: shared.SearchParameterTypeNumber,
                },
                shared.SearchParameter{
                    DefaultValue: cribl.String("commodi"),
                    Name: "Woodrow Fadel",
                    Type: shared.SearchParameterTypeBoolean,
                },
            },
            SearchParameterValues: cribl.String("cum"),
            Status: shared.SearchJobStatusNew,
            TableConfig: &shared.TableViewSettings{
                ColumnFilterSettings: &shared.ColumnFilterSettings{
                    Contains: "quo",
                },
                ColumnFormatSettings: &shared.ColumnFormatSettings{
                    Palette: "officiis",
                    Precision: "laudantium",
                    Prefix: "est",
                    Suffix: "fuga",
                },
                ColumnOrderSettings: &shared.ColumnOrderSettings{
                    Order: "autem",
                },
                ColumnSortSettings: &shared.ColumnSortSettings{
                    Sort: "quis",
                },
                RowNumberColumnWidth: cribl.Int64(265507),
                ShowColumnTotals: false,
                ShowColumnTotalsPinned: false,
                ShowRowNumbers: false,
                ShowRowTotals: false,
                ShowRowTotalsPinned: false,
                ViewMode: shared.TableViewSettingsViewModeEvent,
            },
            TargetEventTime: cribl.Int64(184401),
            TimeCompleted: cribl.Int64(642434),
            TimeCreated: 585645,
            TimeStarted: cribl.Int64(525089),
            Type: shared.SearchJobTypeDatatypePreview.ToPointer(),
            User: "debitis",
        },
        ID: "b7e14ca5-6408-4915-8097-019a48f88ece",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateSearchJobRequest](../../models/operations/updatesearchjobrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.UpdateSearchJobResponse](../../models/operations/updatesearchjobresponse.md), error**

