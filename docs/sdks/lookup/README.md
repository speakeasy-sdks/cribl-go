# Lookup

### Available Operations

* [Create](#create) - Create LookupFile
* [Delete](#delete) - Delete LookupFile
* [Get](#get) - Get LookupFile by ID
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
    res, err := s.Lookup.Create(ctx, operations.CreateLookupRequestBody1{
        Description: cribl.String("dolores"),
        FileInfo: &operations.CreateLookupRequestBody1FileInfo{
            Filename: "enim",
        },
        ID: "7b992c8d-bda6-4a61-afa2-198258fd0a9e",
        Size: cribl.Int64(723623),
        Tags: cribl.String("animi"),
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
    res, err := s.Lookup.Delete(ctx, operations.DeleteLookupRequest{
        ID: "47f7d3ef-0496-440d-aa18-31c87adf596f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LookupFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteLookupRequest](../../models/operations/deletelookuprequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


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
    res, err := s.Lookup.Get(ctx, operations.GetLookupRequest{
        ID: "df1ad837-ae80-4c1c-99c9-5ba998678fa3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LookupFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetLookupRequest](../../models/operations/getlookuprequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetLookupResponse](../../models/operations/getlookupresponse.md), error**


## Update

Update LookupFile

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
    res, err := s.Lookup.Update(ctx, operations.UpdateLookupRequest{
        RequestBody: &operations.UpdateLookupRequestBody2{
            Content: cribl.String("iure"),
            Description: cribl.String("sint"),
            ID: "6991af38-8ce0-4361-8448-c7977a0ef2f5",
            Size: cribl.Int64(217552),
            Tags: cribl.String("laboriosam"),
        },
        ID: "028efeef-9341-452e-97e2-53f4c157deaa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LookupFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UpdateLookupRequest](../../models/operations/updatelookuprequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


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
    res, err := s.Lookup.Upload(ctx, operations.UploadLookupRequest{
        Filename: cribl.String("in"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LookupFileInfoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UploadLookupRequest](../../models/operations/uploadlookuprequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.UploadLookupResponse](../../models/operations/uploadlookupresponse.md), error**

