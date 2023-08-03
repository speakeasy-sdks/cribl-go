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
        Content: "nihil",
        ID: "411faf4b-7544-4e47-ae80-2857a5b40463",
        Size: 652125,
        Tags: cribl.String("dignissimos"),
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
    id := "fugiat"

    ctx := context.Background()
    res, err := s.GrokFile.Delete(ctx, id)
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
    id := "nostrum"

    ctx := context.Background()
    res, err := s.GrokFile.Get(ctx, id)
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
    id := "molestiae"
    grokFile := &shared.GrokFile{
        Content: "veniam",
        ID: "f1400e76-4ad7-4334-ac1b-781b36a08088",
        Size: 832239,
        Tags: cribl.String("veritatis"),
    }

    ctx := context.Background()
    res, err := s.GrokFile.Update(ctx, id, grokFile)
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

