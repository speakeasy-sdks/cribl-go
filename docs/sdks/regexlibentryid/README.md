# RegexLibEntryID

### Available Operations

* [Get](#get) - Get RegexLibEntry by ID

## Get

Get RegexLibEntry by ID

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
    res, err := s.RegexLibEntryID.Get(ctx, operations.GetRegexLibEntryIDRequest{
        ID: "d62fcdac-e1f0-4121-ace2-239e8f25cd0d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RegexLibEntries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetRegexLibEntryIDRequest](../../models/operations/getregexlibentryidrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetRegexLibEntryIDResponse](../../models/operations/getregexlibentryidresponse.md), error**

