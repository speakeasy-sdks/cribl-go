# DatabaseConnections

### Available Operations

* [TestDatabaseConnection](#testdatabaseconnection) - Test a database connection given a type and connectionString
* [Delete](#delete) - Delete DatabaseConnectionConfig
* [Get](#get) - Get DatabaseConnectionConfig by ID
* [ListDatabaseConnection](#listdatabaseconnection) - Get a list of DatabaseConnection objects
* [Post](#post) - Create DatabaseConnectionConfig
* [Update](#update) - Update DatabaseConnectionConfig

## TestDatabaseConnection

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
    res, err := s.DatabaseConnections.TestDatabaseConnection(ctx, shared.DatabaseConnectionTest{
        AuthType: "earum",
        ConfigObj: cribl.String("soluta"),
        ConnectionString: cribl.String("hic"),
        ConnectionTimeout: cribl.Int64(848151),
        DatabaseType: "eaque",
        TextSecret: cribl.String("earum"),
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

**[*operations.TestDatabaseConnectionResponse](../../models/operations/testdatabaseconnectionresponse.md), error**


## Delete

Delete DatabaseConnectionConfig

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
    id := "perspiciatis"

    ctx := context.Background()
    res, err := s.DatabaseConnections.Delete(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


### Response

**[*operations.DeleteDatabaseConnectionConfigIDResponse](../../models/operations/deletedatabaseconnectionconfigidresponse.md), error**


## Get

Get DatabaseConnectionConfig by ID

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
    id := "maiores"

    ctx := context.Background()
    res, err := s.DatabaseConnections.Get(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


### Response

**[*operations.GetDatabaseConnectionConfigIDResponse](../../models/operations/getdatabaseconnectionconfigidresponse.md), error**


## ListDatabaseConnection

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
    databaseType := "debitis"

    ctx := context.Background()
    res, err := s.DatabaseConnections.ListDatabaseConnection(ctx, databaseType)
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

**[*operations.ListDatabaseConnectionResponse](../../models/operations/listdatabaseconnectionresponse.md), error**


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
    res, err := s.DatabaseConnections.Post(ctx, shared.DatabaseConnectionConfig{
        AuthType: "aliquid",
        ConfigObj: cribl.String("porro"),
        ConnectionString: cribl.String("suscipit"),
        ConnectionTimeout: cribl.Int64(211534),
        DatabaseType: &shared.DatabaseConnectionType{},
        Description: "fugit",
        ID: "ca3aed01-1799-4631-afde-04771778ff61",
        RequestTimeout: cribl.Int64(857125),
        Tags: cribl.String("doloremque"),
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


## Update

Update DatabaseConnectionConfig

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
    id := "dicta"
    databaseConnectionConfig := &shared.DatabaseConnectionConfig{
        AuthType: "odio",
        ConfigObj: cribl.String("tempora"),
        ConnectionString: cribl.String("esse"),
        ConnectionTimeout: cribl.Int64(403793),
        DatabaseType: &shared.DatabaseConnectionType{},
        Description: "consectetur",
        ID: "60a15db6-a660-4659-a1ad-eaab5851d6c6",
        RequestTimeout: cribl.Int64(281153),
        Tags: cribl.String("ad"),
    }

    ctx := context.Background()
    res, err := s.DatabaseConnections.Update(ctx, id, databaseConnectionConfig)
    if err != nil {
        log.Fatal(err)
    }

    if res.DatabaseConnectionConfigs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |
| `id`                                                                                | *string*                                                                            | :heavy_check_mark:                                                                  | Unique ID                                                                           |
| `databaseConnectionConfig`                                                          | [*shared.DatabaseConnectionConfig](../../models/shared/databaseconnectionconfig.md) | :heavy_minus_sign:                                                                  | DatabaseConnectionConfig object to be updated                                       |


### Response

**[*operations.UpdateDatabaseConnectionConfigIDResponse](../../models/operations/updatedatabaseconnectionconfigidresponse.md), error**

