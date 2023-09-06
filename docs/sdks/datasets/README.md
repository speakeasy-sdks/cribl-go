# Datasets

### Available Operations

* [Create](#create) - Create Dataset
* [CreateProviderType](#createprovidertype) - Create DatasetProviderType
* [Delete](#delete) - Delete Dataset
* [DeleteProviderType](#deleteprovidertype) - Delete DatasetProviderType
* [Get](#get) - Get Dataset by ID
* [GetProviderType](#getprovidertype) - Get DatasetProviderType by ID
* [ListDatasetObjects](#listdatasetobjects) - Get a list of Dataset objects
* [ListProviderTypes](#listprovidertypes) - Get a list of DatasetProviderType objects
* [Update](#update) - Update Dataset
* [UpdateProviderType](#updateprovidertype) - Update DatasetProviderType

## Create

Create Dataset

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
    res, err := s.Datasets.Create(ctx, "provident")
    if err != nil {
        log.Fatal(err)
    }

    if res.Dataset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [interface{}](../../models//.md)                      | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateDatasetObjectResponse](../../models/operations/createdatasetobjectresponse.md), error**


## CreateProviderType

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
    res, err := s.Datasets.CreateProviderType(ctx, shared.DatasetProviderType{
        Description: cribl.String("ipsa"),
        ID: shared.ProviderTypeQuicksort,
        Locality: &shared.OriginConfig{
            FilterExpression: cribl.String("magnam"),
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

Delete Dataset

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
    res, err := s.Datasets.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Dataset != nil {
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

**[*operations.DeleteDatasetObjectResponse](../../models/operations/deletedatasetobjectresponse.md), error**


## DeleteProviderType

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
    id := "eius"

    ctx := context.Background()
    res, err := s.Datasets.DeleteProviderType(ctx, id)
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

Get Dataset by ID

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
    id := "esse"

    ctx := context.Background()
    res, err := s.Datasets.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Dataset != nil {
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

**[*operations.GetDatasetObjectResponse](../../models/operations/getdatasetobjectresponse.md), error**


## GetProviderType

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
    id := "esse"

    ctx := context.Background()
    res, err := s.Datasets.GetProviderType(ctx, id)
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


## ListDatasetObjects

Get a list of Dataset objects

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
    res, err := s.Datasets.ListDatasetObjects(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Datasets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListDatasetObjectsResponse](../../models/operations/listdatasetobjectsresponse.md), error**


## ListProviderTypes

Get a list of DatasetProviderType objects

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
    res, err := s.Datasets.ListProviderTypes(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.DatasetProviderTypes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListDatasetProviderTypesResponse](../../models/operations/listdatasetprovidertypesresponse.md), error**


## Update

Update Dataset

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
    id := "rem"
    requestBody := "fuga"

    ctx := context.Background()
    res, err := s.Datasets.Update(ctx, id, requestBody)
    if err != nil {
        log.Fatal(err)
    }

    if res.Dataset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |
| `requestBody`                                         | *interface{}*                                         | :heavy_minus_sign:                                    | Dataset object to be updated                          |


### Response

**[*operations.UpdateDatasetObjectResponse](../../models/operations/updatedatasetobjectresponse.md), error**


## UpdateProviderType

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
    id := "reprehenderit"
    datasetProviderType := &shared.DatasetProviderType{
        Description: cribl.String("quidem"),
        ID: shared.ProviderTypeInvalid,
        Locality: &shared.OriginConfig{
            FilterExpression: cribl.String("ut"),
            Origin: shared.DatasetOrigin{},
        },
    }

    ctx := context.Background()
    res, err := s.Datasets.UpdateProviderType(ctx, id, datasetProviderType)
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

