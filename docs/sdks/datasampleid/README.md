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
    id := "nemo"

    ctx := context.Background()
    res, err := s.DataSampleID.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.DataSamples != nil {
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
    id := "quos"

    ctx := context.Background()
    res, err := s.DataSampleID.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.DataSamples != nil {
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
    id := "eius"
    requestBody := map[string]interface{}{
        "ducimus": "nesciunt",
    }

    ctx := context.Background()
    res, err := s.DataSampleID.Update(ctx, id, requestBody)
    if err != nil {
        log.Fatal(err)
    }

    if res.DataSamples != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |
| `requestBody`                                         | map[string]*interface{}*                              | :heavy_minus_sign:                                    | DataSample object to be updated                       |


### Response

**[*operations.UpdateDataSampleIDResponse](../../models/operations/updatedatasampleidresponse.md), error**

