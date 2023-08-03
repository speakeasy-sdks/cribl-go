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
    id := "nesciunt"

    ctx := context.Background()
    res, err := s.NotificationTarget.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.NotificationTargets != nil {
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

**[*operations.DeletetNotificationTargetResponse](../../models/operations/deletetnotificationtargetresponse.md), error**


## Get

Get NotificationTarget by ID

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
    id := "delectus"

    ctx := context.Background()
    res, err := s.NotificationTarget.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.NotificationTargets != nil {
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

**[*operations.GetNotificationTargetResponse](../../models/operations/getnotificationtargetresponse.md), error**


## Update

Update NotificationTarget

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
    id := "laborum"
    notificationTarget := shared.NotificationTarget{}

    ctx := context.Background()
    res, err := s.NotificationTarget.Update(ctx, id, notificationTarget)
    if err != nil {
        log.Fatal(err)
    }

    if res.NotificationTargets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |
| `id`                                                                    | *string*                                                                | :heavy_check_mark:                                                      | Unique ID                                                               |
| `notificationTarget`                                                    | [*shared.NotificationTarget](../../models/shared/notificationtarget.md) | :heavy_minus_sign:                                                      | NotificationTarget object to be updated                                 |


### Response

**[*operations.UpdatetNotificationTargetResponse](../../models/operations/updatetnotificationtargetresponse.md), error**

