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
    res, err := s.KeyMetadataEntity.Create(ctx, shared.KeyMetadataEntity{
        Algorithm: shared.KeyMetadataEntityEncryptionAlgorithmAes256Gcm,
        CipherKey: cribl.String("at"),
        Created: cribl.Int64(986594),
        Description: cribl.String("omnis"),
        Expires: cribl.Int64(463695),
        KeyID: "exercitationem",
        Keyclass: 915145,
        Kms: shared.KeyMetadataEntityKMSForThisKeyLocal,
        PlainKey: cribl.String("sequi"),
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
    res, err := s.KeyMetadataEntity.Delete(ctx, operations.DeleteKeyMetadataEntityRequest{
        ID: "56686092-e9c3-4ddc-9f11-1dea1026d541",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.DeleteKeyMetadataEntityRequest](../../models/operations/deletekeymetadataentityrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


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
    res, err := s.KeyMetadataEntity.Get(ctx, operations.GetKeyMetadataEntityRequest{
        ID: "a4d190fe-b217-480b-8cc0-dbbddb484708",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetKeyMetadataEntityRequest](../../models/operations/getkeymetadataentityrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


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
    res, err := s.KeyMetadataEntity.Update(ctx, operations.UpdateKeyMetadataEntityRequest{
        KeyMetadataEntity: &shared.KeyMetadataEntity{
            Algorithm: shared.KeyMetadataEntityEncryptionAlgorithmAes256Gcm,
            CipherKey: cribl.String("facilis"),
            Created: cribl.Int64(306023),
            Description: cribl.String("repudiandae"),
            Expires: cribl.Int64(227424),
            KeyID: "natus",
            Keyclass: 68253,
            Kms: shared.KeyMetadataEntityKMSForThisKeyLocal,
            PlainKey: cribl.String("officiis"),
            UseIV: cribl.Bool(false),
        },
        ID: "6bc158c4-c4e5-4459-9ea3-42260e9b200c",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateKeyMetadataEntityRequest](../../models/operations/updatekeymetadataentityrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.UpdateKeyMetadataEntityResponse](../../models/operations/updatekeymetadataentityresponse.md), error**

