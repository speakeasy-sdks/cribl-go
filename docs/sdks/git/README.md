# Git

### Available Operations

* [CountFiles](#countfiles) - get the count of files of changed
* [Create](#create) - create a new commit containing the current configs the given log message describing the changes.
* [Get](#get) - Get info about versioning availability
* [GetLogandTextualDiff](#getlogandtextualdiff) - get the log message and textual diff for given commit
* [GetTextualDiff](#gettextualdiff) - get the textual diff for given commit
* [GetWorkingTree](#getworkingtree) - get the the working tree status
* [ListBranches](#listbranches) - get the list of branches
* [ListChangedFiles](#listchangedfiles) - get the files changed
* [ListGitSettings](#listgitsettings) - Get git settings
* [Push](#push) - push the current configs to the remote repository.
* [Sync](#sync) - syncs with remote repo via POST requests
* [Update](#update) - Update git settings

## CountFiles

get the count of files of changed

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
    group := "libero"

    ctx := context.Background()
    res, err := s.Git.CountFiles(ctx, group)
    if err != nil {
        log.Fatal(err)
    }

    if res.CountFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `group`                                               | **string*                                             | :heavy_minus_sign:                                    | Group ID                                              |


### Response

**[*operations.GetCountFileResponse](../../models/operations/getcountfileresponse.md), error**


## Create

create a new commit containing the current configs the given log message describing the changes.

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
    res, err := s.Git.Create(ctx, shared.GitCommitParams{
        Effective: cribl.Bool(false),
        Group: cribl.String("architecto"),
        Message: "cupiditate",
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


## Get

Get info about versioning availability

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
    res, err := s.Git.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GitInfos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetVersioningResponse](../../models/operations/getversioningresponse.md), error**


## GetLogandTextualDiff

get the log message and textual diff for given commit

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
    commit := "molestiae"
    group := "eligendi"

    ctx := context.Background()
    res, err := s.Git.GetLogandTextualDiff(ctx, commit, group)
    if err != nil {
        log.Fatal(err)
    }

    if res.TextualDiff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `commit`                                              | **string*                                             | :heavy_minus_sign:                                    | Commit hash (default is HEAD)                         |
| `group`                                               | **string*                                             | :heavy_minus_sign:                                    | Group ID                                              |


### Response

**[*operations.GetLogandTextualDiffResponse](../../models/operations/getlogandtextualdiffresponse.md), error**


## GetTextualDiff

get the textual diff for given commit

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
    commit := "possimus"
    group := "non"

    ctx := context.Background()
    res, err := s.Git.GetTextualDiff(ctx, commit, group)
    if err != nil {
        log.Fatal(err)
    }

    if res.TextualDiff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `commit`                                              | **string*                                             | :heavy_minus_sign:                                    | Commit hash (default is HEAD)                         |
| `group`                                               | **string*                                             | :heavy_minus_sign:                                    | Group ID                                              |


### Response

**[*operations.GetTextualDiffResponse](../../models/operations/gettextualdiffresponse.md), error**


## GetWorkingTree

get the the working tree status

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
    group := "magnam"

    ctx := context.Background()
    res, err := s.Git.GetWorkingTree(ctx, group)
    if err != nil {
        log.Fatal(err)
    }

    if res.GitStatusResults != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `group`                                               | **string*                                             | :heavy_minus_sign:                                    | Group ID                                              |


### Response

**[*operations.GetWorkingTreeResponse](../../models/operations/getworkingtreeresponse.md), error**


## ListBranches

get the list of branches

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
    res, err := s.Git.ListBranches(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Branches != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListBranchesResponse](../../models/operations/listbranchesresponse.md), error**


## ListChangedFiles

get the files changed

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
    id := "itaque"
    group := "sed"

    ctx := context.Background()
    res, err := s.Git.ListChangedFiles(ctx, id, group)
    if err != nil {
        log.Fatal(err)
    }

    if res.ChangedFiles != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | **string*                                             | :heavy_minus_sign:                                    | Commit ID                                             |
| `group`                                               | **string*                                             | :heavy_minus_sign:                                    | Group ID                                              |


### Response

**[*operations.ListChangedFilesResponse](../../models/operations/listchangedfilesresponse.md), error**


## ListGitSettings

Get git settings

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
    res, err := s.Git.ListGitSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GitSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListGitSettingsResponse](../../models/operations/listgitsettingsresponse.md), error**


## Push

push the current configs to the remote repository.

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
    res, err := s.Git.Push(ctx)
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


## Sync

syncs with remote repo via POST requests

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
    res, err := s.Git.Sync(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GitStatusResults != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.SyncRemoteRepoResponse](../../models/operations/syncremotereporesponse.md), error**


## Update

Update git settings

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
    res, err := s.Git.Update(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GitSettings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.UpdateGitSettingsResponse](../../models/operations/updategitsettingsresponse.md), error**

