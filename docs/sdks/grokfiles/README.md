# Grokfiles

### Available Operations

* [Create](#create) - Create GrokFile
* [Delete](#delete) - Delete GrokFile
* [Get](#get) - Get GrokFile by ID
* [ListGrokFiles](#listgrokfiles) - Get a list of GrokFile objects
* [Update](#update) - Update GrokFile

## Create

Create GrokFile

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
    res, err := s.Grokfiles.Create(ctx, shared.GrokFile{
        Content: "ab",
        ID: "94a276b2-6916-4fe1-b08f-4294e3698f44",
        Size: 455444,
        Tags: cribl.String("reiciendis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GrokFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.GrokFile](../../models/shared/grokfile.md)    | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateGrokFileResponse](../../models/operations/creategrokfileresponse.md), error**


## Delete

Delete GrokFile

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
    id := "ex"

    ctx := context.Background()
    res, err := s.Grokfiles.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GrokFile != nil {
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

**[*operations.DeleteGrokFileResponse](../../models/operations/deletegrokfileresponse.md), error**


## Get

Get GrokFile by ID

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
    res, err := s.Grokfiles.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GrokFile != nil {
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

**[*operations.GetGrokFileResponse](../../models/operations/getgrokfileresponse.md), error**


## ListGrokFiles

Get a list of GrokFile objects

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
    res, err := s.Grokfiles.ListGrokFiles(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GrokFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListGrokFilesResponse](../../models/operations/listgrokfilesresponse.md), error**


## Update

Update GrokFile

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
    id := "non"
    grokFile := &shared.GrokFile{
        Content: "officiis",
        ID: "8b445e80-ca55-4efd-a0e4-57e1858b6a89",
        Size: 944708,
        Tags: cribl.String("expedita"),
    }

    ctx := context.Background()
    res, err := s.Grokfiles.Update(ctx, id, grokFile)
    if err != nil {
        log.Fatal(err)
    }

    if res.GrokFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |
| `grokFile`                                            | [*shared.GrokFile](../../models/shared/grokfile.md)   | :heavy_minus_sign:                                    | GrokFile object to be updated                         |


### Response

**[*operations.UpdateGrokFileResponse](../../models/operations/updategrokfileresponse.md), error**

