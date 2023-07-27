# DatasetObject

### Available Operations

* [Create](#create) - Create Dataset
* [Delete](#delete) - Delete Dataset
* [Get](#get) - Get Dataset by ID
* [Update](#update) - Update Dataset

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
    res, err := s.DatasetObject.Create(ctx, "fuga")
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

    ctx := context.Background()
    res, err := s.DatasetObject.Delete(ctx, operations.DeleteDatasetObjectRequest{
        ID: "51262438-35bb-4c05-a23a-45cefc5fde10",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Dataset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteDatasetObjectRequest](../../models/operations/deletedatasetobjectrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.DeleteDatasetObjectResponse](../../models/operations/deletedatasetobjectresponse.md), error**


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

    ctx := context.Background()
    res, err := s.DatasetObject.Get(ctx, operations.GetDatasetObjectRequest{
        ID: "a0ce2169-e510-4019-86dc-5e34762799bf",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Dataset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetDatasetObjectRequest](../../models/operations/getdatasetobjectrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetDatasetObjectResponse](../../models/operations/getdatasetobjectresponse.md), error**


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

    ctx := context.Background()
    res, err := s.DatasetObject.Update(ctx, operations.UpdateDatasetObjectRequest{
        RequestBody: cribl.String("facilis"),
        ID: "be6949fb-2bb4-4eca-a6c3-d5db3adebd5d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Dataset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateDatasetObjectRequest](../../models/operations/updatedatasetobjectrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UpdateDatasetObjectResponse](../../models/operations/updatedatasetobjectresponse.md), error**

