# Messages

## Overview

Actions related to messages

### Available Operations

* [Create](#create) - Create BulletinMessage
* [Delete](#delete) - Delete BulletinMessage
* [Get](#get) - Get BulletinMessage by ID
* [ListBulletinMessages](#listbulletinmessages) - Get a list of BulletinMessage objects

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
    res, err := s.Messages.Create(ctx, shared.BulletinMessage{
        Group: cribl.String("nihil"),
        ID: "5a60a04c-495c-4c69-9171-b51c1bdb1cf4",
        Metadata: []shared.BulletinMessageMetadata{
            shared.BulletinMessageMetadata{},
            shared.BulletinMessageMetadata{},
            shared.BulletinMessageMetadata{},
        },
        Severity: shared.BulletinMessageSeverityError.ToPointer(),
        Text: "rem",
        Time: cribl.Int64(543353),
        Title: cribl.String("Dr."),
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
    id := "cum"

    ctx := context.Background()
    res, err := s.Messages.Delete(ctx, id)
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
    id := "pariatur"

    ctx := context.Background()
    res, err := s.Messages.Get(ctx, id)
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


## ListBulletinMessages

Get a list of BulletinMessage objects

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
    res, err := s.Messages.ListBulletinMessages(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.BulletinMessages != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListBulletinMessagesResponse](../../models/operations/listbulletinmessagesresponse.md), error**

