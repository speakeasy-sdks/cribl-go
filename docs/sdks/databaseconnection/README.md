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
        AuthType: "a",
        ConfigObj: cribl.String("dolor"),
        ConnectionString: cribl.String("occaecati"),
        ConnectionTimeout: cribl.Int64(539813),
        DatabaseType: &shared.DatabaseConnectionType{},
        Description: "beatae",
        ID: "d4c700b6-07f3-4c93-873b-9da3f2ceda7e",
        RequestTimeout: cribl.Int64(185348),
        Tags: cribl.String("consectetur"),
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

