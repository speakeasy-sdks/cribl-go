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
        ID: "0003eb22-d9b3-4a70-994f-aa741c57d1fe",
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
        ID: "dc2050d3-8dc3-4ce1-8547-2f9ee69166a8",
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
            "saepe": "adipisci",
            "dolore": "tempora",
            "quaerat": "debitis",
        },
        ID: "ac8b3a28-75c6-4c1f-a606-d07d2a9c87ae",
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

