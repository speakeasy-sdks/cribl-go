# GlobalVariable

### Available Operations

* [Post](#post) - Create Global Variable

## Post

Create Global Variable

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
    res, err := s.GlobalVariable.Post(ctx, shared.GlobalVar{
        Description: cribl.String("expedita"),
        ID: "7c15dfba-ce18-48b1-84ee-2c8c6ce611fe",
        Lib: cribl.String("vero"),
        Tags: cribl.String("quidem"),
        Type: shared.GlobalVarTypeString,
        Value: "quo",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GlobalVars != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.GlobalVar](../../models/shared/globalvar.md)  | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostGlobalVariableResponse](../../models/operations/postglobalvariableresponse.md), error**

