# BulletinMessage

### Available Operations

* [Create](#create) - Create BulletinMessage
* [Delete](#delete) - Delete BulletinMessage
* [Get](#get) - Get BulletinMessage by ID

## Create

Create BulletinMessage

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
    res, err := s.BulletinMessage.Create(ctx, shared.BulletinMessage{
        Group: cribl.String("error"),
        ID: "eee9526f-8d98-46e8-81ea-d4f0e1012563",
        Metadata: []shared.BulletinMessageMetadata{
            shared.BulletinMessageMetadata{},
            shared.BulletinMessageMetadata{},
            shared.BulletinMessageMetadata{},
            shared.BulletinMessageMetadata{},
        },
        Severity: shared.BulletinMessageSeverityError.ToPointer(),
        Text: "magnam",
        Time: cribl.Int64(906355),
        Title: cribl.String("Mr."),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BulletinMessage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [shared.BulletinMessage](../../models/shared/bulletinmessage.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |


### Response

**[*operations.CreateBulletinMessageResponse](../../models/operations/createbulletinmessageresponse.md), error**


## Delete

Delete BulletinMessage

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
    id := "occaecati"

    ctx := context.Background()
    res, err := s.BulletinMessage.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.BulletinMessage != nil {
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

**[*operations.DeleteBulletinMessageResponse](../../models/operations/deletebulletinmessageresponse.md), error**


## Get

Get BulletinMessage by ID

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
    id := "officiis"

    ctx := context.Background()
    res, err := s.BulletinMessage.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.BulletinMessage != nil {
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

**[*operations.GetBulletinMessageResponse](../../models/operations/getbulletinmessageresponse.md), error**

