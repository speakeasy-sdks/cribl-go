# InternalSystemMetrics

### Available Operations

* [Post](#post) - Aggregate raw internal system metrics

## Post

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
    res, err := s.InternalSystemMetrics.Post(ctx, shared.MetricsAggOpts{
        Aggs: shared.AggregationMgrOptions{
            Aggregations: []string{
                "ducimus",
                "libero",
            },
            Cumulative: false,
            FlushEventLimit: 562065,
            FlushMemLimit: 900676,
            Hostname: "eager-satire.name",
            IdleTimeLimitSeconds: cribl.Int64(849690),
            LagToleranceSeconds: cribl.Int64(717853),
            MetricsMode: false,
            Prefix: cribl.String("voluptatem"),
            PreserveSplitByStructure: cribl.Bool(false),
            SearchAggMode: cribl.String("non"),
            SplitBys: []string{
                "consequatur",
                "laudantium",
            },
            SufficientStatsOnly: false,
            TimeWindowSeconds: 831067,
        },
        AlwaysBounds: cribl.Bool(false),
        Metrics: &shared.MetricsStore{},
        Where: cribl.String("commodi"),
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

