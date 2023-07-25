# Script

### Available Operations

* [Create](#create) - Create Script
* [Delete](#delete) - Delete Script
* [Get](#get) - Get Script by ID
* [Update](#update) - Update Script

## Create

Create Script

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
    res, err := s.Script.Create(ctx, map[string]interface{}{
        "voluptatem": "quod",
        "vitae": "vel",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ScriptLibEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [map[string]interface{}](../../models//.md)           | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateScriptResponse](../../models/operations/createscriptresponse.md), error**


## Delete

Delete Script

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
    res, err := s.Script.Delete(ctx, operations.DeleteScriptRequest{
        ID: "661a1d91-36a7-4e8d-9321-3f3f658752db",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ScriptLibEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteScriptRequest](../../models/operations/deletescriptrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.DeleteScriptResponse](../../models/operations/deletescriptresponse.md), error**


## Get

Get Script by ID

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
    res, err := s.Script.Get(ctx, operations.GetScriptRequest{
        ID: "764c59f0-a56c-4ebc-ada2-9ca79181c956",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ScriptLibEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetScriptRequest](../../models/operations/getscriptrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetScriptResponse](../../models/operations/getscriptresponse.md), error**


## Update

Update Script

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
    res, err := s.Script.Update(ctx, operations.UpdateScriptRequest{
        RequestBody: map[string]interface{}{
            "ab": "ex",
            "iure": "dolorem",
        },
        ID: "c530b566-5163-4a36-b851-2ab2521b9f2e",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ScriptLibEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UpdateScriptRequest](../../models/operations/updatescriptrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.UpdateScriptResponse](../../models/operations/updatescriptresponse.md), error**

