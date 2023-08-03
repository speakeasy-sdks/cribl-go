# AppscopeLibEntry

### Available Operations

* [Create](#create) - Create AppscopeLibEntry
* [Delete](#delete) - Delete AppscopeLibEntry
* [Get](#get) - Get AppscopeLibEntry by ID
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
    res, err := s.AppscopeLibEntry.Create(ctx, shared.AppscopeLibEntry{
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
    id := "voluptatibus"

    ctx := context.Background()
    res, err := s.AppscopeLibEntry.Delete(ctx, id)
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
    id := "quisquam"

    ctx := context.Background()
    res, err := s.AppscopeLibEntry.Get(ctx, id)
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
    id := "vero"
    appscopeLibEntry := &shared.AppscopeLibEntry{
        Config: shared.AppscopeConfigWithCustom{
            Cribl: &shared.AppscopeConfigWithCustomCribl{
                Authtoken: cribl.String("omnis"),
                Enable: cribl.Bool(false),
                Transport: &shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                    Host: cribl.String("ipsum"),
                    Path: cribl.String("delectus"),
                    Port: cribl.Int64(455169),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("consectetur"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("vero"),
                },
                UseScopeSourceTransport: cribl.Bool(false),
            },
            Custom: []shared.AppscopeCustom{
                shared.AppscopeCustom{
                    Ancestor: cribl.String("dignissimos"),
                    Arg: cribl.String("hic"),
                    Config: shared.AppscopeConfig{
                        Cribl: &shared.AppscopeConfigCribl{
                            Authtoken: cribl.String("distinctio"),
                            Enable: cribl.Bool(false),
                            Transport: &shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("odio"),
                                Path: cribl.String("similique"),
                                Port: cribl.Int64(708548),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("vero"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("ducimus"),
                            },
                            UseScopeSourceTransport: cribl.Bool(false),
                        },
                        Event: &shared.AppscopeConfigEvent{
                            Enable: false,
                            Format: shared.AppscopeConfigEventFormat{
                                Enhancefs: false,
                                Maxeventpersec: 293020,
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("illum"),
                                Path: cribl.String("sequi"),
                                Port: cribl.Int64(617877),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("impedit"),
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
                                    Field: cribl.String("exercitationem"),
                                    Headers: cribl.String("nulla"),
                                    Name: cribl.String("Johnnie Wunsch"),
                                    Type: "eligendi",
                                    Value: cribl.String("ducimus"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("alias"),
                                    Headers: cribl.String("officia"),
                                    Name: cribl.String("Roberta Jenkins"),
                                    Type: "possimus",
                                    Value: cribl.String("magnam"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("ratione"),
                                    Headers: cribl.String("ex"),
                                    Name: cribl.String("Willie Fahey III"),
                                    Type: "nulla",
                                    Value: cribl.String("excepturi"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("voluptatibus"),
                                    Headers: cribl.String("nostrum"),
                                    Name: cribl.String("Devin Ullrich"),
                                    Type: "corporis",
                                    Value: cribl.String("veniam"),
                                },
                            },
                        },
                        Libscope: &shared.AppscopeConfigLibscope{
                            Commanddir: cribl.String("aliquid"),
                            Configevent: cribl.Bool(false),
                            Log: &shared.AppscopeConfigLibscopeLog{
                                Level: shared.AppscopeConfigLibscopeLogLevelDebug.ToPointer(),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("ea"),
                                    Path: cribl.String("quo"),
                                    Port: cribl.Int64(232234),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("recusandae"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("aspernatur"),
                                },
                            },
                            Summaryperiod: cribl.Int64(325310),
                        },
                        Metric: &shared.AppscopeConfigMetric{
                            Enable: false,
                            Format: shared.AppscopeConfigMetricFormat{
                                Statsdmaxlen: cribl.Int64(53427),
                                Statsdprefix: cribl.String("a"),
                                Type: cribl.String("libero"),
                                Verbosity: cribl.Int64(13948),
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("deleniti"),
                                Path: cribl.String("impedit"),
                                Port: cribl.Int64(304582),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("fugit"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("accusamus"),
                            },
                            Watch: []string{
                                "non",
                            },
                        },
                        Payload: &shared.AppscopeConfigPayload{
                            Dir: "et",
                            Enable: false,
                        },
                        Protocol: []shared.AppscopeConfigProtocol{
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 672048,
                                Name: "Lee Kemmer",
                                Payload: false,
                                Regex: "quas",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 829603,
                                Name: "Mrs. Shane Reinger",
                                Payload: false,
                                Regex: "explicabo",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 591935,
                                Name: "Minnie Gutkowski",
                                Payload: false,
                                Regex: "esse",
                            },
                        },
                        Tags: []shared.AppscopeConfigTags{
                            shared.AppscopeConfigTags{
                                Key: "rem",
                                Value: "fuga",
                            },
                            shared.AppscopeConfigTags{
                                Key: "reprehenderit",
                                Value: "quidem",
                            },
                        },
                    },
                    Env: cribl.String("fugiat"),
                    Hostname: cribl.String("firm-honoree.info"),
                    Procname: cribl.String("assumenda"),
                    Username: cribl.String("Chet.Lang5"),
                },
                shared.AppscopeCustom{
                    Ancestor: cribl.String("id"),
                    Arg: cribl.String("quidem"),
                    Config: shared.AppscopeConfig{
                        Cribl: &shared.AppscopeConfigCribl{
                            Authtoken: cribl.String("neque"),
                            Enable: cribl.Bool(false),
                            Transport: &shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("illum"),
                                Path: cribl.String("quo"),
                                Port: cribl.Int64(681359),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("eius"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("eos"),
                            },
                            UseScopeSourceTransport: cribl.Bool(false),
                        },
                        Event: &shared.AppscopeConfigEvent{
                            Enable: false,
                            Format: shared.AppscopeConfigEventFormat{
                                Enhancefs: false,
                                Maxeventpersec: 373813,
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("cupiditate"),
                                Path: cribl.String("consequatur"),
                                Port: cribl.Int64(272822),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("debitis"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("ipsam"),
                            },
                            Type: shared.AppscopeConfigEventTypeNdjson,
                            Watch: []shared.AppscopeConfigEventWatch{
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("sequi"),
                                    Headers: cribl.String("quo"),
                                    Name: cribl.String("Sophie Bayer"),
                                    Type: "dignissimos",
                                    Value: cribl.String("inventore"),
                                },
                            },
                        },
                        Libscope: &shared.AppscopeConfigLibscope{
                            Commanddir: cribl.String("nihil"),
                            Configevent: cribl.Bool(false),
                            Log: &shared.AppscopeConfigLibscopeLog{
                                Level: shared.AppscopeConfigLibscopeLogLevelWarning.ToPointer(),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                    Host: cribl.String("aliquam"),
                                    Path: cribl.String("odio"),
                                    Port: cribl.Int64(577543),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("commodi"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("sapiente"),
                                },
                            },
                            Summaryperiod: cribl.Int64(174112),
                        },
                        Metric: &shared.AppscopeConfigMetric{
                            Enable: false,
                            Format: shared.AppscopeConfigMetricFormat{
                                Statsdmaxlen: cribl.Int64(645570),
                                Statsdprefix: cribl.String("molestiae"),
                                Type: cribl.String("accusantium"),
                                Verbosity: cribl.Int64(783648),
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("quas"),
                                Path: cribl.String("praesentium"),
                                Port: cribl.Int64(159867),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("deleniti"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("fugit"),
                            },
                            Watch: []string{
                                "mollitia",
                                "incidunt",
                                "atque",
                            },
                        },
                        Payload: &shared.AppscopeConfigPayload{
                            Dir: "explicabo",
                            Enable: false,
                        },
                        Protocol: []shared.AppscopeConfigProtocol{
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 392676,
                                Name: "Jeannie Cronin",
                                Payload: false,
                                Regex: "saepe",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 578922,
                                Name: "Carl Koch",
                                Payload: false,
                                Regex: "veritatis",
                            },
                        },
                        Tags: []shared.AppscopeConfigTags{
                            shared.AppscopeConfigTags{
                                Key: "quod",
                                Value: "nam",
                            },
                            shared.AppscopeConfigTags{
                                Key: "vero",
                                Value: "aliquid",
                            },
                        },
                    },
                    Env: cribl.String("quasi"),
                    Hostname: cribl.String("untidy-heterosexual.net"),
                    Procname: cribl.String("molestiae"),
                    Username: cribl.String("Maximus71"),
                },
                shared.AppscopeCustom{
                    Ancestor: cribl.String("eligendi"),
                    Arg: cribl.String("sit"),
                    Config: shared.AppscopeConfig{
                        Cribl: &shared.AppscopeConfigCribl{
                            Authtoken: cribl.String("culpa"),
                            Enable: cribl.Bool(false),
                            Transport: &shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("adipisci"),
                                Path: cribl.String("cumque"),
                                Port: cribl.Int64(160538),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("consequatur"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("minus"),
                            },
                            UseScopeSourceTransport: cribl.Bool(false),
                        },
                        Event: &shared.AppscopeConfigEvent{
                            Enable: false,
                            Format: shared.AppscopeConfigEventFormat{
                                Enhancefs: false,
                                Maxeventpersec: 308286,
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("consectetur"),
                                Path: cribl.String("esse"),
                                Port: cribl.Int64(503427),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("provident"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("a"),
                            },
                            Type: shared.AppscopeConfigEventTypeNdjson,
                            Watch: []shared.AppscopeConfigEventWatch{
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("quas"),
                                    Headers: cribl.String("esse"),
                                    Name: cribl.String("Lorene Mueller"),
                                    Type: "possimus",
                                    Value: cribl.String("quia"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("eveniet"),
                                    Headers: cribl.String("asperiores"),
                                    Name: cribl.String("Miss Peter Cronin"),
                                    Type: "aliquid",
                                    Value: cribl.String("tenetur"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("quae"),
                                    Headers: cribl.String("earum"),
                                    Name: cribl.String("Pearl Gerlach"),
                                    Type: "soluta",
                                    Value: cribl.String("accusantium"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("aliquam"),
                                    Headers: cribl.String("sapiente"),
                                    Name: cribl.String("Marion Kihn"),
                                    Type: "aut",
                                    Value: cribl.String("voluptatum"),
                                },
                            },
                        },
                        Libscope: &shared.AppscopeConfigLibscope{
                            Commanddir: cribl.String("qui"),
                            Configevent: cribl.Bool(false),
                            Log: &shared.AppscopeConfigLibscopeLog{
                                Level: shared.AppscopeConfigLibscopeLogLevelNone.ToPointer(),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("deleniti"),
                                    Path: cribl.String("itaque"),
                                    Port: cribl.Int64(680270),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("architecto"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("omnis"),
                                },
                            },
                            Summaryperiod: cribl.Int64(945302),
                        },
                        Metric: &shared.AppscopeConfigMetric{
                            Enable: false,
                            Format: shared.AppscopeConfigMetricFormat{
                                Statsdmaxlen: cribl.Int64(98478),
                                Statsdprefix: cribl.String("at"),
                                Type: cribl.String("et"),
                                Verbosity: cribl.Int64(454162),
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("minima"),
                                Path: cribl.String("veritatis"),
                                Port: cribl.Int64(232744),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("adipisci"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("iste"),
                            },
                            Watch: []string{
                                "accusantium",
                                "rem",
                                "aut",
                                "laudantium",
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
                                Len: 68074,
                                Name: "Kyle Bartoletti",
                                Payload: false,
                                Regex: "numquam",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 771089,
                                Name: "Loretta Anderson DVM",
                                Payload: false,
                                Regex: "natus",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 244651,
                                Name: "Ms. Glen Zboncak",
                                Payload: false,
                                Regex: "consequuntur",
                            },
                        },
                        Tags: []shared.AppscopeConfigTags{
                            shared.AppscopeConfigTags{
                                Key: "officia",
                                Value: "maxime",
                            },
                            shared.AppscopeConfigTags{
                                Key: "dignissimos",
                                Value: "officia",
                            },
                            shared.AppscopeConfigTags{
                                Key: "asperiores",
                                Value: "nemo",
                            },
                            shared.AppscopeConfigTags{
                                Key: "quae",
                                Value: "quaerat",
                            },
                        },
                    },
                    Env: cribl.String("porro"),
                    Hostname: cribl.String("steep-dumbwaiter.com"),
                    Procname: cribl.String("adipisci"),
                    Username: cribl.String("Mark.Ondricka"),
                },
                shared.AppscopeCustom{
                    Ancestor: cribl.String("culpa"),
                    Arg: cribl.String("est"),
                    Config: shared.AppscopeConfig{
                        Cribl: &shared.AppscopeConfigCribl{
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
                        Event: &shared.AppscopeConfigEvent{
                            Enable: false,
                            Format: shared.AppscopeConfigEventFormat{
                                Enhancefs: false,
                                Maxeventpersec: 287051,
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("facilis"),
                                Path: cribl.String("cum"),
                                Port: cribl.Int64(414857),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("in"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("corporis"),
                            },
                            Type: shared.AppscopeConfigEventTypeNdjson,
                            Watch: []shared.AppscopeConfigEventWatch{
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("assumenda"),
                                    Headers: cribl.String("nemo"),
                                    Name: cribl.String("Gilbert Bayer"),
                                    Type: "in",
                                    Value: cribl.String("exercitationem"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("earum"),
                                    Headers: cribl.String("facere"),
                                    Name: cribl.String("Melba Homenick"),
                                    Type: "saepe",
                                    Value: cribl.String("necessitatibus"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("dolore"),
                                    Headers: cribl.String("sunt"),
                                    Name: cribl.String("Chad Franey IV"),
                                    Type: "a",
                                    Value: cribl.String("debitis"),
                                },
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("consectetur"),
                                    Headers: cribl.String("corporis"),
                                    Name: cribl.String("Rick Beer"),
                                    Type: "vitae",
                                    Value: cribl.String("accusamus"),
                                },
                            },
                        },
                        Libscope: &shared.AppscopeConfigLibscope{
                            Commanddir: cribl.String("similique"),
                            Configevent: cribl.Bool(false),
                            Log: &shared.AppscopeConfigLibscopeLog{
                                Level: shared.AppscopeConfigLibscopeLogLevelInfo.ToPointer(),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("voluptas"),
                                    Path: cribl.String("voluptas"),
                                    Port: cribl.Int64(374296),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("minima"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("nobis"),
                                },
                            },
                            Summaryperiod: cribl.Int64(680116),
                        },
                        Metric: &shared.AppscopeConfigMetric{
                            Enable: false,
                            Format: shared.AppscopeConfigMetricFormat{
                                Statsdmaxlen: cribl.Int64(237807),
                                Statsdprefix: cribl.String("minus"),
                                Type: cribl.String("dolores"),
                                Verbosity: cribl.Int64(503934),
                            },
                            Transport: shared.AppscopeTransport{
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
                            Watch: []string{
                                "cum",
                            },
                        },
                        Payload: &shared.AppscopeConfigPayload{
                            Dir: "blanditiis",
                            Enable: false,
                        },
                        Protocol: []shared.AppscopeConfigProtocol{
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 942584,
                                Name: "Molly Lowe",
                                Payload: false,
                                Regex: "hic",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 348783,
                                Name: "Paul Price",
                                Payload: false,
                                Regex: "explicabo",
                            },
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 994401,
                                Name: "Miss Jared Quitzon",
                                Payload: false,
                                Regex: "laborum",
                            },
                        },
                        Tags: []shared.AppscopeConfigTags{
                            shared.AppscopeConfigTags{
                                Key: "in",
                                Value: "commodi",
                            },
                        },
                    },
                    Env: cribl.String("quidem"),
                    Hostname: cribl.String("common-gather.name"),
                    Procname: cribl.String("architecto"),
                    Username: cribl.String("Geovany.Willms96"),
                },
            },
            Event: &shared.AppscopeConfigWithCustomEvent{
                Enable: false,
                Format: shared.AppscopeConfigWithCustomEventFormat{
                    Enhancefs: false,
                    Maxeventpersec: 19300,
                },
                Transport: shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                    Host: cribl.String("maiores"),
                    Path: cribl.String("incidunt"),
                    Port: cribl.Int64(148478),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("provident"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("eius"),
                },
                Type: shared.AppscopeConfigWithCustomEventTypeNdjson,
                Watch: []shared.AppscopeConfigWithCustomEventWatch{
                    shared.AppscopeConfigWithCustomEventWatch{
                        Allowbinary: cribl.Bool(false),
                        Enabled: cribl.Bool(false),
                        Field: cribl.String("ipsum"),
                        Headers: cribl.String("ea"),
                        Name: cribl.String("Isaac Wolf"),
                        Type: "voluptate",
                        Value: cribl.String("reiciendis"),
                    },
                    shared.AppscopeConfigWithCustomEventWatch{
                        Allowbinary: cribl.Bool(false),
                        Enabled: cribl.Bool(false),
                        Field: cribl.String("ex"),
                        Headers: cribl.String("sit"),
                        Name: cribl.String("Cecelia Lakin"),
                        Type: "incidunt",
                        Value: cribl.String("ipsam"),
                    },
                    shared.AppscopeConfigWithCustomEventWatch{
                        Allowbinary: cribl.Bool(false),
                        Enabled: cribl.Bool(false),
                        Field: cribl.String("debitis"),
                        Headers: cribl.String("rem"),
                        Name: cribl.String("Della Muller"),
                        Type: "recusandae",
                        Value: cribl.String("reiciendis"),
                    },
                    shared.AppscopeConfigWithCustomEventWatch{
                        Allowbinary: cribl.Bool(false),
                        Enabled: cribl.Bool(false),
                        Field: cribl.String("nulla"),
                        Headers: cribl.String("magni"),
                        Name: cribl.String("Gwen Fritsch"),
                        Type: "officiis",
                        Value: cribl.String("beatae"),
                    },
                },
            },
            Libscope: &shared.AppscopeConfigWithCustomLibscope{
                Commanddir: cribl.String("laudantium"),
                Configevent: cribl.Bool(false),
                Log: &shared.AppscopeConfigWithCustomLibscopeLog{
                    Level: shared.AppscopeConfigWithCustomLibscopeLogLevelInfo.ToPointer(),
                    Transport: &shared.AppscopeTransport{
                        Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                        Host: cribl.String("cum"),
                        Path: cribl.String("laboriosam"),
                        Port: cribl.Int64(680515),
                        TLS: &shared.AppscopeTransportTLS{
                            Cacertpath: cribl.String("voluptatum"),
                            Enable: cribl.Bool(false),
                            Validateserver: cribl.Bool(false),
                        },
                        Type: cribl.String("error"),
                    },
                },
                Summaryperiod: cribl.Int64(944708),
            },
            Metric: &shared.AppscopeConfigWithCustomMetric{
                Enable: false,
                Format: shared.AppscopeConfigWithCustomMetricFormat{
                    Statsdmaxlen: cribl.Int64(710529),
                    Statsdprefix: cribl.String("debitis"),
                    Type: cribl.String("neque"),
                    Verbosity: cribl.Int64(677115),
                },
                Transport: shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                    Host: cribl.String("officia"),
                    Path: cribl.String("dolorum"),
                    Port: cribl.Int64(548361),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("accusamus"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("tempora"),
                },
                Watch: []string{
                    "fugit",
                    "ut",
                    "fugiat",
                },
            },
            Payload: &shared.AppscopeConfigWithCustomPayload{
                Dir: "voluptatem",
                Enable: false,
            },
            Protocol: []shared.AppscopeConfigWithCustomProtocol{
                shared.AppscopeConfigWithCustomProtocol{
                    Binary: false,
                    Detect: false,
                    Len: 710337,
                    Name: "Barbara Koelpin IV",
                    Payload: false,
                    Regex: "quas",
                },
                shared.AppscopeConfigWithCustomProtocol{
                    Binary: false,
                    Detect: false,
                    Len: 922112,
                    Name: "Janet Kuvalis",
                    Payload: false,
                    Regex: "sit",
                },
                shared.AppscopeConfigWithCustomProtocol{
                    Binary: false,
                    Detect: false,
                    Len: 425508,
                    Name: "Mrs. Tricia Mueller",
                    Payload: false,
                    Regex: "dolorem",
                },
            },
            Tags: []shared.AppscopeConfigWithCustomTags{
                shared.AppscopeConfigWithCustomTags{
                    Key: "dicta",
                    Value: "architecto",
                },
                shared.AppscopeConfigWithCustomTags{
                    Key: "occaecati",
                    Value: "labore",
                },
                shared.AppscopeConfigWithCustomTags{
                    Key: "quidem",
                    Value: "atque",
                },
            },
        },
        Description: "laborum",
        ID: "bf603a79-f9df-4e0a-b7da-8a50ce187f86",
        Lib: shared.CriblLibCustom,
        Tags: cribl.String("maxime"),
    }

    ctx := context.Background()
    res, err := s.AppscopeLibEntry.Update(ctx, id, appscopeLibEntry)
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

