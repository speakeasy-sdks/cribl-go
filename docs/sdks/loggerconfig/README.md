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
        ID: "f1ad837a-e80c-41c1-9c95-ba998678fa3f",
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
        ID: "696991af-388c-4e03-a144-48c7977a0ef2",
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
                    ID: "536028ef-eef9-4341-92ed-7e253f4c157d",
                    Level: "voluptates",
                },
                shared.LoggerEntry{
                    CanDelete: cribl.Bool(false),
                    ID: "aa7170f4-45ac-4cf6-a7aa-f9bbad185fe4",
                    Level: "nesciunt",
                },
                shared.LoggerEntry{
                    CanDelete: cribl.Bool(false),
                    ID: "1d6bf5c8-38fb-4b8c-a0cb-67fc4b425e99",
                    Level: "debitis",
                },
                shared.LoggerEntry{
                    CanDelete: cribl.Bool(false),
                    ID: "6234c9f7-b79d-4feb-b7a5-c38d4baf91e5",
                    Level: "eaque",
                },
            },
            DefaultRedactFields: []string{
                "eveniet",
                "delectus",
            },
            ID: "890a54b4-75f1-46f5-ad38-5a3c4ac631b9",
            RedactFields: []string{
                "necessitatibus",
                "fugit",
                "autem",
            },
            RedactLabel: "optio",
        },
        ID: "ed8f9fdb-9410-4f63-bbf8-17837b01afdd",
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

