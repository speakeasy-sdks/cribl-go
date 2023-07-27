# RegexLibEntry

### Available Operations

* [Delete](#delete) - Delete RegexLibEntry
* [Post](#post) - Create RegexLibEntry
* [Update](#update) - Update RegexLibEntry

## Delete

Delete RegexLibEntry

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RegexLibEntry.Delete(ctx, operations.DeleteRegexLibEntryRequest{
        ID: "b7537cd9-222c-49ff-9749-1aabfa2e761f",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteRegexLibEntryRequest](../../models/operations/deleteregexlibentryrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.DeleteRegexLibEntryResponse](../../models/operations/deleteregexlibentryresponse.md), error**


## Post

Create RegexLibEntry

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
    res, err := s.RegexLibEntry.Post(ctx, shared.RegexLibEntry{
        Description: cribl.String("accusantium"),
        ID: "ca4d456e-f103-41e6-899f-0c2001e22cd5",
        Lib: cribl.String("exercitationem"),
        Regex: "quod",
        SampleData: cribl.String("quod"),
        Tags: cribl.String("alias"),
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

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.RegexLibEntry](../../models/shared/regexlibentry.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |


### Response

**[*operations.PostRegexLibEntryResponse](../../models/operations/postregexlibentryresponse.md), error**


## Update

Update RegexLibEntry

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RegexLibEntry.Update(ctx, operations.UpdateRegexLibEntryRequest{
        RegexLibEntry: &shared.RegexLibEntry{
            Description: cribl.String("nemo"),
            ID: "84a184d7-6d97-41fc-820c-65b037bb8e0c",
            Lib: cribl.String("impedit"),
            Regex: "corrupti",
            SampleData: cribl.String("quas"),
            Tags: cribl.String("ullam"),
        },
        ID: "187e4de0-4af2-48c5-9ddb-46aa1cfd6d82",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateRegexLibEntryRequest](../../models/operations/updateregexlibentryrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UpdateRegexLibEntryResponse](../../models/operations/updateregexlibentryresponse.md), error**

