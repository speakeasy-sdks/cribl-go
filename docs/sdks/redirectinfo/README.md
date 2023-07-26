# RedirectInfo

### Available Operations

* [Get](#get) - Obtain redirect information

## Get

Obtain information needed to redirect users to the configured SSO IDP for authentication.

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
    res, err := s.RedirectInfo.Get(ctx)
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

