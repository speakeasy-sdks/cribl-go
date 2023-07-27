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
    databaseType := "ducimus"

    ctx := context.Background()
    res, err := s.ListDatabaseConnection.Get(ctx, databaseType)
    if err != nil {
        log.Fatal(err)
    }

    if res.DatabaseConnectionConfigs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `databaseType`                                        | **string*                                             | :heavy_minus_sign:                                    | type of database connection                           |


### Response

**[*operations.GetListDatabaseConnectionResponse](../../models/operations/getlistdatabaseconnectionresponse.md), error**

