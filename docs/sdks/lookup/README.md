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
    res, err := s.Lookup.Create(ctx, operations.CreateLookupRequestBody2{
        Content: cribl.String("minus"),
        Description: cribl.String("quia"),
        ID: "47c88373-a40e-4194-af32-e55055756f5d",
        Size: cribl.Int64(366480),
        Tags: cribl.String("suscipit"),
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
    id := "pariatur"

    ctx := context.Background()
    res, err := s.Lookup.Delete(ctx, id)
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
    id := "sit"

    ctx := context.Background()
    res, err := s.Lookup.Get(ctx, id)
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
    id := "quidem"
    requestBody := operations.UpdateLookupRequestBody2{
        Content: cribl.String("perferendis"),
        Description: cribl.String("id"),
        ID: "f2dfe13d-b4f6-42cb-a3f8-941aebc0b80a",
        Size: cribl.Int64(396223),
        Tags: cribl.String("excepturi"),
    }

    ctx := context.Background()
    res, err := s.Lookup.Update(ctx, id, requestBody)
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
    filename := "magni"

    ctx := context.Background()
    res, err := s.Lookup.Upload(ctx, filename)
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

