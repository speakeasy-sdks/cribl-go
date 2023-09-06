# AppscopeConfigs

### Available Operations

* [Create](#create) - Create AppscopeLibEntry
* [Delete](#delete) - Delete AppscopeLibEntry
* [Get](#get) - Get AppscopeLibEntry by ID
* [ListAppscopeLibEntries](#listappscopelibentries) - Get a list of AppscopeLibEntry objects
* [Update](#update) - Update AppscopeLibEntry

## Create

Create AppscopeLibEntry

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
    res, err := s.AppscopeConfigs.Create(ctx, shared.AppscopeLibEntry{
        Config: shared.AppscopeConfigWithCustom{
            Cribl: &shared.AppscopeConfigWithCustomCribl{
                Authtoken: cribl.String("dicta"),
                Enable: cribl.Bool(false),
                Transport: &shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                    Host: cribl.String("dolore"),
                    Path: cribl.String("iusto"),
                    Port: cribl.Int64(118727),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("harum"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("enim"),
                },
                UseScopeSourceTransport: cribl.Bool(false),
            },
            Custom: []shared.AppscopeCustom{
                shared.AppscopeCustom{
                    Ancestor: cribl.String("accusamus"),
                    Arg: cribl.String("commodi"),
                    Config: shared.AppscopeConfig{
                        Cribl: &shared.AppscopeConfigCribl{
                            Authtoken: cribl.String("repudiandae"),
                            Enable: cribl.Bool(false),
                            Transport: &shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("ipsum"),
                                Path: cribl.String("quidem"),
                                Port: cribl.Int64(565189),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("excepturi"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("pariatur"),
                            },
                            UseScopeSourceTransport: cribl.Bool(false),
                        },
                        Event: &shared.AppscopeConfigEvent{
                            Enable: false,
                            Format: shared.AppscopeConfigEventFormat{
                                Enhancefs: false,
                                Maxeventpersec: 265389,
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("rem"),
                                Path: cribl.String("voluptates"),
                                Port: cribl.Int64(93940),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("repudiandae"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("sint"),
                            },
                            Type: shared.AppscopeConfigEventTypeNdjson,
                            Watch: []shared.AppscopeConfigEventWatch{
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("veritatis"),
                                    Headers: cribl.String("itaque"),
                                    Name: cribl.String("Erin Altenwerth"),
                                    Type: "explicabo",
                                    Value: cribl.String("deserunt"),
                                },
                            },
                        },
                        Libscope: &shared.AppscopeConfigLibscope{
                            Commanddir: cribl.String("distinctio"),
                            Configevent: cribl.Bool(false),
                            Log: &shared.AppscopeConfigLibscopeLog{
                                Level: shared.AppscopeConfigLibscopeLogLevelNone.ToPointer(),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("modi"),
                                    Path: cribl.String("qui"),
                                    Port: cribl.Int64(397821),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("cupiditate"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("quos"),
                                },
                            },
                            Summaryperiod: cribl.Int64(20107),
                        },
                        Metric: &shared.AppscopeConfigMetric{
                            Enable: false,
                            Format: shared.AppscopeConfigMetricFormat{
                                Statsdmaxlen: cribl.Int64(164940),
                                Statsdprefix: cribl.String("assumenda"),
                                Type: cribl.String("ipsam"),
                                Verbosity: cribl.Int64(4695),
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("dolorum"),
                                Path: cribl.String("excepturi"),
                                Port: cribl.Int64(270008),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("facilis"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("tempore"),
                            },
                            Watch: []string{
                                "labore",
                            },
                        },
                        Payload: &shared.AppscopeConfigPayload{
                            Dir: "delectus",
                            Enable: false,
                        },
                        Protocol: []shared.AppscopeConfigProtocol{
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 433288,
                                Name: "Sheri Mayer",
                                Payload: false,
                                Regex: "necessitatibus",
                            },
                        },
                        Tags: []shared.AppscopeConfigTags{
                            shared.AppscopeConfigTags{
                                Key: "sint",
                                Value: "officia",
                            },
                        },
                    },
                    Env: cribl.String("dolor"),
                    Hostname: cribl.String("unimportant-venture.net"),
                    Procname: cribl.String("in"),
                    Username: cribl.String("Isobel_Stamm11"),
                },
            },
            Event: &shared.AppscopeConfigWithCustomEvent{
                Enable: false,
                Format: shared.AppscopeConfigWithCustomEventFormat{
                    Enhancefs: false,
                    Maxeventpersec: 297437,
                },
                Transport: shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                    Host: cribl.String("facere"),
                    Path: cribl.String("ea"),
                    Port: cribl.Int64(396506),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("laborum"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("accusamus"),
                },
                Type: shared.AppscopeConfigWithCustomEventTypeNdjson,
                Watch: []shared.AppscopeConfigWithCustomEventWatch{
                    shared.AppscopeConfigWithCustomEventWatch{
                        Allowbinary: cribl.Bool(false),
                        Enabled: cribl.Bool(false),
                        Field: cribl.String("non"),
                        Headers: cribl.String("occaecati"),
                        Name: cribl.String("Sophia Wintheiser"),
                        Type: "nam",
                        Value: cribl.String("id"),
                    },
                },
            },
            Libscope: &shared.AppscopeConfigWithCustomLibscope{
                Commanddir: cribl.String("blanditiis"),
                Configevent: cribl.Bool(false),
                Log: &shared.AppscopeConfigWithCustomLibscopeLog{
                    Level: shared.AppscopeConfigWithCustomLibscopeLogLevelWarning.ToPointer(),
                    Transport: &shared.AppscopeTransport{
                        Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                        Host: cribl.String("amet"),
                        Path: cribl.String("deserunt"),
                        Port: cribl.Int64(394869),
                        TLS: &shared.AppscopeTransportTLS{
                            Cacertpath: cribl.String("vel"),
                            Enable: cribl.Bool(false),
                            Validateserver: cribl.Bool(false),
                        },
                        Type: cribl.String("natus"),
                    },
                },
                Summaryperiod: cribl.Int64(606393),
            },
            Metric: &shared.AppscopeConfigWithCustomMetric{
                Enable: false,
                Format: shared.AppscopeConfigWithCustomMetricFormat{
                    Statsdmaxlen: cribl.Int64(474867),
                    Statsdprefix: cribl.String("perferendis"),
                    Type: cribl.String("nihil"),
                    Verbosity: cribl.Int64(301575),
                },
                Transport: shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                    Host: cribl.String("id"),
                    Path: cribl.String("labore"),
                    Port: cribl.Int64(290077),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("suscipit"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("natus"),
                },
                Watch: []string{
                    "nobis",
                },
            },
            Payload: &shared.AppscopeConfigWithCustomPayload{
                Dir: "eum",
                Enable: false,
            },
            Protocol: []shared.AppscopeConfigWithCustomProtocol{
                shared.AppscopeConfigWithCustomProtocol{
                    Binary: false,
                    Detect: false,
                    Len: 878453,
                    Name: "Ms. Julie Gusikowski",
                    Payload: false,
                    Regex: "provident",
                },
            },
            Tags: []shared.AppscopeConfigWithCustomTags{
                shared.AppscopeConfigWithCustomTags{
                    Key: "quos",
                    Value: "sint",
                },
            },
        },
        Description: "accusantium",
        ID: "afa563e2-516f-4e4c-8b71-1e5b7fd2ed02",
        Lib: shared.CriblLibCriblCustom,
        Tags: cribl.String("natus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppscopeLibEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [shared.AppscopeLibEntry](../../models/shared/appscopelibentry.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[*operations.CreateAppscopeLibEntryResponse](../../models/operations/createappscopelibentryresponse.md), error**


## Delete

Delete AppscopeLibEntry

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
    id := "magni"

    ctx := context.Background()
    res, err := s.AppscopeConfigs.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.AppscopeLibEntry != nil {
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

**[*operations.DeleteAppscopeLibEntryResponse](../../models/operations/deleteappscopelibentryresponse.md), error**


## Get

Get AppscopeLibEntry by ID

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
    id := "sunt"

    ctx := context.Background()
    res, err := s.AppscopeConfigs.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.AppscopeLibEntry != nil {
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

**[*operations.GetAppscopeLibEntryResponse](../../models/operations/getappscopelibentryresponse.md), error**


## ListAppscopeLibEntries

Get a list of AppscopeLibEntry objects

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
    res, err := s.AppscopeConfigs.ListAppscopeLibEntries(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.AppScopeLibEntries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListAppscopeLibEntriesResponse](../../models/operations/listappscopelibentriesresponse.md), error**


## Update

Update AppscopeLibEntry

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
    id := "quo"
    appscopeLibEntry := &shared.AppscopeLibEntry{
        Config: shared.AppscopeConfigWithCustom{
            Cribl: &shared.AppscopeConfigWithCustomCribl{
                Authtoken: cribl.String("illum"),
                Enable: cribl.Bool(false),
                Transport: &shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                    Host: cribl.String("maxime"),
                    Path: cribl.String("ea"),
                    Port: cribl.Int64(569101),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("odit"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("ea"),
                },
                UseScopeSourceTransport: cribl.Bool(false),
            },
            Custom: []shared.AppscopeCustom{
                shared.AppscopeCustom{
                    Ancestor: cribl.String("accusantium"),
                    Arg: cribl.String("ab"),
                    Config: shared.AppscopeConfig{
                        Cribl: &shared.AppscopeConfigCribl{
                            Authtoken: cribl.String("maiores"),
                            Enable: cribl.Bool(false),
                            Transport: &shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("ipsam"),
                                Path: cribl.String("voluptate"),
                                Port: cribl.Int64(420075),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("nam"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("eaque"),
                            },
                            UseScopeSourceTransport: cribl.Bool(false),
                        },
                        Event: &shared.AppscopeConfigEvent{
                            Enable: false,
                            Format: shared.AppscopeConfigEventFormat{
                                Enhancefs: false,
                                Maxeventpersec: 866383,
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("voluptatibus"),
                                Path: cribl.String("perferendis"),
                                Port: cribl.Int64(855804),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("amet"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("aut"),
                            },
                            Type: shared.AppscopeConfigEventTypeNdjson,
                            Watch: []shared.AppscopeConfigEventWatch{
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("cumque"),
                                    Headers: cribl.String("corporis"),
                                    Name: cribl.String("Pete Rohan"),
                                    Type: "totam",
                                    Value: cribl.String("dignissimos"),
                                },
                            },
                        },
                        Libscope: &shared.AppscopeConfigLibscope{
                            Commanddir: cribl.String("eaque"),
                            Configevent: cribl.Bool(false),
                            Log: &shared.AppscopeConfigLibscopeLog{
                                Level: shared.AppscopeConfigLibscopeLogLevelInfo.ToPointer(),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("eos"),
                                    Path: cribl.String("perferendis"),
                                    Port: cribl.Int64(170986),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("minus"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("quam"),
                                },
                            },
                            Summaryperiod: cribl.Int64(223924),
                        },
                        Metric: &shared.AppscopeConfigMetric{
                            Enable: false,
                            Format: shared.AppscopeConfigMetricFormat{
                                Statsdmaxlen: cribl.Int64(874573),
                                Statsdprefix: cribl.String("nostrum"),
                                Type: cribl.String("hic"),
                                Verbosity: cribl.Int64(928082),
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("facilis"),
                                Path: cribl.String("perspiciatis"),
                                Port: cribl.Int64(31838),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("porro"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("consequuntur"),
                            },
                            Watch: []string{
                                "blanditiis",
                            },
                        },
                        Payload: &shared.AppscopeConfigPayload{
                            Dir: "error",
                            Enable: false,
                        },
                        Protocol: []shared.AppscopeConfigProtocol{
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 50370,
                                Name: "Jean Ferry",
                                Payload: false,
                                Regex: "modi",
                            },
                        },
                        Tags: []shared.AppscopeConfigTags{
                            shared.AppscopeConfigTags{
                                Key: "iste",
                                Value: "dolorum",
                            },
                        },
                    },
                    Env: cribl.String("deleniti"),
                    Hostname: cribl.String("traumatic-neighbourhood.net"),
                    Procname: cribl.String("libero"),
                    Username: cribl.String("Violet.Hahn"),
                },
            },
            Event: &shared.AppscopeConfigWithCustomEvent{
                Enable: false,
                Format: shared.AppscopeConfigWithCustomEventFormat{
                    Enhancefs: false,
                    Maxeventpersec: 212390,
                },
                Transport: shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                    Host: cribl.String("dolor"),
                    Path: cribl.String("qui"),
                    Port: cribl.Int64(218749),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("hic"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("excepturi"),
                },
                Type: shared.AppscopeConfigWithCustomEventTypeNdjson,
                Watch: []shared.AppscopeConfigWithCustomEventWatch{
                    shared.AppscopeConfigWithCustomEventWatch{
                        Allowbinary: cribl.Bool(false),
                        Enabled: cribl.Bool(false),
                        Field: cribl.String("cum"),
                        Headers: cribl.String("voluptate"),
                        Name: cribl.String("Johanna Farrell"),
                        Type: "veritatis",
                        Value: cribl.String("ipsa"),
                    },
                },
            },
            Libscope: &shared.AppscopeConfigWithCustomLibscope{
                Commanddir: cribl.String("ipsa"),
                Configevent: cribl.Bool(false),
                Log: &shared.AppscopeConfigWithCustomLibscopeLog{
                    Level: shared.AppscopeConfigWithCustomLibscopeLogLevelWarning.ToPointer(),
                    Transport: &shared.AppscopeTransport{
                        Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                        Host: cribl.String("quaerat"),
                        Path: cribl.String("accusamus"),
                        Port: cribl.Int64(696344),
                        TLS: &shared.AppscopeTransportTLS{
                            Cacertpath: cribl.String("voluptatibus"),
                            Enable: cribl.Bool(false),
                            Validateserver: cribl.Bool(false),
                        },
                        Type: cribl.String("voluptas"),
                    },
                },
                Summaryperiod: cribl.Int64(617658),
            },
            Metric: &shared.AppscopeConfigWithCustomMetric{
                Enable: false,
                Format: shared.AppscopeConfigWithCustomMetricFormat{
                    Statsdmaxlen: cribl.Int64(179603),
                    Statsdprefix: cribl.String("atque"),
                    Type: cribl.String("sit"),
                    Verbosity: cribl.Int64(854614),
                },
                Transport: shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                    Host: cribl.String("soluta"),
                    Path: cribl.String("dolorum"),
                    Port: cribl.Int64(478596),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("voluptate"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("dolorum"),
                },
                Watch: []string{
                    "deleniti",
                },
            },
            Payload: &shared.AppscopeConfigWithCustomPayload{
                Dir: "omnis",
                Enable: false,
            },
            Protocol: []shared.AppscopeConfigWithCustomProtocol{
                shared.AppscopeConfigWithCustomProtocol{
                    Binary: false,
                    Detect: false,
                    Len: 896672,
                    Name: "Emmett Kovacek",
                    Payload: false,
                    Regex: "id",
                },
            },
            Tags: []shared.AppscopeConfigWithCustomTags{
                shared.AppscopeConfigWithCustomTags{
                    Key: "saepe",
                    Value: "eius",
                },
            },
        },
        Description: "aspernatur",
        ID: "03ce5e6a-95d8-4a0d-846c-e2af7a73cf3b",
        Lib: shared.CriblLibCustom,
        Tags: cribl.String("numquam"),
    }

    ctx := context.Background()
    res, err := s.AppscopeConfigs.Update(ctx, id, appscopeLibEntry)
    if err != nil {
        log.Fatal(err)
    }

    if res.AppscopeLibEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ctx`                                                               | [context.Context](https://pkg.go.dev/context#Context)               | :heavy_check_mark:                                                  | The context to use for the request.                                 |
| `id`                                                                | *string*                                                            | :heavy_check_mark:                                                  | Unique ID                                                           |
| `appscopeLibEntry`                                                  | [*shared.AppscopeLibEntry](../../models/shared/appscopelibentry.md) | :heavy_minus_sign:                                                  | AppscopeLibEntry object to be updated                               |


### Response

**[*operations.UpdateAppscopeLibEntryResponse](../../models/operations/updateappscopelibentryresponse.md), error**

