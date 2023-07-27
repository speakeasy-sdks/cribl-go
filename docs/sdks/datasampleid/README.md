# DataSampleID

### Available Operations

* [Delete](#delete) - Delete DataSample
* [Get](#get) - Get DataSample by ID
* [Update](#update) - Update DataSample

## Delete

Delete DataSample

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
    res, err := s.DataSampleID.Delete(ctx, operations.DeleteDataSampleIDRequest{
        ID: "319c177d-525f-477b-914e-eb52ff785fc3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DataSamples != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteDataSampleIDRequest](../../models/operations/deletedatasampleidrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.DeleteDataSampleIDResponse](../../models/operations/deletedatasampleidresponse.md), error**


## Get

Get DataSample by ID

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
    res, err := s.DataSampleID.Get(ctx, operations.GetDataSampleIDRequest{
        ID: "7814d4c9-8e0c-42bb-89eb-75dad636c600",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DataSamples != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetDataSampleIDRequest](../../models/operations/getdatasampleidrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetDataSampleIDResponse](../../models/operations/getdatasampleidresponse.md), error**


## Update

Update DataSample

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
    res, err := s.DataSampleID.Update(ctx, operations.UpdateDataSampleIDRequest{
        RequestBody: map[string]interface{}{
            "quae": "amet",
            "illum": "praesentium",
        },
        ID: "bb31180f-739a-4e9e-857e-b809e2810331",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DataSamples != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateDataSampleIDRequest](../../models/operations/updatedatasampleidrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateDataSampleIDResponse](../../models/operations/updatedatasampleidresponse.md), error**

