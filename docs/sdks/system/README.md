# System

## Overview

Actions related to system settings

### Available Operations

* [ReloadCriblSettings](#reloadcriblsettings) - Reload Cribl settings from the filesystem
* [CancelRunningGroup](#cancelrunninggroup) - Cancel a running group upgrade
* [ExecuteGroupUpgrade](#executegroupupgrade) - Execute distributed group upgrade
* [~~Get~~](#get) - Get Cribl system settings. Deprecated: use specific endpoints /system/limits, /system/job-limits, /system/redis-cache-limits, /system/services-limits, /system/settings/git-settings, and /system/settings/conf respectively :warning: **Deprecated**
* [GetPreviousPackage](#getpreviouspackage) - Get the previously downloaded Cribl package
* [ListAuthenticationSettings](#listauthenticationsettings) - Get authentication settings
* [ListCriblVersion](#listcriblversion) - Get a list of Cribl versions available for upgrade
* [ListSearchLimits](#listsearchlimits) - Get search limits
* [ListSettings](#listsettings) - Get Cribl system settings
* [Restart](#restart) - Restart Cribl server
* [StageGroupUpgrade](#stagegroupupgrade) - Stage distributed group upgrade
* [~~Update~~](#update) - Update Cribl system settings. Deprecated: use specific endpoints /system/limits, /system/job-limits, /system/settings/git-settings, /system/settings/auth and /system/settings/conf respectively :warning: **Deprecated**
* [UpdateAuthSettings](#updateauthsettings) - Update authentication settings
* [UpdateChangelogViewState](#updatechangelogviewstate) - Update changelog viewed state
* [UpdateCriblSettings](#updatecriblsettings) - Update Cribl system settings
* [UpgradeCribl](#upgradecribl) - Upgrade Cribl to a given version
* [UpgradeMaster](#upgrademaster) - Upgrade master node with the provided package

## ReloadCriblSettings

Reload Cribl settings from the filesystem

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
    res, err := s.System.ReloadCriblSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ReloadCriblSettingsResponse](../../models/operations/reloadcriblsettingsresponse.md), error**


## CancelRunningGroup

Cancel a running group upgrade

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
    group := "quis"

    ctx := context.Background()
    res, err := s.System.CancelRunningGroup(ctx, group)
    if err != nil {
        log.Fatal(err)
    }

    if res.CancelRunningGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `group`                                               | *string*                                              | :heavy_check_mark:                                    | id of the group                                       |


### Response

**[*operations.CancelRunningGroupResponse](../../models/operations/cancelrunninggroupresponse.md), error**


## ExecuteGroupUpgrade

Execute distributed group upgrade

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
    group := "accusantium"
    distributedUpgradeRequest := &shared.DistributedUpgradeRequest{
        PackageUrls: []shared.DistributedUpgradeRequestPackageUrls{
            shared.DistributedUpgradeRequestPackageUrls{
                PackageHashURL: cribl.String("consectetur"),
                PackageURL: "asperiores",
            },
        },
        UpgradeMode: shared.DistributedUpgradeRequestUpgradeModeRolling.ToPointer(),
        UpgradePercentage: cribl.Int64(571498),
        WorkerRetries: cribl.Int64(821993),
        WorkerRetryDelay: cribl.Int64(732815),
    }

    ctx := context.Background()
    res, err := s.System.ExecuteGroupUpgrade(ctx, group, distributedUpgradeRequest)
    if err != nil {
        log.Fatal(err)
    }

    if res.CriblPackage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `group`                                                                               | *string*                                                                              | :heavy_check_mark:                                                                    | Group to upgrade                                                                      |
| `distributedUpgradeRequest`                                                           | [*shared.DistributedUpgradeRequest](../../models/shared/distributedupgraderequest.md) | :heavy_minus_sign:                                                                    | distributedUpgrade object                                                             |


### Response

**[*operations.ExecuteDistributedUpgradeResponse](../../models/operations/executedistributedupgraderesponse.md), error**


## ~~Get~~

Get Cribl system settings. Deprecated: use specific endpoints /system/limits, /system/job-limits, /system/redis-cache-limits, /system/services-limits, /system/settings/git-settings, and /system/settings/conf respectively

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
    res, err := s.System.Get(ctx)
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

**[*operations.GetCriblsSettingsResponse](../../models/operations/getcriblssettingsresponse.md), error**


## GetPreviousPackage

Get the previously downloaded Cribl package

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
    file := "asperiores"

    ctx := context.Background()
    res, err := s.System.GetPreviousPackage(ctx, file)
    if err != nil {
        log.Fatal(err)
    }

    if res.CriblPackage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `file`                                                | *string*                                              | :heavy_check_mark:                                    | Name of the file to be downloaded                     |


### Response

**[*operations.GetPreviousCriblPackageResponse](../../models/operations/getpreviouscriblpackageresponse.md), error**


## ListAuthenticationSettings

Get authentication settings

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
    res, err := s.System.ListAuthenticationSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.AuthConfigs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListAuthenticationSettingsResponse](../../models/operations/listauthenticationsettingsresponse.md), error**


## ListCriblVersion

Get a list of Cribl versions available for upgrade

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
    res, err := s.System.ListCriblVersion(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpgradeResults != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListCriblVersionResponse](../../models/operations/listcriblversionresponse.md), error**


## ListSearchLimits

Get search limits

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
    res, err := s.System.ListSearchLimits(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchSettingses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListSearchLimitsResponse](../../models/operations/listsearchlimitsresponse.md), error**


## ListSettings

Get Cribl system settings

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
    res, err := s.System.ListSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.SystemSettings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListSettingsResponse](../../models/operations/listsettingsresponse.md), error**


## Restart

Restart Cribl server

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
    res, err := s.System.Restart(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.RestartCriblSettingsResponse](../../models/operations/restartcriblsettingsresponse.md), error**


## StageGroupUpgrade

Stage distributed group upgrade

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
    group := "quasi"
    upgradePercentage := "consequuntur"

    ctx := context.Background()
    res, err := s.System.StageGroupUpgrade(ctx, group, upgradePercentage)
    if err != nil {
        log.Fatal(err)
    }

    if res.CriblPackage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `group`                                                        | *string*                                                       | :heavy_check_mark:                                             | Group to upgrade                                               |
| `upgradePercentage`                                            | **string*                                                      | :heavy_minus_sign:                                             | body number percentage of nodes on the worker group to upgrade |


### Response

**[*operations.PostStageDistributedPackageResponse](../../models/operations/poststagedistributedpackageresponse.md), error**


## ~~Update~~

Update Cribl system settings. Deprecated: use specific endpoints /system/limits, /system/job-limits, /system/settings/git-settings, /system/settings/auth and /system/settings/conf respectively

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
    res, err := s.System.Update(ctx)
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

**[*operations.UpdateCriblsSettingsResponse](../../models/operations/updatecriblssettingsresponse.md), error**


## UpdateAuthSettings

Update authentication settings

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
    res, err := s.System.UpdateAuthSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.AuthConfigs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.UpdateAuthenticationSettingsResponse](../../models/operations/updateauthenticationsettingsresponse.md), error**


## UpdateChangelogViewState

Update changelog viewed state

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
    res, err := s.System.UpdateChangelogViewState(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ChangelogStateses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.UpdateChangelogViewStateResponse](../../models/operations/updatechangelogviewstateresponse.md), error**


## UpdateCriblSettings

Update Cribl system settings

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
    res, err := s.System.UpdateCriblSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.SystemSettingses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.UpdateCriblSystemSettingsResponse](../../models/operations/updatecriblsystemsettingsresponse.md), error**


## UpgradeCribl

Upgrade Cribl to a given version

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
    version := "nemo"

    ctx := context.Background()
    res, err := s.System.UpgradeCribl(ctx, version)
    if err != nil {
        log.Fatal(err)
    }

    if res.Cribl != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `version`                                             | *string*                                              | :heavy_check_mark:                                    | Version number                                        |


### Response

**[*operations.PostGiveCriblVersionResponse](../../models/operations/postgivecriblversionresponse.md), error**


## UpgradeMaster

Upgrade master node with the provided package

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
    res, err := s.System.UpgradeMaster(ctx, shared.UpgradeMasterRequest{
        Packages: []shared.UpgradeMasterRequestPackages{
            shared.UpgradeMasterRequestPackages{
                PackageHashURL: cribl.String("debitis"),
                PackageURL: "labore",
            },
            shared.UpgradeMasterRequestPackages{
                PackageHashURL: cribl.String("veritatis"),
                PackageURL: "minima",
            },
            shared.UpgradeMasterRequestPackages{
                PackageHashURL: cribl.String("magni"),
                PackageURL: "itaque",
            },
            shared.UpgradeMasterRequestPackages{
                PackageHashURL: cribl.String("error"),
                PackageURL: "expedita",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Cribl != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.UpgradeMasterRequest](../../models/shared/upgrademasterrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.PostMasterNodePackageResponse](../../models/operations/postmasternodepackageresponse.md), error**

