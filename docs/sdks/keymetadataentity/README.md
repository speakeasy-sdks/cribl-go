# KeyMetadataEntity

### Available Operations

* [Create](#create) - Create KeyMetadataEntity
* [Delete](#delete) - Delete KeyMetadataEntity
* [Get](#get) - Get KeyMetadataEntity by ID
* [Update](#update) - Update KeyMetadataEntity

## Create

Create KeyMetadataEntity

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
    res, err := s.KeyMetadataEntity.Create(ctx, shared.KeyMetadataEntity{
        Algorithm: shared.KeyMetadataEntityEncryptionAlgorithmAes256Gcm,
        CipherKey: cribl.String("accusamus"),
        Created: cribl.Int64(518990),
        Description: cribl.String("reiciendis"),
        Expires: cribl.Int64(66074),
        KeyID: "sint",
        Keyclass: 472414,
        Kms: shared.KeyMetadataEntityKMSForThisKeyLocal,
        PlainKey: cribl.String("esse"),
        UseIV: cribl.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.KeyMetadataEntities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.KeyMetadataEntity](../../models/shared/keymetadataentity.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.CreateKeyMetadataEntityResponse](../../models/operations/createkeymetadataentityresponse.md), error**


## Delete

Delete KeyMetadataEntity

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
    id := "iure"

    ctx := context.Background()
    res, err := s.KeyMetadataEntity.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.KeyMetadataEntities != nil {
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

**[*operations.DeleteKeyMetadataEntityResponse](../../models/operations/deletekeymetadataentityresponse.md), error**


## Get

Get KeyMetadataEntity by ID

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
    id := "odio"

    ctx := context.Background()
    res, err := s.KeyMetadataEntity.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.KeyMetadataEntities != nil {
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

**[*operations.GetKeyMetadataEntityResponse](../../models/operations/getkeymetadataentityresponse.md), error**


## Update

Update KeyMetadataEntity

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
    keyMetadataEntity := &shared.KeyMetadataEntity{
        Algorithm: shared.KeyMetadataEntityEncryptionAlgorithmAes256Gcm,
        CipherKey: cribl.String("vel"),
        Created: cribl.Int64(208683),
        Description: cribl.String("corporis"),
        Expires: cribl.Int64(375350),
        KeyID: "consequuntur",
        Keyclass: 641133,
        Kms: shared.KeyMetadataEntityKMSForThisKeyLocal,
        PlainKey: cribl.String("reprehenderit"),
        UseIV: cribl.Bool(false),
    }

    ctx := context.Background()
    res, err := s.KeyMetadataEntity.Update(ctx, id, keyMetadataEntity)
    if err != nil {
        log.Fatal(err)
    }

    if res.KeyMetadataEntities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |
| `id`                                                                  | *string*                                                              | :heavy_check_mark:                                                    | Unique ID                                                             |
| `keyMetadataEntity`                                                   | [*shared.KeyMetadataEntity](../../models/shared/keymetadataentity.md) | :heavy_minus_sign:                                                    | KeyMetadataEntity object to be updated                                |


### Response

**[*operations.UpdateKeyMetadataEntityResponse](../../models/operations/updatekeymetadataentityresponse.md), error**

