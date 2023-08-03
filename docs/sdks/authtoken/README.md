# AuthToken

### Available Operations

* [Login](#login) - Log in and obtain Auth token

## Login

This endpoint is unavailable on Cribl.Cloud. Instead, get the Bearer token as detailed here: https://docs.cribl.io/stream/api-tutorials/#criblcloud-free-tier

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
    res, err := s.AuthToken.Login(ctx, shared.LoginInfo{
        Password: "et",
        Username: "Jairo.Farrell53",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AuthToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.LoginInfo](../../models/shared/logininfo.md)  | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.LoginAuthTokenResponse](../../models/operations/loginauthtokenresponse.md), error**

