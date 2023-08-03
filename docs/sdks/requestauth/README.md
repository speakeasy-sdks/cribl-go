# RequestAuth

### Available Operations

* [Get](#get) - Accepts an authentication request from an IDP and authenticates the user
* [Post](#post) - API call that the IDP should use for an authentication request

## Get

Accepts an authentication request from an IDP and authenticates the user

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
    relayState := "architecto"
    samlResponse := "a"

    ctx := context.Background()
    res, err := s.RequestAuth.Get(ctx, relayState, samlResponse)
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
| `samlResponse`                                        | **string*                                             | :heavy_minus_sign:                                    | Authentication request object                         |


### Response

**[*operations.GetRequestAuthResponse](../../models/operations/getrequestauthresponse.md), error**


## Post

API call that the IDP should use for an authentication request

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

    ctx := context.Background()
    res, err := s.RequestAuth.Post(ctx, operations.PostRequestAuthRequestBody{
        RelayState: cribl.String("laborum"),
        SAMLResponse: cribl.String("veritatis"),
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PostRequestAuthRequestBody](../../models/operations/postrequestauthrequestbody.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.PostRequestAuthResponse](../../models/operations/postrequestauthresponse.md), error**

