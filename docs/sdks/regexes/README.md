# Regexes

## Overview

Actions related to regular expressions

### Available Operations

* [Delete](#delete) - Delete RegexLibEntry
* [Get](#get) - Get RegexLibEntry by ID
* [ListRegexLibEntryObject](#listregexlibentryobject) - Get a list of RegexLibEntry objects
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
    id := "totam"

    ctx := context.Background()
    res, err := s.Regexes.Delete(ctx, id)
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


## Get

Get RegexLibEntry by ID

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
    id := "dicta"

    ctx := context.Background()
    res, err := s.Regexes.Get(ctx, id)
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

**[*operations.GetRegexLibEntryIDResponse](../../models/operations/getregexlibentryidresponse.md), error**


## ListRegexLibEntryObject

Get a list of RegexLibEntry objects

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
    res, err := s.Regexes.ListRegexLibEntryObject(ctx)
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


### Response

**[*operations.ListRegexLibEntryObjectResponse](../../models/operations/listregexlibentryobjectresponse.md), error**


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
    res, err := s.Regexes.Post(ctx, shared.RegexLibEntry{
        Description: cribl.String("voluptatem"),
        ID: "331f3981-d4c7-400b-a07f-3c93c73b9da3",
        Lib: cribl.String("tenetur"),
        Regex: "aspernatur",
        SampleData: cribl.String("quo"),
        Tags: cribl.String("itaque"),
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
    id := "illum"
    regexLibEntry := &shared.RegexLibEntry{
        Description: cribl.String("laborum"),
        ID: "7e23f225-7411-4faf-8b75-44e472e80285",
        Lib: cribl.String("voluptate"),
        Regex: "similique",
        SampleData: cribl.String("minima"),
        Tags: cribl.String("libero"),
    }

    ctx := context.Background()
    res, err := s.Regexes.Update(ctx, id, regexLibEntry)
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

