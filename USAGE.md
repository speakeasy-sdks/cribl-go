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