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
    res, err := s.GlobalVariableID.Delete(ctx, operations.DeleteGlobalVariableIDRequest{
        ID: "7cbdb6ee-c743-478b-a253-17747dc915ad",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeleteGlobalVariableIDRequest](../../models/operations/deleteglobalvariableidrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


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
    res, err := s.GlobalVariableID.Get(ctx, operations.GetGlobalVariableIDRequest{
        ID: "2caf5dd6-723d-4c0f-9ae2-f3a6b7008787",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetGlobalVariableIDRequest](../../models/operations/getglobalvariableidrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


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
    res, err := s.GlobalVariableID.Update(ctx, operations.UpdateGlobalVariableIDRequest{
        GlobalVar: &shared.GlobalVar{
            Description: cribl.String("quis"),
            ID: "6143f5a6-c98b-4555-9408-0d40bcacc6cb",
            Lib: cribl.String("fugiat"),
            Tags: cribl.String("laboriosam"),
            Type: shared.GlobalVarTypeExpression,
            Value: "enim",
        },
        ID: "f3ec9093-04f9-426b-ad25-53819b474b0e",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateGlobalVariableIDRequest](../../models/operations/updateglobalvariableidrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.UpdateGlobalVariableIDResponse](../../models/operations/updateglobalvariableidresponse.md), error**

