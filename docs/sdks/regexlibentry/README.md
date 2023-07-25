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
    res, err := s.RegexLibEntry.Delete(ctx, operations.DeleteRegexLibEntryRequest{
        ID: "c060b06a-1287-4764-aef6-d0c6d6ed9c73",
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
    res, err := s.RegexLibEntry.Post(ctx, shared.RegexLibEntry{
        Description: cribl.String("temporibus"),
        ID: "d6345715-09a8-4e87-8d3c-5a1f9c242c7b",
        Lib: cribl.String("iure"),
        Regex: "aliquid",
        SampleData: cribl.String("culpa"),
        Tags: cribl.String("illo"),
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
    res, err := s.RegexLibEntry.Update(ctx, operations.UpdateRegexLibEntryRequest{
        RegexLibEntry: &shared.RegexLibEntry{
            Description: cribl.String("reiciendis"),
            ID: "30c73df5-b671-4989-8f42-a4bb438d85b2",
            Lib: cribl.String("laboriosam"),
            Regex: "aperiam",
            SampleData: cribl.String("minima"),
            Tags: cribl.String("perspiciatis"),
        },
        ID: "1d745e3c-2059-4c9c-bf56-7e0e252765b1",
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

