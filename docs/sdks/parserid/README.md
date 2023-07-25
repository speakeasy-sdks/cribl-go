# ParserID

### Available Operations

* [Delete](#delete) - Delete Parser
* [Get](#get) - Get Parser by ID
* [Update](#update) - Update Parser

## Delete

Delete Parser

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
    res, err := s.ParserID.Delete(ctx, operations.DeleteParserIDRequest{
        ID: "a197f6de-9221-451f-a171-2099853e9f54",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ParserLibEntries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteParserIDRequest](../../models/operations/deleteparseridrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DeleteParserIDResponse](../../models/operations/deleteparseridresponse.md), error**


## Get

Get Parser by ID

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
    res, err := s.ParserID.Get(ctx, operations.GetParserIDRequest{
        ID: "3d854439-ee22-4446-8443-bc154188c2f5",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ParserLibEntries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetParserIDRequest](../../models/operations/getparseridrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetParserIDResponse](../../models/operations/getparseridresponse.md), error**


## Update

Update Parser

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
    res, err := s.ParserID.Update(ctx, operations.UpdateParserIDRequest{
        RequestBody: map[string]interface{}{
            "accusamus": "totam",
            "ipsam": "at",
        },
        ID: "a7832eab-d617-4c3b-8d51-a44bf01bad87",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ParserLibEntries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateParserIDRequest](../../models/operations/updateparseridrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateParserIDResponse](../../models/operations/updateparseridresponse.md), error**

