# GrokFile

### Available Operations

* [Create](#create) - Create GrokFile
* [Delete](#delete) - Delete GrokFile
* [Get](#get) - Get GrokFile by ID
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
    res, err := s.GrokFile.Create(ctx, shared.GrokFile{
        Content: "possimus",
        ID: "20e56248-fff6-439a-910a-bdcab6267669",
        Size: 385760,
        Tags: cribl.String("accusamus"),
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

    ctx := context.Background()
    res, err := s.GrokFile.Delete(ctx, operations.DeleteGrokFileRequest{
        ID: "1ec00221-b335-4d89-acb3-ecfda8d0c549",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteGrokFileRequest](../../models/operations/deletegrokfilerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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

    ctx := context.Background()
    res, err := s.GrokFile.Get(ctx, operations.GetGrokFileRequest{
        ID: "ef030049-78a6-41fa-9cf2-0688f77c1ffc",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetGrokFileRequest](../../models/operations/getgrokfilerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetGrokFileResponse](../../models/operations/getgrokfileresponse.md), error**


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

    ctx := context.Background()
    res, err := s.GrokFile.Update(ctx, operations.UpdateGrokFileRequest{
        GrokFile: &shared.GrokFile{
            Content: "voluptate",
            ID: "1dca163f-2a3c-480a-97ff-334cddf857a9",
            Size: 915647,
            Tags: cribl.String("eum"),
        },
        ID: "1876c6ab-21d2-49df-894d-6fecd7993900",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateGrokFileRequest](../../models/operations/updategrokfilerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateGrokFileResponse](../../models/operations/updategrokfileresponse.md), error**

