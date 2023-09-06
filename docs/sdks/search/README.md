# Search

## Overview

Actions related to search

### Available Operations

* [Apply](#apply) - Applies a query snippet on a set of input events for preview
* [Create](#create) - Create SearchJob
* [Delete](#delete) - Delete SearchJob
* [DispatchSearch](#dispatchsearch) - Dispatch search *id* to worker nodes filtered by worker node filter using RPC
* [Get](#get) - Get Search documentation
* [GetJob](#getjob) - Get SearchJob by ID
* [GetTimeline](#gettimeline) - Get search timeline
* [ListFieldSummaries](#listfieldsummaries) - List field summaries
* [ListJobStatus](#listjobstatus) - Get job status
* [ListSearchJobMetrics](#listsearchjobmetrics) - Get search job metrics
* [ListSearchJobs](#listsearchjobs) - Get a list of SearchJob objects
* [ListSearchLogs](#listsearchlogs) - Get search logs
* [Post](#post) - Runs an event breaker rule on the specified data
* [Update](#update) - Update SearchJob

## Apply

Applies a query snippet on a set of input events for preview

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
    res, err := s.Search.Apply(ctx, shared.PreviewRequestBody{
        Events: []interface{}{
            "provident",
        },
        Options: &shared.PreviewOptions{},
        Query: "cumque",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PreviewResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.PreviewRequestBody](../../models/shared/previewrequestbody.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.ApplyQuerySnippetResponse](../../models/operations/applyquerysnippetresponse.md), error**


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
    res, err := s.Search.Create(ctx, shared.SearchJob{
        ChartConfig: &shared.ChartConfig{
            Axis: &shared.ChartConfigAxis{
                XAxis: cribl.String("iure"),
                YAxis: []string{
                    "quibusdam",
                },
            },
            Legend: &shared.Legend{},
            Series: []shared.ChartSeries{
                shared.ChartSeries{
                    Color: cribl.String("quod"),
                    Data: []map[string]interface{}{
                        map[string]interface{}{
                            "nemo": "recusandae",
                        },
                    },
                    DataExpression: "velit",
                    DataFilter: map[string]interface{}{
                        "magnam": "dignissimos",
                    },
                    Name: "Jane Kuhlman",
                    Type: &shared.ChartType{},
                },
            },
            Settings: &shared.Settings{
                Color: cribl.String("cum"),
                ColorPalette: 983854,
                Type: "facilis",
            },
            SingleValue: shared.SingleValue{
                Color: cribl.String("quidem"),
                Decimals: cribl.Int64(932080),
                Label: cribl.String("laboriosam"),
                Prefix: cribl.String("unde"),
                Suffix: cribl.String("modi"),
                Type: cribl.String("perspiciatis"),
            },
            XAxis: &shared.Axis{
                AxisLabel: &shared.AxisLabel{
                    Formatter: &shared.AxisLabelFormatter{},
                    FormatterData: []int64{
                        944626,
                    },
                },
                Type: cribl.String("cum"),
            },
        },
        CompatibilityChecks: &shared.SearchJobCompatibilityChecks{
            Datatypes: cribl.Bool(false),
        },
        CorrelationID: cribl.String("aspernatur"),
        CPUMetrics: &shared.CPUTimeMetric{
            TotalCPUSeconds: 725784,
            TotalExecCPUSeconds: 720266,
        },
        DatatypeOverrides: &shared.DatatypeOverrides{
            BreakerRulesets: []shared.EventBreakerRuleset{
                shared.EventBreakerRuleset{
                    Description: cribl.String("incidunt"),
                    ID: "ecae6c3d-5db3-4ade-bd5d-aea4c506a8aa",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(577413),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "labore",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("quo"),
                                FieldFilterExpr: cribl.String("perferendis"),
                                Fields: []string{
                                    "fugit",
                                },
                                Keep: []string{
                                    "aliquid",
                                },
                                Remove: []string{
                                    "magnam",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Lynne Weissnat"),
                                    Value: "unde",
                                },
                            },
                            MaxEventBytes: cribl.Int64(860311),
                            Name: "Shaun Gusikowski",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("corrupti"),
                                Length: cribl.Int64(684799),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                            },
                            TimestampAnchorRegex: "impedit",
                            TimestampEarliest: cribl.String("quasi"),
                            TimestampLatest: cribl.String("deserunt"),
                            TimestampTimezone: cribl.String("quod"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                        },
                    },
                    Tags: cribl.String("doloremque"),
                },
            },
            DisableBreakers: false,
        },
        EarliestEpoch: cribl.Int64(31574),
        ErrorStateConfig: &shared.SearchJobErrorStateConfig{
            Coordinated: false,
            ErrorMessages: []string{
                "facere",
            },
        },
        Group: "necessitatibus",
        ID: "c001ac80-2e2e-4c09-bf8f-0f816ff3477c",
        LatestEpoch: cribl.Int64(94697),
        NumEventsAfter: cribl.Int64(203621),
        NumEventsBefore: cribl.Int64(876649),
        Query: "excepturi",
        SampleRate: cribl.Int64(37534),
        SearchConfig: shared.SearchConfig{
            Datasets: []string{
                "qui",
            },
            HasSendOperator: false,
            OrderedFieldNames: []string{
                "impedit",
            },
            SearchTerms: []shared.SearchTerm{
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "beatae",
                },
            },
            SortFields: []shared.SortByField{
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionAscending,
                    FieldName: "dicta",
                    NullPosition: shared.SortByFieldNullPositionNullsFirst,
                },
            },
            SuppressEventMarking: false,
            UseFormattedVisualization: false,
        },
        SearchParameterDeclarations: []shared.SearchParameter{
            shared.SearchParameter{
                DefaultValue: cribl.String("corporis"),
                Name: "Robert Muller MD",
                Type: shared.SearchParameterTypeNumber,
            },
        },
        SearchParameterValues: cribl.String("ex"),
        Status: shared.SearchJobStatusFailed,
        TableConfig: &shared.TableViewSettings{
            ColumnFilterSettings: &shared.ColumnFilterSettings{
                Contains: "veritatis",
            },
            ColumnFormatSettings: &shared.ColumnFormatSettings{
                Palette: "ullam",
                Precision: "quae",
                Prefix: "similique",
                Suffix: "incidunt",
            },
            ColumnOrderSettings: &shared.ColumnOrderSettings{
                Order: "quam",
            },
            ColumnSortSettings: &shared.ColumnSortSettings{
                Sort: "magni",
            },
            RowNumberColumnWidth: cribl.Int64(646329),
            ShowColumnTotals: false,
            ShowColumnTotalsPinned: false,
            ShowRowNumbers: false,
            ShowRowTotals: false,
            ShowRowTotalsPinned: false,
            ViewMode: shared.TableViewSettingsViewModeTable,
        },
        TargetEventTime: cribl.Int64(609537),
        TimeCompleted: cribl.Int64(151230),
        TimeCreated: 200950,
        TimeStarted: cribl.Int64(805463),
        Type: shared.SearchJobTypeDatatypePreview.ToPointer(),
        User: "cupiditate",
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
    id := "aliquam"

    ctx := context.Background()
    res, err := s.Search.Delete(ctx, id)
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


## DispatchSearch

Dispatch search *id* to worker nodes filtered by worker node filter using RPC

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
    id := "excepturi"

    ctx := context.Background()
    res, err := s.Search.DispatchSearch(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchID != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | search job id                                         |


### Response

**[*operations.DispatchSearchResponse](../../models/operations/dispatchsearchresponse.md), error**


## Get

Get Search documentation

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
    res, err := s.Search.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSearchDoc200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetSearchDocResponse](../../models/operations/getsearchdocresponse.md), error**


## GetJob

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
    id := "maiores"

    ctx := context.Background()
    res, err := s.Search.GetJob(ctx, id)
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


## GetTimeline

Get search timeline

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
    id := "laudantium"

    ctx := context.Background()
    res, err := s.Search.GetTimeline(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchTimeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | search job id                                         |


### Response

**[*operations.GetSearchTimelineResponse](../../models/operations/getsearchtimelineresponse.md), error**


## ListFieldSummaries

List field summaries

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
    id := "velit"

    ctx := context.Background()
    res, err := s.Search.ListFieldSummaries(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.FieldSummaries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | search job id                                         |


### Response

**[*operations.ListFieldSummariesResponse](../../models/operations/listfieldsummariesresponse.md), error**


## ListJobStatus

Get job status

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
    id := "reiciendis"

    ctx := context.Background()
    res, err := s.Search.ListJobStatus(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchJobStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | search job id                                         |


### Response

**[*operations.ListJobStatusResponse](../../models/operations/listjobstatusresponse.md), error**


## ListSearchJobMetrics

Get search job metrics

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
    id := "amet"

    ctx := context.Background()
    res, err := s.Search.ListSearchJobMetrics(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSearchJobMetrics200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | search job id                                         |


### Response

**[*operations.ListSearchJobMetricsResponse](../../models/operations/listsearchjobmetricsresponse.md), error**


## ListSearchJobs

Get a list of SearchJob objects

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
    res, err := s.Search.ListSearchJobs(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchJobs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListSearchJobsResponse](../../models/operations/listsearchjobsresponse.md), error**


## ListSearchLogs

Get search logs

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
    id := "nemo"

    ctx := context.Background()
    res, err := s.Search.ListSearchLogs(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchLogs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | search job id                                         |


### Response

**[*operations.ListSearchLogsResponse](../../models/operations/listsearchlogsresponse.md), error**


## Post

Runs an event breaker rule on the specified data

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
    res, err := s.Search.Post(ctx, shared.DatatypePreviewRequestBody{
        EventBreakerRule: &shared.EventBreakerRule{
            CleanFields: cribl.Bool(false),
            Condition: "ipsa",
            Delimiter: cribl.String("quisquam"),
            DelimiterRegex: cribl.String("tenetur"),
            Disabled: cribl.Bool(false),
            EscapeChar: cribl.String("quas"),
            EventBreakerRegex: cribl.String("molestiae"),
            Fields: []shared.EventBreakerRuleFields{
                shared.EventBreakerRuleFields{
                    Name: "Nettie Wilkinson",
                    Value: "accusantium",
                },
            },
            FieldsLineRegex: cribl.String("dicta"),
            HeaderLineRegex: cribl.String("minus"),
            JSONArrayField: cribl.String("commodi"),
            JSONExtractAll: cribl.Bool(false),
            JSONTimeField: cribl.String("eveniet"),
            MaxEventBytes: 781582,
            Name: "Abraham Gleason",
            NullFieldVal: cribl.String("eius"),
            Parser: cribl.String("sequi"),
            ParserEnabled: cribl.Bool(false),
            QuoteChar: cribl.String("eligendi"),
            TimeField: cribl.String("asperiores"),
            Timestamp: shared.EventBreakerRuleTimestamp{
                Format: cribl.String("esse"),
                Length: cribl.Int64(500021),
                Type: shared.EventBreakerRuleTimestampTypeFormat,
            },
            TimestampAnchorRegex: "repellat",
            TimestampEarliest: cribl.String("a"),
            TimestampLatest: cribl.String("animi"),
            TimestampTimezone: "maiores",
            Type: shared.EventBreakerRuleTypeCsv.ToPointer(),
        },
        Input: "nulla",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PreviewResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [shared.DatatypePreviewRequestBody](../../models/shared/datatypepreviewrequestbody.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PostEventBreakerOnDataResponse](../../models/operations/posteventbreakerondataresponse.md), error**


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
    id := "deserunt"
    searchJob := &shared.SearchJob{
        ChartConfig: &shared.ChartConfig{
            Axis: &shared.ChartConfigAxis{
                XAxis: cribl.String("corporis"),
                YAxis: []string{
                    "velit",
                },
            },
            Legend: &shared.Legend{},
            Series: []shared.ChartSeries{
                shared.ChartSeries{
                    Color: cribl.String("officiis"),
                    Data: []map[string]interface{}{
                        map[string]interface{}{
                            "enim": "officia",
                        },
                    },
                    DataExpression: "saepe",
                    DataFilter: map[string]interface{}{
                        "eum": "repudiandae",
                    },
                    Name: "Ms. Kristi Russel",
                    Type: &shared.ChartType{},
                },
            },
            Settings: &shared.Settings{
                Color: cribl.String("quisquam"),
                ColorPalette: 177929,
                Type: "nobis",
            },
            SingleValue: shared.SingleValue{
                Color: cribl.String("natus"),
                Decimals: cribl.Int64(793568),
                Label: cribl.String("quia"),
                Prefix: cribl.String("magnam"),
                Suffix: cribl.String("reprehenderit"),
                Type: cribl.String("quod"),
            },
            XAxis: &shared.Axis{
                AxisLabel: &shared.AxisLabel{
                    Formatter: &shared.AxisLabelFormatter{},
                    FormatterData: []int64{
                        552440,
                    },
                },
                Type: cribl.String("corrupti"),
            },
        },
        CompatibilityChecks: &shared.SearchJobCompatibilityChecks{
            Datatypes: cribl.Bool(false),
        },
        CorrelationID: cribl.String("amet"),
        CPUMetrics: &shared.CPUTimeMetric{
            TotalCPUSeconds: 473326,
            TotalExecCPUSeconds: 227156,
        },
        DatatypeOverrides: &shared.DatatypeOverrides{
            BreakerRulesets: []shared.EventBreakerRuleset{
                shared.EventBreakerRuleset{
                    Description: cribl.String("laborum"),
                    ID: "40e1942f-32e5-4505-9756-f5d56d0bd0af",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(152364),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "possimus",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("repellat"),
                                FieldFilterExpr: cribl.String("repudiandae"),
                                Fields: []string{
                                    "architecto",
                                },
                                Keep: []string{
                                    "adipisci",
                                },
                                Remove: []string{
                                    "pariatur",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Theodore Wolff DDS"),
                                    Value: "soluta",
                                },
                            },
                            MaxEventBytes: cribl.Int64(677509),
                            Name: "Lila Langosh",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("quasi"),
                                Length: cribl.Int64(649032),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                            },
                            TimestampAnchorRegex: "harum",
                            TimestampEarliest: cribl.String("cumque"),
                            TimestampLatest: cribl.String("doloremque"),
                            TimestampTimezone: cribl.String("expedita"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeHeader,
                        },
                    },
                    Tags: cribl.String("eaque"),
                },
            },
            DisableBreakers: false,
        },
        EarliestEpoch: cribl.Int64(643199),
        ErrorStateConfig: &shared.SearchJobErrorStateConfig{
            Coordinated: false,
            ErrorMessages: []string{
                "aliquid",
            },
        },
        Group: "excepturi",
        ID: "24d3b2ec-fcc8-4f89-9010-f5dd3d6fa180",
        LatestEpoch: cribl.Int64(251200),
        NumEventsAfter: cribl.Int64(913284),
        NumEventsBefore: cribl.Int64(324052),
        Query: "aliquam",
        SampleRate: cribl.Int64(788165),
        SearchConfig: shared.SearchConfig{
            Datasets: []string{
                "quas",
            },
            HasSendOperator: false,
            OrderedFieldNames: []string{
                "consequuntur",
            },
            SearchTerms: []shared.SearchTerm{
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "maiores",
                },
            },
            SortFields: []shared.SortByField{
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionAscending,
                    FieldName: "aliquid",
                    NullPosition: shared.SortByFieldNullPositionNullsLast,
                },
            },
            SuppressEventMarking: false,
            UseFormattedVisualization: false,
        },
        SearchParameterDeclarations: []shared.SearchParameter{
            shared.SearchParameter{
                DefaultValue: cribl.String("est"),
                Name: "Vicki Feeney",
                Type: shared.SearchParameterTypeNumber,
            },
        },
        SearchParameterValues: cribl.String("ducimus"),
        Status: shared.SearchJobStatusRunning,
        TableConfig: &shared.TableViewSettings{
            ColumnFilterSettings: &shared.ColumnFilterSettings{
                Contains: "recusandae",
            },
            ColumnFormatSettings: &shared.ColumnFormatSettings{
                Palette: "tempora",
                Precision: "blanditiis",
                Prefix: "numquam",
                Suffix: "sequi",
            },
            ColumnOrderSettings: &shared.ColumnOrderSettings{
                Order: "voluptatum",
            },
            ColumnSortSettings: &shared.ColumnSortSettings{
                Sort: "sit",
            },
            RowNumberColumnWidth: cribl.Int64(703189),
            ShowColumnTotals: false,
            ShowColumnTotalsPinned: false,
            ShowRowNumbers: false,
            ShowRowTotals: false,
            ShowRowTotalsPinned: false,
            ViewMode: shared.TableViewSettingsViewModeEvent,
        },
        TargetEventTime: cribl.Int64(949280),
        TimeCompleted: cribl.Int64(419858),
        TimeCreated: 694088,
        TimeStarted: cribl.Int64(517121),
        Type: shared.SearchJobTypeScheduled.ToPointer(),
        User: "deserunt",
    }

    ctx := context.Background()
    res, err := s.Search.Update(ctx, id, searchJob)
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

