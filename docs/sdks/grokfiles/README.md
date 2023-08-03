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
        Content: "sequi",
        ID: "7814d4c9-8e0c-42bb-89eb-75dad636c600",
        Size: 322829,
        Tags: cribl.String("quae"),
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
    id := "amet"

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
    id := "illum"

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
    id := "praesentium"
    grokFile := &shared.GrokFile{
        Content: "quidem",
        ID: "b31180f7-39ae-49e0-97eb-809e2810331f",
        Size: 223291,
        Tags: cribl.String("occaecati"),
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

