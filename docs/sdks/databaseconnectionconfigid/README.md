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
    id := "modi"

    ctx := context.Background()
    res, err := s.DatabaseConnectionConfigID.Delete(ctx, id)
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
    id := "nam"

    ctx := context.Background()
    res, err := s.DatabaseConnectionConfigID.Get(ctx, id)
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
    id := "vero"
    databaseConnectionConfig := &shared.DatabaseConnectionConfig{
        AuthType: "voluptatem",
        ConfigObj: cribl.String("ipsam"),
        ConnectionString: cribl.String("vel"),
        ConnectionTimeout: cribl.Int64(1383),
        DatabaseType: &shared.DatabaseConnectionType{},
        Description: "quasi",
        ID: "3f59da75-7a59-4ecf-af66-ef1caa3383c2",
        RequestTimeout: cribl.Int64(746585),
        Tags: cribl.String("repudiandae"),
    }

    ctx := context.Background()
    res, err := s.DatabaseConnectionConfigID.Update(ctx, id, databaseConnectionConfig)
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

