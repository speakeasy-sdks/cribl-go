# Logging

## Overview

Actions related to logging

### Available Operations

* [Get](#get) - Get contents of the log file
* [ListLogFileContents](#listlogfilecontents) - Get contents of the log file
* [ListLogFiles](#listlogfiles) - Get a list of log files
* [ListLogFilesContents](#listlogfilescontents) - Get contents of the log file

## Get

Get contents of the log file

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

    ctx := context.Background()
    res, err := s.Logging.Get(ctx, operations.GetLogFileContentRequest{
        EndOffset: cribl.Int64(260588),
        Et: cribl.Int64(458212),
        Filter: cribl.String("in"),
        ID: "c13e902c-1412-45b0-960a-668151a472af",
        Limit: cribl.Int64(609537),
        Lt: cribl.Int64(151230),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LogFileContents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetLogFileContentRequest](../../models/operations/getlogfilecontentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetLogFileContentResponse](../../models/operations/getlogfilecontentresponse.md), error**


## ListLogFileContents

Get contents of the log file

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

    ctx := context.Background()
    res, err := s.Logging.ListLogFileContents(ctx, operations.ListLogFileContentsRequest{
        EndOffset: cribl.Int64(200950),
        Et: cribl.Int64(805463),
        Filter: cribl.String("quis"),
        GroupID: "cupiditate",
        ID: "49f83f35-0cf8-476f-bb90-1c6ecbb4e243",
        Limit: cribl.Int64(758194),
        Lt: cribl.Int64(992887),
        Offset: cribl.Int64(459875),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LogFileContents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListLogFileContentsRequest](../../models/operations/listlogfilecontentsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ListLogFileContentsResponse](../../models/operations/listlogfilecontentsresponse.md), error**


## ListLogFiles

Get a list of log files

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
    res, err := s.Logging.ListLogFiles(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.LogFilesInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListLogFilesResponse](../../models/operations/listlogfilesresponse.md), error**


## ListLogFilesContents

Get contents of the log file

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

    ctx := context.Background()
    res, err := s.Logging.ListLogFilesContents(ctx, operations.ListLogFilesContentsRequest{
        Et: cribl.Int64(500021),
        Files: cribl.String("sint"),
        Filter: cribl.String("repellat"),
        GroupID: cribl.String("a"),
        Limit: cribl.Int64(658604),
        Lt: cribl.Int64(979287),
        Type: "itaque",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LogFileContents != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListLogFilesContentsRequest](../../models/operations/listlogfilescontentsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.ListLogFilesContentsResponse](../../models/operations/listlogfilescontentsresponse.md), error**

