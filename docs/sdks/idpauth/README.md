# IDPAuth

### Available Operations

* [Get](#get) - Get IDP used for an authorization code callback

## Get

Get IDP used for an authorization code callback

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
    code := "mollitia"
    state := "nulla"

    ctx := context.Background()
    res, err := s.IDPAuth.Get(ctx, code, state)
    if err != nil {
        log.Fatal(err)
    }

    if res.Success != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `code`                                                | **string*                                             | :heavy_minus_sign:                                    | Authorization Code                                    |
| `state`                                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |


### Response

**[*operations.GetIDPAuthResponse](../../models/operations/getidpauthresponse.md), error**

