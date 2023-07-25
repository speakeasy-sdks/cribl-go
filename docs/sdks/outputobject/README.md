# OutputObject

### Available Operations

* [Create](#create) - Create Output

## Create

Create Output

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
    res, err := s.OutputObject.Create(ctx, shared.Output{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Outputs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Output](../../models/shared/output.md)        | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateOutputObjectResponse](../../models/operations/createoutputobjectresponse.md), error**

