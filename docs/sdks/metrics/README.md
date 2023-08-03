# Metrics

## Overview

Actions related to metrics

### Available Operations

* [Aggregate](#aggregate) - Aggregate raw internal system metrics
* [Post](#post) - Enumerate all internal system metrics
* [Query](#query) - Query raw internal system metrics

## Aggregate

Aggregate raw internal system metrics

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
    res, err := s.Metrics.Aggregate(ctx, shared.MetricsAggOpts{
        Aggs: shared.AggregationMgrOptions{
            Aggregations: []string{
                "quo",
                "incidunt",
                "quod",
                "minus",
            },
            Cumulative: false,
            FlushEventLimit: 781491,
            FlushMemLimit: 659971,
            Hostname: "nimble-mouse.net",
            IdleTimeLimitSeconds: cribl.Int64(774880),
            LagToleranceSeconds: cribl.Int64(457150),
            MetricsMode: false,
            Prefix: cribl.String("hic"),
            PreserveSplitByStructure: cribl.Bool(false),
            SearchAggMode: cribl.String("maxime"),
            SplitBys: []string{
                "soluta",
            },
            SufficientStatsOnly: false,
            TimeWindowSeconds: 147400,
        },
        AlwaysBounds: cribl.Bool(false),
        Metrics: &shared.MetricsStore{},
        Where: cribl.String("pariatur"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MetricsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [shared.MetricsAggOpts](../../models/shared/metricsaggopts.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |


### Response

**[*operations.PostInternalSystemMetricsResponse](../../models/operations/postinternalsystemmetricsresponse.md), error**


## Post

Enumerate all internal system metrics

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
    res, err := s.Metrics.Post(ctx, shared.GetNamesOpts{
        DimKeyFilter: cribl.String("eligendi"),
        DimValueFilter: cribl.String("recusandae"),
        Earliest: cribl.Int64(83791),
        FilterExpr: &shared.Expression{
            MaxCache: 13637,
            Cache: shared.Map{
                Op: shared.MapOp{},
            },
            DeclaredVariables: []string{
                "iusto",
                "dolor",
                "voluptates",
            },
            IsSafe: false,
            ModifiedExpression: "tempora",
            Opt: shared.ExpressionOptions{
                AllowInternal: cribl.Bool(false),
                Args: []string{
                    "rerum",
                },
                AstTraverseCallback: &shared.TraverseCallback{},
                Context: cribl.String("doloremque"),
                DisallowAssign: cribl.Bool(false),
                MaxNumOfAllowedIterations: cribl.Int64(30192),
                PartialEval: &shared.PartialEvalRewrite{
                    FieldFilter: shared.PartialEvalFieldFilter{},
                    NullObj: "eum",
                },
                ReplaceIdentifiers: cribl.Bool(false),
                ReplaceLiterals: cribl.Bool(false),
                Unprotected: cribl.Bool(false),
            },
            OriginalExpression: "at",
            PartialExpression: "eum",
            ReferencedCriblExport: false,
            ReplaceIdentifiersExpression: "reprehenderit",
            ReplaceLiteralExpression: "voluptatum",
        },
        MaxValues: cribl.Int64(502393),
        MetricNameFilter: cribl.String("nihil"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MetricsInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [shared.GetNamesOpts](../../models/shared/getnamesopts.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |


### Response

**[*operations.PostMetricsResponse](../../models/operations/postmetricsresponse.md), error**


## Query

Query raw internal system metrics

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

    ctx := context.Background()
    res, err := s.Metrics.Query(ctx, operations.QueryMetricsRequest{
        Earliest: cribl.Int64(540048),
        FilterExpr: cribl.String("rerum"),
        Latest: cribl.Int64(645544),
        MetricNameFilter: cribl.String("atque"),
        NumBuckets: cribl.Int64(344856),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MetricsInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.QueryMetricsRequest](../../models/operations/querymetricsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.QueryMetricsResponse](../../models/operations/querymetricsresponse.md), error**

