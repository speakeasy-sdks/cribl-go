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
        "optio": "minima",
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
        ID: "7c1e4981-e8aa-4257-9dc1-912ebde64bfc",
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
        ID: "c5469d40-15df-4a79-a206-bef2b0a3e42c",
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
            "deserunt": "fuga",
        },
        ID: "010e9aac-2e91-4355-86d1-8f9f97a4bfad",
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

