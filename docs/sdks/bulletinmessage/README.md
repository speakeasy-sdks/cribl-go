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
	"cribl"
	"cribl/pkg/models/shared"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.BulletinMessage.Create(ctx, shared.BulletinMessage{
        Group: cribl.String("provident"),
        ID: "dfe0ab7d-a8a5-40ce-987f-86bc173d689e",
        Metadata: []shared.BulletinMessageMetadata{
            shared.BulletinMessageMetadata{},
            shared.BulletinMessageMetadata{},
            shared.BulletinMessageMetadata{},
            shared.BulletinMessageMetadata{},
        },
        Severity: shared.BulletinMessageSeverityFatal.ToPointer(),
        Text: "natus",
        Time: cribl.Int64(328303),
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
    res, err := s.BulletinMessage.Delete(ctx, operations.DeleteBulletinMessageRequest{
        ID: "6f8d986e-881e-4ad4-b0e1-012563f94e29",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.DeleteBulletinMessageRequest](../../models/operations/deletebulletinmessagerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


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
    res, err := s.BulletinMessage.Get(ctx, operations.GetBulletinMessageRequest{
        ID: "e973e922-a57a-415b-a3e0-60807e2b6e3a",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetBulletinMessageRequest](../../models/operations/getbulletinmessagerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetBulletinMessageResponse](../../models/operations/getbulletinmessageresponse.md), error**

