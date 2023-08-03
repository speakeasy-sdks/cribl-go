# Auth

## Overview

Actions related to authentication. Do not use the /auth endpoints in Cribl.Cloud deployments. Instead, obtain a Bearer token as described here: https://docs.cribl.io/stream/api-tutorials/#criblcloud-free-tier

### Available Operations

* [IDPlogout](#idplogout) - Accepts a logout request from an IDP and logs out the user
* [AcceptIDPRequest](#acceptidprequest) - Accepts an authentication request from an IDP and authenticates the user
* [Get](#get) - Get IDP used for an authorization code callback
* [GetMetadata](#getmetadata) - Obtain metadata which Cribl Stream/Edge uses when acting as a Service Provider
* [GetRedirect](#getredirect) - Obtain redirect information
* [ListAuthGroup](#listauthgroup) - List the external authentication system's groups
* [Login](#login) - Log in and obtain Auth token
* [Logout](#logout) - Log current user out
* [LogoutRedirect](#logoutredirect) - Redirect user to IDP with logout request
* [Post](#post) - API call that the IDP should use for an authentication request
* [RequestLogout](#requestlogout) - API call that the IDP should use for a logout request

## IDPlogout

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
    relayState := "qui"
    samlResponse := "cum"

    ctx := context.Background()
    res, err := s.Auth.IDPlogout(ctx, relayState, samlResponse)
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


## AcceptIDPRequest

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
    relayState := "iure"
    samlResponse := "necessitatibus"

    ctx := context.Background()
    res, err := s.Auth.AcceptIDPRequest(ctx, relayState, samlResponse)
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
    code := "ratione"
    state := "laborum"

    ctx := context.Background()
    res, err := s.Auth.Get(ctx, code, state)
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


## GetMetadata

Obtain metadata which Cribl Stream/Edge uses when acting as a Service Provider

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
    res, err := s.Auth.GetMetadata(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCriblMetadata200TextXMLString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetCriblMetadataResponse](../../models/operations/getcriblmetadataresponse.md), error**


## GetRedirect

Obtain information needed to redirect users to the configured SSO IDP for authentication.

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
    res, err := s.Auth.GetRedirect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.RedirectInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetRedirectInfoResponse](../../models/operations/getredirectinforesponse.md), error**


## ListAuthGroup

List the external authentication system's groups

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
    res, err := s.Auth.ListAuthGroup(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.CrudEntityBases != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListAuthGroupResponse](../../models/operations/listauthgroupresponse.md), error**


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
    res, err := s.Auth.Login(ctx, shared.LoginInfo{
        Password: "distinctio",
        Username: "Junior32",
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


## Logout

Log current user out

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
    res, err := s.Auth.Logout(ctx)
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


### Response

**[*operations.LogoutUserAuthResponse](../../models/operations/logoutuserauthresponse.md), error**


## LogoutRedirect

Redirect user to IDP with logout request

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
    res, err := s.Auth.LogoutRedirect(ctx)
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


### Response

**[*operations.LogoutRedirectUserAuthResponse](../../models/operations/logoutredirectuserauthresponse.md), error**


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
    res, err := s.Auth.Post(ctx, operations.PostRequestAuthRequestBody{
        RelayState: cribl.String("repellat"),
        SAMLResponse: cribl.String("alias"),
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


## RequestLogout

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
    res, err := s.Auth.RequestLogout(ctx, shared.LogoutRequest{
        RelayState: cribl.String("corporis"),
        SAMLResponse: cribl.String("perspiciatis"),
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

