# Metrics

### Available Operations

* [Post](#post) - Enumerate all internal system metrics
* [Query](#query) - Query raw internal system metrics

## Post

Enumerate all internal system metrics

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
    res, err := s.Metrics.Post(ctx, shared.GetNamesOpts{
        DimKeyFilter: cribl.String("nulla"),
        DimValueFilter: cribl.String("quod"),
        Earliest: cribl.Int64(247457),
        FilterExpr: &shared.Expression{
            MaxCache: 293176,
            Cache: shared.Map{
                Op: shared.MapOp{},
            },
            DeclaredVariables: []string{
                "alias",
                "blanditiis",
            },
            IsSafe: false,
            ModifiedExpression: "modi",
            Opt: shared.ExpressionOptions{
                AllowInternal: cribl.Bool(false),
                Args: []string{
                    "a",
                },
                AstTraverseCallback: &shared.TraverseCallback{},
                Context: cribl.String("et"),
                DisallowAssign: cribl.Bool(false),
                MaxNumOfAllowedIterations: cribl.Int64(474267),
                PartialEval: &shared.PartialEvalRewrite{
                    FieldFilter: shared.PartialEvalFieldFilter{},
                    NullObj: "autem",
                },
                ReplaceIdentifiers: cribl.Bool(false),
                ReplaceLiterals: cribl.Bool(false),
                Unprotected: cribl.Bool(false),
            },
            OriginalExpression: "dolore",
            PartialExpression: "eius",
            ReferencedCriblExport: false,
            ReplaceIdentifiersExpression: "nostrum",
            ReplaceLiteralExpression: "ex",
        },
        MaxValues: cribl.Int64(229197),
        MetricNameFilter: cribl.String("voluptate"),
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
    res, err := s.Metrics.Query(ctx, operations.QueryMetricsRequest{
        Earliest: cribl.Int64(565304),
        FilterExpr: cribl.String("voluptatibus"),
        Latest: cribl.Int64(217167),
        MetricNameFilter: cribl.String("hic"),
        NumBuckets: cribl.Int64(695892),
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

