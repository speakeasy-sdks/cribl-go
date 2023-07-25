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
        ID: "2f0ea930-b69f-47ac-af72-f88500904911",
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
        ID: "60820788-8ec6-4618-bbfe-9659eb40ec16",
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
        ID: "faf75b0b-532a-44da-b7cb-aaf4452c4842",
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

