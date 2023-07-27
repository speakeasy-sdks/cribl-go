# DatabaseConnection

### Available Operations

* [Post](#post) - Create DatabaseConnectionConfig

## Post

Create DatabaseConnectionConfig

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
    res, err := s.DatabaseConnection.Post(ctx, shared.DatabaseConnectionConfig{
        AuthType: "fuga",
        ConfigObj: cribl.String("laudantium"),
        ConnectionString: cribl.String("incidunt"),
        ConnectionTimeout: cribl.Int64(97493),
        DatabaseType: &shared.DatabaseConnectionType{},
        Description: "rem",
        ID: "d162309f-b092-4992-9aef-b9f58c4d86e6",
        RequestTimeout: cribl.Int64(520761),
        Tags: cribl.String("earum"),
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.DatabaseConnectionConfig](../../models/shared/databaseconnectionconfig.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.PostDatabaseConnectionResponse](../../models/operations/postdatabaseconnectionresponse.md), error**

