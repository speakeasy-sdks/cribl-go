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
        DimKeyFilter: cribl.String("officiis"),
        DimValueFilter: cribl.String("mollitia"),
        Earliest: cribl.Int64(765070),
        FilterExpr: &shared.Expression{
            MaxCache: 337081,
            Cache: shared.Map{
                Op: shared.MapOp{},
            },
            DeclaredVariables: []string{
                "eum",
                "nemo",
            },
            IsSafe: false,
            ModifiedExpression: "illum",
            Opt: shared.ExpressionOptions{
                AllowInternal: cribl.Bool(false),
                Args: []string{
                    "sit",
                },
                AstTraverseCallback: &shared.TraverseCallback{},
                Context: cribl.String("odio"),
                DisallowAssign: cribl.Bool(false),
                MaxNumOfAllowedIterations: cribl.Int64(794306),
                PartialEval: &shared.PartialEvalRewrite{
                    FieldFilter: shared.PartialEvalFieldFilter{},
                    NullObj: "asperiores",
                },
                ReplaceIdentifiers: cribl.Bool(false),
                ReplaceLiterals: cribl.Bool(false),
                Unprotected: cribl.Bool(false),
            },
            OriginalExpression: "recusandae",
            PartialExpression: "voluptates",
            ReferencedCriblExport: false,
            ReplaceIdentifiersExpression: "praesentium",
            ReplaceLiteralExpression: "dicta",
        },
        MaxValues: cribl.Int64(144179),
        MetricNameFilter: cribl.String("sit"),
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
        Earliest: cribl.Int64(396234),
        FilterExpr: cribl.String("necessitatibus"),
        Latest: cribl.Int64(148975),
        MetricNameFilter: cribl.String("deleniti"),
        NumBuckets: cribl.Int64(122744),
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

