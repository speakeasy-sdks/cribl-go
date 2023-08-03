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
            "totam",
            "cupiditate",
            "at",
        },
        Options: &shared.PreviewOptions{},
        Query: "doloribus",
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
                XAxis: cribl.String("omnis"),
                YAxis: []string{
                    "exercitationem",
                    "voluptates",
                },
            },
            Legend: &shared.Legend{},
            Series: []shared.ChartSeries{
                shared.ChartSeries{
                    Color: cribl.String("quis"),
                    Data: []map[string]interface{}{
                        map[string]interface{}{
                            "rem": "aliquid",
                            "aperiam": "perspiciatis",
                        },
                        map[string]interface{}{
                            "itaque": "unde",
                        },
                    },
                    DataExpression: "cumque",
                    DataFilter: map[string]interface{}{
                        "repellendus": "temporibus",
                    },
                    Name: "Mr. Warren Wilderman DDS",
                    Type: &shared.ChartType{},
                },
            },
            Settings: &shared.Settings{
                Color: cribl.String("earum"),
                ColorPalette: 672273,
                Type: "dicta",
            },
            SingleValue: shared.SingleValue{
                Color: cribl.String("voluptatem"),
                Decimals: cribl.Int64(140909),
                Label: cribl.String("aliquid"),
                Prefix: cribl.String("pariatur"),
                Suffix: cribl.String("enim"),
                Type: cribl.String("numquam"),
            },
            XAxis: &shared.Axis{
                AxisLabel: &shared.AxisLabel{
                    Formatter: &shared.AxisLabelFormatter{},
                    FormatterData: []int64{
                        665952,
                    },
                },
                Type: cribl.String("quaerat"),
            },
        },
        CompatibilityChecks: &shared.SearchJobCompatibilityChecks{
            Datatypes: cribl.Bool(false),
        },
        CorrelationID: cribl.String("facere"),
        CPUMetrics: &shared.CPUTimeMetric{
            TotalCPUSeconds: 112391,
            TotalExecCPUSeconds: 607181,
        },
        DatatypeOverrides: &shared.DatatypeOverrides{
            BreakerRulesets: []shared.EventBreakerRuleset{
                shared.EventBreakerRuleset{
                    Description: cribl.String("sapiente"),
                    ID: "eb21780b-ccc0-4dbb-9db4-84708fb4e391",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(887835),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "rerum",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("placeat"),
                                FieldFilterExpr: cribl.String("ab"),
                                Fields: []string{
                                    "blanditiis",
                                    "porro",
                                },
                                Keep: []string{
                                    "impedit",
                                    "ut",
                                },
                                Remove: []string{
                                    "ullam",
                                    "numquam",
                                    "enim",
                                    "cupiditate",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Blake Feil"),
                                    Value: "explicabo",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Melissa Vandervort"),
                                    Value: "qui",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Elizabeth Rutherford"),
                                    Value: "deleniti",
                                },
                            },
                            MaxEventBytes: cribl.Int64(666321),
                            Name: "Mamie Spinka",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("quidem"),
                                Length: cribl.Int64(448601),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "doloremque",
                            TimestampEarliest: cribl.String("fuga"),
                            TimestampLatest: cribl.String("dicta"),
                            TimestampTimezone: cribl.String("architecto"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "eligendi",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("officiis"),
                                FieldFilterExpr: cribl.String("dignissimos"),
                                Fields: []string{
                                    "ratione",
                                },
                                Keep: []string{
                                    "quaerat",
                                    "aut",
                                    "natus",
                                    "esse",
                                },
                                Remove: []string{
                                    "deserunt",
                                    "ratione",
                                    "ipsa",
                                    "debitis",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Irvin Kuphal"),
                                    Value: "libero",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Belinda Cartwright"),
                                    Value: "consequatur",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Elizabeth Stehr"),
                                    Value: "hic",
                                },
                            },
                            MaxEventBytes: cribl.Int64(356485),
                            Name: "Dexter Ratke",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("provident"),
                                Length: cribl.Int64(581889),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                            },
                            TimestampAnchorRegex: "sunt",
                            TimestampEarliest: cribl.String("odit"),
                            TimestampLatest: cribl.String("vero"),
                            TimestampTimezone: cribl.String("deleniti"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeTimestamp,
                        },
                    },
                    Tags: cribl.String("quasi"),
                },
            },
            DisableBreakers: false,
        },
        EarliestEpoch: cribl.Int64(998199),
        ErrorStateConfig: &shared.SearchJobErrorStateConfig{
            Coordinated: false,
            ErrorMessages: []string{
                "magnam",
                "perspiciatis",
                "amet",
            },
        },
        Group: "corrupti",
        ID: "25fdc42c-876c-42c2-9fb4-cfc1c76230f8",
        LatestEpoch: cribl.Int64(256310),
        NumEventsAfter: cribl.Int64(113894),
        NumEventsBefore: cribl.Int64(980376),
        Query: "nam",
        SampleRate: cribl.Int64(102184),
        SearchConfig: shared.SearchConfig{
            Datasets: []string{
                "assumenda",
                "eos",
                "dolorem",
            },
            HasSendOperator: false,
            OrderedFieldNames: []string{
                "repellendus",
                "nam",
                "ab",
                "magnam",
            },
            SearchTerms: []shared.SearchTerm{
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "expedita",
                },
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "autem",
                },
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "tempore",
                },
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "recusandae",
                },
            },
            SortFields: []shared.SortByField{
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionDescending,
                    FieldName: "voluptas",
                    NullPosition: shared.SortByFieldNullPositionNullsLast,
                },
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionAscending,
                    FieldName: "excepturi",
                    NullPosition: shared.SortByFieldNullPositionNullsLast,
                },
            },
            SuppressEventMarking: false,
            UseFormattedVisualization: false,
        },
        SearchParameterDeclarations: []shared.SearchParameter{
            shared.SearchParameter{
                DefaultValue: cribl.String("necessitatibus"),
                Name: "Wanda Osinski",
                Type: shared.SearchParameterTypeString,
            },
            shared.SearchParameter{
                DefaultValue: cribl.String("temporibus"),
                Name: "Scott Johns",
                Type: shared.SearchParameterTypeString,
            },
            shared.SearchParameter{
                DefaultValue: cribl.String("nam"),
                Name: "Heidi Bernier",
                Type: shared.SearchParameterTypeNumber,
            },
        },
        SearchParameterValues: cribl.String("sint"),
        Status: shared.SearchJobStatusCanceled,
        TableConfig: &shared.TableViewSettings{
            ColumnFilterSettings: &shared.ColumnFilterSettings{
                Contains: "ullam",
            },
            ColumnFormatSettings: &shared.ColumnFormatSettings{
                Palette: "molestiae",
                Precision: "itaque",
                Prefix: "rem",
                Suffix: "nemo",
            },
            ColumnOrderSettings: &shared.ColumnOrderSettings{
                Order: "non",
            },
            ColumnSortSettings: &shared.ColumnSortSettings{
                Sort: "recusandae",
            },
            RowNumberColumnWidth: cribl.Int64(605043),
            ShowColumnTotals: false,
            ShowColumnTotalsPinned: false,
            ShowRowNumbers: false,
            ShowRowTotals: false,
            ShowRowTotalsPinned: false,
            ViewMode: shared.TableViewSettingsViewModeEvent,
        },
        TargetEventTime: cribl.Int64(302905),
        TimeCompleted: cribl.Int64(219940),
        TimeCreated: 577284,
        TimeStarted: cribl.Int64(842678),
        Type: shared.SearchJobTypeStandard.ToPointer(),
        User: "consequuntur",
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
    id := "consequuntur"

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
    id := "eius"

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
    id := "commodi"

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
    id := "ipsam"

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
    id := "vel"

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
    id := "cupiditate"

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
    id := "modi"

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
    id := "nisi"

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
            Condition: "explicabo",
            Delimiter: cribl.String("modi"),
            DelimiterRegex: cribl.String("doloremque"),
            Disabled: cribl.Bool(false),
            EscapeChar: cribl.String("odio"),
            EventBreakerRegex: cribl.String("sit"),
            Fields: []shared.EventBreakerRuleFields{
                shared.EventBreakerRuleFields{
                    Name: "Lucia Koepp",
                    Value: "ratione",
                },
                shared.EventBreakerRuleFields{
                    Name: "Kara Turner Sr.",
                    Value: "eos",
                },
                shared.EventBreakerRuleFields{
                    Name: "Beatrice Buckridge",
                    Value: "fugiat",
                },
            },
            FieldsLineRegex: cribl.String("quidem"),
            HeaderLineRegex: cribl.String("exercitationem"),
            JSONArrayField: cribl.String("veniam"),
            JSONExtractAll: cribl.Bool(false),
            JSONTimeField: cribl.String("modi"),
            MaxEventBytes: 95456,
            Name: "Hattie Schuster",
            NullFieldVal: cribl.String("ex"),
            Parser: cribl.String("sint"),
            ParserEnabled: cribl.Bool(false),
            QuoteChar: cribl.String("est"),
            TimeField: cribl.String("doloribus"),
            Timestamp: shared.EventBreakerRuleTimestamp{
                Format: cribl.String("provident"),
                Length: cribl.Int64(1594),
                Type: shared.EventBreakerRuleTimestampTypeFormat,
            },
            TimestampAnchorRegex: "fugit",
            TimestampEarliest: cribl.String("autem"),
            TimestampLatest: cribl.String("quo"),
            TimestampTimezone: "molestiae",
            Type: shared.EventBreakerRuleTypeTimestamp.ToPointer(),
        },
        Input: "facere",
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
    id := "impedit"
    searchJob := &shared.SearchJob{
        ChartConfig: &shared.ChartConfig{
            Axis: &shared.ChartConfigAxis{
                XAxis: cribl.String("cupiditate"),
                YAxis: []string{
                    "quasi",
                    "maiores",
                    "voluptatem",
                },
            },
            Legend: &shared.Legend{},
            Series: []shared.ChartSeries{
                shared.ChartSeries{
                    Color: cribl.String("laudantium"),
                    Data: []map[string]interface{}{
                        map[string]interface{}{
                            "quae": "facere",
                            "ea": "libero",
                            "nam": "amet",
                        },
                        map[string]interface{}{
                            "minus": "hic",
                        },
                        map[string]interface{}{
                            "fuga": "consectetur",
                            "labore": "laudantium",
                            "cumque": "adipisci",
                        },
                    },
                    DataExpression: "veritatis",
                    DataFilter: map[string]interface{}{
                        "voluptatibus": "magnam",
                        "aperiam": "ducimus",
                        "itaque": "necessitatibus",
                    },
                    Name: "Ora Rosenbaum DDS",
                    Type: &shared.ChartType{},
                },
                shared.ChartSeries{
                    Color: cribl.String("numquam"),
                    Data: []map[string]interface{}{
                        map[string]interface{}{
                            "nihil": "voluptatum",
                            "reiciendis": "vitae",
                            "ullam": "nisi",
                        },
                    },
                    DataExpression: "consequuntur",
                    DataFilter: map[string]interface{}{
                        "ratione": "excepturi",
                        "corrupti": "est",
                    },
                    Name: "Mona Russel",
                    Type: &shared.ChartType{},
                },
            },
            Settings: &shared.Settings{
                Color: cribl.String("nisi"),
                ColorPalette: 223358,
                Type: "fugit",
            },
            SingleValue: shared.SingleValue{
                Color: cribl.String("dolore"),
                Decimals: cribl.Int64(804823),
                Label: cribl.String("maxime"),
                Prefix: cribl.String("expedita"),
                Suffix: cribl.String("accusantium"),
                Type: cribl.String("ea"),
            },
            XAxis: &shared.Axis{
                AxisLabel: &shared.AxisLabel{
                    Formatter: &shared.AxisLabelFormatter{},
                    FormatterData: []int64{
                        518432,
                        762297,
                        665183,
                        81541,
                    },
                },
                Type: cribl.String("consequuntur"),
            },
        },
        CompatibilityChecks: &shared.SearchJobCompatibilityChecks{
            Datatypes: cribl.Bool(false),
        },
        CorrelationID: cribl.String("repellendus"),
        CPUMetrics: &shared.CPUTimeMetric{
            TotalCPUSeconds: 27197,
            TotalExecCPUSeconds: 172042,
        },
        DatatypeOverrides: &shared.DatatypeOverrides{
            BreakerRulesets: []shared.EventBreakerRuleset{
                shared.EventBreakerRuleset{
                    Description: cribl.String("aspernatur"),
                    ID: "9270b8d5-722d-4d89-9b8b-cf24db959693",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(221319),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "qui",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("tenetur"),
                                FieldFilterExpr: cribl.String("molestiae"),
                                Fields: []string{
                                    "ullam",
                                    "velit",
                                },
                                Keep: []string{
                                    "cupiditate",
                                },
                                Remove: []string{
                                    "numquam",
                                    "fugiat",
                                    "molestiae",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Elbert Fay"),
                                    Value: "necessitatibus",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Danny Lockman"),
                                    Value: "ullam",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Gerardo Ritchie"),
                                    Value: "commodi",
                                },
                            },
                            MaxEventBytes: cribl.Int64(414438),
                            Name: "Miss Marion Hermann",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("praesentium"),
                                Length: cribl.Int64(237189),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "qui",
                            TimestampEarliest: cribl.String("deserunt"),
                            TimestampLatest: cribl.String("eligendi"),
                            TimestampTimezone: cribl.String("incidunt"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeHeader,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "dolor",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("est"),
                                FieldFilterExpr: cribl.String("reiciendis"),
                                Fields: []string{
                                    "odit",
                                    "consectetur",
                                    "inventore",
                                    "minima",
                                },
                                Keep: []string{
                                    "facilis",
                                    "deserunt",
                                    "nisi",
                                },
                                Remove: []string{
                                    "voluptatem",
                                    "illo",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Laverne Berge"),
                                    Value: "veniam",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Aubrey Kemmer"),
                                    Value: "nemo",
                                },
                            },
                            MaxEventBytes: cribl.Int64(565446),
                            Name: "Mamie Schaefer",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("assumenda"),
                                Length: cribl.Int64(921916),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeCurrent,
                            },
                            TimestampAnchorRegex: "ipsum",
                            TimestampEarliest: cribl.String("commodi"),
                            TimestampLatest: cribl.String("vitae"),
                            TimestampTimezone: cribl.String("fugit"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeTimestamp,
                        },
                    },
                    Tags: cribl.String("ex"),
                },
                shared.EventBreakerRuleset{
                    Description: cribl.String("neque"),
                    ID: "c205fda8-4077-44a6-8a9a-35d086b6f66f",
                    Lib: shared.EventBreakerRulesetLibraryCustom.ToPointer(),
                    MinRawLength: cribl.Int64(881067),
                    Rules: []shared.EventBreakerRulesetRules{
                        shared.EventBreakerRulesetRules{
                            Condition: "consequatur",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("sed"),
                                FieldFilterExpr: cribl.String("accusantium"),
                                Fields: []string{
                                    "provident",
                                    "maiores",
                                    "quaerat",
                                    "numquam",
                                },
                                Keep: []string{
                                    "cum",
                                },
                                Remove: []string{
                                    "dolores",
                                    "enim",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Arturo Marks"),
                                    Value: "corrupti",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Felipe Spinka"),
                                    Value: "similique",
                                },
                            },
                            MaxEventBytes: cribl.Int64(421273),
                            Name: "Kate Will",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("inventore"),
                                Length: cribl.Int64(612238),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "explicabo",
                            TimestampEarliest: cribl.String("ullam"),
                            TimestampLatest: cribl.String("atque"),
                            TimestampTimezone: cribl.String("doloribus"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeCsv,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "aut",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("similique"),
                                FieldFilterExpr: cribl.String("iste"),
                                Fields: []string{
                                    "nam",
                                    "animi",
                                    "labore",
                                    "voluptate",
                                },
                                Keep: []string{
                                    "quam",
                                    "nulla",
                                    "dolorem",
                                    "voluptates",
                                },
                                Remove: []string{
                                    "perferendis",
                                    "quaerat",
                                    "excepturi",
                                    "aliquid",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Desiree Howell IV"),
                                    Value: "ratione",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Rosalie Lindgren"),
                                    Value: "temporibus",
                                },
                            },
                            MaxEventBytes: cribl.Int64(955913),
                            Name: "Mae Hoppe",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("delectus"),
                                Length: cribl.Int64(67631),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "possimus",
                            TimestampEarliest: cribl.String("praesentium"),
                            TimestampLatest: cribl.String("neque"),
                            TimestampTimezone: cribl.String("quam"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeHeader,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "debitis",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("voluptatum"),
                                FieldFilterExpr: cribl.String("voluptatem"),
                                Fields: []string{
                                    "vitae",
                                    "cumque",
                                    "architecto",
                                    "sint",
                                },
                                Keep: []string{
                                    "occaecati",
                                    "quis",
                                    "tempore",
                                    "officia",
                                },
                                Remove: []string{
                                    "unde",
                                    "quas",
                                    "laboriosam",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Timmy Oberbrunner"),
                                    Value: "iure",
                                },
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Milton Morar MD"),
                                    Value: "a",
                                },
                            },
                            MaxEventBytes: cribl.Int64(222771),
                            Name: "Daryl Schmitt I",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("autem"),
                                Length: cribl.Int64(110804),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeAuto,
                            },
                            TimestampAnchorRegex: "incidunt",
                            TimestampEarliest: cribl.String("modi"),
                            TimestampLatest: cribl.String("quos"),
                            TimestampTimezone: cribl.String("minus"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeJSONArray,
                        },
                        shared.EventBreakerRulesetRules{
                            Condition: "error",
                            Definitions: &shared.EventBreakerRulesetRulesDefinitions{
                                DstField: cribl.String("reprehenderit"),
                                FieldFilterExpr: cribl.String("reprehenderit"),
                                Fields: []string{
                                    "consequatur",
                                    "voluptates",
                                    "delectus",
                                },
                                Keep: []string{
                                    "delectus",
                                },
                                Remove: []string{
                                    "ipsum",
                                    "laboriosam",
                                },
                            },
                            Disabled: cribl.Bool(false),
                            Fields: []shared.EventBreakerRulesetRulesFields{
                                shared.EventBreakerRulesetRulesFields{
                                    Name: cribl.String("Nora Toy"),
                                    Value: "eveniet",
                                },
                            },
                            MaxEventBytes: cribl.Int64(999208),
                            Name: "Mrs. Travis Gutmann",
                            ParserEnabled: cribl.Bool(false),
                            Timestamp: shared.EventBreakerRulesetRulesTimestampFormat{
                                Format: cribl.String("necessitatibus"),
                                Length: cribl.Int64(820866),
                                Type: shared.EventBreakerRulesetRulesTimestampFormatTimestampTypeFormat,
                            },
                            TimestampAnchorRegex: "itaque",
                            TimestampEarliest: cribl.String("explicabo"),
                            TimestampLatest: cribl.String("ullam"),
                            TimestampTimezone: cribl.String("non"),
                            Type: shared.EventBreakerRulesetRulesEventBreakerTypeCsv,
                        },
                    },
                    Tags: cribl.String("incidunt"),
                },
            },
            DisableBreakers: false,
        },
        EarliestEpoch: cribl.Int64(796952),
        ErrorStateConfig: &shared.SearchJobErrorStateConfig{
            Coordinated: false,
            ErrorMessages: []string{
                "ullam",
            },
        },
        Group: "quam",
        ID: "deaa7170-f445-4acc-b667-aaf9bbad185f",
        LatestEpoch: cribl.Int64(913393),
        NumEventsAfter: cribl.Int64(281135),
        NumEventsBefore: cribl.Int64(201005),
        Query: "ab",
        SampleRate: cribl.Int64(843591),
        SearchConfig: shared.SearchConfig{
            Datasets: []string{
                "quidem",
                "delectus",
            },
            HasSendOperator: false,
            OrderedFieldNames: []string{
                "cumque",
                "voluptatum",
            },
            SearchTerms: []shared.SearchTerm{
                shared.SearchTerm{
                    IsCaseSensitive: false,
                    Term: "atque",
                },
            },
            SortFields: []shared.SortByField{
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionDescending,
                    FieldName: "rerum",
                    NullPosition: shared.SortByFieldNullPositionNullsLast,
                },
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionDescending,
                    FieldName: "aspernatur",
                    NullPosition: shared.SortByFieldNullPositionNullsFirst,
                },
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionDescending,
                    FieldName: "nam",
                    NullPosition: shared.SortByFieldNullPositionNullsFirst,
                },
                shared.SortByField{
                    Direction: shared.SortByFieldDirectionAscending,
                    FieldName: "delectus",
                    NullPosition: shared.SortByFieldNullPositionNullsLast,
                },
            },
            SuppressEventMarking: false,
            UseFormattedVisualization: false,
        },
        SearchParameterDeclarations: []shared.SearchParameter{
            shared.SearchParameter{
                DefaultValue: cribl.String("distinctio"),
                Name: "Julia Hartmann",
                Type: shared.SearchParameterTypeNumber,
            },
            shared.SearchParameter{
                DefaultValue: cribl.String("debitis"),
                Name: "Julia Farrell",
                Type: shared.SearchParameterTypeNumber,
            },
        },
        SearchParameterValues: cribl.String("reiciendis"),
        Status: shared.SearchJobStatusCompleted,
        TableConfig: &shared.TableViewSettings{
            ColumnFilterSettings: &shared.ColumnFilterSettings{
                Contains: "tempore",
            },
            ColumnFormatSettings: &shared.ColumnFormatSettings{
                Palette: "in",
                Precision: "omnis",
                Prefix: "possimus",
                Suffix: "tenetur",
            },
            ColumnOrderSettings: &shared.ColumnOrderSettings{
                Order: "recusandae",
            },
            ColumnSortSettings: &shared.ColumnSortSettings{
                Sort: "expedita",
            },
            RowNumberColumnWidth: cribl.Int64(481307),
            ShowColumnTotals: false,
            ShowColumnTotalsPinned: false,
            ShowRowNumbers: false,
            ShowRowTotals: false,
            ShowRowTotalsPinned: false,
            ViewMode: shared.TableViewSettingsViewModeEvent,
        },
        TargetEventTime: cribl.Int64(687352),
        TimeCompleted: cribl.Int64(321654),
        TimeCreated: 801059,
        TimeStarted: cribl.Int64(188705),
        Type: shared.SearchJobTypeDatatypePreview.ToPointer(),
        User: "vero",
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

