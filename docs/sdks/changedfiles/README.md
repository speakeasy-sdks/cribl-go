# ChangedFiles

### Available Operations

* [Get](#get) - get the files changed

## Get

get the files changed

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
    res, err := s.ChangedFiles.Get(ctx, operations.GetChangedFilesRequest{
        ID: cribl.String("6c645b08-b618-491b-aa0f-e1ade008e6f8"),
        Group: cribl.String("minus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ChangedFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetChangedFilesRequest](../../models/operations/getchangedfilesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetChangedFilesResponse](../../models/operations/getchangedfilesresponse.md), error**

