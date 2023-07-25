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
                XAxis: cribl.String("pariatur"),
                YAxis: []string{
                    "aperiam",
                },
            },
            Legend: &shared.Legend{},
            Series: []shared.ChartSeries{
                shared.ChartSeries{
                    Color: cribl.String("quo"),
                    Data: []map[string]interface{}{
                        map[string]interface{}{
                            "debitis": "animi",
                        },
                        map[string]interface{}{
                            "quasi": "repellendus",
                            "architecto": "hic",
                        },
                        map[string]interface{}{
                            "nisi": "voluptas",
                            "saepe": "praesentium",
                        },
                    },
                    DataExpression: "veniam",
                    DataFilter: map[string]interface{}{
                        "minima": "nisi",
                    },
                    Name: "Timmy Wilderman",
                    Type: &shared.ChartType{},
                },
                shared.ChartSeries{
                    Color: cribl.String("voluptatibus"),
                    Data: []map[string]interface{}{
                        map[string]interface{}{
                            "nostrum": "tempora",
                            "maiores": "facere",
                            "illum": "exercitationem",
                            "saepe": "deserunt",
                        },
                        map[string]interface{}{
                            "minima": "dolore",
                            "sequi": "sequi",
                            "occaecati": "voluptatum",
                        },
                        map[string]interface{}{
                            "laborum": "sapiente",
                            "soluta": "eius",
                            "dolores": "dolorum",
                            "deleniti": "assumenda",
                        },
                        map[string]interface{}{
                            "dolorem": "neque",
                            "laudantium": "quos",
                        },
                    },
                    DataExpression: "saepe",
                    DataFilter: map[string]interface{}{
                        "pariatur": "deleniti",
                        "aut": "nesciunt",
                    },
                    Name: "Santiago Nicolas",
                    Type: &shared.ChartType{},
                },
                shared.ChartSeries{
                    Color: cribl.String("iste"),
                    Data: []map[string]interface{}{
                        map[string]interface{}{
                            "atque": "similique",
                        },
                        map[string]interface{}{
                            "tempora": "modi",
                        },
                        map[string]interface{}{
                            "possimus": "nisi",
                            "ab": "omnis",
                            "aut": "ipsum",
                            "iste": "nulla",
                        },
                    },
                    DataExpression: "dolorum",
                    DataFilter: map[string]interface{}{
                        "assumenda": "velit",
                        "voluptatum": "eveniet",
                        "quibusdam": "doloremque",
                        "assumenda": "nobis",
                    },
                    Name: "Tamara Borer",
                    Type: &shared.ChartType{},
                },
                shared.ChartSeries{
                    Color: cribl.String("quam"),
                    Data: []map[string]interface{}{
                        map[string]interface{}{
                            "vero": "amet",
                        },
                        map[string]interface{}{
                            "asperiores": "quasi",
                            "veniam": "provident",
                            "consequuntur": "aut",
                        },
                        map[string]interface{}{
                            "provident": "aperiam",
                            "repellendus": "ab",
                            "cum": "magnam",
                            "natus": "aperiam",
                        },
                        map[string]interface{}{
                            "repellat": "dolores",
                        },
                    },
                    DataExpression: "harum",
                    DataFilter: map[string]interface{}{
                        "totam": "provident",
                        "maxime": "totam",
                        "id": "neque",
                        "dolores": "vel",
                    },
                    Name: "Toni Stroman",
                    Type: &shared.ChartType{},
                },
            },
            Settings: &shared.Settings{
                Color: cribl.String("facilis"),
                ColorPalette: 447640,
                Type: "nobis",
            },
            SingleValue: shared.SingleValue{
                Color: cribl.String("nisi"),
                Decimals: cribl.Int64(620858),
                Label: cribl.String("eaque"),
                Prefix: cribl.String("quia"),
                Suffix: cribl.String("facilis"),
                Type: cribl.String("atque"),
            },
            XAxis: &shared.Axis{
                AxisLabel: &shared.AxisLabel{
                    Formatter: &shared.AxisLabelFormatter{},
                    FormatterData: []int64{
                        74814,
                        676140,
                        590800,
                    },
                },
                Type: cribl.String("incidunt"),
            },
        },
        CompatibilityChecks: &shared.SearchJobCompatibilityChecks{
            Datatypes: cribl.Bool(false),
        },
        CorrelationID: cribl.String("maiores"),
        CPUMetrics: &shared.CPUTimeMetric{
            TotalCPUSeconds: 406914,
            TotalExecCPUSeconds: 267247,
        },
        DatatypeOverrides: &shared.DatatypeOverrides{
            BreakerRulesets: []shared.EventBreakerRuleset{
                shared.EventBreakerRuleset{
                    Description: cribl.String("nisi"),
                    ID: "64a8f0af-8c69-41d7-b2d9-fbaf9476a2ae",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(555313),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "optio",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("quo"),
                                FieldFilterExpr: cribl.String("ullam"),
                                Fields: []string{
                                    "placeat",
                                },
                                Keep: []string{
                                    "culpa",
                                    "consectetur",
                                    "nostrum",
                                },
                                Remove: []string{
                                    "eos",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Connie Kling"),
                                    Value: "rem",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Leonard Abshire"),
                                    Value: "eaque",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("James Balistreri"),
                                    Value: "commodi",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Tricia Rosenbaum"),
                                    Value: "eum",
                                },
                            },
                            MaxEventBytes: cribl.Int64(821579),
                            Name: "Rita Bode",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("nesciunt"),
                                Length: cribl.Int64(608470),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "optio",
                            TimestampEarliest: cribl.String("molestiae"),
                            TimestampLatest: cribl.String("atque"),
                            TimestampTimezone: cribl.String("dolor"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeTimestamp,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "cupiditate",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("quia"),
                                FieldFilterExpr: cribl.String("dolor"),
                                Fields: []string{
                                    "praesentium",
                                    "accusamus",
                                    "fugiat",
                                },
                                Keep: []string{
                                    "pariatur",
                                },
                                Remove: []string{
                                    "deserunt",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Chelsea Ortiz"),
                                    Value: "veniam",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Otis Lowe"),
                                    Value: "error",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Kurt Auer"),
                                    Value: "quibusdam",
                                },
                            },
                            MaxEventBytes: cribl.Int64(329700),
                            Name: "Edgar Moore",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("distinctio"),
                                Length: cribl.Int64(449647),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                            },
                            TimestampAnchorRegex: "voluptatem",
                            TimestampEarliest: cribl.String("voluptas"),
                            TimestampLatest: cribl.String("magnam"),
                            TimestampTimezone: cribl.String("ad"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeRegex,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "ipsa",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("iure"),
                                FieldFilterExpr: cribl.String("voluptate"),
                                Fields: []string{
                                    "et",
                                    "perspiciatis",
                                    "accusamus",
                                    "laborum",
                                },
                                Keep: []string{
                                    "ratione",
                                    "facere",
                                    "eius",
                                },
                                Remove: []string{
                                    "consequuntur",
                                    "earum",
                                    "quibusdam",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Hannah MacGyver"),
                                    Value: "minus",
                                },
                            },
                            MaxEventBytes: cribl.Int64(123989),
                            Name: "Greg Gutkowski",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("minima"),
                                Length: cribl.Int64(887620),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "ullam",
                            TimestampEarliest: cribl.String("enim"),
                            TimestampLatest: cribl.String("facere"),
                            TimestampTimezone: cribl.String("cumque"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeTimestamp,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "et",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("praesentium"),
                                FieldFilterExpr: cribl.String("minima"),
                                Fields: []string{
                                    "est",
                                    "magnam",
                                    "unde",
                                    "consequatur",
                                },
                                Keep: []string{
                                    "eligendi",
                                },
                                Remove: []string{
                                    "cumque",
                                    "quaerat",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Carroll Christiansen"),
                                    Value: "animi",
                                },
                            },
                            MaxEventBytes: cribl.Int64(463832),
                            Name: "Troy Murphy",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("ratione"),
                                Length: cribl.Int64(848953),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                            },
                            TimestampAnchorRegex: "amet",
                            TimestampEarliest: cribl.String("aut"),
                            TimestampLatest: cribl.String("voluptates"),
                            TimestampTimezone: cribl.String("nulla"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeCsv,
                        },
                    },
                    Tags: cribl.String("dignissimos"),
                },
            },
            DisableBreakers: false,
        },
        EarliestEpoch: cribl.Int64(220049),
        ErrorStateConfig: &shared.SearchJobErrorStateConfig{
            Coordinated: false,
            ErrorMessages: []string{
                "beatae",
                "vitae",
                "laborum",
            },
        },
        Group: "beatae",
        ID: "15382bd7-ed56-4507-a21c-58f4d7396564",
        LatestEpoch: cribl.Int64(763869),
        NumEventsAfter: cribl.Int64(145964),
        NumEventsBefore: cribl.Int64(33457),
        Query: "fuga",
        SampleRate: cribl.Int64(38622),
        SearchConfig: shared.SearchConfig{
            Datasets: []string{
                "dicta",
                "veritatis",
            },
            HasSendOperator: false,
            OrderedFieldNames: []string{
                "omnis",
                "commodi",
                "dicta",
            },
            SearchTerms: []shared.SearchTerm{
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "aspernatur",
                },
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "ut",
                },
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "deserunt",
                },
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "dignissimos",
                },
            },
            SortFields: []shared.SortByField{
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionDescending,
                    FieldName: "libero",
                    NullPosition: shared.SortByFieldNullPositionNullsLast,
                },
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionDescending,
                    FieldName: "enim",
                    NullPosition: shared.SortByFieldNullPositionNullsFirst,
                },
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionAscending,
                    FieldName: "assumenda",
                    NullPosition: shared.SortByFieldNullPositionNullsLast,
                },
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionDescending,
                    FieldName: "odit",
                    NullPosition: shared.SortByFieldNullPositionNullsLast,
                },
            },
            SuppressEventMarking: false,
            UseFormattedVisualization: false,
        },
        SearchParameterDeclarations: []shared.SearchParameter{
            shared.SearchParameter{
                DefaultValue: cribl.String("iusto"),
                Name: "Lawrence Daniel",
                Type: shared.SearchParameterTypeNumber,
            },
            shared.SearchParameter{
                DefaultValue: cribl.String("architecto"),
                Name: "Leticia Lesch",
                Type: shared.SearchParameterTypeString,
            },
            shared.SearchParameter{
                DefaultValue: cribl.String("eius"),
                Name: "Gayle Zemlak",
                Type: shared.SearchParameterTypeNumber,
            },
            shared.SearchParameter{
                DefaultValue: cribl.String("a"),
                Name: "Nelson Welch",
                Type: shared.SearchParameterTypeString,
            },
        },
        SearchParameterValues: cribl.String("harum"),
        Status: shared.SearchJobStatusQueued,
        TableConfig: &shared.TableViewSettings{
            ColumnFilterSettings: &shared.ColumnFilterSettings{
                Contains: "alias",
            },
            ColumnFormatSettings: &shared.ColumnFormatSettings{
                Palette: "tempore",
                Precision: "quod",
                Prefix: "totam",
                Suffix: "officiis",
            },
            ColumnOrderSettings: &shared.ColumnOrderSettings{
                Order: "dicta",
            },
            ColumnSortSettings: &shared.ColumnSortSettings{
                Sort: "maiores",
            },
            RowNumberColumnWidth: cribl.Int64(175676),
            ShowColumnTotals: false,
            ShowColumnTotalsPinned: false,
            ShowRowNumbers: false,
            ShowRowTotals: false,
            ShowRowTotalsPinned: false,
            ViewMode: shared.TableViewSettingsViewModeEvent,
        },
        TargetEventTime: cribl.Int64(432167),
        TimeCompleted: cribl.Int64(859604),
        TimeCreated: 362259,
        TimeStarted: cribl.Int64(866638),
        Type: shared.SearchJobTypeDatatypePreview.ToPointer(),
        User: "sequi",
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
        ID: "1d008109-0f67-4066-b3f3-a681c5768dce",
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
        ID: "742409a2-15e0-4860-9489-a5f63e3af3dd",
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
                    XAxis: cribl.String("natus"),
                    YAxis: []string{
                        "quibusdam",
                        "officia",
                        "dolorem",
                        "velit",
                    },
                },
                Legend: &shared.Legend{},
                Series: []shared.ChartSeries{
                    shared.ChartSeries{
                        Color: cribl.String("placeat"),
                        Data: []map[string]interface{}{
                            map[string]interface{}{
                                "non": "incidunt",
                                "praesentium": "ipsum",
                            },
                            map[string]interface{}{
                                "non": "dolorum",
                                "esse": "id",
                                "natus": "quas",
                                "saepe": "modi",
                            },
                            map[string]interface{}{
                                "maiores": "neque",
                                "in": "debitis",
                                "quaerat": "nostrum",
                                "libero": "totam",
                            },
                            map[string]interface{}{
                                "veniam": "nostrum",
                                "facere": "aliquam",
                                "vitae": "ipsum",
                            },
                        },
                        DataExpression: "recusandae",
                        DataFilter: map[string]interface{}{
                            "ipsum": "error",
                        },
                        Name: "Terry Dare Jr.",
                        Type: &shared.ChartType{},
                    },
                    shared.ChartSeries{
                        Color: cribl.String("occaecati"),
                        Data: []map[string]interface{}{
                            map[string]interface{}{
                                "cum": "facere",
                                "ratione": "quis",
                            },
                        },
                        DataExpression: "modi",
                        DataFilter: map[string]interface{}{
                            "aut": "cupiditate",
                            "sed": "harum",
                            "vero": "nihil",
                            "velit": "incidunt",
                        },
                        Name: "Michael Cruickshank",
                        Type: &shared.ChartType{},
                    },
                    shared.ChartSeries{
                        Color: cribl.String("excepturi"),
                        Data: []map[string]interface{}{
                            map[string]interface{}{
                                "vel": "delectus",
                                "modi": "expedita",
                                "quidem": "consequuntur",
                            },
                            map[string]interface{}{
                                "asperiores": "debitis",
                            },
                            map[string]interface{}{
                                "quibusdam": "provident",
                                "veritatis": "sunt",
                            },
                            map[string]interface{}{
                                "expedita": "sapiente",
                                "itaque": "dignissimos",
                                "magnam": "excepturi",
                                "placeat": "dolorum",
                            },
                        },
                        DataExpression: "voluptatibus",
                        DataFilter: map[string]interface{}{
                            "enim": "mollitia",
                            "sed": "molestiae",
                        },
                        Name: "Ruben Muller",
                        Type: &shared.ChartType{},
                    },
                    shared.ChartSeries{
                        Color: cribl.String("quod"),
                        Data: []map[string]interface{}{
                            map[string]interface{}{
                                "eum": "temporibus",
                                "quae": "doloremque",
                                "debitis": "omnis",
                                "temporibus": "quidem",
                            },
                            map[string]interface{}{
                                "est": "repellendus",
                            },
                            map[string]interface{}{
                                "quisquam": "vel",
                                "cum": "doloremque",
                            },
                        },
                        DataExpression: "adipisci",
                        DataFilter: map[string]interface{}{
                            "accusantium": "molestias",
                        },
                        Name: "Seth Schaefer",
                        Type: &shared.ChartType{},
                    },
                },
                Settings: &shared.Settings{
                    Color: cribl.String("in"),
                    ColorPalette: 264002,
                    Type: "nihil",
                },
                SingleValue: shared.SingleValue{
                    Color: cribl.String("velit"),
                    Decimals: cribl.Int64(11336),
                    Label: cribl.String("rem"),
                    Prefix: cribl.String("odit"),
                    Suffix: cribl.String("libero"),
                    Type: cribl.String("error"),
                },
                XAxis: &shared.Axis{
                    AxisLabel: &shared.AxisLabel{
                        Formatter: &shared.AxisLabelFormatter{},
                        FormatterData: []int64{
                            992821,
                            181311,
                        },
                    },
                    Type: cribl.String("mollitia"),
                },
            },
            CompatibilityChecks: &shared.SearchJobCompatibilityChecks{
                Datatypes: cribl.Bool(false),
            },
            CorrelationID: cribl.String("distinctio"),
            CPUMetrics: &shared.CPUTimeMetric{
                TotalCPUSeconds: 120216,
                TotalExecCPUSeconds: 937687,
            },
            DatatypeOverrides: &shared.DatatypeOverrides{
                BreakerRulesets: []shared.EventBreakerRuleset{
                    shared.EventBreakerRuleset{
                        Description: cribl.String("nemo"),
                        ID: "671e9c32-6350-4a46-b143-789ce0e99159",
                        Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                        MinRawLength: cribl.Int64(257873),
                        Rules: []shared.EventBreakerRulesetRules{
                            shared.EventBreakerRulesetRules{
                                Condition: "omnis",
                                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                    DstField: cribl.String("amet"),
                                    FieldFilterExpr: cribl.String("deserunt"),
                                    Fields: []string{
                                        "modi",
                                        "impedit",
                                    },
                                    Keep: []string{
                                        "dolores",
                                    },
                                    Remove: []string{
                                        "sed",
                                        "a",
                                    },
                                },
                                Disabled: cribl.Bool(false),
                                Fields: []shared.EventBreakerRulesetRulesFields{
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Yvette Hagenes"),
                                        Value: "quibusdam",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Alberto Rice"),
                                        Value: "quas",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Rudy Reinger"),
                                        Value: "quasi",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Brandon Rutherford"),
                                        Value: "eaque",
                                    },
                                },
                                MaxEventBytes: cribl.Int64(136851),
                                Name: "Neal Welch",
                                ParserEnabled: cribl.Bool(false),
                                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                    Format: cribl.String("eligendi"),
                                    Length: cribl.Int64(722265),
                                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                                },
                                TimestampAnchorRegex: "aliquam",
                                TimestampEarliest: cribl.String("aliquid"),
                                TimestampLatest: cribl.String("adipisci"),
                                TimestampTimezone: cribl.String("ipsam"),
                                Type: shared.EventBreakerRulesetRulesEventBreakerTypeCsv,
                            },
                            shared.EventBreakerRulesetRules{
                                Condition: "enim",
                                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                    DstField: cribl.String("eveniet"),
                                    FieldFilterExpr: cribl.String("eum"),
                                    Fields: []string{
                                        "at",
                                        "culpa",
                                    },
                                    Keep: []string{
                                        "eos",
                                    },
                                    Remove: []string{
                                        "quisquam",
                                        "dolor",
                                        "accusamus",
                                    },
                                },
                                Disabled: cribl.Bool(false),
                                Fields: []shared.EventBreakerRulesetRulesFields{
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Dr. Marie O'Hara"),
                                        Value: "aut",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Laurence Oberbrunner"),
                                        Value: "commodi",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Patsy McClure"),
                                        Value: "nam",
                                    },
                                },
                                MaxEventBytes: cribl.Int64(462268),
                                Name: "Rick Kuhlman",
                                ParserEnabled: cribl.Bool(false),
                                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                    Format: cribl.String("veritatis"),
                                    Length: cribl.Int64(232342),
                                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                                },
                                TimestampAnchorRegex: "quasi",
                                TimestampEarliest: cribl.String("eos"),
                                TimestampLatest: cribl.String("dolorum"),
                                TimestampTimezone: cribl.String("vel"),
                                Type: shared.EventBreakerRulesetRulesEventBreakerTypeTimestamp,
                            },
                            shared.EventBreakerRulesetRules{
                                Condition: "cupiditate",
                                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                    DstField: cribl.String("excepturi"),
                                    FieldFilterExpr: cribl.String("fugit"),
                                    Fields: []string{
                                        "perspiciatis",
                                        "dolore",
                                    },
                                    Keep: []string{
                                        "natus",
                                        "numquam",
                                    },
                                    Remove: []string{
                                        "corrupti",
                                        "ducimus",
                                    },
                                },
                                Disabled: cribl.Bool(false),
                                Fields: []shared.EventBreakerRulesetRulesFields{
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Myra Langosh"),
                                        Value: "blanditiis",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Rosemary Rice"),
                                        Value: "harum",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Adrienne Stokes"),
                                        Value: "dolore",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Mrs. Cathy Quigley"),
                                        Value: "molestias",
                                    },
                                },
                                MaxEventBytes: cribl.Int64(941710),
                                Name: "Mr. Marcos Windler",
                                ParserEnabled: cribl.Bool(false),
                                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                    Format: cribl.String("ut"),
                                    Length: cribl.Int64(887911),
                                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                                },
                                TimestampAnchorRegex: "nulla",
                                TimestampEarliest: cribl.String("distinctio"),
                                TimestampLatest: cribl.String("recusandae"),
                                TimestampTimezone: cribl.String("in"),
                                Type: shared.EventBreakerRulesetRulesEventBreakerTypeHeader,
                            },
                            shared.EventBreakerRulesetRules{
                                Condition: "tempore",
                                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                    DstField: cribl.String("reiciendis"),
                                    FieldFilterExpr: cribl.String("commodi"),
                                    Fields: []string{
                                        "ea",
                                    },
                                    Keep: []string{
                                        "quia",
                                        "ipsam",
                                        "rem",
                                    },
                                    Remove: []string{
                                        "eius",
                                        "necessitatibus",
                                        "culpa",
                                    },
                                },
                                Disabled: cribl.Bool(false),
                                Fields: []shared.EventBreakerRulesetRulesFields{
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Grace Stehr"),
                                        Value: "in",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Allison McLaughlin"),
                                        Value: "quam",
                                    },
                                },
                                MaxEventBytes: cribl.Int64(514441),
                                Name: "Evelyn Gutkowski",
                                ParserEnabled: cribl.Bool(false),
                                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                    Format: cribl.String("eum"),
                                    Length: cribl.Int64(845164),
                                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                                },
                                TimestampAnchorRegex: "ut",
                                TimestampEarliest: cribl.String("cupiditate"),
                                TimestampLatest: cribl.String("debitis"),
                                TimestampTimezone: cribl.String("minima"),
                                Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                            },
                        },
                        Tags: cribl.String("consectetur"),
                    },
                    shared.EventBreakerRuleset{
                        Description: cribl.String("nostrum"),
                        ID: "b33bc0f9-70c4-42fc-9f48-44225e75b796",
                        Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                        MinRawLength: cribl.Int64(41800),
                        Rules: []shared.EventBreakerRulesetRules{
                            shared.EventBreakerRulesetRules{
                                Condition: "nostrum",
                                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                    DstField: cribl.String("cumque"),
                                    FieldFilterExpr: cribl.String("eaque"),
                                    Fields: []string{
                                        "earum",
                                        "culpa",
                                        "vel",
                                        "sapiente",
                                    },
                                    Keep: []string{
                                        "ratione",
                                        "libero",
                                        "molestias",
                                    },
                                    Remove: []string{
                                        "fuga",
                                    },
                                },
                                Disabled: cribl.Bool(false),
                                Fields: []shared.EventBreakerRulesetRulesFields{
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Alfredo Roob"),
                                        Value: "nam",
                                    },
                                },
                                MaxEventBytes: cribl.Int64(882195),
                                Name: "Janice Hermiston",
                                ParserEnabled: cribl.Bool(false),
                                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                    Format: cribl.String("nihil"),
                                    Length: cribl.Int64(239185),
                                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                                },
                                TimestampAnchorRegex: "reiciendis",
                                TimestampEarliest: cribl.String("dolore"),
                                TimestampLatest: cribl.String("voluptatibus"),
                                TimestampTimezone: cribl.String("eveniet"),
                                Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                            },
                            shared.EventBreakerRulesetRules{
                                Condition: "voluptate",
                                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                    DstField: cribl.String("explicabo"),
                                    FieldFilterExpr: cribl.String("architecto"),
                                    Fields: []string{
                                        "possimus",
                                    },
                                    Keep: []string{
                                        "delectus",
                                    },
                                    Remove: []string{
                                        "exercitationem",
                                        "quis",
                                    },
                                },
                                Disabled: cribl.Bool(false),
                                Fields: []shared.EventBreakerRulesetRulesFields{
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Seth Marks"),
                                        Value: "sed",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Mona Considine"),
                                        Value: "accusantium",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Ms. Andy Gerhold"),
                                        Value: "possimus",
                                    },
                                },
                                MaxEventBytes: cribl.Int64(599413),
                                Name: "Trevor Nolan II",
                                ParserEnabled: cribl.Bool(false),
                                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                    Format: cribl.String("quia"),
                                    Length: cribl.Int64(844566),
                                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                                },
                                TimestampAnchorRegex: "voluptate",
                                TimestampEarliest: cribl.String("placeat"),
                                TimestampLatest: cribl.String("est"),
                                TimestampTimezone: cribl.String("est"),
                                Type: shared.EventBreakerRulesetRulesEventBreakerTypeTimestamp,
                            },
                        },
                        Tags: cribl.String("occaecati"),
                    },
                    shared.EventBreakerRuleset{
                        Description: cribl.String("nam"),
                        ID: "4caa1cfe-9e15-4df9-8390-7f37831983d4",
                        Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                        MinRawLength: cribl.Int64(140972),
                        Rules: []shared.EventBreakerRulesetRules{
                            shared.EventBreakerRulesetRules{
                                Condition: "minima",
                                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                    DstField: cribl.String("aliquam"),
                                    FieldFilterExpr: cribl.String("similique"),
                                    Fields: []string{
                                        "corporis",
                                        "dolore",
                                        "commodi",
                                    },
                                    Keep: []string{
                                        "quis",
                                        "provident",
                                    },
                                    Remove: []string{
                                        "optio",
                                        "ipsam",
                                    },
                                },
                                Disabled: cribl.Bool(false),
                                Fields: []shared.EventBreakerRulesetRulesFields{
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Florence Dooley I"),
                                        Value: "nihil",
                                    },
                                },
                                MaxEventBytes: cribl.Int64(112399),
                                Name: "Floyd Borer",
                                ParserEnabled: cribl.Bool(false),
                                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                    Format: cribl.String("mollitia"),
                                    Length: cribl.Int64(420223),
                                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                                },
                                TimestampAnchorRegex: "at",
                                TimestampEarliest: cribl.String("hic"),
                                TimestampLatest: cribl.String("nemo"),
                                TimestampTimezone: cribl.String("laborum"),
                                Type: shared.EventBreakerRulesetRulesEventBreakerTypeTimestamp,
                            },
                            shared.EventBreakerRulesetRules{
                                Condition: "nulla",
                                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                    DstField: cribl.String("aliquid"),
                                    FieldFilterExpr: cribl.String("eius"),
                                    Fields: []string{
                                        "reprehenderit",
                                        "quo",
                                        "enim",
                                    },
                                    Keep: []string{
                                        "impedit",
                                        "aspernatur",
                                        "nobis",
                                        "voluptatum",
                                    },
                                    Remove: []string{
                                        "aspernatur",
                                        "est",
                                    },
                                },
                                Disabled: cribl.Bool(false),
                                Fields: []shared.EventBreakerRulesetRulesFields{
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Shelia Ullrich"),
                                        Value: "perspiciatis",
                                    },
                                },
                                MaxEventBytes: cribl.Int64(906293),
                                Name: "Mrs. Sarah Berge",
                                ParserEnabled: cribl.Bool(false),
                                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                    Format: cribl.String("suscipit"),
                                    Length: cribl.Int64(222109),
                                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                                },
                                TimestampAnchorRegex: "libero",
                                TimestampEarliest: cribl.String("quibusdam"),
                                TimestampLatest: cribl.String("fuga"),
                                TimestampTimezone: cribl.String("nihil"),
                                Type: shared.EventBreakerRulesetRulesEventBreakerTypeHeader,
                            },
                            shared.EventBreakerRulesetRules{
                                Condition: "repellat",
                                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                    DstField: cribl.String("quibusdam"),
                                    FieldFilterExpr: cribl.String("accusamus"),
                                    Fields: []string{
                                        "blanditiis",
                                        "tempora",
                                        "similique",
                                        "dolor",
                                    },
                                    Keep: []string{
                                        "dolorum",
                                        "aliquam",
                                    },
                                    Remove: []string{
                                        "explicabo",
                                    },
                                },
                                Disabled: cribl.Bool(false),
                                Fields: []shared.EventBreakerRulesetRulesFields{
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Stewart Bode"),
                                        Value: "sequi",
                                    },
                                },
                                MaxEventBytes: cribl.Int64(321724),
                                Name: "Sylvester Dibbert",
                                ParserEnabled: cribl.Bool(false),
                                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                    Format: cribl.String("saepe"),
                                    Length: cribl.Int64(216972),
                                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                                },
                                TimestampAnchorRegex: "expedita",
                                TimestampEarliest: cribl.String("itaque"),
                                TimestampLatest: cribl.String("maiores"),
                                TimestampTimezone: cribl.String("provident"),
                                Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                            },
                            shared.EventBreakerRulesetRules{
                                Condition: "dicta",
                                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                    DstField: cribl.String("id"),
                                    FieldFilterExpr: cribl.String("blanditiis"),
                                    Fields: []string{
                                        "modi",
                                        "ex",
                                        "nobis",
                                        "placeat",
                                    },
                                    Keep: []string{
                                        "architecto",
                                        "quae",
                                        "aut",
                                    },
                                    Remove: []string{
                                        "delectus",
                                        "officiis",
                                    },
                                },
                                Disabled: cribl.Bool(false),
                                Fields: []shared.EventBreakerRulesetRulesFields{
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Roberta Roberts Sr."),
                                        Value: "pariatur",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Mattie Roob"),
                                        Value: "deleniti",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Rodolfo Moore"),
                                        Value: "libero",
                                    },
                                },
                                MaxEventBytes: cribl.Int64(596393),
                                Name: "Homer Hilll V",
                                ParserEnabled: cribl.Bool(false),
                                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                    Format: cribl.String("eum"),
                                    Length: cribl.Int64(339333),
                                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                                },
                                TimestampAnchorRegex: "eveniet",
                                TimestampEarliest: cribl.String("possimus"),
                                TimestampLatest: cribl.String("dolor"),
                                TimestampTimezone: cribl.String("ratione"),
                                Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSON,
                            },
                        },
                        Tags: cribl.String("soluta"),
                    },
                    shared.EventBreakerRuleset{
                        Description: cribl.String("cum"),
                        ID: "0ce8aa65-432a-4986-ab7e-14ca56408915",
                        Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                        MinRawLength: cribl.Int64(53057),
                        Rules: []shared.EventBreakerRulesetRules{
                            shared.EventBreakerRulesetRules{
                                Condition: "natus",
                                Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                    DstField: cribl.String("nihil"),
                                    FieldFilterExpr: cribl.String("accusantium"),
                                    Fields: []string{
                                        "unde",
                                    },
                                    Keep: []string{
                                        "eius",
                                        "rem",
                                        "delectus",
                                    },
                                    Remove: []string{
                                        "laudantium",
                                        "earum",
                                        "impedit",
                                    },
                                },
                                Disabled: cribl.Bool(false),
                                Fields: []shared.EventBreakerRulesetRulesFields{
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Lula Weimann I"),
                                        Value: "saepe",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Mrs. Alice Carter"),
                                        Value: "sequi",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Salvador Abshire III"),
                                        Value: "aspernatur",
                                    },
                                    shared.EventBreakerRulesetRulesFields{
                                        Name: cribl.String("Shane Renner"),
                                        Value: "nisi",
                                    },
                                },
                                MaxEventBytes: cribl.Int64(502675),
                                Name: "Anthony Yost",
                                ParserEnabled: cribl.Bool(false),
                                Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                    Format: cribl.String("ducimus"),
                                    Length: cribl.Int64(692401),
                                    Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                                },
                                TimestampAnchorRegex: "at",
                                TimestampEarliest: cribl.String("alias"),
                                TimestampLatest: cribl.String("consectetur"),
                                TimestampTimezone: cribl.String("mollitia"),
                                Type: shared.EventBreakerRulesetRulesEventBreakerTypeRegex,
                            },
                        },
                        Tags: cribl.String("numquam"),
                    },
                },
                DisableBreakers: false,
            },
            EarliestEpoch: cribl.Int64(508520),
            ErrorStateConfig: &shared.SearchJobErrorStateConfig{
                Coordinated: false,
                ErrorMessages: []string{
                    "hic",
                },
            },
            Group: "blanditiis",
            ID: "de30f069-d810-4618-997e-152297510da8",
            LatestEpoch: cribl.Int64(32352),
            NumEventsAfter: cribl.Int64(228977),
            NumEventsBefore: cribl.Int64(110280),
            Query: "explicabo",
            SampleRate: cribl.Int64(158768),
            SearchConfig: shared.SearchConfig{
                Datasets: []string{
                    "fugit",
                    "nobis",
                    "optio",
                },
                HasSendOperator: false,
                OrderedFieldNames: []string{
                    "quasi",
                    "porro",
                },
                SearchTerms: []shared.SearchTerm{
                    shared.SearchTerm{
                        IsCaseSensitive: false,
                        Term: "laborum",
                    },
                },
                SortFields: []shared.SortByField{
                    shared.SortByField{
                        Direction: shared.SortByFieldDirectionAscending,
                        FieldName: "odit",
                        NullPosition: shared.SortByFieldNullPositionNullsLast,
                    },
                    shared.SortByField{
                        Direction: shared.SortByFieldDirectionDescending,
                        FieldName: "sint",
                        NullPosition: shared.SortByFieldNullPositionNullsFirst,
                    },
                },
                SuppressEventMarking: false,
                UseFormattedVisualization: false,
            },
            SearchParameterDeclarations: []shared.SearchParameter{
                shared.SearchParameter{
                    DefaultValue: cribl.String("accusamus"),
                    Name: "Sharon Daniel",
                    Type: shared.SearchParameterTypeString,
                },
                shared.SearchParameter{
                    DefaultValue: cribl.String("fugiat"),
                    Name: "Mike Hansen",
                    Type: shared.SearchParameterTypeBoolean,
                },
                shared.SearchParameter{
                    DefaultValue: cribl.String("sit"),
                    Name: "Pam West",
                    Type: shared.SearchParameterTypeBoolean,
                },
                shared.SearchParameter{
                    DefaultValue: cribl.String("similique"),
                    Name: "Willis Graham",
                    Type: shared.SearchParameterTypeString,
                },
            },
            SearchParameterValues: cribl.String("dolores"),
            Status: shared.SearchJobStatusCanceled,
            TableConfig: &shared.TableViewSettings{
                ColumnFilterSettings: &shared.ColumnFilterSettings{
                    Contains: "quod",
                },
                ColumnFormatSettings: &shared.ColumnFormatSettings{
                    Palette: "sunt",
                    Precision: "iure",
                    Prefix: "voluptatem",
                    Suffix: "incidunt",
                },
                ColumnOrderSettings: &shared.ColumnOrderSettings{
                    Order: "soluta",
                },
                ColumnSortSettings: &shared.ColumnSortSettings{
                    Sort: "a",
                },
                RowNumberColumnWidth: cribl.Int64(110823),
                ShowColumnTotals: false,
                ShowColumnTotalsPinned: false,
                ShowRowNumbers: false,
                ShowRowTotals: false,
                ShowRowTotalsPinned: false,
                ViewMode: shared.TableViewSettingsViewModeTable,
            },
            TargetEventTime: cribl.Int64(811167),
            TimeCompleted: cribl.Int64(586303),
            TimeCreated: 989079,
            TimeStarted: cribl.Int64(802449),
            Type: shared.SearchJobTypeDatatypePreview.ToPointer(),
            User: "dicta",
        },
        ID: "aae5eb5f-0c49-42b5-b44d-08a2267aaee7",
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

