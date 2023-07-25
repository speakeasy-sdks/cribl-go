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
    res, err := s.InternalSystemMetrics.Post(ctx, shared.MetricsAggOpts{
        Aggs: shared.AggregationMgrOptions{
            Aggregations: []string{
                "dolorem",
            },
            Cumulative: false,
            FlushEventLimit: 696291,
            FlushMemLimit: 476614,
            Hostname: "uniform-buddy.biz",
            IdleTimeLimitSeconds: cribl.Int64(321422),
            LagToleranceSeconds: cribl.Int64(539886),
            MetricsMode: false,
            Prefix: cribl.String("ut"),
            PreserveSplitByStructure: cribl.Bool(false),
            SearchAggMode: cribl.String("asperiores"),
            SplitBys: []string{
                "deserunt",
                "itaque",
            },
            SufficientStatsOnly: false,
            TimeWindowSeconds: 93254,
        },
        AlwaysBounds: cribl.Bool(false),
        Metrics: &shared.MetricsStore{},
        Where: cribl.String("eos"),
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

