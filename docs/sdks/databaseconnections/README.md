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
        AuthType: "facilis",
        ConfigObj: cribl.String("vero"),
        ConnectionString: cribl.String("ducimus"),
        ConnectionTimeout: cribl.Int64(293020),
        DatabaseType: "quibusdam",
        TextSecret: cribl.String("illum"),
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
    id := "sequi"

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
    id := "natus"

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
    databaseType := "impedit"

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
        AuthType: "aut",
        ConfigObj: cribl.String("voluptatibus"),
        ConnectionString: cribl.String("exercitationem"),
        ConnectionTimeout: cribl.Int64(862310),
        DatabaseType: &shared.DatabaseConnectionType{},
        Description: "fugit",
        ID: "cff7c70a-4562-46d4-b681-3f16d9f5fce6",
        RequestTimeout: cribl.Int64(774048),
        Tags: cribl.String("corporis"),
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
    id := "veniam"
    databaseConnectionConfig := &shared.DatabaseConnectionConfig{
        AuthType: "aliquid",
        ConfigObj: cribl.String("inventore"),
        ConnectionString: cribl.String("magnam"),
        ConnectionTimeout: cribl.Int64(407241),
        DatabaseType: &shared.DatabaseConnectionType{},
        Description: "quo",
        ID: "3e250fb0-08c4-42e1-81aa-c366c8dd6b14",
        RequestTimeout: cribl.Int64(256139),
        Tags: cribl.String("explicabo"),
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

