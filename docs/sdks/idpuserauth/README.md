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
    res, err := s.IDPUserAuth.Logout(ctx, operations.LogoutIDPUserAuthRequest{
        RelayState: cribl.String("molestias"),
        SAMLResponse: cribl.String("beatae"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Success != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.LogoutIDPUserAuthRequest](../../models/operations/logoutidpuserauthrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.LogoutIDPUserAuthResponse](../../models/operations/logoutidpuserauthresponse.md), error**

