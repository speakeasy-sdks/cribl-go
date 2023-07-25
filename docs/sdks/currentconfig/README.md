# CurrentConfig

### Available Operations

* [Push](#push) - push the current configs to the remote repository.

## Push

push the current configs to the remote repository.

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
    res, err := s.CurrentConfig.Push(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.PushConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.PushCurrentConfigResponse](../../models/operations/pushcurrentconfigresponse.md), error**

