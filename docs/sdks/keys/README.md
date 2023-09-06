# Keys

## Overview

Actions related to encryption keys

### Available Operations

* [Create](#create) - Create KeyMetadataEntity
* [Delete](#delete) - Delete KeyMetadataEntity
* [Get](#get) - Get KeyMetadataEntity by ID
* [ListKeyMetadataEntities](#listkeymetadataentities) - Get a list of KeyMetadataEntity objects
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
    res, err := s.Keys.Create(ctx, shared.KeyMetadataEntity{
        Algorithm: shared.KeyMetadataEntityEncryptionAlgorithmAes256Gcm,
        CipherKey: cribl.String("nihil"),
        Created: cribl.Int64(649078),
        Description: cribl.String("voluptas"),
        Expires: cribl.Int64(5189),
        KeyID: "maiores",
        Keyclass: 970222,
        Kms: shared.KeyMetadataEntityKMSForThisKeyLocal,
        PlainKey: cribl.String("dolores"),
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
    id := "id"

    ctx := context.Background()
    res, err := s.Keys.Delete(ctx, id)
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
    id := "minima"

    ctx := context.Background()
    res, err := s.Keys.Get(ctx, id)
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


## ListKeyMetadataEntities

Get a list of KeyMetadataEntity objects

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
    res, err := s.Keys.ListKeyMetadataEntities(ctx)
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


### Response

**[*operations.ListKeyMetadataEntitiesResponse](../../models/operations/listkeymetadataentitiesresponse.md), error**


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
    id := "dolore"
    keyMetadataEntity := &shared.KeyMetadataEntity{
        Algorithm: shared.KeyMetadataEntityEncryptionAlgorithmAes256Gcm,
        CipherKey: cribl.String("nesciunt"),
        Created: cribl.Int64(63207),
        Description: cribl.String("recusandae"),
        Expires: cribl.Int64(607249),
        KeyID: "quaerat",
        Keyclass: 477646,
        Kms: shared.KeyMetadataEntityKMSForThisKeyLocal,
        PlainKey: cribl.String("ex"),
        UseIV: cribl.Bool(false),
    }

    ctx := context.Background()
    res, err := s.Keys.Update(ctx, id, keyMetadataEntity)
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

