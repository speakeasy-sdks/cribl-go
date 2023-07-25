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
    res, err := s.LoggerConfig.Delete(ctx, operations.DeleteLoggerConfigRequest{
        ID: "faa348c3-1bf4-407e-a4fc-f0c42b78f156",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggerConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteLoggerConfigRequest](../../models/operations/deleteloggerconfigrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


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
    res, err := s.LoggerConfig.Get(ctx, operations.GetLoggerConfigRequest{
        ID: "26398a0d-c766-4324-8cb0-6c8ca12d0252",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggerConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetLoggerConfigRequest](../../models/operations/getloggerconfigrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


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
    res, err := s.LoggerConfig.Update(ctx, operations.UpdateLoggerConfigRequest{
        LoggerConfig: &shared.LoggerConfig{
            Channels: []shared.LoggerEntry{
                shared.LoggerEntry{
                    CanDelete: cribl.Bool(false),
                    ID: "270b8d57-22dd-4895-b8bc-f24db9596933",
                    Level: "nostrum",
                },
                shared.LoggerEntry{
                    CanDelete: cribl.Bool(false),
                    ID: "2f745339-94d7-48de-bb6e-9389f5abb7f6",
                    Level: "commodi",
                },
                shared.LoggerEntry{
                    CanDelete: cribl.Bool(false),
                    ID: "2550a283-82ac-4483-afd2-315bba650164",
                    Level: "eveniet",
                },
            },
            DefaultRedactFields: []string{
                "ea",
            },
            ID: "f5bf6ae5-91bc-48bd-af36-12b63c205fda",
            RedactFields: []string{
                "quaerat",
                "aperiam",
                "dignissimos",
            },
            RedactLabel: "quam",
        },
        ID: "4a68a9a3-5d08-46b6-b66f-ef020e9f443b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoggerConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateLoggerConfigRequest](../../models/operations/updateloggerconfigrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateLoggerConfigResponse](../../models/operations/updateloggerconfigresponse.md), error**

