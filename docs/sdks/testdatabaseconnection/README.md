# TestDatabaseConnection

### Available Operations

* [Post](#post) - Test a database connection given a type and connectionString

## Post

Test a database connection given a type and connectionString

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
    res, err := s.TestDatabaseConnection.Post(ctx, shared.DatabaseConnectionTest{
        AuthType: "beatae",
        ConfigObj: cribl.String("quos"),
        ConnectionString: cribl.String("consectetur"),
        ConnectionTimeout: cribl.Int64(742010),
        DatabaseType: "tenetur",
        TextSecret: cribl.String("necessitatibus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DatabaseConnectionTestResults != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.DatabaseConnectionTest](../../models/shared/databaseconnectiontest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.PostTestDatabaseConnectionResponse](../../models/operations/posttestdatabaseconnectionresponse.md), error**

