# CriblMetadata

### Available Operations

* [Get](#get) - Obtain metadata which Cribl Stream/Edge uses when acting as a Service Provider

## Get

Obtain metadata which Cribl Stream/Edge uses when acting as a Service Provider

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.CriblMetadata.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCriblMetadata200TextXMLString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetCriblMetadataResponse](../../models/operations/getcriblmetadataresponse.md), error**

