# Lookups

## Overview

Actions related to lookups

### Available Operations

* [Create](#create) - Create LookupFile
* [Delete](#delete) - Delete LookupFile
* [Get](#get) - Get LookupFile by ID
* [ListLookups](#listlookups) - Get a list of LookupFile objects
* [Update](#update) - Update LookupFile
* [Upload](#upload) - Upload LookupFile

## Create

Create LookupFile

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
    res, err := s.Lookups.Create(ctx, operations.CreateLookupRequestBody1{
        Description: cribl.String("fugit"),
        FileInfo: &operations.CreateLookupRequestBody1FileInfo{
            Filename: "ab",
        },
        ID: "813d5208-ece7-4e25-bb66-8451c6c6e205",
        Size: cribl.Int64(895692),
        Tags: cribl.String("quasi"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LookupFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [interface{}](../../models//.md)                      | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateLookupResponse](../../models/operations/createlookupresponse.md), error**


## Delete

Delete LookupFile

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

    ctx := context.Background()
    res, err := s.Lookups.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.LookupFile != nil {
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

**[*operations.DeleteLookupResponse](../../models/operations/deletelookupresponse.md), error**


## Get

Get LookupFile by ID

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
    id := "at"

    ctx := context.Background()
    res, err := s.Lookups.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.LookupFile != nil {
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

**[*operations.GetLookupResponse](../../models/operations/getlookupresponse.md), error**


## ListLookups

Get a list of LookupFile objects

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
    res, err := s.Lookups.ListLookups(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.LookupFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListLookupsResponse](../../models/operations/listlookupsresponse.md), error**


## Update

Update LookupFile

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
    requestBody := operations.UpdateLookupRequestBody2{
        Content: cribl.String("harum"),
        Description: cribl.String("sequi"),
        ID: "fec9578a-6458-4427-ba84-18d162309fb0",
        Size: cribl.Int64(608989),
        Tags: cribl.String("eos"),
    }

    ctx := context.Background()
    res, err := s.Lookups.Update(ctx, id, requestBody)
    if err != nil {
        log.Fatal(err)
    }

    if res.LookupFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |
| `requestBody`                                         | *interface{}*                                         | :heavy_minus_sign:                                    | LookupFile object to be updated                       |


### Response

**[*operations.UpdateLookupResponse](../../models/operations/updatelookupresponse.md), error**


## Upload

Upload LookupFile

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
    filename := "occaecati"

    ctx := context.Background()
    res, err := s.Lookups.Upload(ctx, filename)
    if err != nil {
        log.Fatal(err)
    }

    if res.LookupFileInfoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `filename`                                            | **string*                                             | :heavy_minus_sign:                                    | query LookupFilenameSchema required Filename          |


### Response

**[*operations.UploadLookupResponse](../../models/operations/uploadlookupresponse.md), error**

