# RequestUserAuth

### Available Operations

* [Logout](#logout) - API call that the IDP should use for a logout request

## Logout

API call that the IDP should use for a logout request

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
    res, err := s.RequestUserAuth.Logout(ctx, shared.LogoutRequest{
        RelayState: cribl.String("culpa"),
        SAMLResponse: cribl.String("reiciendis"),
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

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.LogoutRequest](../../models/shared/logoutrequest.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |


### Response

**[*operations.LogoutRequestUserAuthResponse](../../models/operations/logoutrequestuserauthresponse.md), error**

