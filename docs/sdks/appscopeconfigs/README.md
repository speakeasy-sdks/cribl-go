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
                Authtoken: cribl.String("voluptatibus"),
                Enable: cribl.Bool(false),
                Transport: &shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                    Host: cribl.String("vero"),
                    Path: cribl.String("omnis"),
                    Port: cribl.Int64(338159),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("ipsum"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("delectus"),
                },
                UseScopeSourceTransport: cribl.Bool(false),
            },
            Custom: []shared.AppscopeCustom{
                shared.AppscopeCustom{
                    Ancestor: cribl.String("consectetur"),
                    Arg: cribl.String("vero"),
                    Config: shared.AppscopeConfig{
                        Cribl: &shared.AppscopeConfigCribl{
                            Authtoken: cribl.String("tenetur"),
                            Enable: cribl.Bool(false),
                            Transport: &shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("hic"),
                                Path: cribl.String("distinctio"),
                                Port: cribl.Int64(799203),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("odio"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("similique"),
                            },
                            UseScopeSourceTransport: cribl.Bool(false),
                        },
                        Event: &shared.AppscopeConfigEvent{
                            Enable: false,
                            Format: shared.AppscopeConfigEventFormat{
                                Enhancefs: false,
                                Maxeventpersec: 708548,
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("ducimus"),
                                Path: cribl.String("dolore"),
                                Port: cribl.Int64(844550),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("illum"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("sequi"),
                            },
                            Type: shared.AppscopeConfigEventTypeNdjson,
                            Watch: []shared.AppscopeConfigEventWatch{
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("impedit"),
                                    Headers: cribl.String("aut"),
                                    Name: cribl.String("Greg Stoltenberg"),
                                    Type: "maiores",
                                    Value: cribl.String("doloribus"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("iusto"),
                                    Headers: cribl.String("eligendi"),
                                    Name: cribl.String("Linda Nikolaus"),
                                    Type: "ea",
                                    Value: cribl.String("aspernatur"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("vel"),
                                    Headers: cribl.String("possimus"),
                                    Name: cribl.String("Paula Jacobs I"),
                                    Type: "maiores",
                                    Value: cribl.String("quasi"),
                                },
                            },
                        },
                        Libscope: &shared.AppscopeConfigLibscope{
                            Commanddir: cribl.String("ex"),
                            Configevent: cribl.Bool(false),
                            Log: &shared.AppscopeConfigLibscopeLog{
                                Level: shared.AppscopeConfigLibscopeLogLevelNone.ToPointer(),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                    Host: cribl.String("voluptatibus"),
                                    Path: cribl.String("nostrum"),
                                    Port: cribl.Int64(960835),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("quisquam"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("saepe"),
                                },
                            },
                            Summaryperiod: cribl.Int64(411372),
                        },
                        Metric: &shared.AppscopeConfigMetric{
                            Enable: false,
                            Format: shared.AppscopeConfigMetricFormat{
                                Statsdmaxlen: cribl.Int64(774048),
                                Statsdprefix: cribl.String("corporis"),
                                Type: cribl.String("veniam"),
                                Verbosity: cribl.Int64(399499),
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("magnam"),
                                Path: cribl.String("ea"),
                                Port: cribl.Int64(775220),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("consectetur"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("recusandae"),
                            },
                            Watch: []string{
                                "minima",
                            },
                        },
                        Payload: &shared.AppscopeConfigPayload{
                            Dir: "eaque",
                            Enable: false,
                        },
                        Protocol: []shared.AppscopeConfigProtocol{
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 725595,
                                Name: "Jennifer Lesch",
                                Payload: false,
                                Regex: "fugit",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 882860,
                                Name: "Ellen Borer",
                                Payload: false,
                                Regex: "placeat",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 245367,
                                Name: "Stacey Rolfson",
                                Payload: false,
                                Regex: "nulla",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 379034,
                                Name: "Ryan Glover",
                                Payload: false,
                                Regex: "provident",
                            },
                        },
                        Tags: []shared.AppscopeConfigTags{
                            shared.AppscopeConfigTags{
                                Key: "molestiae",
                                Value: "magnam",
                            },
                        },
                    },
                    Env: cribl.String("odio"),
                    Hostname: cribl.String("fair-infection.info"),
                    Procname: cribl.String("rem"),
                    Username: cribl.String("Marjory_Kihn28"),
                },
                shared.AppscopeCustom{
                    Ancestor: cribl.String("eum"),
                    Arg: cribl.String("suscipit"),
                    Config: shared.AppscopeConfig{
                        Cribl: &shared.AppscopeConfigCribl{
                            Authtoken: cribl.String("assumenda"),
                            Enable: cribl.Bool(false),
                            Transport: &shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("praesentium"),
                                Path: cribl.String("quisquam"),
                                Port: cribl.Int64(86377),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("ipsa"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("id"),
                            },
                            UseScopeSourceTransport: cribl.Bool(false),
                        },
                        Event: &shared.AppscopeConfigEvent{
                            Enable: false,
                            Format: shared.AppscopeConfigEventFormat{
                                Enhancefs: false,
                                Maxeventpersec: 696997,
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("quo"),
                                Path: cribl.String("illum"),
                                Port: cribl.Int64(777408),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("fuga"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("eius"),
                            },
                            Type: shared.AppscopeConfigEventTypeNdjson,
                            Watch: []shared.AppscopeConfigEventWatch{
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("voluptas"),
                                    Headers: cribl.String("ab"),
                                    Name: cribl.String("William Goodwin"),
                                    Type: "aspernatur",
                                    Value: cribl.String("sequi"),
                                },
                            },
                        },
                        Libscope: &shared.AppscopeConfigLibscope{
                            Commanddir: cribl.String("quo"),
                            Configevent: cribl.Bool(false),
                            Log: &shared.AppscopeConfigLibscopeLog{
                                Level: shared.AppscopeConfigLibscopeLogLevelWarning.ToPointer(),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                    Host: cribl.String("aperiam"),
                                    Path: cribl.String("distinctio"),
                                    Port: cribl.Int64(799796),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("dignissimos"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("inventore"),
                                },
                            },
                            Summaryperiod: cribl.Int64(469498),
                        },
                        Metric: &shared.AppscopeConfigMetric{
                            Enable: false,
                            Format: shared.AppscopeConfigMetricFormat{
                                Statsdmaxlen: cribl.Int64(518835),
                                Statsdprefix: cribl.String("accusamus"),
                                Type: cribl.String("aliquam"),
                                Verbosity: cribl.Int64(488410),
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("commodi"),
                                Path: cribl.String("sapiente"),
                                Port: cribl.Int64(174112),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("deserunt"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("molestiae"),
                            },
                            Watch: []string{
                                "porro",
                            },
                        },
                        Payload: &shared.AppscopeConfigPayload{
                            Dir: "eum",
                            Enable: false,
                        },
                        Protocol: []shared.AppscopeConfigProtocol{
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 510017,
                                Name: "Cassandra Considine",
                                Payload: false,
                                Regex: "incidunt",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 539224,
                                Name: "Cathy Huel",
                                Payload: false,
                                Regex: "consequuntur",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 187131,
                                Name: "Kerry Mayert IV",
                                Payload: false,
                                Regex: "eveniet",
                            },
                        },
                        Tags: []shared.AppscopeConfigTags{
                            shared.AppscopeConfigTags{
                                Key: "veritatis",
                                Value: "esse",
                            },
                            shared.AppscopeConfigTags{
                                Key: "quod",
                                Value: "nam",
                            },
                            shared.AppscopeConfigTags{
                                Key: "vero",
                                Value: "aliquid",
                            },
                            shared.AppscopeConfigTags{
                                Key: "quasi",
                                Value: "saepe",
                            },
                        },
                    },
                    Env: cribl.String("vel"),
                    Hostname: cribl.String("relieved-investigator.net"),
                    Procname: cribl.String("occaecati"),
                    Username: cribl.String("Erika.Rau63"),
                },
            },
            Event: &shared.AppscopeConfigWithCustomEvent{
                Enable: false,
                Format: shared.AppscopeConfigWithCustomEventFormat{
                    Enhancefs: false,
                    Maxeventpersec: 731398,
                },
                Transport: shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                    Host: cribl.String("cumque"),
                    Path: cribl.String("consequuntur"),
                    Port: cribl.Int64(9766),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("minus"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("quaerat"),
                },
                Type: shared.AppscopeConfigWithCustomEventTypeNdjson,
                Watch: []shared.AppscopeConfigWithCustomEventWatch{
                    shared.AppscopeConfigWithCustomEventWatch{
                        Allowbinary: cribl.Bool(false),
                        Enabled: cribl.Bool(false),
                        Field: cribl.String("consectetur"),
                        Headers: cribl.String("esse"),
                        Name: cribl.String("Eduardo Wilkinson"),
                        Type: "esse",
                        Value: cribl.String("quasi"),
                    },
                    shared.AppscopeConfigWithCustomEventWatch{
                        Allowbinary: cribl.Bool(false),
                        Enabled: cribl.Bool(false),
                        Field: cribl.String("a"),
                        Headers: cribl.String("error"),
                        Name: cribl.String("Jody Schuster"),
                        Type: "asperiores",
                        Value: cribl.String("facere"),
                    },
                    shared.AppscopeConfigWithCustomEventWatch{
                        Allowbinary: cribl.Bool(false),
                        Enabled: cribl.Bool(false),
                        Field: cribl.String("veritatis"),
                        Headers: cribl.String("consequuntur"),
                        Name: cribl.String("Hattie Nader"),
                        Type: "quae",
                        Value: cribl.String("earum"),
                    },
                    shared.AppscopeConfigWithCustomEventWatch{
                        Allowbinary: cribl.Bool(false),
                        Enabled: cribl.Bool(false),
                        Field: cribl.String("vel"),
                        Headers: cribl.String("in"),
                        Name: cribl.String("Jeannette Stanton II"),
                        Type: "sapiente",
                        Value: cribl.String("dicta"),
                    },
                },
            },
            Libscope: &shared.AppscopeConfigWithCustomLibscope{
                Commanddir: cribl.String("ullam"),
                Configevent: cribl.Bool(false),
                Log: &shared.AppscopeConfigWithCustomLibscopeLog{
                    Level: shared.AppscopeConfigWithCustomLibscopeLogLevelWarning.ToPointer(),
                    Transport: &shared.AppscopeTransport{
                        Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                        Host: cribl.String("nisi"),
                        Path: cribl.String("aut"),
                        Port: cribl.Int64(531849),
                        TLS: &shared.AppscopeTransportTLS{
                            Cacertpath: cribl.String("qui"),
                            Enable: cribl.Bool(false),
                            Validateserver: cribl.Bool(false),
                        },
                        Type: cribl.String("quibusdam"),
                    },
                },
                Summaryperiod: cribl.Int64(401259),
            },
            Metric: &shared.AppscopeConfigWithCustomMetric{
                Enable: false,
                Format: shared.AppscopeConfigWithCustomMetricFormat{
                    Statsdmaxlen: cribl.Int64(536275),
                    Statsdprefix: cribl.String("itaque"),
                    Type: cribl.String("dolorum"),
                    Verbosity: cribl.Int64(99615),
                },
                Transport: shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                    Host: cribl.String("tenetur"),
                    Path: cribl.String("quasi"),
                    Port: cribl.Int64(869489),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("et"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("voluptate"),
                },
                Watch: []string{
                    "minima",
                },
            },
            Payload: &shared.AppscopeConfigWithCustomPayload{
                Dir: "veritatis",
                Enable: false,
            },
            Protocol: []shared.AppscopeConfigWithCustomProtocol{
                shared.AppscopeConfigWithCustomProtocol{
                    Binary: false,
                    Detect: false,
                    Len: 237173,
                    Name: "Lionel Bartoletti IV",
                    Payload: false,
                    Regex: "eum",
                },
            },
            Tags: []shared.AppscopeConfigWithCustomTags{
                shared.AppscopeConfigWithCustomTags{
                    Key: "ab",
                    Value: "corrupti",
                },
                shared.AppscopeConfigWithCustomTags{
                    Key: "non",
                    Value: "voluptatem",
                },
                shared.AppscopeConfigWithCustomTags{
                    Key: "dolor",
                    Value: "occaecati",
                },
            },
        },
        Description: "numquam",
        ID: "c26071f9-3f5f-4064-adac-7af515cc413a",
        Lib: shared.CriblLibCriblCustom,
        Tags: cribl.String("suscipit"),
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
    id := "velit"

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
    id := "culpa"

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
    id := "est"
    appscopeLibEntry := &shared.AppscopeLibEntry{
        Config: shared.AppscopeConfigWithCustom{
            Cribl: &shared.AppscopeConfigWithCustomCribl{
                Authtoken: cribl.String("recusandae"),
                Enable: cribl.Bool(false),
                Transport: &shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                    Host: cribl.String("fugiat"),
                    Path: cribl.String("vel"),
                    Port: cribl.Int64(497678),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("quos"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("vel"),
                },
                UseScopeSourceTransport: cribl.Bool(false),
            },
            Custom: []shared.AppscopeCustom{
                shared.AppscopeCustom{
                    Ancestor: cribl.String("possimus"),
                    Arg: cribl.String("facilis"),
                    Config: shared.AppscopeConfig{
                        Cribl: &shared.AppscopeConfigCribl{
                            Authtoken: cribl.String("cum"),
                            Enable: cribl.Bool(false),
                            Transport: &shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("in"),
                                Path: cribl.String("corporis"),
                                Port: cribl.Int64(968904),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("assumenda"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("nemo"),
                            },
                            UseScopeSourceTransport: cribl.Bool(false),
                        },
                        Event: &shared.AppscopeConfigEvent{
                            Enable: false,
                            Format: shared.AppscopeConfigEventFormat{
                                Enhancefs: false,
                                Maxeventpersec: 924967,
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("aperiam"),
                                Path: cribl.String("cum"),
                                Port: cribl.Int64(232627),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("in"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("exercitationem"),
                            },
                            Type: shared.AppscopeConfigEventTypeNdjson,
                            Watch: []shared.AppscopeConfigEventWatch{
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("facere"),
                                    Headers: cribl.String("numquam"),
                                    Name: cribl.String("Ricardo Wisozk"),
                                    Type: "necessitatibus",
                                    Value: cribl.String("dolore"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("sunt"),
                                    Headers: cribl.String("asperiores"),
                                    Name: cribl.String("Ms. Ethel Farrell"),
                                    Type: "debitis",
                                    Value: cribl.String("consectetur"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("corporis"),
                                    Headers: cribl.String("harum"),
                                    Name: cribl.String("Melissa Von PhD"),
                                    Type: "similique",
                                    Value: cribl.String("tempora"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("aspernatur"),
                                    Headers: cribl.String("voluptas"),
                                    Name: cribl.String("Melanie Hane"),
                                    Type: "adipisci",
                                    Value: cribl.String("minus"),
                                },
                            },
                        },
                        Libscope: &shared.AppscopeConfigLibscope{
                            Commanddir: cribl.String("dolores"),
                            Configevent: cribl.Bool(false),
                            Log: &shared.AppscopeConfigLibscopeLog{
                                Level: shared.AppscopeConfigLibscopeLogLevelWarning.ToPointer(),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("dolore"),
                                    Path: cribl.String("aliquam"),
                                    Port: cribl.Int64(885963),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("temporibus"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("ullam"),
                                },
                            },
                            Summaryperiod: cribl.Int64(237742),
                        },
                        Metric: &shared.AppscopeConfigMetric{
                            Enable: false,
                            Format: shared.AppscopeConfigMetricFormat{
                                Statsdmaxlen: cribl.Int64(738391),
                                Statsdprefix: cribl.String("blanditiis"),
                                Type: cribl.String("quas"),
                                Verbosity: cribl.Int64(942584),
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("culpa"),
                                Path: cribl.String("corrupti"),
                                Port: cribl.Int64(867290),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("totam"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("hic"),
                            },
                            Watch: []string{
                                "nobis",
                                "sit",
                            },
                        },
                        Payload: &shared.AppscopeConfigPayload{
                            Dir: "rerum",
                            Enable: false,
                        },
                        Protocol: []shared.AppscopeConfigProtocol{
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 967966,
                                Name: "Jaime Quigley",
                                Payload: false,
                                Regex: "ab",
                            },
                        },
                        Tags: []shared.AppscopeConfigTags{
                            shared.AppscopeConfigTags{
                                Key: "dolore",
                                Value: "laborum",
                            },
                            shared.AppscopeConfigTags{
                                Key: "sed",
                                Value: "in",
                            },
                            shared.AppscopeConfigTags{
                                Key: "commodi",
                                Value: "quidem",
                            },
                        },
                    },
                    Env: cribl.String("explicabo"),
                    Hostname: cribl.String("hearty-oats.com"),
                    Procname: cribl.String("suscipit"),
                    Username: cribl.String("Vickie96"),
                },
                shared.AppscopeCustom{
                    Ancestor: cribl.String("perferendis"),
                    Arg: cribl.String("corrupti"),
                    Config: shared.AppscopeConfig{
                        Cribl: &shared.AppscopeConfigCribl{
                            Authtoken: cribl.String("maiores"),
                            Enable: cribl.Bool(false),
                            Transport: &shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("sed"),
                                Path: cribl.String("provident"),
                                Port: cribl.Int64(258702),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("necessitatibus"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("ipsum"),
                            },
                            UseScopeSourceTransport: cribl.Bool(false),
                        },
                        Event: &shared.AppscopeConfigEvent{
                            Enable: false,
                            Format: shared.AppscopeConfigEventFormat{
                                Enhancefs: false,
                                Maxeventpersec: 406733,
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("quos"),
                                Path: cribl.String("voluptatibus"),
                                Port: cribl.Int64(271653),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("tempora"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("voluptate"),
                            },
                            Type: shared.AppscopeConfigEventTypeNdjson,
                            Watch: []shared.AppscopeConfigEventWatch{
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("ex"),
                                    Headers: cribl.String("sit"),
                                    Name: cribl.String("Cecelia Lakin"),
                                    Type: "incidunt",
                                    Value: cribl.String("ipsam"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("debitis"),
                                    Headers: cribl.String("rem"),
                                    Name: cribl.String("Della Muller"),
                                    Type: "recusandae",
                                    Value: cribl.String("reiciendis"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("nulla"),
                                    Headers: cribl.String("magni"),
                                    Name: cribl.String("Gwen Fritsch"),
                                    Type: "officiis",
                                    Value: cribl.String("beatae"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("laudantium"),
                                    Headers: cribl.String("exercitationem"),
                                    Name: cribl.String("Bennie Howe"),
                                    Type: "error",
                                    Value: cribl.String("hic"),
                                },
                            },
                        },
                        Libscope: &shared.AppscopeConfigLibscope{
                            Commanddir: cribl.String("expedita"),
                            Configevent: cribl.Bool(false),
                            Log: &shared.AppscopeConfigLibscopeLog{
                                Level: shared.AppscopeConfigLibscopeLogLevelNone.ToPointer(),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("dolorum"),
                                    Path: cribl.String("nostrum"),
                                    Port: cribl.Int64(639028),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("dolorum"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("corrupti"),
                                },
                            },
                            Summaryperiod: cribl.Int64(879235),
                        },
                        Metric: &shared.AppscopeConfigMetric{
                            Enable: false,
                            Format: shared.AppscopeConfigMetricFormat{
                                Statsdmaxlen: cribl.Int64(272683),
                                Statsdprefix: cribl.String("atque"),
                                Type: cribl.String("fugit"),
                                Verbosity: cribl.Int64(282699),
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("voluptatem"),
                                Path: cribl.String("culpa"),
                                Port: cribl.Int64(710337),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("magnam"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("consequatur"),
                            },
                            Watch: []string{
                                "ipsam",
                                "sit",
                            },
                        },
                        Payload: &shared.AppscopeConfigPayload{
                            Dir: "voluptatum",
                            Enable: false,
                        },
                        Protocol: []shared.AppscopeConfigProtocol{
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 922112,
                                Name: "Janet Kuvalis",
                                Payload: false,
                                Regex: "sit",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 425508,
                                Name: "Mrs. Tricia Mueller",
                                Payload: false,
                                Regex: "dolorem",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 690894,
                                Name: "Diane Mayer",
                                Payload: false,
                                Regex: "atque",
                            },
                        },
                        Tags: []shared.AppscopeConfigTags{
                            shared.AppscopeConfigTags{
                                Key: "nam",
                                Value: "tenetur",
                            },
                            shared.AppscopeConfigTags{
                                Key: "laboriosam",
                                Value: "alias",
                            },
                            shared.AppscopeConfigTags{
                                Key: "amet",
                                Value: "deserunt",
                            },
                        },
                    },
                    Env: cribl.String("voluptate"),
                    Hostname: cribl.String("outlying-watercress.name"),
                    Procname: cribl.String("repellendus"),
                    Username: cribl.String("Vincent66"),
                },
            },
            Event: &shared.AppscopeConfigWithCustomEvent{
                Enable: false,
                Format: shared.AppscopeConfigWithCustomEventFormat{
                    Enhancefs: false,
                    Maxeventpersec: 696483,
                },
                Transport: shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                    Host: cribl.String("facere"),
                    Path: cribl.String("fuga"),
                    Port: cribl.Int64(509807),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("mollitia"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("veniam"),
                },
                Type: shared.AppscopeConfigWithCustomEventTypeNdjson,
                Watch: []shared.AppscopeConfigWithCustomEventWatch{
                    shared.AppscopeConfigWithCustomEventWatch{
                        Allowbinary: cribl.Bool(false),
                        Enabled: cribl.Bool(false),
                        Field: cribl.String("quisquam"),
                        Headers: cribl.String("repudiandae"),
                        Name: cribl.String("Kay Kihn"),
                        Type: "suscipit",
                        Value: cribl.String("quidem"),
                    },
                },
            },
            Libscope: &shared.AppscopeConfigWithCustomLibscope{
                Commanddir: cribl.String("maxime"),
                Configevent: cribl.Bool(false),
                Log: &shared.AppscopeConfigWithCustomLibscopeLog{
                    Level: shared.AppscopeConfigWithCustomLibscopeLogLevelDebug.ToPointer(),
                    Transport: &shared.AppscopeTransport{
                        Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                        Host: cribl.String("amet"),
                        Path: cribl.String("assumenda"),
                        Port: cribl.Int64(410301),
                        TLS: &shared.AppscopeTransportTLS{
                            Cacertpath: cribl.String("atque"),
                            Enable: cribl.Bool(false),
                            Validateserver: cribl.Bool(false),
                        },
                        Type: cribl.String("error"),
                    },
                },
                Summaryperiod: cribl.Int64(887265),
            },
            Metric: &shared.AppscopeConfigWithCustomMetric{
                Enable: false,
                Format: shared.AppscopeConfigWithCustomMetricFormat{
                    Statsdmaxlen: cribl.Int64(886961),
                    Statsdprefix: cribl.String("accusamus"),
                    Type: cribl.String("natus"),
                    Verbosity: cribl.Int64(328303),
                },
                Transport: shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                    Host: cribl.String("ex"),
                    Path: cribl.String("maiores"),
                    Port: cribl.Int64(544647),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("at"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("error"),
                },
                Watch: []string{
                    "suscipit",
                    "repudiandae",
                    "atque",
                },
            },
            Payload: &shared.AppscopeConfigWithCustomPayload{
                Dir: "atque",
                Enable: false,
            },
            Protocol: []shared.AppscopeConfigWithCustomProtocol{
                shared.AppscopeConfigWithCustomProtocol{
                    Binary: false,
                    Detect: false,
                    Len: 923306,
                    Name: "Mack Grant DVM",
                    Payload: false,
                    Regex: "dicta",
                },
            },
            Tags: []shared.AppscopeConfigWithCustomTags{
                shared.AppscopeConfigWithCustomTags{
                    Key: "beatae",
                    Value: "dolores",
                },
            },
        },
        Description: "enim",
        ID: "63f94e29-e973-4e92-aa57-a15be3e06080",
        Lib: shared.CriblLibCriblCustom,
        Tags: cribl.String("eveniet"),
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

