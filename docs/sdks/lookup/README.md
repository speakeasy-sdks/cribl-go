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
        Description: cribl.String("deleniti"),
        FileInfo: &operations.CreateLookupRequestBody1FileInfo{
            Filename: "rem",
        },
        ID: "624189eb-4487-43f5-833f-19dbf125ce41",
        Size: cribl.Int64(328744),
        Tags: cribl.String("magni"),
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
        ID: "eab9cd7e-5224-4a6a-8e12-3b7847ec59e1",
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
        ID: "f67f3c4c-ce4b-46d7-a96f-f3c574750135",
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
        RequestBody: &operations.UpdateLookupRequestBody1{
            Description: cribl.String("officiis"),
            FileInfo: &operations.UpdateLookupRequestBody1FileInfo{
                Filename: "dolore",
            },
            ID: "4f51f8b0-84c3-4197-a193-a245467f9487",
            Size: cribl.Int64(260046),
            Tags: cribl.String("placeat"),
        },
        ID: "2d5cc497-2233-4e66-bd8f-e5d00b979ef2",
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
        Filename: cribl.String("sit"),
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

