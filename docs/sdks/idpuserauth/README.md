# IDPUserAuth

### Available Operations

* [Logout](#logout) - Accepts a logout request from an IDP and logs out the user

## Logout

Accepts a logout request from an IDP and logs out the user

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
    relayState := "officia"
    samlResponse := "sed"

    ctx := context.Background()
    res, err := s.IDPUserAuth.Logout(ctx, relayState, samlResponse)
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
| `relayState`                                          | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `samlResponse`                                        | **string*                                             | :heavy_minus_sign:                                    | Logout request object                                 |


### Response

**[*operations.LogoutIDPUserAuthResponse](../../models/operations/logoutidpuserauthresponse.md), error**

