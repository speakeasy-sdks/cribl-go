# GlobalVariables

### Available Operations

* [Delete](#delete) - Delete Global Variable
* [Get](#get) - Get Global Variable by ID
* [ListGlobalVariable](#listglobalvariable) - Get a list of Global Variable objects
* [Post](#post) - Create Global Variable
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
    id := "asperiores"

    ctx := context.Background()
    res, err := s.GlobalVariables.Delete(ctx, id)
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
    id := "veniam"

    ctx := context.Background()
    res, err := s.GlobalVariables.Get(ctx, id)
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


## ListGlobalVariable

Get a list of Global Variable objects

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
    res, err := s.GlobalVariables.ListGlobalVariable(ctx)
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


### Response

**[*operations.ListGlobalVariableResponse](../../models/operations/listglobalvariableresponse.md), error**


## Post

Create Global Variable

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
    res, err := s.GlobalVariables.Post(ctx, shared.GlobalVar{
        Description: cribl.String("consequuntur"),
        ID: "d82d3513-bb6f-448b-a56b-cdb35ff2e4b2",
        Lib: cribl.String("reprehenderit"),
        Tags: cribl.String("nostrum"),
        Type: shared.GlobalVarTypeNumber,
        Value: "iusto",
    })
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
| `request`                                             | [shared.GlobalVar](../../models/shared/globalvar.md)  | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostGlobalVariableResponse](../../models/operations/postglobalvariableresponse.md), error**


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
    id := "est"
    globalVar := &shared.GlobalVar{
        Description: cribl.String("rem"),
        ID: "cd9e7319-c177-4d52-9f77-b114eeb52ff7",
        Lib: cribl.String("praesentium"),
        Tags: cribl.String("nemo"),
        Type: shared.GlobalVarTypeAny,
        Value: "quisquam",
    }

    ctx := context.Background()
    res, err := s.GlobalVariables.Update(ctx, id, globalVar)
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

