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
    res, err := s.AppscopeLibEntry.Delete(ctx, operations.DeleteAppscopeLibEntryRequest{
        ID: "fce953f7-3ef7-4fbc-babd-74dd39c0f5d2",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeleteAppscopeLibEntryRequest](../../models/operations/deleteappscopelibentryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


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
    res, err := s.AppscopeLibEntry.Get(ctx, operations.GetAppscopeLibEntryRequest{
        ID: "cff7c70a-4562-46d4-b681-3f16d9f5fce6",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetAppscopeLibEntryRequest](../../models/operations/getappscopelibentryrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


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
    res, err := s.AppscopeLibEntry.Update(ctx, operations.UpdateAppscopeLibEntryRequest{
        AppscopeLibEntry: &shared.AppscopeLibEntry{
            Config: shared.AppscopeConfigWithCustom{
                Cribl: &shared.AppscopeConfigWithCustomCribl{
                    Authtoken: cribl.String("impedit"),
                    Enable: cribl.Bool(false),
                    Transport: &shared.AppscopeTransport{
                        Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                        Host: cribl.String("veniam"),
                        Path: cribl.String("aliquid"),
                        Port: cribl.Int64(81101),
                        TLS: &shared.AppscopeTransportTLS{
                            Cacertpath: cribl.String("magnam"),
                            Enable: cribl.Bool(false),
                            Validateserver: cribl.Bool(false),
                        },
                        Type: cribl.String("ea"),
                    },
                    UseScopeSourceTransport: cribl.Bool(false),
                },
                Custom: []shared.AppscopeCustom{
                    shared.AppscopeCustom{
                        Ancestor: cribl.String("consectetur"),
                        Arg: cribl.String("recusandae"),
                        Config: shared.AppscopeConfig{
                            Cribl: &shared.AppscopeConfigCribl{
                                Authtoken: cribl.String("aspernatur"),
                                Enable: cribl.Bool(false),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("eaque"),
                                    Path: cribl.String("a"),
                                    Port: cribl.Int64(725595),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("aut"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("aut"),
                                },
                                UseScopeSourceTransport: cribl.Bool(false),
                            },
                            Event: &shared.AppscopeConfigEvent{
                                Enable: false,
                                Format: shared.AppscopeConfigEventFormat{
                                    Enhancefs: false,
                                    Maxeventpersec: 533466,
                                },
                                Transport: shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                    Host: cribl.String("aliquam"),
                                    Path: cribl.String("fugit"),
                                    Port: cribl.Int64(882860),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("inventore"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("non"),
                                },
                                Type: shared.AppscopeConfigEventTypeNdjson,
                                Watch: []shared.AppscopeConfigEventWatch{
                                    shared.AppscopeConfigEventWatch{
                                        Allowbinary: cribl.Bool(false),
                                        Enabled: cribl.Bool(false),
                                        Field: cribl.String("dolorum"),
                                        Headers: cribl.String("laborum"),
                                        Name: cribl.String("Lee Kemmer"),
                                        Type: "quas",
                                        Value: cribl.String("assumenda"),
                                    },
                                },
                            },
                            Libscope: &shared.AppscopeConfigLibscope{
                                Commanddir: cribl.String("nulla"),
                                Configevent: cribl.Bool(false),
                                Log: &shared.AppscopeConfigLibscopeLog{
                                    Level: shared.AppscopeConfigLibscopeLogLevelInfo.ToPointer(),
                                    Transport: &shared.AppscopeTransport{
                                        Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                        Host: cribl.String("quasi"),
                                        Path: cribl.String("tempora"),
                                        Port: cribl.Int64(256139),
                                        TLS: &shared.AppscopeTransportTLS{
                                            Cacertpath: cribl.String("explicabo"),
                                            Enable: cribl.Bool(false),
                                            Validateserver: cribl.Bool(false),
                                        },
                                        Type: cribl.String("provident"),
                                    },
                                },
                                Summaryperiod: cribl.Int64(55374),
                            },
                            Metric: &shared.AppscopeConfigMetric{
                                Enable: false,
                                Format: shared.AppscopeConfigMetricFormat{
                                    Statsdmaxlen: cribl.Int64(476477),
                                    Statsdprefix: cribl.String("magnam"),
                                    Type: cribl.String("odio"),
                                    Verbosity: cribl.Int64(262118),
                                },
                                Transport: shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("esse"),
                                    Path: cribl.String("rem"),
                                    Port: cribl.Int64(683282),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("reprehenderit"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("quidem"),
                                },
                                Watch: []string{
                                    "ut",
                                    "eum",
                                    "suscipit",
                                    "assumenda",
                                },
                            },
                            Payload: &shared.AppscopeConfigPayload{
                                Dir: "eos",
                                Enable: false,
                            },
                            Protocol: []shared.AppscopeConfigProtocol{
                                shared.AppscopeConfigProtocol{
                                    Binary: false,
                                    Detect: false,
                                    Len: 788546,
                                    Name: "Angela Olson",
                                    Payload: false,
                                    Regex: "quo",
                                },
                                shared.AppscopeConfigProtocol{
                                    Binary: false,
                                    Detect: false,
                                    Len: 847276,
                                    Name: "Wilbur Gerlach",
                                    Payload: false,
                                    Regex: "ab",
                                },
                                shared.AppscopeConfigProtocol{
                                    Binary: false,
                                    Detect: false,
                                    Len: 587600,
                                    Name: "Rhonda Toy",
                                    Payload: false,
                                    Regex: "sequi",
                                },
                            },
                            Tags: []shared.AppscopeConfigTags{
                                shared.AppscopeConfigTags{
                                    Key: "esse",
                                    Value: "recusandae",
                                },
                                shared.AppscopeConfigTags{
                                    Key: "aperiam",
                                    Value: "distinctio",
                                },
                                shared.AppscopeConfigTags{
                                    Key: "quod",
                                    Value: "dignissimos",
                                },
                                shared.AppscopeConfigTags{
                                    Key: "inventore",
                                    Value: "nihil",
                                },
                            },
                        },
                        Env: cribl.String("totam"),
                        Hostname: cribl.String("uncommon-encounter.info"),
                        Procname: cribl.String("occaecati"),
                        Username: cribl.String("Harvey64"),
                    },
                    shared.AppscopeCustom{
                        Ancestor: cribl.String("molestiae"),
                        Arg: cribl.String("accusantium"),
                        Config: shared.AppscopeConfig{
                            Cribl: &shared.AppscopeConfigCribl{
                                Authtoken: cribl.String("porro"),
                                Enable: cribl.Bool(false),
                                Transport: &shared.AppscopeTransport{
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
                                UseScopeSourceTransport: cribl.Bool(false),
                            },
                            Event: &shared.AppscopeConfigEvent{
                                Enable: false,
                                Format: shared.AppscopeConfigEventFormat{
                                    Enhancefs: false,
                                    Maxeventpersec: 681393,
                                },
                                Transport: shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                    Host: cribl.String("incidunt"),
                                    Path: cribl.String("atque"),
                                    Port: cribl.Int64(128860),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("minima"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("nisi"),
                                },
                                Type: shared.AppscopeConfigEventTypeNdjson,
                                Watch: []shared.AppscopeConfigEventWatch{
                                    shared.AppscopeConfigEventWatch{
                                        Allowbinary: cribl.Bool(false),
                                        Enabled: cribl.Bool(false),
                                        Field: cribl.String("sapiente"),
                                        Headers: cribl.String("consequuntur"),
                                        Name: cribl.String("Rose Turner"),
                                        Type: "et",
                                        Value: cribl.String("esse"),
                                    },
                                },
                            },
                            Libscope: &shared.AppscopeConfigLibscope{
                                Commanddir: cribl.String("eveniet"),
                                Configevent: cribl.Bool(false),
                                Log: &shared.AppscopeConfigLibscopeLog{
                                    Level: shared.AppscopeConfigLibscopeLogLevelNone.ToPointer(),
                                    Transport: &shared.AppscopeTransport{
                                        Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                        Host: cribl.String("esse"),
                                        Path: cribl.String("quod"),
                                        Port: cribl.Int64(724168),
                                        TLS: &shared.AppscopeTransportTLS{
                                            Cacertpath: cribl.String("vero"),
                                            Enable: cribl.Bool(false),
                                            Validateserver: cribl.Bool(false),
                                        },
                                        Type: cribl.String("aliquid"),
                                    },
                                },
                                Summaryperiod: cribl.Int64(93459),
                            },
                            Metric: &shared.AppscopeConfigMetric{
                                Enable: false,
                                Format: shared.AppscopeConfigMetricFormat{
                                    Statsdmaxlen: cribl.Int64(904045),
                                    Statsdprefix: cribl.String("vel"),
                                    Type: cribl.String("harum"),
                                    Verbosity: cribl.Int64(473221),
                                },
                                Transport: shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                    Host: cribl.String("occaecati"),
                                    Path: cribl.String("minima"),
                                    Port: cribl.Int64(716244),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("eligendi"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("sit"),
                                },
                                Watch: []string{
                                    "tempore",
                                    "adipisci",
                                    "cumque",
                                },
                            },
                            Payload: &shared.AppscopeConfigPayload{
                                Dir: "consequuntur",
                                Enable: false,
                            },
                            Protocol: []shared.AppscopeConfigProtocol{
                                shared.AppscopeConfigProtocol{
                                    Binary: false,
                                    Detect: false,
                                    Len: 796392,
                                    Name: "Miranda Feest",
                                    Payload: false,
                                    Regex: "provident",
                                },
                            },
                            Tags: []shared.AppscopeConfigTags{
                                shared.AppscopeConfigTags{
                                    Key: "nulla",
                                    Value: "quas",
                                },
                                shared.AppscopeConfigTags{
                                    Key: "esse",
                                    Value: "quasi",
                                },
                                shared.AppscopeConfigTags{
                                    Key: "a",
                                    Value: "error",
                                },
                                shared.AppscopeConfigTags{
                                    Key: "sint",
                                    Value: "pariatur",
                                },
                            },
                        },
                        Env: cribl.String("possimus"),
                        Hostname: cribl.String("cuddly-timpani.org"),
                        Procname: cribl.String("facere"),
                        Username: cribl.String("Arvel62"),
                    },
                    shared.AppscopeCustom{
                        Ancestor: cribl.String("culpa"),
                        Arg: cribl.String("aliquid"),
                        Config: shared.AppscopeConfig{
                            Cribl: &shared.AppscopeConfigCribl{
                                Authtoken: cribl.String("tenetur"),
                                Enable: cribl.Bool(false),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("earum"),
                                    Path: cribl.String("vel"),
                                    Port: cribl.Int64(447378),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("eius"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("libero"),
                                },
                                UseScopeSourceTransport: cribl.Bool(false),
                            },
                            Event: &shared.AppscopeConfigEvent{
                                Enable: false,
                                Format: shared.AppscopeConfigEventFormat{
                                    Enhancefs: false,
                                    Maxeventpersec: 849039,
                                },
                                Transport: shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                    Host: cribl.String("accusantium"),
                                    Path: cribl.String("aliquam"),
                                    Port: cribl.Int64(958983),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("dicta"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("ullam"),
                                },
                                Type: shared.AppscopeConfigEventTypeNdjson,
                                Watch: []shared.AppscopeConfigEventWatch{
                                    shared.AppscopeConfigEventWatch{
                                        Allowbinary: cribl.Bool(false),
                                        Enabled: cribl.Bool(false),
                                        Field: cribl.String("ullam"),
                                        Headers: cribl.String("nisi"),
                                        Name: cribl.String("Nora Denesik"),
                                        Type: "deleniti",
                                        Value: cribl.String("itaque"),
                                    },
                                    shared.AppscopeConfigEventWatch{
                                        Allowbinary: cribl.Bool(false),
                                        Enabled: cribl.Bool(false),
                                        Field: cribl.String("dolorum"),
                                        Headers: cribl.String("architecto"),
                                        Name: cribl.String("Irvin Boyle III"),
                                        Type: "ipsa",
                                        Value: cribl.String("minima"),
                                    },
                                },
                            },
                            Libscope: &shared.AppscopeConfigLibscope{
                                Commanddir: cribl.String("veritatis"),
                                Configevent: cribl.Bool(false),
                                Log: &shared.AppscopeConfigLibscopeLog{
                                    Level: shared.AppscopeConfigLibscopeLogLevelInfo.ToPointer(),
                                    Transport: &shared.AppscopeTransport{
                                        Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                        Host: cribl.String("iste"),
                                        Path: cribl.String("temporibus"),
                                        Port: cribl.Int64(33074),
                                        TLS: &shared.AppscopeTransportTLS{
                                            Cacertpath: cribl.String("rem"),
                                            Enable: cribl.Bool(false),
                                            Validateserver: cribl.Bool(false),
                                        },
                                        Type: cribl.String("aut"),
                                    },
                                },
                                Summaryperiod: cribl.Int64(513075),
                            },
                            Metric: &shared.AppscopeConfigMetric{
                                Enable: false,
                                Format: shared.AppscopeConfigMetricFormat{
                                    Statsdmaxlen: cribl.Int64(428796),
                                    Statsdprefix: cribl.String("mollitia"),
                                    Type: cribl.String("ab"),
                                    Verbosity: cribl.Int64(544591),
                                },
                                Transport: shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("voluptatem"),
                                    Path: cribl.String("dolor"),
                                    Port: cribl.Int64(580152),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("numquam"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("impedit"),
                                },
                                Watch: []string{
                                    "voluptas",
                                },
                            },
                            Payload: &shared.AppscopeConfigPayload{
                                Dir: "aut",
                                Enable: false,
                            },
                            Protocol: []shared.AppscopeConfigProtocol{
                                shared.AppscopeConfigProtocol{
                                    Binary: false,
                                    Detect: false,
                                    Len: 115484,
                                    Name: "Wendell Frami",
                                    Payload: false,
                                    Regex: "asperiores",
                                },
                                shared.AppscopeConfigProtocol{
                                    Binary: false,
                                    Detect: false,
                                    Len: 45659,
                                    Name: "Bertha Cruickshank",
                                    Payload: false,
                                    Regex: "maxime",
                                },
                            },
                            Tags: []shared.AppscopeConfigTags{
                                shared.AppscopeConfigTags{
                                    Key: "officia",
                                    Value: "asperiores",
                                },
                                shared.AppscopeConfigTags{
                                    Key: "nemo",
                                    Value: "quae",
                                },
                            },
                        },
                        Env: cribl.String("quaerat"),
                        Hostname: cribl.String("spanish-show-stopper.biz"),
                        Procname: cribl.String("ab"),
                        Username: cribl.String("Dayton.Paucek"),
                    },
                    shared.AppscopeCustom{
                        Ancestor: cribl.String("velit"),
                        Arg: cribl.String("culpa"),
                        Config: shared.AppscopeConfig{
                            Cribl: &shared.AppscopeConfigCribl{
                                Authtoken: cribl.String("est"),
                                Enable: cribl.Bool(false),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                    Host: cribl.String("totam"),
                                    Path: cribl.String("fugiat"),
                                    Port: cribl.Int64(424089),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("ducimus"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("quos"),
                                },
                                UseScopeSourceTransport: cribl.Bool(false),
                            },
                            Event: &shared.AppscopeConfigEvent{
                                Enable: false,
                                Format: shared.AppscopeConfigEventFormat{
                                    Enhancefs: false,
                                    Maxeventpersec: 427834,
                                },
                                Transport: shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("possimus"),
                                    Path: cribl.String("facilis"),
                                    Port: cribl.Int64(738227),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("commodi"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("in"),
                                },
                                Type: shared.AppscopeConfigEventTypeNdjson,
                                Watch: []shared.AppscopeConfigEventWatch{
                                    shared.AppscopeConfigEventWatch{
                                        Allowbinary: cribl.Bool(false),
                                        Enabled: cribl.Bool(false),
                                        Field: cribl.String("reiciendis"),
                                        Headers: cribl.String("assumenda"),
                                        Name: cribl.String("Miss Sophie Jacobi"),
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
                                },
                            },
                            Libscope: &shared.AppscopeConfigLibscope{
                                Commanddir: cribl.String("dolore"),
                                Configevent: cribl.Bool(false),
                                Log: &shared.AppscopeConfigLibscopeLog{
                                    Level: shared.AppscopeConfigLibscopeLogLevelDebug.ToPointer(),
                                    Transport: &shared.AppscopeTransport{
                                        Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                        Host: cribl.String("adipisci"),
                                        Path: cribl.String("non"),
                                        Port: cribl.Int64(228263),
                                        TLS: &shared.AppscopeTransportTLS{
                                            Cacertpath: cribl.String("beatae"),
                                            Enable: cribl.Bool(false),
                                            Validateserver: cribl.Bool(false),
                                        },
                                        Type: cribl.String("dignissimos"),
                                    },
                                },
                                Summaryperiod: cribl.Int64(950953),
                            },
                            Metric: &shared.AppscopeConfigMetric{
                                Enable: false,
                                Format: shared.AppscopeConfigMetricFormat{
                                    Statsdmaxlen: cribl.Int64(891523),
                                    Statsdprefix: cribl.String("consectetur"),
                                    Type: cribl.String("corporis"),
                                    Verbosity: cribl.Int64(689768),
                                },
                                Transport: shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                    Host: cribl.String("ipsa"),
                                    Path: cribl.String("voluptates"),
                                    Port: cribl.Int64(730709),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("vitae"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("accusamus"),
                                },
                                Watch: []string{
                                    "tempora",
                                    "aspernatur",
                                    "voluptas",
                                },
                            },
                            Payload: &shared.AppscopeConfigPayload{
                                Dir: "voluptas",
                                Enable: false,
                            },
                            Protocol: []shared.AppscopeConfigProtocol{
                                shared.AppscopeConfigProtocol{
                                    Binary: false,
                                    Detect: false,
                                    Len: 324405,
                                    Name: "Wilbur Ferry",
                                    Payload: false,
                                    Regex: "blanditiis",
                                },
                                shared.AppscopeConfigProtocol{
                                    Binary: false,
                                    Detect: false,
                                    Len: 449292,
                                    Name: "Suzanne Torphy",
                                    Payload: false,
                                    Regex: "adipisci",
                                },
                            },
                            Tags: []shared.AppscopeConfigTags{
                                shared.AppscopeConfigTags{
                                    Key: "blanditiis",
                                    Value: "quas",
                                },
                                shared.AppscopeConfigTags{
                                    Key: "hic",
                                    Value: "nesciunt",
                                },
                                shared.AppscopeConfigTags{
                                    Key: "culpa",
                                    Value: "corrupti",
                                },
                            },
                        },
                        Env: cribl.String("pariatur"),
                        Hostname: cribl.String("mature-underneath.info"),
                        Procname: cribl.String("nobis"),
                        Username: cribl.String("Alberto96"),
                    },
                },
                Event: &shared.AppscopeConfigWithCustomEvent{
                    Enable: false,
                    Format: shared.AppscopeConfigWithCustomEventFormat{
                        Enhancefs: false,
                        Maxeventpersec: 131852,
                    },
                    Transport: shared.AppscopeTransport{
                        Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                        Host: cribl.String("facilis"),
                        Path: cribl.String("voluptate"),
                        Port: cribl.Int64(709072),
                        TLS: &shared.AppscopeTransportTLS{
                            Cacertpath: cribl.String("ab"),
                            Enable: cribl.Bool(false),
                            Validateserver: cribl.Bool(false),
                        },
                        Type: cribl.String("iste"),
                    },
                    Type: shared.AppscopeConfigWithCustomEventTypeNdjson,
                    Watch: []shared.AppscopeConfigWithCustomEventWatch{
                        shared.AppscopeConfigWithCustomEventWatch{
                            Allowbinary: cribl.Bool(false),
                            Enabled: cribl.Bool(false),
                            Field: cribl.String("laborum"),
                            Headers: cribl.String("sed"),
                            Name: cribl.String("Tonya Predovic"),
                            Type: "unde",
                            Value: cribl.String("architecto"),
                        },
                        shared.AppscopeConfigWithCustomEventWatch{
                            Allowbinary: cribl.Bool(false),
                            Enabled: cribl.Bool(false),
                            Field: cribl.String("suscipit"),
                            Headers: cribl.String("sapiente"),
                            Name: cribl.String("Ms. Gregory Wisoky"),
                            Type: "incidunt",
                            Value: cribl.String("sed"),
                        },
                    },
                },
                Libscope: &shared.AppscopeConfigWithCustomLibscope{
                    Commanddir: cribl.String("provident"),
                    Configevent: cribl.Bool(false),
                    Log: &shared.AppscopeConfigWithCustomLibscopeLog{
                        Level: shared.AppscopeConfigWithCustomLibscopeLogLevelInfo.ToPointer(),
                        Transport: &shared.AppscopeTransport{
                            Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                            Host: cribl.String("ipsum"),
                            Path: cribl.String("ea"),
                            Port: cribl.Int64(579912),
                            TLS: &shared.AppscopeTransportTLS{
                                Cacertpath: cribl.String("quos"),
                                Enable: cribl.Bool(false),
                                Validateserver: cribl.Bool(false),
                            },
                            Type: cribl.String("voluptatibus"),
                        },
                    },
                    Summaryperiod: cribl.Int64(271653),
                },
                Metric: &shared.AppscopeConfigWithCustomMetric{
                    Enable: false,
                    Format: shared.AppscopeConfigWithCustomMetricFormat{
                        Statsdmaxlen: cribl.Int64(273009),
                        Statsdprefix: cribl.String("voluptate"),
                        Type: cribl.String("reiciendis"),
                        Verbosity: cribl.Int64(401713),
                    },
                    Transport: shared.AppscopeTransport{
                        Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                        Host: cribl.String("non"),
                        Path: cribl.String("officiis"),
                        Port: cribl.Int64(505866),
                        TLS: &shared.AppscopeTransportTLS{
                            Cacertpath: cribl.String("facilis"),
                            Enable: cribl.Bool(false),
                            Validateserver: cribl.Bool(false),
                        },
                        Type: cribl.String("quaerat"),
                    },
                    Watch: []string{
                        "ipsam",
                        "debitis",
                    },
                },
                Payload: &shared.AppscopeConfigWithCustomPayload{
                    Dir: "rem",
                    Enable: false,
                },
                Protocol: []shared.AppscopeConfigWithCustomProtocol{
                    shared.AppscopeConfigWithCustomProtocol{
                        Binary: false,
                        Detect: false,
                        Len: 750595,
                        Name: "Floyd Harber",
                        Payload: false,
                        Regex: "nulla",
                    },
                },
                Tags: []shared.AppscopeConfigWithCustomTags{
                    shared.AppscopeConfigWithCustomTags{
                        Key: "aperiam",
                        Value: "saepe",
                    },
                },
            },
            Description: "numquam",
            ID: "57e1858b-6a89-4fbe-ba5a-a8e4824d0ab4",
            Lib: shared.CriblLibCribl,
            Tags: cribl.String("esse"),
        },
        ID: "5088e518-6206-45e9-84f3-b1194b8abf60",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.UpdateAppscopeLibEntryRequest](../../models/operations/updateappscopelibentryrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.UpdateAppscopeLibEntryResponse](../../models/operations/updateappscopelibentryresponse.md), error**

