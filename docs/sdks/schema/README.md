# Schema

### Available Operations

* [Create](#create) - Create Schema
* [Delete](#delete) - Delete Schema
* [Get](#get) - Get Schema by ID
* [Post](#post) - Create Schema
* [Update](#update) - Update Schema

## Create

Create Schema

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
    res, err := s.Schema.Create(ctx, map[string]interface{}{
        "odit": "voluptatem",
        "voluptas": "libero",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SchemaLibEntry != nil {
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

**[*operations.CreateSchemaResponse](../../models/operations/createschemaresponse.md), error**


## Delete

Delete Schema

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
    res, err := s.Schema.Delete(ctx, operations.DeleteSchemaRequest{
        ID: "ef2b0a3e-42c1-4aa0-90e9-aac2e9135586",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SchemaLibEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteSchemaRequest](../../models/operations/deleteschemarequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.DeleteSchemaResponse](../../models/operations/deleteschemaresponse.md), error**


## Get

Get Schema by ID

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
    res, err := s.Schema.Get(ctx, operations.GetSchemaRequest{
        ID: "d18f9f97-a4bf-4ad2-bf7d-67ca84ad99b4",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SchemaLibEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetSchemaRequest](../../models/operations/getschemarequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetSchemaResponse](../../models/operations/getschemaresponse.md), error**


## Post

Create Schema

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
    res, err := s.Schema.Post(ctx, map[string]interface{}{
        "quibusdam": "voluptas",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SchemaLibEntries != nil {
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

**[*operations.PostSchemaResponse](../../models/operations/postschemaresponse.md), error**


## Update

Update Schema

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
    res, err := s.Schema.Update(ctx, operations.UpdateSchemaRequest{
        RequestBody: map[string]interface{}{
            "magni": "numquam",
        },
        ID: "3531870c-f68b-403a-9421-bd43d1f0cb0a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SchemaLibEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UpdateSchemaRequest](../../models/operations/updateschemarequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.UpdateSchemaResponse](../../models/operations/updateschemaresponse.md), error**

