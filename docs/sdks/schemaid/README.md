# SchemaID

### Available Operations

* [Delete](#delete) - Delete Schema
* [Get](#get) - Get Schema by ID
* [Update](#update) - Update Schema

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
    res, err := s.SchemaID.Delete(ctx, operations.DeleteSchemaIDRequest{
        ID: "aae03f33-ca79-4fb9-9e40-32ba26fd368b",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteSchemaIDRequest](../../models/operations/deleteschemaidrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DeleteSchemaIDResponse](../../models/operations/deleteschemaidresponse.md), error**


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
    res, err := s.SchemaID.Get(ctx, operations.GetSchemaIDRequest{
        ID: "a9216bcb-4158-435c-b364-1723133edc04",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetSchemaIDRequest](../../models/operations/getschemaidrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetSchemaIDResponse](../../models/operations/getschemaidresponse.md), error**


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
    res, err := s.SchemaID.Update(ctx, operations.UpdateSchemaIDRequest{
        RequestBody: map[string]interface{}{
            "rerum": "quod",
            "nemo": "architecto",
        },
        ID: "63bbca49-227c-442c-a2c5-5350495c5dbb",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateSchemaIDRequest](../../models/operations/updateschemaidrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateSchemaIDResponse](../../models/operations/updateschemaidresponse.md), error**

