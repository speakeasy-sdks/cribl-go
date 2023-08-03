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
    id := "beatae"

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
    id := "hic"

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
    id := "nisi"
    loggerConfig := &shared.LoggerConfig{
        Channels: []shared.LoggerEntry{
            shared.LoggerEntry{
                CanDelete: cribl.Bool(false),
                ID: "37a51262-4383-45bb-805a-23a45cefc5fd",
                Level: "officiis",
            },
            shared.LoggerEntry{
                CanDelete: cribl.Bool(false),
                ID: "10a0ce21-69e5-4100-99c6-dc5e34762799",
                Level: "cum",
            },
            shared.LoggerEntry{
                CanDelete: cribl.Bool(false),
                ID: "fbbe6949-fb2b-4b4e-8ae6-c3d5db3adebd",
                Level: "ad",
            },
            shared.LoggerEntry{
                CanDelete: cribl.Bool(false),
                ID: "daea4c50-6a8a-4a94-8026-44cf5e9d9a45",
                Level: "esse",
            },
        },
        DefaultRedactFields: []string{
            "fuga",
            "facere",
            "impedit",
        },
        ID: "1ac600de-c001-4ac8-82e2-ec09ff8f0f81",
        RedactFields: []string{
            "earum",
            "doloribus",
        },
        RedactLabel: "velit",
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

