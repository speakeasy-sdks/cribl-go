# DatabaseConnectionConfigID

### Available Operations

* [Delete](#delete) - Delete DatabaseConnectionConfig
* [Get](#get) - Get DatabaseConnectionConfig by ID
* [Update](#update) - Update DatabaseConnectionConfig

## Delete

Delete DatabaseConnectionConfig

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
    res, err := s.DatabaseConnectionConfigID.Delete(ctx, operations.DeleteDatabaseConnectionConfigIDRequest{
        ID: "f2257411-faf4-4b75-84e4-72e802857a5b",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.DeleteDatabaseConnectionConfigIDRequest](../../models/operations/deletedatabaseconnectionconfigidrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


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
    res, err := s.DatabaseConnectionConfigID.Get(ctx, operations.GetDatabaseConnectionConfigIDRequest{
        ID: "40463a7d-575f-4140-8e76-4ad7334ec1b7",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetDatabaseConnectionConfigIDRequest](../../models/operations/getdatabaseconnectionconfigidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetDatabaseConnectionConfigIDResponse](../../models/operations/getdatabaseconnectionconfigidresponse.md), error**


## Update

Update DatabaseConnectionConfig

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
    res, err := s.DatabaseConnectionConfigID.Update(ctx, operations.UpdateDatabaseConnectionConfigIDRequest{
        DatabaseConnectionConfig: &shared.DatabaseConnectionConfig{
            AuthType: "quas",
            ConfigObj: cribl.String("et"),
            ConnectionString: cribl.String("facilis"),
            ConnectionTimeout: cribl.Int64(229276),
            DatabaseType: &shared.DatabaseConnectionType{},
            Description: "autem",
            ID: "a08088d1-00ef-4ada-a00e-f0422eb2164c",
            RequestTimeout: cribl.Int64(975095),
            Tags: cribl.String("molestias"),
        },
        ID: "ab8366c7-23ff-4da9-a06b-ee4825c1fc0e",
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.UpdateDatabaseConnectionConfigIDRequest](../../models/operations/updatedatabaseconnectionconfigidrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.UpdateDatabaseConnectionConfigIDResponse](../../models/operations/updatedatabaseconnectionconfigidresponse.md), error**

