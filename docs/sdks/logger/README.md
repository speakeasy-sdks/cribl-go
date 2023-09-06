# Logger

### Available Operations

* [Delete](#delete) - Delete LoggerConfig
* [Get](#get) - Get LoggerConfig by ID
* [ListLoggerConfigs](#listloggerconfigs) - Get a list of LoggerConfig objects
* [Update](#update) - Update LoggerConfig

## Delete

Delete LoggerConfig

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
    id := "velit"

    ctx := context.Background()
    res, err := s.Logger.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggerConfig != nil {
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

**[*operations.DeleteLoggerConfigResponse](../../models/operations/deleteloggerconfigresponse.md), error**


## Get

Get LoggerConfig by ID

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
    id := "aspernatur"

    ctx := context.Background()
    res, err := s.Logger.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggerConfig != nil {
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

**[*operations.GetLoggerConfigResponse](../../models/operations/getloggerconfigresponse.md), error**


## ListLoggerConfigs

Get a list of LoggerConfig objects

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
    res, err := s.Logger.ListLoggerConfigs(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggerConfigs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListLoggerConfigsResponse](../../models/operations/listloggerconfigsresponse.md), error**


## Update

Update LoggerConfig

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
    id := "eum"
    loggerConfig := &shared.LoggerConfig{
        Channels: []shared.LoggerEntry{
            shared.LoggerEntry{
                CanDelete: cribl.Bool(false),
                ID: "48dc2f61-5199-4ebf-90e9-fe6c632ca3ae",
                Level: "nulla",
            },
        },
        DefaultRedactFields: []string{
            "consequatur",
        },
        ID: "11799631-2fde-4047-b177-8ff61d017476",
        RedactFields: []string{
            "consectetur",
        },
        RedactLabel: "aliquid",
    }

    ctx := context.Background()
    res, err := s.Logger.Update(ctx, id, loggerConfig)
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggerConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |
| `id`                                                        | *string*                                                    | :heavy_check_mark:                                          | Unique ID                                                   |
| `loggerConfig`                                              | [*shared.LoggerConfig](../../models/shared/loggerconfig.md) | :heavy_minus_sign:                                          | LoggerConfig object to be updated                           |


### Response

**[*operations.UpdateLoggerConfigResponse](../../models/operations/updateloggerconfigresponse.md), error**

