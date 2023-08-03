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
    res, err := s.Lookups.Create(ctx, operations.CreateLookupRequestBody2{
        Content: cribl.String("deserunt"),
        Description: cribl.String("corporis"),
        ID: "3e5ae6e0-ac18-44c2-b9c2-47c88373a40e",
        Size: cribl.Int64(102316),
        Tags: cribl.String("molestias"),
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
    id := "dolore"

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
    id := "sunt"

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
    id := "maiores"
    requestBody := operations.UpdateLookupRequestBody1{
        Description: cribl.String("odit"),
        FileInfo: &operations.UpdateLookupRequestBody1FileInfo{
            Filename: "earum",
        },
        ID: "55055756-f5d5-46d0-bd0a-f2dfe13db4f6",
        Size: cribl.Int64(127087),
        Tags: cribl.String("minus"),
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
    filename := "soluta"

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

