# AggregationMgrOptions


## Fields

| Field                      | Type                       | Required                   | Description                |
| -------------------------- | -------------------------- | -------------------------- | -------------------------- |
| `Aggregations`             | []*string*                 | :heavy_check_mark:         | N/A                        |
| `Cumulative`               | *bool*                     | :heavy_check_mark:         | N/A                        |
| `FlushEventLimit`          | *int64*                    | :heavy_check_mark:         | N/A                        |
| `FlushMemLimit`            | *int64*                    | :heavy_check_mark:         | N/A                        |
| `Hostname`                 | *string*                   | :heavy_check_mark:         | N/A                        |
| `IdleTimeLimitSeconds`     | **int64*                   | :heavy_minus_sign:         | N/A                        |
| `LagToleranceSeconds`      | **int64*                   | :heavy_minus_sign:         | N/A                        |
| `MetricsMode`              | *bool*                     | :heavy_check_mark:         | N/A                        |
| `Prefix`                   | **string*                  | :heavy_minus_sign:         | N/A                        |
| `PreserveSplitByStructure` | **bool*                    | :heavy_minus_sign:         | N/A                        |
| `SearchAggMode`            | *interface{}*              | :heavy_minus_sign:         | N/A                        |
| `SplitBys`                 | []*string*                 | :heavy_minus_sign:         | N/A                        |
| `SufficientStatsOnly`      | *bool*                     | :heavy_check_mark:         | N/A                        |
| `TimeWindowSeconds`        | *int64*                    | :heavy_check_mark:         | N/A                        |