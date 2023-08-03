# GlobalVariableID

### Available Operations

* [Delete](#delete) - Delete Global Variable
* [Get](#get) - Get Global Variable by ID
* [Update](#update) - Update Global Variable

## Delete

Delete Global Variable

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
    id := "atque"

    ctx := context.Background()
    res, err := s.GlobalVariableID.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GlobalVars != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


### Response

**[*operations.DeleteGlobalVariableIDResponse](../../models/operations/deleteglobalvariableidresponse.md), error**


## Get

Get Global Variable by ID

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
    id := "beatae"

    ctx := context.Background()
    res, err := s.GlobalVariableID.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GlobalVars != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


### Response

**[*operations.GetGlobalVariableIDResponse](../../models/operations/getglobalvariableidresponse.md), error**


## Update

Update Global Variable

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
    id := "at"
    globalVar := &shared.GlobalVar{
        Description: cribl.String("labore"),
        ID: "c700b607-f3c9-43c7-bb9d-a3f2ceda7e23",
        Lib: cribl.String("repellat"),
        Tags: cribl.String("explicabo"),
        Type: shared.GlobalVarTypeString,
        Value: "exercitationem",
    }

    ctx := context.Background()
    res, err := s.GlobalVariableID.Update(ctx, id, globalVar)
    if err != nil {
        log.Fatal(err)
    }

    if res.GlobalVars != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |
| `globalVar`                                           | [*shared.GlobalVar](../../models/shared/globalvar.md) | :heavy_minus_sign:                                    | Global Variable object to be updated                  |


### Response

**[*operations.UpdateGlobalVariableIDResponse](../../models/operations/updateglobalvariableidresponse.md), error**

