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
    res, err := s.SearchJob.Create(ctx, shared.SearchJob{
        ChartConfig: &shared.ChartConfig{
            Axis: &shared.ChartConfigAxis{
                XAxis: cribl.String("enim"),
                YAxis: []string{
                    "unde",
                    "quae",
                    "eum",
                },
            },
            Legend: &shared.Legend{},
            Series: []shared.ChartSeries{
                shared.ChartSeries{
                    Color: cribl.String("eveniet"),
                    Data: []map[string]interface{}{
                        map[string]interface{}{
                            "blanditiis": "quidem",
                        },
                        map[string]interface{}{
                            "reiciendis": "placeat",
                            "dolores": "consequatur",
                            "nesciunt": "quia",
                            "quidem": "voluptas",
                        },
                    },
                    DataExpression: "quo",
                    DataFilter: map[string]interface{}{
                        "dignissimos": "omnis",
                        "omnis": "fugit",
                        "dolorem": "quidem",
                    },
                    Name: "Raquel Bruen",
                    Type: &shared.ChartType{},
                },
                shared.ChartSeries{
                    Color: cribl.String("atque"),
                    Data: []map[string]interface{}{
                        map[string]interface{}{
                            "reprehenderit": "deserunt",
                            "itaque": "et",
                            "eos": "impedit",
                            "ex": "praesentium",
                        },
                        map[string]interface{}{
                            "vitae": "tenetur",
                            "laudantium": "aspernatur",
                            "eligendi": "repudiandae",
                        },
                    },
                    DataExpression: "dicta",
                    DataFilter: map[string]interface{}{
                        "ullam": "iusto",
                    },
                    Name: "Allison Corwin II",
                    Type: &shared.ChartType{},
                },
            },
            Settings: &shared.Settings{
                Color: cribl.String("amet"),
                ColorPalette: 454165,
                Type: "voluptate",
            },
            SingleValue: shared.SingleValue{
                Color: cribl.String("pariatur"),
                Decimals: cribl.Int64(791421),
                Label: cribl.String("a"),
                Prefix: cribl.String("fuga"),
                Suffix: cribl.String("totam"),
                Type: cribl.String("cupiditate"),
            },
            XAxis: &shared.Axis{
                AxisLabel: &shared.AxisLabel{
                    Formatter: &shared.AxisLabelFormatter{},
                    FormatterData: []int64{
                        986594,
                        608172,
                        463695,
                        349087,
                    },
                },
                Type: cribl.String("voluptates"),
            },
        },
        CompatibilityChecks: &shared.SearchJobCompatibilityChecks{
            Datatypes: cribl.Bool(false),
        },
        CorrelationID: cribl.String("sequi"),
        CPUMetrics: &shared.CPUTimeMetric{
            TotalCPUSeconds: 336640,
            TotalExecCPUSeconds: 413135,
        },
        DatatypeOverrides: &shared.DatatypeOverrides{
            BreakerRulesets: []shared.EventBreakerRuleset{
                shared.EventBreakerRuleset{
                    Description: cribl.String("rem"),
                    ID: "6092e9c3-ddc5-4f11-9dea-1026d541a4d1",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(607181),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "sapiente",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("officiis"),
                                FieldFilterExpr: cribl.String("expedita"),
                                Fields: []string{
                                    "vitae",
                                },
                                Keep: []string{
                                    "quas",
                                    "ipsa",
                                },
                                Remove: []string{
                                    "placeat",
                                    "quod",
                                    "eligendi",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Preston Rau"),
                                    Value: "soluta",
                                },
                            },
                            MaxEventBytes: cribl.Int64(297325),
                            Name: "Ms. Ricky Koch",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("facilis"),
                                Length: cribl.Int64(306023),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                            },
                            TimestampAnchorRegex: "amet",
                            TimestampEarliest: cribl.String("natus"),
                            TimestampLatest: cribl.String("ab"),
                            TimestampTimezone: cribl.String("officiis"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                        },
                    },
                    Tags: cribl.String("rerum"),
                },
                shared.EventBreakerRuleset{
                    Description: cribl.String("placeat"),
                    ID: "158c4c4e-5459-49ea-b422-60e9b200ce78",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(666321),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "expedita",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("quibusdam"),
                                FieldFilterExpr: cribl.String("quos"),
                                Fields: []string{
                                    "quidem",
                                    "in",
                                    "culpa",
                                    "doloremque",
                                },
                                Keep: []string{
                                    "dicta",
                                    "architecto",
                                    "suscipit",
                                },
                                Remove: []string{
                                    "officiis",
                                    "dignissimos",
                                    "fugit",
                                    "ratione",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Susan Mraz"),
                                    Value: "deserunt",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Brenda Tremblay"),
                                    Value: "hic",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Andrea Hilpert"),
                                    Value: "omnis",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Mrs. Phyllis Christiansen PhD"),
                                    Value: "voluptatum",
                                },
                            },
                            MaxEventBytes: cribl.Int64(245116),
                            Name: "Dan Nolan",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("iusto"),
                                Length: cribl.Int64(493865),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "occaecati",
                            TimestampEarliest: cribl.String("assumenda"),
                            TimestampLatest: cribl.String("sunt"),
                            TimestampTimezone: cribl.String("odit"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeCsv,
                        },
                    },
                    Tags: cribl.String("deleniti"),
                },
            },
            DisableBreakers: false,
        },
        EarliestEpoch: cribl.Int64(760793),
        ErrorStateConfig: &shared.SearchJobErrorStateConfig{
            Coordinated: false,
            ErrorMessages: []string{
                "repellat",
            },
        },
        Group: "atque",
        ID: "493825fd-c42c-4876-82c2-dfb4cfc1c762",
        LatestEpoch: cribl.Int64(197441),
        NumEventsAfter: cribl.Int64(4787),
        NumEventsBefore: cribl.Int64(970131),
        Query: "quos",
        SampleRate: cribl.Int64(256310),
        SearchConfig: shared.SearchConfig{
            Datasets: []string{
                "maiores",
            },
            HasSendOperator: false,
            OrderedFieldNames: []string{
                "architecto",
                "rerum",
                "assumenda",
            },
            SearchTerms: []shared.SearchTerm{
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "dolorem",
                },
            },
            SortFields: []shared.SortByField{
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionDescending,
                    FieldName: "nam",
                    NullPosition: shared.SortByFieldNullPositionNullsFirst,
                },
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionAscending,
                    FieldName: "pariatur",
                    NullPosition: shared.SortByFieldNullPositionNullsLast,
                },
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionAscending,
                    FieldName: "tempore",
                    NullPosition: shared.SortByFieldNullPositionNullsLast,
                },
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionAscending,
                    FieldName: "officia",
                    NullPosition: shared.SortByFieldNullPositionNullsFirst,
                },
            },
            SuppressEventMarking: false,
            UseFormattedVisualization: false,
        },
        SearchParameterDeclarations: []shared.SearchParameter{
            shared.SearchParameter{
                DefaultValue: cribl.String("corporis"),
                Name: "Luther Leuschke",
                Type: shared.SearchParameterTypeString,
            },
            shared.SearchParameter{
                DefaultValue: cribl.String("laborum"),
                Name: "Craig Barrows",
                Type: shared.SearchParameterTypeString,
            },
            shared.SearchParameter{
                DefaultValue: cribl.String("commodi"),
                Name: "Sylvester Cormier",
                Type: shared.SearchParameterTypeNumber,
            },
        },
        SearchParameterValues: cribl.String("ab"),
        Status: shared.SearchJobStatusFailed,
        TableConfig: &shared.TableViewSettings{
            ColumnFilterSettings: &shared.ColumnFilterSettings{
                Contains: "sed",
            },
            ColumnFormatSettings: &shared.ColumnFormatSettings{
                Palette: "blanditiis",
                Precision: "sint",
                Prefix: "placeat",
                Suffix: "ullam",
            },
            ColumnOrderSettings: &shared.ColumnOrderSettings{
                Order: "molestiae",
            },
            ColumnSortSettings: &shared.ColumnSortSettings{
                Sort: "itaque",
            },
            RowNumberColumnWidth: cribl.Int64(523156),
            ShowColumnTotals: false,
            ShowColumnTotalsPinned: false,
            ShowRowNumbers: false,
            ShowRowTotals: false,
            ShowRowTotalsPinned: false,
            ViewMode: shared.TableViewSettingsViewModeEvent,
        },
        TargetEventTime: cribl.Int64(250520),
        TimeCompleted: cribl.Int64(926142),
        TimeCreated: 605043,
        TimeStarted: cribl.Int64(58567),
        Type: shared.SearchJobTypeStandard.ToPointer(),
        User: "dolor",
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
    id := "occaecati"

    ctx := context.Background()
    res, err := s.SearchJob.Delete(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


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
    res, err := s.SearchJob.Get(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


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
    id := "magni"
    searchJob := &shared.SearchJob{
        ChartConfig: &shared.ChartConfig{
            Axis: &shared.ChartConfigAxis{
                XAxis: cribl.String("consequuntur"),
                YAxis: []string{
                    "eius",
                },
            },
            Legend: &shared.Legend{},
            Series: []shared.ChartSeries{
                shared.ChartSeries{
                    Color: cribl.String("ipsam"),
                    Data: []map[string]interface{}{
                        map[string]interface{}{
                            "modi": "nisi",
                            "explicabo": "modi",
                            "doloremque": "odio",
                        },
                        map[string]interface{}{
                            "voluptatum": "tempora",
                        },
                    },
                    DataExpression: "delectus",
                    DataFilter: map[string]interface{}{
                        "dolorum": "libero",
                        "ratione": "molestiae",
                    },
                    Name: "Mr. Bert Wunsch",
                    Type: &shared.ChartType{},
                },
                shared.ChartSeries{
                    Color: cribl.String("consequuntur"),
                    Data: []map[string]interface{}{
                        map[string]interface{}{
                            "occaecati": "labore",
                        },
                        map[string]interface{}{
                            "quidem": "exercitationem",
                            "veniam": "modi",
                            "quasi": "quae",
                            "similique": "possimus",
                        },
                    },
                    DataExpression: "quo",
                    DataFilter: map[string]interface{}{
                        "ex": "sint",
                        "est": "doloribus",
                    },
                    Name: "James Nolan",
                    Type: &shared.ChartType{},
                },
            },
            Settings: &shared.Settings{
                Color: cribl.String("quo"),
                ColorPalette: 477099,
                Type: "maxime",
            },
            SingleValue: shared.SingleValue{
                Color: cribl.String("facere"),
                Decimals: cribl.Int64(769247),
                Label: cribl.String("cupiditate"),
                Prefix: cribl.String("deleniti"),
                Suffix: cribl.String("quasi"),
                Type: cribl.String("maiores"),
            },
            XAxis: &shared.Axis{
                AxisLabel: &shared.AxisLabel{
                    Formatter: &shared.AxisLabelFormatter{},
                    FormatterData: []int64{
                        396184,
                    },
                },
                Type: cribl.String("laudantium"),
            },
        },
        CompatibilityChecks: &shared.SearchJobCompatibilityChecks{
            Datatypes: cribl.Bool(false),
        },
        CorrelationID: cribl.String("unde"),
        CPUMetrics: &shared.CPUTimeMetric{
            TotalCPUSeconds: 544660,
            TotalExecCPUSeconds: 63369,
        },
        DatatypeOverrides: &shared.DatatypeOverrides{
            BreakerRulesets: []shared.EventBreakerRuleset{
                shared.EventBreakerRuleset{
                    Description: cribl.String("ea"),
                    ID: "bb33cfaa-348c-431b-b407-ee4fcf0c42b7",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(532336),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "vitae",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("ullam"),
                                FieldFilterExpr: cribl.String("nisi"),
                                Fields: []string{
                                    "voluptas",
                                },
                                Keep: []string{
                                    "excepturi",
                                },
                                Remove: []string{
                                    "est",
                                    "perferendis",
                                    "quibusdam",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Sue Huel"),
                                    Value: "dolore",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Ms. Wilbert Ratke"),
                                    Value: "totam",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Rex Bode"),
                                    Value: "sit",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Jill Collins"),
                                    Value: "odio",
                                },
                            },
                            MaxEventBytes: cribl.Int64(2758),
                            Name: "Jaime Schuppe",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("fugit"),
                                Length: cribl.Int64(134267),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                            },
                            TimestampAnchorRegex: "illum",
                            TimestampEarliest: cribl.String("praesentium"),
                            TimestampLatest: cribl.String("sint"),
                            TimestampTimezone: cribl.String("exercitationem"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeTimestamp,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "voluptatum",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("facilis"),
                                FieldFilterExpr: cribl.String("placeat"),
                                Fields: []string{
                                    "dolores",
                                    "dolore",
                                    "pariatur",
                                    "facilis",
                                },
                                Keep: []string{
                                    "nemo",
                                    "natus",
                                    "nisi",
                                },
                                Remove: []string{
                                    "amet",
                                    "dolor",
                                    "nostrum",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Clinton Greenholt"),
                                    Value: "adipisci",
                                },
                            },
                            MaxEventBytes: cribl.Int64(583138),
                            Name: "Herbert Steuber",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("repellendus"),
                                Length: cribl.Int64(906524),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                            },
                            TimestampAnchorRegex: "distinctio",
                            TimestampEarliest: cribl.String("vel"),
                            TimestampLatest: cribl.String("necessitatibus"),
                            TimestampTimezone: cribl.String("iste"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSON,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "corrupti",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("cupiditate"),
                                FieldFilterExpr: cribl.String("voluptatibus"),
                                Fields: []string{
                                    "dolorum",
                                    "soluta",
                                },
                                Keep: []string{
                                    "in",
                                    "delectus",
                                    "commodi",
                                },
                                Remove: []string{
                                    "fugit",
                                    "ullam",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Monique Denesik"),
                                    Value: "totam",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Iris Roob"),
                                    Value: "dolor",
                                },
                            },
                            MaxEventBytes: cribl.Int64(667634),
                            Name: "Rufus Conn II",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("facilis"),
                                Length: cribl.Int64(708360),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "nisi",
                            TimestampEarliest: cribl.String("ipsam"),
                            TimestampLatest: cribl.String("voluptatem"),
                            TimestampTimezone: cribl.String("illo"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "incidunt",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("eveniet"),
                                FieldFilterExpr: cribl.String("quae"),
                                Fields: []string{
                                    "asperiores",
                                    "veniam",
                                },
                                Keep: []string{
                                    "asperiores",
                                    "eum",
                                    "deserunt",
                                },
                                Remove: []string{
                                    "nemo",
                                    "molestias",
                                    "architecto",
                                    "expedita",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Malcolm Shanahan"),
                                    Value: "ipsum",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Jean Corkery"),
                                    Value: "neque",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Jesse Abbott"),
                                    Value: "facere",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Ms. Alberto Hagenes"),
                                    Value: "modi",
                                },
                            },
                            MaxEventBytes: cribl.Int64(683980),
                            Name: "Leah Nienow",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("dolor"),
                                Length: cribl.Int64(322773),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                            },
                            TimestampAnchorRegex: "sit",
                            TimestampEarliest: cribl.String("molestias"),
                            TimestampLatest: cribl.String("voluptas"),
                            TimestampTimezone: cribl.String("expedita"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                        },
                    },
                    Tags: cribl.String("maiores"),
                },
                shared.EventBreakerRuleset{
                    Description: cribl.String("ea"),
                    ID: "6fef020e-9f44-43b4-a57b-992c8dbda6a6",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(120432),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "sapiente",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("id"),
                                FieldFilterExpr: cribl.String("odit"),
                                Fields: []string{
                                    "iste",
                                },
                                Keep: []string{
                                    "explicabo",
                                    "ullam",
                                    "atque",
                                },
                                Remove: []string{
                                    "pariatur",
                                    "aut",
                                    "similique",
                                    "iste",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Lynn Green"),
                                    Value: "quam",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Stanley Von II"),
                                    Value: "excepturi",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Eleanor Bartell"),
                                    Value: "culpa",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Miss Leah Dicki"),
                                    Value: "nihil",
                                },
                            },
                            MaxEventBytes: cribl.Int64(631904),
                            Name: "Tommie Hand",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("reiciendis"),
                                Length: cribl.Int64(832135),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                            },
                            TimestampAnchorRegex: "ab",
                            TimestampEarliest: cribl.String("mollitia"),
                            TimestampLatest: cribl.String("possimus"),
                            TimestampTimezone: cribl.String("praesentium"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSON,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "quam",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("animi"),
                                FieldFilterExpr: cribl.String("debitis"),
                                Fields: []string{
                                    "voluptatem",
                                    "quisquam",
                                    "vitae",
                                },
                                Keep: []string{
                                    "architecto",
                                    "sint",
                                    "eligendi",
                                    "occaecati",
                                },
                                Remove: []string{
                                    "tempore",
                                    "officia",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Willard Howe"),
                                    Value: "sapiente",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Rodney Weissnat"),
                                    Value: "autem",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Terrance Bernhard"),
                                    Value: "dolor",
                                },
                            },
                            MaxEventBytes: cribl.Int64(528082),
                            Name: "Mrs. Forrest Waelchi",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("vitae"),
                                Length: cribl.Int64(256975),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                            },
                            TimestampAnchorRegex: "modi",
                            TimestampEarliest: cribl.String("quos"),
                            TimestampLatest: cribl.String("minus"),
                            TimestampTimezone: cribl.String("voluptate"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeHeader,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "reprehenderit",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("reprehenderit"),
                                FieldFilterExpr: cribl.String("animi"),
                                Fields: []string{
                                    "voluptates",
                                },
                                Keep: []string{
                                    "qui",
                                    "delectus",
                                    "exercitationem",
                                    "ipsum",
                                },
                                Remove: []string{
                                    "doloremque",
                                    "sed",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Darnell Watsica"),
                                    Value: "cupiditate",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Alicia Bernier"),
                                    Value: "necessitatibus",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Cody Ward"),
                                    Value: "non",
                                },
                            },
                            MaxEventBytes: cribl.Int64(962497),
                            Name: "Sadie Casper",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("illum"),
                                Length: cribl.Int64(916051),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "est",
                            TimestampEarliest: cribl.String("in"),
                            TimestampLatest: cribl.String("illo"),
                            TimestampTimezone: cribl.String("voluptate"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeRegex,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "delectus",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("incidunt"),
                                FieldFilterExpr: cribl.String("dolore"),
                                Fields: []string{
                                    "est",
                                    "quo",
                                },
                                Keep: []string{
                                    "delectus",
                                    "laboriosam",
                                    "laboriosam",
                                    "quam",
                                },
                                Remove: []string{
                                    "officia",
                                    "repellat",
                                    "cupiditate",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Ms. Oliver Steuber"),
                                    Value: "sapiente",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Dr. Alexander Douglas"),
                                    Value: "quidem",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Maurice Runolfsdottir"),
                                    Value: "atque",
                                },
                            },
                            MaxEventBytes: cribl.Int64(979706),
                            Name: "Andres Larson",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("eaque"),
                                Length: cribl.Int64(773455),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                            },
                            TimestampAnchorRegex: "ex",
                            TimestampEarliest: cribl.String("odio"),
                            TimestampLatest: cribl.String("delectus"),
                            TimestampTimezone: cribl.String("minus"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSON,
                        },
                    },
                    Tags: cribl.String("distinctio"),
                },
                shared.EventBreakerRuleset{
                    Description: cribl.String("eius"),
                    ID: "25e99e62-34c9-4f7b-b9df-eb77a5c38d4b",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(638085),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "cupiditate",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("illo"),
                                FieldFilterExpr: cribl.String("saepe"),
                                Fields: []string{
                                    "eaque",
                                    "ex",
                                },
                                Keep: []string{
                                    "delectus",
                                    "deleniti",
                                    "provident",
                                    "aut",
                                },
                                Remove: []string{
                                    "nostrum",
                                    "tempora",
                                    "nam",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Ms. Bernice Wuckert"),
                                    Value: "minima",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Marcella Dooley"),
                                    Value: "fuga",
                                },
                            },
                            MaxEventBytes: cribl.Int64(195789),
                            Name: "Bradley Osinski",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("adipisci"),
                                Length: cribl.Int64(124796),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                            },
                            TimestampAnchorRegex: "occaecati",
                            TimestampEarliest: cribl.String("provident"),
                            TimestampLatest: cribl.String("necessitatibus"),
                            TimestampTimezone: cribl.String("fugit"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "optio",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("eveniet"),
                                FieldFilterExpr: cribl.String("fugiat"),
                                Fields: []string{
                                    "a",
                                    "natus",
                                    "sapiente",
                                },
                                Keep: []string{
                                    "facilis",
                                    "molestias",
                                    "dolore",
                                    "et",
                                },
                                Remove: []string{
                                    "maiores",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Hannah Rath"),
                                    Value: "vitae",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Kay Frami"),
                                    Value: "aperiam",
                                },
                            },
                            MaxEventBytes: cribl.Int64(112788),
                            Name: "Emmett Swaniawski",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("deleniti"),
                                Length: cribl.Int64(523266),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "eos",
                            TimestampEarliest: cribl.String("labore"),
                            TimestampLatest: cribl.String("sunt"),
                            TimestampTimezone: cribl.String("blanditiis"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeHeader,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "accusamus",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("distinctio"),
                                FieldFilterExpr: cribl.String("incidunt"),
                                Fields: []string{
                                    "blanditiis",
                                    "ducimus",
                                },
                                Keep: []string{
                                    "sapiente",
                                },
                                Remove: []string{
                                    "accusantium",
                                    "ratione",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Juan Marquardt"),
                                    Value: "asperiores",
                                },
                            },
                            MaxEventBytes: cribl.Int64(98250),
                            Name: "Vivian Rolfson",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("veritatis"),
                                Length: cribl.Int64(328744),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                            },
                            TimestampAnchorRegex: "itaque",
                            TimestampEarliest: cribl.String("error"),
                            TimestampLatest: cribl.String("expedita"),
                            TimestampTimezone: cribl.String("error"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeTimestamp,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "temporibus",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("voluptate"),
                                FieldFilterExpr: cribl.String("earum"),
                                Fields: []string{
                                    "odit",
                                    "odit",
                                },
                                Keep: []string{
                                    "error",
                                    "vel",
                                },
                                Remove: []string{
                                    "alias",
                                    "itaque",
                                    "ab",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Bridget Kshlerin"),
                                    Value: "esse",
                                },
                            },
                            MaxEventBytes: cribl.Int64(876682),
                            Name: "Derek Mohr DVM",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("autem"),
                                Length: cribl.Int64(449861),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                            },
                            TimestampAnchorRegex: "amet",
                            TimestampEarliest: cribl.String("cumque"),
                            TimestampLatest: cribl.String("dolore"),
                            TimestampTimezone: cribl.String("optio"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeTimestamp,
                        },
                    },
                    Tags: cribl.String("repudiandae"),
                },
                shared.EventBreakerRuleset{
                    Description: cribl.String("tempora"),
                    ID: "b6d7696f-f3c5-4747-9013-57e44f51f8b0",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(516739),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "quo",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("dolor"),
                                FieldFilterExpr: cribl.String("sunt"),
                                Fields: []string{
                                    "quam",
                                    "officiis",
                                    "dicta",
                                },
                                Keep: []string{
                                    "consectetur",
                                    "deserunt",
                                    "odit",
                                },
                                Remove: []string{
                                    "corporis",
                                    "quaerat",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Essie Mante"),
                                    Value: "nihil",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Rachael Corkery"),
                                    Value: "nobis",
                                },
                            },
                            MaxEventBytes: cribl.Int64(784120),
                            Name: "Katrina Kunze",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("consectetur"),
                                Length: cribl.Int64(188399),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                            },
                            TimestampAnchorRegex: "iure",
                            TimestampEarliest: cribl.String("aliquid"),
                            TimestampLatest: cribl.String("cum"),
                            TimestampTimezone: cribl.String("fugiat"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeHeader,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "voluptatibus",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("officiis"),
                                FieldFilterExpr: cribl.String("corporis"),
                                Fields: []string{
                                    "aut",
                                    "voluptatem",
                                    "libero",
                                    "excepturi",
                                },
                                Keep: []string{
                                    "omnis",
                                    "officiis",
                                },
                                Remove: []string{
                                    "magni",
                                    "sit",
                                    "velit",
                                    "voluptatum",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Nicole Becker"),
                                    Value: "aut",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Mr. Forrest Ryan"),
                                    Value: "ea",
                                },
                            },
                            MaxEventBytes: cribl.Int64(262231),
                            Name: "Mr. Sandra Franey",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("ipsum"),
                                Length: cribl.Int64(891063),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "eaque",
                            TimestampEarliest: cribl.String("incidunt"),
                            TimestampLatest: cribl.String("ut"),
                            TimestampTimezone: cribl.String("delectus"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                        },
                    },
                    Tags: cribl.String("ullam"),
                },
            },
            DisableBreakers: false,
        },
        EarliestEpoch: cribl.Int64(978683),
        ErrorStateConfig: &shared.SearchJobErrorStateConfig{
            Coordinated: false,
            ErrorMessages: []string{
                "reprehenderit",
                "eos",
                "assumenda",
                "cumque",
            },
        },
        Group: "ut",
        ID: "077d0cc3-f408-4efc-95ce-b4d6e1eae0f7",
        LatestEpoch: cribl.Int64(332429),
        NumEventsAfter: cribl.Int64(683192),
        NumEventsBefore: cribl.Int64(929012),
        Query: "possimus",
        SampleRate: cribl.Int64(947561),
        SearchConfig: shared.SearchConfig{
            Datasets: []string{
                "deserunt",
            },
            HasSendOperator: false,
            OrderedFieldNames: []string{
                "id",
                "distinctio",
                "corporis",
                "quas",
            },
            SearchTerms: []shared.SearchTerm{
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "cupiditate",
                },
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "unde",
                },
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "et",
                },
            },
            SortFields: []shared.SortByField{
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionDescending,
                    FieldName: "sed",
                    NullPosition: shared.SortByFieldNullPositionNullsFirst,
                },
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionDescending,
                    FieldName: "pariatur",
                    NullPosition: shared.SortByFieldNullPositionNullsLast,
                },
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionAscending,
                    FieldName: "corrupti",
                    NullPosition: shared.SortByFieldNullPositionNullsLast,
                },
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionAscending,
                    FieldName: "vel",
                    NullPosition: shared.SortByFieldNullPositionNullsFirst,
                },
            },
            SuppressEventMarking: false,
            UseFormattedVisualization: false,
        },
        SearchParameterDeclarations: []shared.SearchParameter{
            shared.SearchParameter{
                DefaultValue: cribl.String("odio"),
                Name: "Beverly Block",
                Type: shared.SearchParameterTypeBoolean,
            },
            shared.SearchParameter{
                DefaultValue: cribl.String("commodi"),
                Name: "Mr. Rene Harris",
                Type: shared.SearchParameterTypeString,
            },
            shared.SearchParameter{
                DefaultValue: cribl.String("vero"),
                Name: "Miss Evan Dibbert",
                Type: shared.SearchParameterTypeNumber,
            },
            shared.SearchParameter{
                DefaultValue: cribl.String("maiores"),
                Name: "Lee Runte",
                Type: shared.SearchParameterTypeNumber,
            },
        },
        SearchParameterValues: cribl.String("explicabo"),
        Status: shared.SearchJobStatusQueued,
        TableConfig: &shared.TableViewSettings{
            ColumnFilterSettings: &shared.ColumnFilterSettings{
                Contains: "quos",
            },
            ColumnFormatSettings: &shared.ColumnFormatSettings{
                Palette: "deleniti",
                Precision: "enim",
                Prefix: "sit",
                Suffix: "voluptatem",
            },
            ColumnOrderSettings: &shared.ColumnOrderSettings{
                Order: "natus",
            },
            ColumnSortSettings: &shared.ColumnSortSettings{
                Sort: "voluptatem",
            },
            RowNumberColumnWidth: cribl.Int64(271216),
            ShowColumnTotals: false,
            ShowColumnTotalsPinned: false,
            ShowRowNumbers: false,
            ShowRowTotals: false,
            ShowRowTotalsPinned: false,
            ViewMode: shared.TableViewSettingsViewModeTable,
        },
        TargetEventTime: cribl.Int64(98846),
        TimeCompleted: cribl.Int64(82915),
        TimeCreated: 402908,
        TimeStarted: cribl.Int64(43928),
        Type: shared.SearchJobTypeDatatypePreview.ToPointer(),
        User: "dolores",
    }

    ctx := context.Background()
    res, err := s.SearchJob.Update(ctx, id, searchJob)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |
| `searchJob`                                           | [*shared.SearchJob](../../models/shared/searchjob.md) | :heavy_minus_sign:                                    | SearchJob object to be updated                        |


### Response

**[*operations.UpdateSearchJobResponse](../../models/operations/updatesearchjobresponse.md), error**

