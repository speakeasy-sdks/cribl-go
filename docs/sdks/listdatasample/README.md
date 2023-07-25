# ListDataSample

### Available Operations

* [Get](#get) - Get a list of DataSample objects

## Get

Get a list of DataSample objects

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
    res, err := s.ListDataSample.Get(ctx)
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


### Response

**[*operations.GetListDataSampleResponse](../../models/operations/getlistdatasampleresponse.md), error**

