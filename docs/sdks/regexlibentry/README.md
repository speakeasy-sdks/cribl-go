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
    id := "voluptatem"

    ctx := context.Background()
    res, err := s.RegexLibEntry.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.RegexLibEntries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


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
        Description: cribl.String("repudiandae"),
        ID: "56248fff-639a-4910-abdc-ab62676696e1",
        Lib: cribl.String("itaque"),
        Regex: "quisquam",
        SampleData: cribl.String("eaque"),
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
    id := "qui"
    regexLibEntry := &shared.RegexLibEntry{
        Description: cribl.String("consequuntur"),
        ID: "1b335d89-acb3-4ecf-9a8d-0c549ef03004",
        Lib: cribl.String("perspiciatis"),
        Regex: "quam",
        SampleData: cribl.String("atque"),
        Tags: cribl.String("officia"),
    }

    ctx := context.Background()
    res, err := s.RegexLibEntry.Update(ctx, id, regexLibEntry)
    if err != nil {
        log.Fatal(err)
    }

    if res.RegexLibEntries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                     | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ctx`                                                         | [context.Context](https://pkg.go.dev/context#Context)         | :heavy_check_mark:                                            | The context to use for the request.                           |
| `id`                                                          | *string*                                                      | :heavy_check_mark:                                            | Unique ID                                                     |
| `regexLibEntry`                                               | [*shared.RegexLibEntry](../../models/shared/regexlibentry.md) | :heavy_minus_sign:                                            | RegexLibEntry object to be updated                            |


### Response

**[*operations.UpdateRegexLibEntryResponse](../../models/operations/updateregexlibentryresponse.md), error**

