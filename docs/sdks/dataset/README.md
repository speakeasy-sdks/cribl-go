# Dataset

### Available Operations

* [Create](#create) - Create DatasetProviderType
* [Delete](#delete) - Delete DatasetProviderType
* [Get](#get) - Get DatasetProviderType by ID
* [Update](#update) - Update DatasetProviderType

## Create

Create DatasetProviderType

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
    res, err := s.Dataset.Create(ctx, shared.DatasetProviderType{
        Description: cribl.String("architecto"),
        ID: shared.ProviderTypeMmHeap,
        Locality: &shared.OriginConfig{
            FilterExpression: cribl.String("enim"),
            Origin: shared.DatasetOrigin{},
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DatasetProviderType != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.DatasetProviderType](../../models/shared/datasetprovidertype.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.CreateDatasetProviderTypeResponse](../../models/operations/createdatasetprovidertyperesponse.md), error**


## Delete

Delete DatasetProviderType

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
    res, err := s.Dataset.Delete(ctx, operations.DeleteDatasetProviderTypeRequest{
        ID: "c80bff91-8544-4ec4-adef-cce8f1977773",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DatasetProviderType != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.DeleteDatasetProviderTypeRequest](../../models/operations/deletedatasetprovidertyperequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.DeleteDatasetProviderTypeResponse](../../models/operations/deletedatasetprovidertyperesponse.md), error**


## Get

Get DatasetProviderType by ID

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
    res, err := s.Dataset.Get(ctx, operations.GetDatasetProviderTypeRequest{
        ID: "e63562a7-b408-4f05-a3d4-8fdaf313a1f5",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DatasetProviderType != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetDatasetProviderTypeRequest](../../models/operations/getdatasetprovidertyperequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.GetDatasetProviderTypeResponse](../../models/operations/getdatasetprovidertyperesponse.md), error**


## Update

Update DatasetProviderType

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
    res, err := s.Dataset.Update(ctx, operations.UpdateDatasetProviderTypeRequest{
        DatasetProviderType: &shared.DatasetProviderType{
            Description: cribl.String("doloribus"),
            ID: shared.ProviderTypeInvalid,
            Locality: &shared.OriginConfig{
                FilterExpression: cribl.String("unde"),
                Origin: shared.DatasetOrigin{},
            },
        },
        ID: "4259c0b3-6f25-4ea9-84f3-b756c11f6c37",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DatasetProviderType != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateDatasetProviderTypeRequest](../../models/operations/updatedatasetprovidertyperequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.UpdateDatasetProviderTypeResponse](../../models/operations/updatedatasetprovidertyperesponse.md), error**

