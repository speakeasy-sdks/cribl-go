# Edge

## Overview

Actions enabled in Edge mode

### Available Operations

* [GetDirectoryListing](#getdirectorylisting) - Get a directory listing of the given path
* [GetMetadata](#getmetadata) - Get the host's metadata structure
* [GetRunDetails](#getrundetails) - Get details of a process running on the edge host
* [ListBytes](#listbytes) - Get some number of bytes from the file at the given path
* [ListConfiguredCollectors](#listconfiguredcollectors) - Get list of configured collectors
* [ListEdgeHostFiles](#listedgehostfiles) - Get details about a file on the edge host.
* [ListLogFiles](#listlogfiles) - list log files
* [ListProcessRunningDetail](#listprocessrunningdetail) - Get a detailed list of processes running on the edge host

## GetDirectoryListing

Get a directory listing of the given path

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
    path := "praesentium"

    ctx := context.Background()
    res, err := s.Edge.GetDirectoryListing(ctx, path)
    if err != nil {
        log.Fatal(err)
    }

    if res.FilesystemEntries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `path`                                                | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


### Response

**[*operations.GetEdgeListingResponse](../../models/operations/getedgelistingresponse.md), error**


## GetMetadata

Get the host's metadata structure

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
    res, err := s.Edge.GetMetadata(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.EdgeMetadatas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetHostMetadataStructureResponse](../../models/operations/gethostmetadatastructureresponse.md), error**


## GetRunDetails

Get details of a process running on the edge host

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
    pid := "quisquam"

    ctx := context.Background()
    res, err := s.Edge.GetRunDetails(ctx, pid)
    if err != nil {
        log.Fatal(err)
    }

    if res.Processes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `pid`                                                 | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


### Response

**[*operations.GetProcessRunningDetailResponse](../../models/operations/getprocessrunningdetailresponse.md), error**


## ListBytes

Get some number of bytes from the file at the given path

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
    path := "veritatis"
    bytesRequested := 56848

    ctx := context.Background()
    res, err := s.Edge.ListBytes(ctx, path, bytesRequested)
    if err != nil {
        log.Fatal(err)
    }

    if res.SampleFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `path`                                                                             | *string*                                                                           | :heavy_check_mark:                                                                 | The path to the file to sample                                                     |
| `bytesRequested`                                                                   | **int64*                                                                           | :heavy_minus_sign:                                                                 | The number of bytes to return;   this value could be constrained by system limits. |


### Response

**[*operations.ListBytesResponse](../../models/operations/listbytesresponse.md), error**


## ListConfiguredCollectors

Get list of configured collectors

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
    res, err := s.Edge.ListConfiguredCollectors(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ConfiguredCollectors != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListConfiguredCollectorsResponse](../../models/operations/listconfiguredcollectorsresponse.md), error**


## ListEdgeHostFiles

Get details about a file on the edge host.

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
    res, err := s.Edge.ListEdgeHostFiles(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.EdgeHostFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListEdgeHostFilesResponse](../../models/operations/listedgehostfilesresponse.md), error**


## ListLogFiles

list log files

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
    allow := "id"
    depth := 696997
    mode := "neque"
    path := "quo"

    ctx := context.Background()
    res, err := s.Edge.ListLogFiles(ctx, allow, depth, mode, path)
    if err != nil {
        log.Fatal(err)
    }

    if res.EdgeFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |
| `allow`                                                                 | **string*                                                               | :heavy_minus_sign:                                                      | query array[string] optional List of allowed-filename wildcard patterns |
| `depth`                                                                 | **int64*                                                                | :heavy_minus_sign:                                                      | Search depth for "manual" mode                                          |
| `mode`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | Discovery Mode (default is "auto")                                      |
| `path`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | Search directory for "manual" mode                                      |


### Response

**[*operations.ListLogFileListResponse](../../models/operations/listlogfilelistresponse.md), error**


## ListProcessRunningDetail

Get a detailed list of processes running on the edge host

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
    res, err := s.Edge.ListProcessRunningDetail(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Processes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListProcessRunningDetailResponse](../../models/operations/listprocessrunningdetailresponse.md), error**

