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
                "laborum",
            },
            Cumulative: false,
            FlushEventLimit: 250398,
            FlushMemLimit: 224467,
            Hostname: "left-amber.com",
            IdleTimeLimitSeconds: cribl.Int64(7468),
            LagToleranceSeconds: cribl.Int64(639705),
            MetricsMode: false,
            Prefix: cribl.String("recusandae"),
            PreserveSplitByStructure: cribl.Bool(false),
            SearchAggMode: cribl.String("ea"),
            SplitBys: []string{
                "quidem",
            },
            SufficientStatsOnly: false,
            TimeWindowSeconds: 377406,
        },
        AlwaysBounds: cribl.Bool(false),
        Metrics: &shared.MetricsStore{},
        Where: cribl.String("facilis"),
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
        DimKeyFilter: cribl.String("placeat"),
        DimValueFilter: cribl.String("perspiciatis"),
        Earliest: cribl.Int64(709036),
        FilterExpr: &shared.Expression{
            MaxCache: 537236,
            Cache: shared.Map{
                Op: shared.MapOp{},
            },
            DeclaredVariables: []string{
                "a",
            },
            IsSafe: false,
            ModifiedExpression: "voluptate",
            Opt: shared.ExpressionOptions{
                AllowInternal: cribl.Bool(false),
                Args: []string{
                    "ullam",
                },
                AstTraverseCallback: &shared.TraverseCallback{},
                Context: cribl.String("unde"),
                DisallowAssign: cribl.Bool(false),
                MaxNumOfAllowedIterations: cribl.Int64(897543),
                PartialEval: &shared.PartialEvalRewrite{
                    FieldFilter: shared.PartialEvalFieldFilter{},
                    NullObj: "animi",
                },
                ReplaceIdentifiers: cribl.Bool(false),
                ReplaceLiterals: cribl.Bool(false),
                Unprotected: cribl.Bool(false),
            },
            OriginalExpression: "impedit",
            PartialExpression: "ipsam",
            ReferencedCriblExport: false,
            ReplaceIdentifiersExpression: "corporis",
            ReplaceLiteralExpression: "est",
        },
        MaxValues: cribl.Int64(621666),
        MetricNameFilter: cribl.String("esse"),
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
        Earliest: cribl.Int64(288570),
        FilterExpr: cribl.String("veritatis"),
        Latest: cribl.Int64(874400),
        MetricNameFilter: cribl.String("consectetur"),
        NumBuckets: cribl.Int64(112427),
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

