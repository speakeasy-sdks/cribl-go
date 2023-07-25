# OutputID

### Available Operations

* [Delete](#delete) - Delete Output
* [Get](#get) - Get Output by ID
* [Update](#update) - Update Output

## Delete

Delete Output

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
    res, err := s.OutputID.Delete(ctx, operations.DeleteOutputIDRequest{
        ID: "4009ef8d-29de-41dd-b097-b5da08c57fa6",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Outputs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteOutputIDRequest](../../models/operations/deleteoutputidrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DeleteOutputIDResponse](../../models/operations/deleteoutputidresponse.md), error**


## Get

Get Output by ID

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
    res, err := s.OutputID.Get(ctx, operations.GetOutputIDRequest{
        ID: "c78a216e-19ba-4fec-a619-1498140b64ff",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Outputs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetOutputIDRequest](../../models/operations/getoutputidrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetOutputIDResponse](../../models/operations/getoutputidresponse.md), error**


## Update

Update Output

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
    res, err := s.OutputID.Update(ctx, operations.UpdateOutputIDRequest{
        Output: &shared.Output{},
        ID: "8ae170ef-03b5-4f37-a4aa-868555966732",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Outputs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateOutputIDRequest](../../models/operations/updateoutputidrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateOutputIDResponse](../../models/operations/updateoutputidresponse.md), error**

