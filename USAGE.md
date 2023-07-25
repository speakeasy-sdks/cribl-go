<!-- Start SDK Example Usage -->


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
    res, err := s.AppscopeLibEntries.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.AppScopeLibEntries != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->