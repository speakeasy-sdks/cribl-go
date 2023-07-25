# UIState

### Available Operations

* [Get](#get) - Get UI state by key
* [Update](#update) - Update UI state by key

## Get

Get UI state by key

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
    res, err := s.UIState.Get(ctx, operations.GetUIStateRequest{
        Key: "pariatur",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UIStates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetUIStateRequest](../../models/operations/getuistaterequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetUIStateResponse](../../models/operations/getuistateresponse.md), error**


## Update

Update UI state by key

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
    res, err := s.UIState.Update(ctx, operations.UpdateUIStateRequest{
        UIStatePatch: &shared.UIStatePatch{
            Args: map[string]interface{}{
                "quod": "minus",
            },
            Op: shared.UIStatePatchOpPushRecent,
            Value: "dolor",
        },
        Key: "amet",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UIStates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateUIStateRequest](../../models/operations/updateuistaterequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateUIStateResponse](../../models/operations/updateuistateresponse.md), error**

