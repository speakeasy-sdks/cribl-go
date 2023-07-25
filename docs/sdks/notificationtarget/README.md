# NotificationTarget

### Available Operations

* [Create](#create) - Create NotificationTarget
* [Delete](#delete) - Delete NotificationTarget
* [Get](#get) - Get NotificationTarget by ID
* [Update](#update) - Update NotificationTarget

## Create

Create NotificationTarget

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
    res, err := s.NotificationTarget.Create(ctx, shared.NotificationTarget{})
    if err != nil {
        log.Fatal(err)
    }

    if res.NotificationTargets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.NotificationTarget](../../models/shared/notificationtarget.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.CreateNotificationTargetResponse](../../models/operations/createnotificationtargetresponse.md), error**


## Delete

Delete NotificationTarget

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
    res, err := s.NotificationTarget.Delete(ctx, operations.DeletetNotificationTargetRequest{
        ID: "27e21f86-2657-4b36-bc6b-9f587ce525c6",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.NotificationTargets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.DeletetNotificationTargetRequest](../../models/operations/deletetnotificationtargetrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.DeletetNotificationTargetResponse](../../models/operations/deletetnotificationtargetresponse.md), error**


## Get

Get NotificationTarget by ID

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
    res, err := s.NotificationTarget.Get(ctx, operations.GetNotificationTargetRequest{
        ID: "7641a831-2e50-447b-8c21-ccb423abcdc9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.NotificationTargets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetNotificationTargetRequest](../../models/operations/getnotificationtargetrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetNotificationTargetResponse](../../models/operations/getnotificationtargetresponse.md), error**


## Update

Update NotificationTarget

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
    res, err := s.NotificationTarget.Update(ctx, operations.UpdatetNotificationTargetRequest{
        NotificationTarget: &shared.NotificationTarget{},
        ID: "1faabdd8-8e71-4f6c-8825-2d7771e7fd07",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.NotificationTargets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdatetNotificationTargetRequest](../../models/operations/updatetnotificationtargetrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.UpdatetNotificationTargetResponse](../../models/operations/updatetnotificationtargetresponse.md), error**

