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
        EndOffset: cribl.Int64(58870),
        Et: cribl.Int64(671384),
        Filter: cribl.String("sunt"),
        ID: "5db6a660-659a-41ad-aaab-5851d6c645b0",
        Limit: cribl.Int64(561577),
        Lt: cribl.Int64(737254),
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
        EndOffset: cribl.Int64(399660),
        Et: cribl.Int64(109784),
        Filter: cribl.String("voluptatum"),
        GroupID: "omnis",
        ID: "1baa0fe1-ade0-408e-af8c-5f350d8cdb5a",
        Limit: cribl.Int64(222864),
        Lt: cribl.Int64(307376),
        Offset: cribl.Int64(80532),
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
        Et: cribl.Int64(537279),
        Files: cribl.String("veritatis"),
        Filter: cribl.String("tempora"),
        GroupID: cribl.String("dolor"),
        Limit: cribl.Int64(8689),
        Lt: cribl.Int64(100014),
        Type: "sit",
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

