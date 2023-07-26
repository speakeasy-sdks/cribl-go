# BulletinMessages

### Available Operations

* [Get](#get) - Get a list of BulletinMessage objects

## Get

Get a list of BulletinMessage objects

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
    res, err := s.BulletinMessages.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.BulletinMessages != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetBulletinMessagesResponse](../../models/operations/getbulletinmessagesresponse.md), error**
