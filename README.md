<div align="center">
    <img src="https://github.com/speakeasy-sdks/cribl-demo-go/assets/68016351/3c85f178-5ab2-4679-b0a7-c31ecdcce367" width="350px">
    <h1>Cribl Go SDK</h1>
   <p></p>
   <a href="https://docs.cribl.io/api/"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
<!--    <a href="https://github.com/speakeasy-sdks/cribl-demo-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/cribl-demo-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a> -->
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/crible-demo-go/releases"><img src="https://img.shields.io/github/v/release/speakeasy-sdks/cribl-demo-go?sort=semver&style=for-the-badge" /></a>
</div>

## Authentication
Please fetch a Bearer token for the Cribl Cloud free tier [here](https://docs.cribl.io/stream/api-tutorials/#criblcloud-free-tier)

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/cribl-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->


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
                Authtoken: cribl.String("corrupti"),
                Enable: cribl.Bool(false),
                Transport: &shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                    Host: cribl.String("distinctio"),
                    Path: cribl.String("quibusdam"),
                    Port: cribl.Int64(602763),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("nulla"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("corrupti"),
                },
                UseScopeSourceTransport: cribl.Bool(false),
            },
            Custom: []shared.AppscopeCustom{
                shared.AppscopeCustom{
                    Ancestor: cribl.String("vel"),
                    Arg: cribl.String("error"),
                    Config: shared.AppscopeConfig{
                        Cribl: &shared.AppscopeConfigCribl{
                            Authtoken: cribl.String("deserunt"),
                            Enable: cribl.Bool(false),
                            Transport: &shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("iure"),
                                Path: cribl.String("magnam"),
                                Port: cribl.Int64(891773),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("ipsa"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("delectus"),
                            },
                            UseScopeSourceTransport: cribl.Bool(false),
                        },
                        Event: &shared.AppscopeConfigEvent{
                            Enable: false,
                            Format: shared.AppscopeConfigEventFormat{
                                Enhancefs: false,
                                Maxeventpersec: 272656,
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("molestiae"),
                                Path: cribl.String("minus"),
                                Port: cribl.Int64(812169),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("voluptatum"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("iusto"),
                            },
                            Type: shared.AppscopeConfigEventTypeNdjson,
                            Watch: []shared.AppscopeConfigEventWatch{
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("nisi"),
                                    Headers: cribl.String("recusandae"),
                                    Name: cribl.String("Miss Raymond Hauck III"),
                                    Type: "repellendus",
                                    Value: cribl.String("sapiente"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("quo"),
                                    Headers: cribl.String("odit"),
                                    Name: cribl.String("Wilfred Wolff"),
                                    Type: "quod",
                                    Value: cribl.String("esse"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("totam"),
                                    Headers: cribl.String("porro"),
                                    Name: cribl.String("Samuel Reichel"),
                                    Type: "fugit",
                                    Value: cribl.String("deleniti"),
                                },
                            },
                        },
                        Libscope: &shared.AppscopeConfigLibscope{
                            Commanddir: cribl.String("hic"),
                            Configevent: cribl.Bool(false),
                            Log: &shared.AppscopeConfigLibscopeLog{
                                Level: shared.AppscopeConfigLibscopeLogLevelError.ToPointer(),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                    Host: cribl.String("beatae"),
                                    Path: cribl.String("commodi"),
                                    Port: cribl.Int64(473600),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("modi"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("qui"),
                                },
                            },
                            Summaryperiod: cribl.Int64(774234),
                        },
                        Metric: &shared.AppscopeConfigMetric{
                            Enable: false,
                            Format: shared.AppscopeConfigMetricFormat{
                                Statsdmaxlen: cribl.Int64(736918),
                                Statsdprefix: cribl.String("esse"),
                                Type: cribl.String("ipsum"),
                                Verbosity: cribl.Int64(568434),
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("perferendis"),
                                Path: cribl.String("ad"),
                                Port: cribl.Int64(617636),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("sed"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("iste"),
                            },
                            Watch: []string{
                                "natus",
                            },
                        },
                        Payload: &shared.AppscopeConfigPayload{
                            Dir: "laboriosam",
                            Enable: false,
                        },
                        Protocol: []shared.AppscopeConfigProtocol{
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 902599,
                                Name: "Harvey Hessel",
                                Payload: false,
                                Regex: "saepe",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 697631,
                                Name: "Brenda Wisozk",
                                Payload: false,
                                Regex: "laborum",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 170909,
                                Name: "Stacy Champlin",
                                Payload: false,
                                Regex: "omnis",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 363711,
                                Name: "Velma Batz",
                                Payload: false,
                                Regex: "doloribus",
                            },
                        },
                        Tags: []shared.AppscopeConfigTags{
                            shared.AppscopeConfigTags{
                                Key: "architecto",
                                Value: "mollitia",
                            },
                            shared.AppscopeConfigTags{
                                Key: "dolorem",
                                Value: "culpa",
                            },
                            shared.AppscopeConfigTags{
                                Key: "consequuntur",
                                Value: "repellat",
                            },
                            shared.AppscopeConfigTags{
                                Key: "mollitia",
                                Value: "occaecati",
                            },
                        },
                    },
                    Env: cribl.String("numquam"),
                    Hostname: cribl.String("immediate-instructor.info"),
                    Procname: cribl.String("velit"),
                    Username: cribl.String("Linda.Cronin"),
                },
                shared.AppscopeCustom{
                    Ancestor: cribl.String("laborum"),
                    Arg: cribl.String("animi"),
                    Config: shared.AppscopeConfig{
                        Cribl: &shared.AppscopeConfigCribl{
                            Authtoken: cribl.String("enim"),
                            Enable: cribl.Bool(false),
                            Transport: &shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("quo"),
                                Path: cribl.String("sequi"),
                                Port: cribl.Int64(949572),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("ipsam"),
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
                                Maxeventpersec: 820994,
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("quasi"),
                                Path: cribl.String("error"),
                                Port: cribl.Int64(837945),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("laborum"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("quasi"),
                            },
                            Type: shared.AppscopeConfigEventTypeNdjson,
                            Watch: []shared.AppscopeConfigEventWatch{
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("voluptatibus"),
                                    Headers: cribl.String("vero"),
                                    Name: cribl.String("Miss Irma Wolff"),
                                    Type: "cum",
                                    Value: cribl.String("perferendis"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("doloremque"),
                                    Headers: cribl.String("reprehenderit"),
                                    Name: cribl.String("Shawna Carter"),
                                    Type: "iusto",
                                    Value: cribl.String("dicta"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("harum"),
                                    Headers: cribl.String("enim"),
                                    Name: cribl.String("Mrs. Leslie VonRueden"),
                                    Type: "molestias",
                                    Value: cribl.String("excepturi"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("pariatur"),
                                    Headers: cribl.String("modi"),
                                    Name: cribl.String("Dr. Jordan Von"),
                                    Type: "veritatis",
                                    Value: cribl.String("itaque"),
                                },
                            },
                        },
                        Libscope: &shared.AppscopeConfigLibscope{
                            Commanddir: cribl.String("incidunt"),
                            Configevent: cribl.Bool(false),
                            Log: &shared.AppscopeConfigLibscopeLog{
                                Level: shared.AppscopeConfigLibscopeLogLevelInfo.ToPointer(),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("est"),
                                    Path: cribl.String("quibusdam"),
                                    Port: cribl.Int64(131797),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("deserunt"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("distinctio"),
                                },
                            },
                            Summaryperiod: cribl.Int64(841386),
                        },
                        Metric: &shared.AppscopeConfigMetric{
                            Enable: false,
                            Format: shared.AppscopeConfigMetricFormat{
                                Statsdmaxlen: cribl.Int64(289406),
                                Statsdprefix: cribl.String("modi"),
                                Type: cribl.String("qui"),
                                Verbosity: cribl.Int64(397821),
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("quos"),
                                Path: cribl.String("perferendis"),
                                Port: cribl.Int64(164940),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("assumenda"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("ipsam"),
                            },
                            Watch: []string{
                                "fugit",
                            },
                        },
                        Payload: &shared.AppscopeConfigPayload{
                            Dir: "dolorum",
                            Enable: false,
                        },
                        Protocol: []shared.AppscopeConfigProtocol{
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 270008,
                                Name: "Geoffrey Green",
                                Payload: false,
                                Regex: "non",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 756107,
                                Name: "Gilbert Medhurst",
                                Payload: false,
                                Regex: "officia",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 223081,
                                Name: "Randal Parisian",
                                Payload: false,
                                Regex: "illum",
                            },
                        },
                        Tags: []shared.AppscopeConfigTags{
                            shared.AppscopeConfigTags{
                                Key: "rerum",
                                Value: "dicta",
                            },
                            shared.AppscopeConfigTags{
                                Key: "magnam",
                                Value: "cumque",
                            },
                            shared.AppscopeConfigTags{
                                Key: "facere",
                                Value: "ea",
                            },
                            shared.AppscopeConfigTags{
                                Key: "aliquid",
                                Value: "laborum",
                            },
                        },
                    },
                    Env: cribl.String("accusamus"),
                    Hostname: cribl.String("exemplary-mover.biz"),
                    Procname: cribl.String("accusamus"),
                    Username: cribl.String("Virgil_Pouros"),
                },
                shared.AppscopeCustom{
                    Ancestor: cribl.String("id"),
                    Arg: cribl.String("blanditiis"),
                    Config: shared.AppscopeConfig{
                        Cribl: &shared.AppscopeConfigCribl{
                            Authtoken: cribl.String("deleniti"),
                            Enable: cribl.Bool(false),
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
                            UseScopeSourceTransport: cribl.Bool(false),
                        },
                        Event: &shared.AppscopeConfigEvent{
                            Enable: false,
                            Format: shared.AppscopeConfigEventFormat{
                                Enhancefs: false,
                                Maxeventpersec: 606393,
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("perferendis"),
                                Path: cribl.String("nihil"),
                                Port: cribl.Int64(301575),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("distinctio"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("id"),
                            },
                            Type: shared.AppscopeConfigEventTypeNdjson,
                            Watch: []shared.AppscopeConfigEventWatch{
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("labore"),
                                    Headers: cribl.String("suscipit"),
                                    Name: cribl.String("Robin Keebler"),
                                    Type: "architecto",
                                    Value: cribl.String("magnam"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("et"),
                                    Headers: cribl.String("excepturi"),
                                    Name: cribl.String("Ramona Lueilwitz MD"),
                                    Type: "reiciendis",
                                    Value: cribl.String("mollitia"),
                                },
                            },
                        },
                        Libscope: &shared.AppscopeConfigLibscope{
                            Commanddir: cribl.String("ad"),
                            Configevent: cribl.Bool(false),
                            Log: &shared.AppscopeConfigLibscopeLog{
                                Level: shared.AppscopeConfigLibscopeLogLevelWarning.ToPointer(),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("necessitatibus"),
                                    Path: cribl.String("odit"),
                                    Port: cribl.Int64(367562),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("quasi"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("iure"),
                                },
                            },
                            Summaryperiod: cribl.Int64(984043),
                        },
                        Metric: &shared.AppscopeConfigMetric{
                            Enable: false,
                            Format: shared.AppscopeConfigMetricFormat{
                                Statsdmaxlen: cribl.Int64(891924),
                                Statsdprefix: cribl.String("eius"),
                                Type: cribl.String("maxime"),
                                Verbosity: cribl.Int64(537023),
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("in"),
                                Path: cribl.String("architecto"),
                                Port: cribl.Int64(99569),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("repudiandae"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("ullam"),
                            },
                            Watch: []string{
                                "nihil",
                                "repellat",
                                "quibusdam",
                            },
                        },
                        Payload: &shared.AppscopeConfigPayload{
                            Dir: "sed",
                            Enable: false,
                        },
                        Protocol: []shared.AppscopeConfigProtocol{
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 868126,
                                Name: "Kathryn Lang",
                                Payload: false,
                                Regex: "sunt",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 779051,
                                Name: "Ervin Schoen",
                                Payload: false,
                                Regex: "odit",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 407183,
                                Name: "Virginia Wunsch",
                                Payload: false,
                                Regex: "voluptate",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 420075,
                                Name: "Gary Streich",
                                Payload: false,
                                Regex: "perferendis",
                            },
                        },
                        Tags: []shared.AppscopeConfigTags{
                            shared.AppscopeConfigTags{
                                Key: "amet",
                                Value: "aut",
                            },
                            shared.AppscopeConfigTags{
                                Key: "cumque",
                                Value: "corporis",
                            },
                            shared.AppscopeConfigTags{
                                Key: "hic",
                                Value: "libero",
                            },
                            shared.AppscopeConfigTags{
                                Key: "nobis",
                                Value: "dolores",
                            },
                        },
                    },
                    Env: cribl.String("quis"),
                    Hostname: cribl.String("mealy-kilometer.com"),
                    Procname: cribl.String("quis"),
                    Username: cribl.String("Cody17"),
                },
                shared.AppscopeCustom{
                    Ancestor: cribl.String("minus"),
                    Arg: cribl.String("quam"),
                    Config: shared.AppscopeConfig{
                        Cribl: &shared.AppscopeConfigCribl{
                            Authtoken: cribl.String("dolor"),
                            Enable: cribl.Bool(false),
                            Transport: &shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("nostrum"),
                                Path: cribl.String("hic"),
                                Port: cribl.Int64(928082),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("omnis"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("facilis"),
                            },
                            UseScopeSourceTransport: cribl.Bool(false),
                        },
                        Event: &shared.AppscopeConfigEvent{
                            Enable: false,
                            Format: shared.AppscopeConfigEventFormat{
                                Enhancefs: false,
                                Maxeventpersec: 596656,
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("porro"),
                                Path: cribl.String("consequuntur"),
                                Port: cribl.Int64(500026),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("error"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("eaque"),
                            },
                            Type: shared.AppscopeConfigEventTypeNdjson,
                            Watch: []shared.AppscopeConfigEventWatch{
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("rerum"),
                                    Headers: cribl.String("adipisci"),
                                    Name: cribl.String("Merle Gleichner"),
                                    Type: "deleniti",
                                    Value: cribl.String("pariatur"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("provident"),
                                    Headers: cribl.String("nobis"),
                                    Name: cribl.String("Toby Hahn"),
                                    Type: "dolorem",
                                    Value: cribl.String("dolorem"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("dolor"),
                                    Headers: cribl.String("qui"),
                                    Name: cribl.String("Mindy Marks"),
                                    Type: "dignissimos",
                                    Value: cribl.String("reiciendis"),
                                },
                            },
                        },
                        Libscope: &shared.AppscopeConfigLibscope{
                            Commanddir: cribl.String("amet"),
                            Configevent: cribl.Bool(false),
                            Log: &shared.AppscopeConfigLibscopeLog{
                                Level: shared.AppscopeConfigLibscopeLogLevelError.ToPointer(),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("veritatis"),
                                    Path: cribl.String("ipsa"),
                                    Port: cribl.Int64(56418),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("iure"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("odio"),
                                },
                            },
                            Summaryperiod: cribl.Int64(311796),
                        },
                        Metric: &shared.AppscopeConfigMetric{
                            Enable: false,
                            Format: shared.AppscopeConfigMetricFormat{
                                Statsdmaxlen: cribl.Int64(881005),
                                Statsdprefix: cribl.String("quidem"),
                                Type: cribl.String("voluptatibus"),
                                Verbosity: cribl.Int64(377752),
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("eos"),
                                Path: cribl.String("atque"),
                                Port: cribl.Int64(24678),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("fugiat"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("ab"),
                            },
                            Watch: []string{
                                "dolorum",
                                "iusto",
                                "voluptate",
                            },
                        },
                        Payload: &shared.AppscopeConfigPayload{
                            Dir: "dolorum",
                            Enable: false,
                        },
                        Protocol: []shared.AppscopeConfigProtocol{
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 607045,
                                Name: "Kelvin Zboncak",
                                Payload: false,
                                Regex: "voluptate",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 663078,
                                Name: "Mrs. Ray Collins",
                                Payload: false,
                                Regex: "accusamus",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 320017,
                                Name: "Sam Oberbrunner",
                                Payload: false,
                                Regex: "repellendus",
                            },
                        },
                        Tags: []shared.AppscopeConfigTags{
                            shared.AppscopeConfigTags{
                                Key: "similique",
                                Value: "alias",
                            },
                            shared.AppscopeConfigTags{
                                Key: "at",
                                Value: "quaerat",
                            },
                            shared.AppscopeConfigTags{
                                Key: "tempora",
                                Value: "vel",
                            },
                        },
                    },
                    Env: cribl.String("quod"),
                    Hostname: cribl.String("uneven-commitment.net"),
                    Procname: cribl.String("a"),
                    Username: cribl.String("Jacky.Pfeffer"),
                },
            },
            Event: &shared.AppscopeConfigWithCustomEvent{
                Enable: false,
                Format: shared.AppscopeConfigWithCustomEventFormat{
                    Enhancefs: false,
                    Maxeventpersec: 788740,
                },
                Transport: shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                    Host: cribl.String("amet"),
                    Path: cribl.String("tempore"),
                    Port: cribl.Int64(880298),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("numquam"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("enim"),
                },
                Type: shared.AppscopeConfigWithCustomEventTypeNdjson,
                Watch: []shared.AppscopeConfigWithCustomEventWatch{
                    shared.AppscopeConfigWithCustomEventWatch{
                        Allowbinary: cribl.Bool(false),
                        Enabled: cribl.Bool(false),
                        Field: cribl.String("sapiente"),
                        Headers: cribl.String("totam"),
                        Name: cribl.String("Karen Rath"),
                        Type: "vel",
                        Value: cribl.String("libero"),
                    },
                },
            },
            Libscope: &shared.AppscopeConfigWithCustomLibscope{
                Commanddir: cribl.String("voluptas"),
                Configevent: cribl.Bool(false),
                Log: &shared.AppscopeConfigWithCustomLibscopeLog{
                    Level: shared.AppscopeConfigWithCustomLibscopeLogLevelError.ToPointer(),
                    Transport: &shared.AppscopeTransport{
                        Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                        Host: cribl.String("ipsum"),
                        Path: cribl.String("incidunt"),
                        Port: cribl.Int64(186458),
                        TLS: &shared.AppscopeTransportTLS{
                            Cacertpath: cribl.String("cupiditate"),
                            Enable: cribl.Bool(false),
                            Validateserver: cribl.Bool(false),
                        },
                        Type: cribl.String("maxime"),
                    },
                },
                Summaryperiod: cribl.Int64(863856),
            },
            Metric: &shared.AppscopeConfigWithCustomMetric{
                Enable: false,
                Format: shared.AppscopeConfigWithCustomMetricFormat{
                    Statsdmaxlen: cribl.Int64(747080),
                    Statsdprefix: cribl.String("dicta"),
                    Type: cribl.String("laborum"),
                    Verbosity: cribl.Int64(517379),
                },
                Transport: shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                    Host: cribl.String("aspernatur"),
                    Path: cribl.String("dolores"),
                    Port: cribl.Int64(716860),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("facilis"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("aliquid"),
                },
                Watch: []string{
                    "molestias",
                    "temporibus",
                },
            },
            Payload: &shared.AppscopeConfigWithCustomPayload{
                Dir: "qui",
                Enable: false,
            },
            Protocol: []shared.AppscopeConfigWithCustomProtocol{
                shared.AppscopeConfigWithCustomProtocol{
                    Binary: false,
                    Detect: false,
                    Len: 144847,
                    Name: "Courtney Cassin",
                    Payload: false,
                    Regex: "hic",
                },
            },
            Tags: []shared.AppscopeConfigWithCustomTags{
                shared.AppscopeConfigWithCustomTags{
                    Key: "cumque",
                    Value: "soluta",
                },
            },
        },
        Description: "nobis",
        ID: "1e31b8b9-0f34-443a-9108-e0adcf4b9218",
        Lib: shared.CriblLibCriblCustom,
        Tags: cribl.String("occaecati"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppscopeLibEntry != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [AppscopeConfigs](docs/sdks/appscopeconfigs/README.md)

* [Create](docs/sdks/appscopeconfigs/README.md#create) - Create AppscopeLibEntry
* [Delete](docs/sdks/appscopeconfigs/README.md#delete) - Delete AppscopeLibEntry
* [Get](docs/sdks/appscopeconfigs/README.md#get) - Get AppscopeLibEntry by ID
* [ListAppscopeLibEntries](docs/sdks/appscopeconfigs/README.md#listappscopelibentries) - Get a list of AppscopeLibEntry objects
* [Update](docs/sdks/appscopeconfigs/README.md#update) - Update AppscopeLibEntry

### [Auth](docs/sdks/auth/README.md)

* [IDPlogout](docs/sdks/auth/README.md#idplogout) - Accepts a logout request from an IDP and logs out the user
* [AcceptIDPRequest](docs/sdks/auth/README.md#acceptidprequest) - Accepts an authentication request from an IDP and authenticates the user
* [Get](docs/sdks/auth/README.md#get) - Get IDP used for an authorization code callback
* [GetMetadata](docs/sdks/auth/README.md#getmetadata) - Obtain metadata which Cribl Stream/Edge uses when acting as a Service Provider
* [GetRedirect](docs/sdks/auth/README.md#getredirect) - Obtain redirect information
* [ListAuthGroup](docs/sdks/auth/README.md#listauthgroup) - List the external authentication system's groups
* [Login](docs/sdks/auth/README.md#login) - Log in and obtain Auth token
* [Logout](docs/sdks/auth/README.md#logout) - Log current user out
* [LogoutRedirect](docs/sdks/auth/README.md#logoutredirect) - Redirect user to IDP with logout request
* [Post](docs/sdks/auth/README.md#post) - API call that the IDP should use for an authentication request
* [RequestLogout](docs/sdks/auth/README.md#requestlogout) - API call that the IDP should use for a logout request

### [Authorize](docs/sdks/authorize/README.md)

* [ListAuthorizations](docs/sdks/authorize/README.md#listauthorizations) - get the client's authorization policy
* [ListClientRoles](docs/sdks/authorize/README.md#listclientroles) - get the client's roles

### [Certificates](docs/sdks/certificates/README.md)

* [Create](docs/sdks/certificates/README.md#create) - Create Certificate
* [Delete](docs/sdks/certificates/README.md#delete) - Delete Certificate
* [Get](docs/sdks/certificates/README.md#get) - Get Certificate by ID
* [ListCertificates](docs/sdks/certificates/README.md#listcertificates) - Get a list of Certificate objects
* [Update](docs/sdks/certificates/README.md#update) - Update Certificate

### [Changelog](docs/sdks/changelog/README.md)

* [ListChangelogs](docs/sdks/changelog/README.md#listchangelogs) - Get changelog viewed state

### [Clui](docs/sdks/clui/README.md)

* [ListCluisResults](docs/sdks/clui/README.md#listcluisresults) - Get CLUI search results

### [Collectors](docs/sdks/collectors/README.md)

* [Get](docs/sdks/collectors/README.md#get) - Get Collector by ID
* [ListCollectorObject](docs/sdks/collectors/README.md#listcollectorobject) - Get a list of Collector objects

### [Conditions](docs/sdks/conditions/README.md)

* [Get](docs/sdks/conditions/README.md#get) - Get Condition by ID
* [ListConditions](docs/sdks/conditions/README.md#listconditions) - Get a list of Condition objects

### [DatabaseConnections](docs/sdks/databaseconnections/README.md)

* [TestDatabaseConnection](docs/sdks/databaseconnections/README.md#testdatabaseconnection) - Test a database connection given a type and connectionString
* [Delete](docs/sdks/databaseconnections/README.md#delete) - Delete DatabaseConnectionConfig
* [Get](docs/sdks/databaseconnections/README.md#get) - Get DatabaseConnectionConfig by ID
* [ListDatabaseConnection](docs/sdks/databaseconnections/README.md#listdatabaseconnection) - Get a list of DatabaseConnection objects
* [Post](docs/sdks/databaseconnections/README.md#post) - Create DatabaseConnectionConfig
* [Update](docs/sdks/databaseconnections/README.md#update) - Update DatabaseConnectionConfig

### [Datasets](docs/sdks/datasets/README.md)

* [Create](docs/sdks/datasets/README.md#create) - Create Dataset
* [CreateProviderType](docs/sdks/datasets/README.md#createprovidertype) - Create DatasetProviderType
* [Delete](docs/sdks/datasets/README.md#delete) - Delete Dataset
* [DeleteProviderType](docs/sdks/datasets/README.md#deleteprovidertype) - Delete DatasetProviderType
* [Get](docs/sdks/datasets/README.md#get) - Get Dataset by ID
* [GetProviderType](docs/sdks/datasets/README.md#getprovidertype) - Get DatasetProviderType by ID
* [ListDatasetObjects](docs/sdks/datasets/README.md#listdatasetobjects) - Get a list of Dataset objects
* [ListProviderTypes](docs/sdks/datasets/README.md#listprovidertypes) - Get a list of DatasetProviderType objects
* [Update](docs/sdks/datasets/README.md#update) - Update Dataset
* [UpdateProviderType](docs/sdks/datasets/README.md#updateprovidertype) - Update DatasetProviderType

### [Diag](docs/sdks/diag/README.md)

* [Delete](docs/sdks/diag/README.md#delete) - Remove diag bundle
* [Get](docs/sdks/diag/README.md#get) - Returns a diag bundle as a tar.gz archive
* [GetSystemInfo](docs/sdks/diag/README.md#getsysteminfo) - Get basic system information
* [ListExistingDiagBundles](docs/sdks/diag/README.md#listexistingdiagbundles) - Get list of existing diag bundles
* [Send](docs/sdks/diag/README.md#send) - Send a diag bundle (tar.gz archive) to Cribl

### [Distributed](docs/sdks/distributed/README.md)

* [Get](docs/sdks/distributed/README.md#get) - Get summary of Distributed deployment
* [ListDiagBundles](docs/sdks/distributed/README.md#listdiagbundles) - Get list of existing diag bundles

### [Edge](docs/sdks/edge/README.md)

* [GetDirectoryListing](docs/sdks/edge/README.md#getdirectorylisting) - Get a directory listing of the given path
* [GetMetadata](docs/sdks/edge/README.md#getmetadata) - Get the host's metadata structure
* [GetRunDetails](docs/sdks/edge/README.md#getrundetails) - Get details of a process running on the edge host
* [ListBytes](docs/sdks/edge/README.md#listbytes) - Get some number of bytes from the file at the given path
* [ListConfiguredCollectors](docs/sdks/edge/README.md#listconfiguredcollectors) - Get list of configured collectors
* [ListEdgeHostFiles](docs/sdks/edge/README.md#listedgehostfiles) - Get details about a file on the edge host.
* [ListLogFiles](docs/sdks/edge/README.md#listlogfiles) - list log files
* [ListProcessRunningDetail](docs/sdks/edge/README.md#listprocessrunningdetail) - Get a detailed list of processes running on the edge host

### [EdgeContainers](docs/sdks/edgecontainers/README.md)

* [Get](docs/sdks/edgecontainers/README.md#get) - Get details for a single container on the edge host. Add stream=true to get a stream instead.
* [ListContainerDetail](docs/sdks/edgecontainers/README.md#listcontainerdetail) - Get a detailed list of containers running on the edge host.

### [EdgeEvents](docs/sdks/edgeevents/README.md)

* [ListConfiguredCollectors](docs/sdks/edgeevents/README.md#listconfiguredcollectors) - Get list of configured collectors

### [EdgeFiles](docs/sdks/edgefiles/README.md)

* [ListEdgeHostFiles](docs/sdks/edgefiles/README.md#listedgehostfiles) - Get details about a file on the edge host.

### [EdgeLs](docs/sdks/edgels/README.md)

* [GetDirectoryListing](docs/sdks/edgels/README.md#getdirectorylisting) - Get a directory listing of the given path

### [EdgeProcesses](docs/sdks/edgeprocesses/README.md)

* [GetRunDetails](docs/sdks/edgeprocesses/README.md#getrundetails) - Get details of a process running on the edge host
* [ListProcessRunningDetail](docs/sdks/edgeprocesses/README.md#listprocessrunningdetail) - Get a detailed list of processes running on the edge host

### [EventBreakerRules](docs/sdks/eventbreakerrules/README.md)

* [Delete](docs/sdks/eventbreakerrules/README.md#delete) - Delete Event Breaker Ruleset
* [Get](docs/sdks/eventbreakerrules/README.md#get) - Get Event Breaker Ruleset by ID
* [ListEventBreaker](docs/sdks/eventbreakerrules/README.md#listeventbreaker) - Get a list of Event Breaker Ruleset objects
* [Post](docs/sdks/eventbreakerrules/README.md#post) - Create Event Breaker Ruleset
* [Update](docs/sdks/eventbreakerrules/README.md#update) - Update Event Breaker Ruleset

### [Events](docs/sdks/events/README.md)

* [ListEvents](docs/sdks/events/README.md#listevents) - Get events generated by a specified source

### [Executors](docs/sdks/executors/README.md)

* [Get](docs/sdks/executors/README.md#get) - Get Executor by ID
* [ListExecutorObject](docs/sdks/executors/README.md#listexecutorobject) - Get a list of Executor objects

### [Expressions](docs/sdks/expressions/README.md)

* [Post](docs/sdks/expressions/README.md#post) - Evaluate JavaScript expression

### [Features](docs/sdks/features/README.md)

* [Get](docs/sdks/features/README.md#get) - Get feature by Id
* [ListFeatures](docs/sdks/features/README.md#listfeatures) - List all features

### [FileSampler](docs/sdks/filesampler/README.md)

* [ListBytes](docs/sdks/filesampler/README.md#listbytes) - Get some number of bytes from the file at the given path

### [FleetMappings](docs/sdks/fleetmappings/README.md)

* [Create](docs/sdks/fleetmappings/README.md#create) - Create MappingRuleset
* [Delete](docs/sdks/fleetmappings/README.md#delete) - Delete MappingRuleset
* [Get](docs/sdks/fleetmappings/README.md#get) - Get MappingRuleset by ID
* [ListFleetMappings](docs/sdks/fleetmappings/README.md#listfleetmappings) - Get a list of MappingRuleset objects
* [Update](docs/sdks/fleetmappings/README.md#update) - Update MappingRuleset

### [Functions](docs/sdks/functions/README.md)

* [Get](docs/sdks/functions/README.md#get) - Get Function by ID
* [ListObjectFunction](docs/sdks/functions/README.md#listobjectfunction) - Get a list of Function objects

### [Git](docs/sdks/git/README.md)

* [CountFiles](docs/sdks/git/README.md#countfiles) - get the count of files of changed
* [Create](docs/sdks/git/README.md#create) - create a new commit containing the current configs the given log message describing the changes.
* [Get](docs/sdks/git/README.md#get) - Get info about versioning availability
* [GetLogandTextualDiff](docs/sdks/git/README.md#getlogandtextualdiff) - get the log message and textual diff for given commit
* [GetTextualDiff](docs/sdks/git/README.md#gettextualdiff) - get the textual diff for given commit
* [GetWorkingTree](docs/sdks/git/README.md#getworkingtree) - get the the working tree status
* [ListBranches](docs/sdks/git/README.md#listbranches) - get the list of branches
* [ListChangedFiles](docs/sdks/git/README.md#listchangedfiles) - get the files changed
* [ListGitSettings](docs/sdks/git/README.md#listgitsettings) - Get git settings
* [Push](docs/sdks/git/README.md#push) - push the current configs to the remote repository.
* [Sync](docs/sdks/git/README.md#sync) - syncs with remote repo via POST requests
* [Update](docs/sdks/git/README.md#update) - Update git settings

### [GlobalVariables](docs/sdks/globalvariables/README.md)

* [Delete](docs/sdks/globalvariables/README.md#delete) - Delete Global Variable
* [Get](docs/sdks/globalvariables/README.md#get) - Get Global Variable by ID
* [ListGlobalVariable](docs/sdks/globalvariables/README.md#listglobalvariable) - Get a list of Global Variable objects
* [Post](docs/sdks/globalvariables/README.md#post) - Create Global Variable
* [Update](docs/sdks/globalvariables/README.md#update) - Update Global Variable

### [Grokfiles](docs/sdks/grokfiles/README.md)

* [Create](docs/sdks/grokfiles/README.md#create) - Create GrokFile
* [Delete](docs/sdks/grokfiles/README.md#delete) - Delete GrokFile
* [Get](docs/sdks/grokfiles/README.md#get) - Get GrokFile by ID
* [ListGrokFiles](docs/sdks/grokfiles/README.md#listgrokfiles) - Get a list of GrokFile objects
* [Update](docs/sdks/grokfiles/README.md#update) - Update GrokFile

### [Groups](docs/sdks/groups/README.md)

* [Create](docs/sdks/groups/README.md#create) - Create ConfigGroup
* [Delete](docs/sdks/groups/README.md#delete) - Delete ConfigGroup
* [Deploy](docs/sdks/groups/README.md#deploy) - Deploy commits for a Fleet or Worker Group
* [Get](docs/sdks/groups/README.md#get) - Get a specific ConfigGroup object
* [GetConfigVersion](docs/sdks/groups/README.md#getconfigversion) - Get effective bundle version for given Group
* [ListGroups](docs/sdks/groups/README.md#listgroups) - Get a list of ConfigGroup objects
* [Update](docs/sdks/groups/README.md#update) - Update ConfigGroup

### [Health](docs/sdks/health/README.md)

* [Get](docs/sdks/health/README.md#get) - Provides health info for REST server

### [Jobs](docs/sdks/jobs/README.md)

* [Cancel](docs/sdks/jobs/README.md#cancel) - Cancel a job by instance id
* [Delete](docs/sdks/jobs/README.md#delete) - Remove job from job inspector by instance id
* [Get](docs/sdks/jobs/README.md#get) - Get job info by instance id
* [GetError](docs/sdks/jobs/README.md#geterror) - Get Task errors for a job by id
* [GetResult](docs/sdks/jobs/README.md#getresult) - Get results for a discover job by instance id
* [ListJobInfos](docs/sdks/jobs/README.md#listjobinfos) - Get info on jobs
* [ListJobResults](docs/sdks/jobs/README.md#listjobresults) - Get results for a discover job by instance id
* [ListTaskErrors](docs/sdks/jobs/README.md#listtaskerrors) - Get Task errors for a job by id
* [PauseJob](docs/sdks/jobs/README.md#pausejob) - Pause a job by instance id
* [Prevent](docs/sdks/jobs/README.md#prevent) - prevent job from being deleted automatically
* [Resume](docs/sdks/jobs/README.md#resume) - Resume a job by instance id
* [RunJob](docs/sdks/jobs/README.md#runjob) - Run or schedule a job

### [Keys](docs/sdks/keys/README.md)

* [Create](docs/sdks/keys/README.md#create) - Create KeyMetadataEntity
* [Delete](docs/sdks/keys/README.md#delete) - Delete KeyMetadataEntity
* [Get](docs/sdks/keys/README.md#get) - Get KeyMetadataEntity by ID
* [ListKeyMetadataEntities](docs/sdks/keys/README.md#listkeymetadataentities) - Get a list of KeyMetadataEntity objects
* [Update](docs/sdks/keys/README.md#update) - Update KeyMetadataEntity

### [Licenses](docs/sdks/licenses/README.md)

* [Create](docs/sdks/licenses/README.md#create) - Create License
* [Delete](docs/sdks/licenses/README.md#delete) - Delete License
* [Get](docs/sdks/licenses/README.md#get) - Get License by ID
* [ListLicenseUsageMetrics](docs/sdks/licenses/README.md#listlicenseusagemetrics) - Get license usage metrics, aggregated by day, up to last 90 days
* [ListLicenses](docs/sdks/licenses/README.md#listlicenses) - Get a list of License objects

### [Logger](docs/sdks/logger/README.md)

* [Delete](docs/sdks/logger/README.md#delete) - Delete LoggerConfig
* [Get](docs/sdks/logger/README.md#get) - Get LoggerConfig by ID
* [ListLoggerConfigs](docs/sdks/logger/README.md#listloggerconfigs) - Get a list of LoggerConfig objects
* [Update](docs/sdks/logger/README.md#update) - Update LoggerConfig

### [Logging](docs/sdks/logging/README.md)

* [Get](docs/sdks/logging/README.md#get) - Get contents of the log file
* [ListLogFileContents](docs/sdks/logging/README.md#listlogfilecontents) - Get contents of the log file
* [ListLogFiles](docs/sdks/logging/README.md#listlogfiles) - Get a list of log files
* [ListLogFilesContents](docs/sdks/logging/README.md#listlogfilescontents) - Get contents of the log file

### [Lookups](docs/sdks/lookups/README.md)

* [Create](docs/sdks/lookups/README.md#create) - Create LookupFile
* [Delete](docs/sdks/lookups/README.md#delete) - Delete LookupFile
* [Get](docs/sdks/lookups/README.md#get) - Get LookupFile by ID
* [ListLookups](docs/sdks/lookups/README.md#listlookups) - Get a list of LookupFile objects
* [Update](docs/sdks/lookups/README.md#update) - Update LookupFile
* [Upload](docs/sdks/lookups/README.md#upload) - Upload LookupFile

### [Mappings](docs/sdks/mappings/README.md)

* [Create](docs/sdks/mappings/README.md#create) - Create MappingRuleset
* [Delete](docs/sdks/mappings/README.md#delete) - Delete MappingRuleset
* [Get](docs/sdks/mappings/README.md#get) - Get MappingRuleset by ID
* [ListMappingRulesets](docs/sdks/mappings/README.md#listmappingrulesets) - Get a list of MappingRuleset objects
* [Update](docs/sdks/mappings/README.md#update) - Update MappingRuleset

### [Messages](docs/sdks/messages/README.md)

* [Create](docs/sdks/messages/README.md#create) - Create BulletinMessage
* [Delete](docs/sdks/messages/README.md#delete) - Delete BulletinMessage
* [Get](docs/sdks/messages/README.md#get) - Get BulletinMessage by ID
* [ListBulletinMessages](docs/sdks/messages/README.md#listbulletinmessages) - Get a list of BulletinMessage objects

### [Metrics](docs/sdks/metrics/README.md)

* [Aggregate](docs/sdks/metrics/README.md#aggregate) - Aggregate raw internal system metrics
* [Post](docs/sdks/metrics/README.md#post) - Enumerate all internal system metrics
* [Query](docs/sdks/metrics/README.md#query) - Query raw internal system metrics

### [NotificationTargets](docs/sdks/notificationtargets/README.md)

* [Create](docs/sdks/notificationtargets/README.md#create) - Create NotificationTarget
* [Delete](docs/sdks/notificationtargets/README.md#delete) - Delete NotificationTarget
* [Get](docs/sdks/notificationtargets/README.md#get) - Get NotificationTarget by ID
* [ListNotificationTargets](docs/sdks/notificationtargets/README.md#listnotificationtargets) - Get a list of NotificationTarget objects
* [Update](docs/sdks/notificationtargets/README.md#update) - Update NotificationTarget

### [Outputs](docs/sdks/outputs/README.md)

* [Create](docs/sdks/outputs/README.md#create) - Create Output
* [Delete](docs/sdks/outputs/README.md#delete) - Delete Output
* [DeletePQ](docs/sdks/outputs/README.md#deletepq) - Delete destination persistent queue
* [Get](docs/sdks/outputs/README.md#get) - Get Output by ID
* [GetLatestPQ](docs/sdks/outputs/README.md#getlatestpq) - Get status of latest clear PQ job for an output
* [GetSamples](docs/sdks/outputs/README.md#getsamples) - Get samples data for the specified output. Used to get sample data for the test action.
* [GetStatus](docs/sdks/outputs/README.md#getstatus) - Get OutputStatus by ID
* [ListOutputObjects](docs/sdks/outputs/README.md#listoutputobjects) - Get a list of Output objects
* [ListOutputStatus](docs/sdks/outputs/README.md#listoutputstatus) - Get a list of OutputStatus objects
* [Post](docs/sdks/outputs/README.md#post) - Send sample data to an output to validate configuration or test connectivity
* [Update](docs/sdks/outputs/README.md#update) - Update Output

### [Packs](docs/sdks/packs/README.md)

* [Clone](docs/sdks/packs/README.md#clone) - Clone Pack
* [Export](docs/sdks/packs/README.md#export) - Export Pack
* [Install](docs/sdks/packs/README.md#install) - Install Pack
* [ListPacks](docs/sdks/packs/README.md#listpacks) - Get info on packs
* [Uninstall](docs/sdks/packs/README.md#uninstall) - Uninstall Pack from the system
* [Upgrade](docs/sdks/packs/README.md#upgrade) - Upgrade Pack
* [Upload](docs/sdks/packs/README.md#upload) - Upload Pack

### [Parquetschemas](docs/sdks/parquetschemas/README.md)

* [Delete](docs/sdks/parquetschemas/README.md#delete) - Delete Schema
* [Get](docs/sdks/parquetschemas/README.md#get) - Get Schema by ID
* [ListSchema](docs/sdks/parquetschemas/README.md#listschema) - Get a list of Schema objects
* [Post](docs/sdks/parquetschemas/README.md#post) - Create Schema
* [Update](docs/sdks/parquetschemas/README.md#update) - Update Schema

### [Parsers](docs/sdks/parsers/README.md)

* [Delete](docs/sdks/parsers/README.md#delete) - Delete Parser
* [Get](docs/sdks/parsers/README.md#get) - Get Parser by ID
* [ListParser](docs/sdks/parsers/README.md#listparser) - Get a list of Parser objects
* [Post](docs/sdks/parsers/README.md#post) - Create Parser
* [Update](docs/sdks/parsers/README.md#update) - Update Parser

### [Pipelines](docs/sdks/pipelines/README.md)

* [Delete](docs/sdks/pipelines/README.md#delete) - Delete Pipeline
* [Get](docs/sdks/pipelines/README.md#get) - Get Pipeline by ID
* [ListPipelineObject](docs/sdks/pipelines/README.md#listpipelineobject) - Get a list of Pipeline objects
* [Post](docs/sdks/pipelines/README.md#post) - Create Pipeline
* [Update](docs/sdks/pipelines/README.md#update) - Update Pipeline

### [Policies](docs/sdks/policies/README.md)

* [Create](docs/sdks/policies/README.md#create) - Create PolicyRule
* [Delete](docs/sdks/policies/README.md#delete) - Delete PolicyRule
* [Get](docs/sdks/policies/README.md#get) - Get PolicyRule by ID
* [ListPolicyRules](docs/sdks/policies/README.md#listpolicyrules) - Get a list of PolicyRule objects
* [Update](docs/sdks/policies/README.md#update) - Update PolicyRule

### [Preview](docs/sdks/preview/README.md)

* [CaptureLiveData](docs/sdks/preview/README.md#capturelivedata) - Capture live incoming data
* [SendEvents](docs/sdks/preview/README.md#sendevents) - Sends sample events through a pipeline and returns the results

### [Processes](docs/sdks/processes/README.md)

* [ListProcesses](docs/sdks/processes/README.md#listprocesses) - Get a list of processes under management

### [Profiler](docs/sdks/profiler/README.md)

* [Create](docs/sdks/profiler/README.md#create) - Create ProfilerItem
* [Delete](docs/sdks/profiler/README.md#delete) - Delete ProfilerItem
* [Get](docs/sdks/profiler/README.md#get) - Get ProfilerItem by ID
* [ListProfilers](docs/sdks/profiler/README.md#listprofilers) - Get a list of ProfilerItem objects
* [Update](docs/sdks/profiler/README.md#update) - Update ProfilerItem

### [Regexes](docs/sdks/regexes/README.md)

* [Delete](docs/sdks/regexes/README.md#delete) - Delete RegexLibEntry
* [Get](docs/sdks/regexes/README.md#get) - Get RegexLibEntry by ID
* [ListRegexLibEntryObject](docs/sdks/regexes/README.md#listregexlibentryobject) - Get a list of RegexLibEntry objects
* [Post](docs/sdks/regexes/README.md#post) - Create RegexLibEntry
* [Update](docs/sdks/regexes/README.md#update) - Update RegexLibEntry

### [Roles](docs/sdks/roles/README.md)

* [Create](docs/sdks/roles/README.md#create) - Create Role
* [Delete](docs/sdks/roles/README.md#delete) - Delete Role
* [Get](docs/sdks/roles/README.md#get) - Get Role by ID
* [ListRoles](docs/sdks/roles/README.md#listroles) - Get a list of Role objects
* [Update](docs/sdks/roles/README.md#update) - Update Role

### [Routes](docs/sdks/routes/README.md)

* [Get](docs/sdks/routes/README.md#get) - List all routes by id
* [ListRouteLists](docs/sdks/routes/README.md#listroutelists) - List all routes
* [Update](docs/sdks/routes/README.md#update) - Add, delete or update the routes with the required content.

### [Samples](docs/sdks/samples/README.md)

* [CaptureLiveData](docs/sdks/samples/README.md#capturelivedata) - Capture live incoming data
* [Delete](docs/sdks/samples/README.md#delete) - Delete DataSample
* [Get](docs/sdks/samples/README.md#get) - Get DataSample by ID
* [GetContent](docs/sdks/samples/README.md#getcontent) - Get sample content by ID
* [ListDataSample](docs/sdks/samples/README.md#listdatasample) - Get a list of DataSample objects
* [Post](docs/sdks/samples/README.md#post) - Create DataSample
* [SendEvents](docs/sdks/samples/README.md#sendevents) - Sends sample events through a pipeline and returns the results
* [Update](docs/sdks/samples/README.md#update) - Update DataSample

### [SavedJobs](docs/sdks/savedjobs/README.md)

* [Create](docs/sdks/savedjobs/README.md#create) - Create SavedJob
* [Delete](docs/sdks/savedjobs/README.md#delete) - Delete SavedJob
* [Get](docs/sdks/savedjobs/README.md#get) - Get SavedJob by ID
* [ListSavedJobs](docs/sdks/savedjobs/README.md#listsavedjobs) - Get a list of SavedJob objects
* [Update](docs/sdks/savedjobs/README.md#update) - Update SavedJob

### [SavedQueries](docs/sdks/savedqueries/README.md)

* [Create](docs/sdks/savedqueries/README.md#create) - Create SavedQuery
* [Delete](docs/sdks/savedqueries/README.md#delete) - Delete SavedQuery
* [Get](docs/sdks/savedqueries/README.md#get) - Get SavedQuery by ID
* [ListSavedQueries](docs/sdks/savedqueries/README.md#listsavedqueries) - Get a list of SavedQuery objects
* [Update](docs/sdks/savedqueries/README.md#update) - Update SavedQuery

### [Schemas](docs/sdks/schemas/README.md)

* [Create](docs/sdks/schemas/README.md#create) - Create Schema
* [Delete](docs/sdks/schemas/README.md#delete) - Delete Schema
* [Get](docs/sdks/schemas/README.md#get) - Get Schema by ID
* [ListSchemas](docs/sdks/schemas/README.md#listschemas) - Get a list of Schema objects
* [Update](docs/sdks/schemas/README.md#update) - Update Schema

### [Scripts](docs/sdks/scripts/README.md)

* [Create](docs/sdks/scripts/README.md#create) - Create Script
* [Delete](docs/sdks/scripts/README.md#delete) - Delete Script
* [Get](docs/sdks/scripts/README.md#get) - Get Script by ID
* [ListScripts](docs/sdks/scripts/README.md#listscripts) - Get a list of Script objects
* [Update](docs/sdks/scripts/README.md#update) - Update Script

### [Search](docs/sdks/search/README.md)

* [Apply](docs/sdks/search/README.md#apply) - Applies a query snippet on a set of input events for preview
* [Create](docs/sdks/search/README.md#create) - Create SearchJob
* [Delete](docs/sdks/search/README.md#delete) - Delete SearchJob
* [DispatchSearch](docs/sdks/search/README.md#dispatchsearch) - Dispatch search *id* to worker nodes filtered by worker node filter using RPC
* [Get](docs/sdks/search/README.md#get) - Get Search documentation
* [GetJob](docs/sdks/search/README.md#getjob) - Get SearchJob by ID
* [GetTimeline](docs/sdks/search/README.md#gettimeline) - Get search timeline
* [ListFieldSummaries](docs/sdks/search/README.md#listfieldsummaries) - List field summaries
* [ListJobStatus](docs/sdks/search/README.md#listjobstatus) - Get job status
* [ListSearchJobMetrics](docs/sdks/search/README.md#listsearchjobmetrics) - Get search job metrics
* [ListSearchJobs](docs/sdks/search/README.md#listsearchjobs) - Get a list of SearchJob objects
* [ListSearchLogs](docs/sdks/search/README.md#listsearchlogs) - Get search logs
* [Post](docs/sdks/search/README.md#post) - Runs an event breaker rule on the specified data
* [Update](docs/sdks/search/README.md#update) - Update SearchJob

### [Secrets](docs/sdks/secrets/README.md)

* [Create](docs/sdks/secrets/README.md#create) - Create RestSecret
* [Delete](docs/sdks/secrets/README.md#delete) - Delete RestSecret
* [Get](docs/sdks/secrets/README.md#get) - Get RestSecret by ID
* [ListRestSecrets](docs/sdks/secrets/README.md#listrestsecrets) - Get a list of RestSecret objects
* [Update](docs/sdks/secrets/README.md#update) - Update RestSecret

### [Security](docs/sdks/security/README.md)

* [GetKMSConfig](docs/sdks/security/README.md#getkmsconfig) - Get Cribl KMS config
* [GetKMSHealth](docs/sdks/security/README.md#getkmshealth) - Get Cribl KMS health
* [Update](docs/sdks/security/README.md#update) - Update Cribl KMS config

### [System](docs/sdks/system/README.md)

* [ReloadCriblSettings](docs/sdks/system/README.md#reloadcriblsettings) - Reload Cribl settings from the filesystem
* [CancelRunningGroup](docs/sdks/system/README.md#cancelrunninggroup) - Cancel a running group upgrade
* [ExecuteGroupUpgrade](docs/sdks/system/README.md#executegroupupgrade) - Execute distributed group upgrade
* [~~Get~~](docs/sdks/system/README.md#get) - Get Cribl system settings. Deprecated: use specific endpoints /system/limits, /system/job-limits, /system/redis-cache-limits, /system/services-limits, /system/settings/git-settings, and /system/settings/conf respectively :warning: **Deprecated**
* [GetPreviousPackage](docs/sdks/system/README.md#getpreviouspackage) - Get the previously downloaded Cribl package
* [ListAuthenticationSettings](docs/sdks/system/README.md#listauthenticationsettings) - Get authentication settings
* [ListCriblVersion](docs/sdks/system/README.md#listcriblversion) - Get a list of Cribl versions available for upgrade
* [ListSearchLimits](docs/sdks/system/README.md#listsearchlimits) - Get search limits
* [ListSettings](docs/sdks/system/README.md#listsettings) - Get Cribl system settings
* [Restart](docs/sdks/system/README.md#restart) - Restart Cribl server
* [StageGroupUpgrade](docs/sdks/system/README.md#stagegroupupgrade) - Stage distributed group upgrade
* [~~Update~~](docs/sdks/system/README.md#update) - Update Cribl system settings. Deprecated: use specific endpoints /system/limits, /system/job-limits, /system/settings/git-settings, /system/settings/auth and /system/settings/conf respectively :warning: **Deprecated**
* [UpdateAuthSettings](docs/sdks/system/README.md#updateauthsettings) - Update authentication settings
* [UpdateChangelogViewState](docs/sdks/system/README.md#updatechangelogviewstate) - Update changelog viewed state
* [UpdateCriblSettings](docs/sdks/system/README.md#updatecriblsettings) - Update Cribl system settings
* [UpgradeCribl](docs/sdks/system/README.md#upgradecribl) - Upgrade Cribl to a given version
* [UpgradeMaster](docs/sdks/system/README.md#upgrademaster) - Upgrade master node with the provided package

### [TrustPolicies](docs/sdks/trustpolicies/README.md)

* [ListTrustPolicies](docs/sdks/trustpolicies/README.md#listtrustpolicies) - Get a list of TrustPolicy objects

### [UIState](docs/sdks/uistate/README.md)

* [Get](docs/sdks/uistate/README.md#get) - Get UI state by key
* [Update](docs/sdks/uistate/README.md#update) - Update UI state by key

### [Users](docs/sdks/users/README.md)

* [CreateUser](docs/sdks/users/README.md#createuser) - Create User
* [Delete](docs/sdks/users/README.md#delete) - Delete User
* [Get](docs/sdks/users/README.md#get) - Get User by ID
* [ListUsers](docs/sdks/users/README.md#listusers) - Get a list of User objects
* [UpdateInfo](docs/sdks/users/README.md#updateinfo) - Update User except for their roles
* [UpdateProperties](docs/sdks/users/README.md#updateproperties) - Update User properties – admin use only

### [Versioning](docs/sdks/versioning/README.md)

* [CountFiles](docs/sdks/versioning/README.md#countfiles) - get the count of files of changed
* [Create](docs/sdks/versioning/README.md#create) - create a new commit containing the current configs the given log message describing the changes.
* [Get](docs/sdks/versioning/README.md#get) - Get info about versioning availability
* [GetLogandTextualDiff](docs/sdks/versioning/README.md#getlogandtextualdiff) - get the log message and textual diff for given commit
* [GetTextualDiff](docs/sdks/versioning/README.md#gettextualdiff) - get the textual diff for given commit
* [GetWorkingTree](docs/sdks/versioning/README.md#getworkingtree) - get the the working tree status
* [ListBranches](docs/sdks/versioning/README.md#listbranches) - get the list of branches
* [ListChangedFiles](docs/sdks/versioning/README.md#listchangedfiles) - get the files changed
* [Push](docs/sdks/versioning/README.md#push) - push the current configs to the remote repository.
* [Sync](docs/sdks/versioning/README.md#sync) - syncs with remote repo via POST requests

### [Workers](docs/sdks/workers/README.md)

* [Get](docs/sdks/workers/README.md#get) - get worker and edge nodes count
* [ListWorkerEdgeNodes](docs/sdks/workers/README.md#listworkeredgenodes) - get worker and edge nodes
* [Restart](docs/sdks/workers/README.md#restart) - restarts worker nodes
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
