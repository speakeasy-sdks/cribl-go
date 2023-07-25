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
        DimKeyFilter: cribl.String("odio"),
        DimValueFilter: cribl.String("omnis"),
        Earliest: cribl.Int64(887422),
        FilterExpr: &shared.Expression{
            MaxCache: 961576,
            Cache: shared.Map{
                Op: shared.MapOp{},
            },
            DeclaredVariables: []string{
                "sit",
            },
            IsSafe: false,
            ModifiedExpression: "velit",
            Opt: shared.ExpressionOptions{
                AllowInternal: cribl.Bool(false),
                Args: []string{
                    "nihil",
                    "neque",
                    "aspernatur",
                },
                AstTraverseCallback: &shared.TraverseCallback{},
                Context: cribl.String("eaque"),
                DisallowAssign: cribl.Bool(false),
                MaxNumOfAllowedIterations: cribl.Int64(360334),
                PartialEval: &shared.PartialEvalRewrite{
                    FieldFilter: shared.PartialEvalFieldFilter{},
                    NullObj: "cupiditate",
                },
                ReplaceIdentifiers: cribl.Bool(false),
                ReplaceLiterals: cribl.Bool(false),
                Unprotected: cribl.Bool(false),
            },
            OriginalExpression: "aut",
            PartialExpression: "impedit",
            ReferencedCriblExport: false,
            ReplaceIdentifiersExpression: "quod",
            ReplaceLiteralExpression: "quo",
        },
        MaxValues: cribl.Int64(100704),
        MetricNameFilter: cribl.String("voluptatem"),
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
        Earliest: cribl.Int64(595858),
        FilterExpr: cribl.String("ea"),
        Latest: cribl.Int64(262231),
        MetricNameFilter: cribl.String("aperiam"),
        NumBuckets: cribl.Int64(30015),
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

