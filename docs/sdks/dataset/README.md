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
    res, err := s.Dataset.Create(ctx, shared.DatasetProviderType{
        Description: cribl.String("nam"),
        ID: shared.ProviderTypeQuicksort,
        Locality: &shared.OriginConfig{
            FilterExpression: cribl.String("iusto"),
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
    id := "voluptate"

    ctx := context.Background()
    res, err := s.Dataset.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.DatasetProviderType != nil {
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

**[*operations.DeleteDatasetProviderTypeResponse](../../models/operations/deletedatasetprovidertyperesponse.md), error**


## Get

Get DatasetProviderType by ID

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
    id := "sequi"

    ctx := context.Background()
    res, err := s.Dataset.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.DatasetProviderType != nil {
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

**[*operations.GetDatasetProviderTypeResponse](../../models/operations/getdatasetprovidertyperesponse.md), error**


## Update

Update DatasetProviderType

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
    id := "dignissimos"
    datasetProviderType := &shared.DatasetProviderType{
        Description: cribl.String("neque"),
        ID: shared.ProviderTypeInvalid,
        Locality: &shared.OriginConfig{
            FilterExpression: cribl.String("deleniti"),
            Origin: shared.DatasetOrigin{},
        },
    }

    ctx := context.Background()
    res, err := s.Dataset.Update(ctx, id, datasetProviderType)
    if err != nil {
        log.Fatal(err)
    }

    if res.DatasetProviderType != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `id`                                                                      | *string*                                                                  | :heavy_check_mark:                                                        | Unique ID                                                                 |
| `datasetProviderType`                                                     | [*shared.DatasetProviderType](../../models/shared/datasetprovidertype.md) | :heavy_minus_sign:                                                        | DatasetProviderType object to be updated                                  |


### Response

**[*operations.UpdateDatasetProviderTypeResponse](../../models/operations/updatedatasetprovidertyperesponse.md), error**

