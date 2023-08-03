# LoggerConfig

### Available Operations

* [Delete](#delete) - Delete LoggerConfig
* [Get](#get) - Get LoggerConfig by ID
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
    id := "ad"

    ctx := context.Background()
    res, err := s.LoggerConfig.Delete(ctx, id)
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
    id := "facere"

    ctx := context.Background()
    res, err := s.LoggerConfig.Get(ctx, id)
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
    id := "laborum"
    loggerConfig := &shared.LoggerConfig{
        Channels: []shared.LoggerEntry{
            shared.LoggerEntry{
                CanDelete: cribl.Bool(false),
                ID: "a4c506a8-aa94-4c02-a44c-f5e9d9a4578a",
                Level: "facere",
            },
            shared.LoggerEntry{
                CanDelete: cribl.Bool(false),
                ID: "c1ac600d-ec00-41ac-802e-2ec09ff8f0f8",
                Level: "dicta",
            },
            shared.LoggerEntry{
                CanDelete: cribl.Bool(false),
                ID: "6ff3477c-13e9-402c-9412-5b0960a66815",
                Level: "quae",
            },
            shared.LoggerEntry{
                CanDelete: cribl.Bool(false),
                ID: "a472af92-3c59-449f-83f3-50cf876ffb90",
                Level: "dicta",
            },
        },
        DefaultRedactFields: []string{
            "commodi",
            "eveniet",
            "porro",
            "tempore",
        },
        ID: "b4e243cf-789f-4faf-ada5-3e5ae6e0ac18",
        RedactFields: []string{
            "quisquam",
            "eos",
        },
        RedactLabel: "nobis",
    }

    ctx := context.Background()
    res, err := s.LoggerConfig.Update(ctx, id, loggerConfig)
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

