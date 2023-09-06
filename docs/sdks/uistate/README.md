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
    key := "nesciunt"

    ctx := context.Background()
    res, err := s.UIState.Get(ctx, key)
    if err != nil {
        log.Fatal(err)
    }

    if res.UIStates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `key`                                                 | *string*                                              | :heavy_check_mark:                                    | UI state key                                          |


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
    key := "sit"
    uiStatePatch := &shared.UIStatePatch{
        Args: map[string]interface{}{
            "odio": "minus",
        },
        Op: shared.UIStatePatchOpPushRecent,
        Value: "recusandae",
    }

    ctx := context.Background()
    res, err := s.UIState.Update(ctx, key, uiStatePatch)
    if err != nil {
        log.Fatal(err)
    }

    if res.UIStates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |
| `key`                                                       | *string*                                                    | :heavy_check_mark:                                          | UI state key                                                |
| `uiStatePatch`                                              | [*shared.UIStatePatch](../../models/shared/uistatepatch.md) | :heavy_minus_sign:                                          | UI State Patch object                                       |


### Response

**[*operations.UpdateUIStateResponse](../../models/operations/updateuistateresponse.md), error**

