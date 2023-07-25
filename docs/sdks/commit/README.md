# Commit

### Available Operations

* [Create](#create) - create a new commit containing the current configs the given log message describing the changes.

## Create

create a new commit containing the current configs the given log message describing the changes.

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
    res, err := s.Commit.Create(ctx, shared.GitCommitParams{
        Effective: cribl.Bool(false),
        Group: cribl.String("eveniet"),
        Message: "impedit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GitCommit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [shared.GitCommitParams](../../models/shared/gitcommitparams.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |


### Response

**[*operations.CreateCommitResponse](../../models/operations/createcommitresponse.md), error**

