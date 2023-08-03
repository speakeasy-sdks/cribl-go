# TrustPolicies

### Available Operations

* [ListTrustPolicies](#listtrustpolicies) - Get a list of TrustPolicy objects

## ListTrustPolicies

Get a list of TrustPolicy objects

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
    res, err := s.TrustPolicies.ListTrustPolicies(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.TrustPolicies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListTrustPoliciesResponse](../../models/operations/listtrustpoliciesresponse.md), error**

