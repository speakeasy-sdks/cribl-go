# ListDatabaseConnection

### Available Operations

* [Get](#get) - Get a list of DatabaseConnection objects

## Get

Get a list of DatabaseConnection objects

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
    res, err := s.ListDatabaseConnection.Get(ctx, operations.GetListDatabaseConnectionRequest{
        DatabaseType: cribl.String("sint"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DatabaseConnectionConfigs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetListDatabaseConnectionRequest](../../models/operations/getlistdatabaseconnectionrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetListDatabaseConnectionResponse](../../models/operations/getlistdatabaseconnectionresponse.md), error**

