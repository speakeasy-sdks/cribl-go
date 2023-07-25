# LicenseUsageMetrics

### Available Operations

* [Get](#get) - Get license usage metrics, aggregated by day, up to last 90 days

## Get

Get license usage metrics, aggregated by day, up to last 90 days

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
    res, err := s.LicenseUsageMetrics.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.DailyMetrics != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetLicenseUsageMetricsResponse](../../models/operations/getlicenseusagemetricsresponse.md), error**

