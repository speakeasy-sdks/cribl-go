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
        CipherKey: cribl.String("saepe"),
        Created: cribl.Int64(981142),
        Description: cribl.String("accusantium"),
        Expires: cribl.Int64(150715),
        KeyID: "eos",
        Keyclass: 162251,
        Kms: shared.KeyMetadataEntityKMSForThisKeyLocal,
        PlainKey: cribl.String("quis"),
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
        ID: "194db554-10ad-4c66-9af9-0a26c7cdc981",
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
        ID: "f068981d-6bb3-43cf-aa34-8c31bf407ee4",
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
            CipherKey: cribl.String("eligendi"),
            Created: cribl.Int64(957899),
            Description: cribl.String("alias"),
            Expires: cribl.Int64(770467),
            KeyID: "numquam",
            Keyclass: 136203,
            Kms: shared.KeyMetadataEntityKMSForThisKeyLocal,
            PlainKey: cribl.String("nobis"),
            UseIV: cribl.Bool(false),
        },
        ID: "78f15626-398a-40dc-b663-24ccb06c8ca1",
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

