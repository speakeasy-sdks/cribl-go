# SpecifiedOutput

### Available Operations

* [Get](#get) - Get samples data for the specified output. Used to get sample data for the test action.

## Get

Get samples data for the specified output. Used to get sample data for the test action.

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
    res, err := s.SpecifiedOutput.Get(ctx, operations.GetSpecifiedOutputRequest{
        ID: "10f8c23d-f931-4da3-adb5-1fad94acc943",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSpecifiedOutput200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetSpecifiedOutputRequest](../../models/operations/getspecifiedoutputrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetSpecifiedOutputResponse](../../models/operations/getspecifiedoutputresponse.md), error**

