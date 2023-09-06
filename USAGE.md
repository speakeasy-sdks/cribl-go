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
                    Ancestor: cribl.String("illum"),
                    Arg: cribl.String("vel"),
                    Config: shared.AppscopeConfig{
                        Cribl: &shared.AppscopeConfigCribl{
                            Authtoken: cribl.String("error"),
                            Enable: cribl.Bool(false),
                            Transport: &shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("suscipit"),
                                Path: cribl.String("iure"),
                                Port: cribl.Int64(297534),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("debitis"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("ipsa"),
                            },
                            UseScopeSourceTransport: cribl.Bool(false),
                        },
                        Event: &shared.AppscopeConfigEvent{
                            Enable: false,
                            Format: shared.AppscopeConfigEventFormat{
                                Enhancefs: false,
                                Maxeventpersec: 963663,
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                                Host: cribl.String("suscipit"),
                                Path: cribl.String("molestiae"),
                                Port: cribl.Int64(791725),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("placeat"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("voluptatum"),
                            },
                            Type: shared.AppscopeConfigEventTypeNdjson,
                            Watch: []shared.AppscopeConfigEventWatch{
                                shared.AppscopeConfigEventWatch{
                                    Allowbinary: cribl.Bool(false),
                                    Enabled: cribl.Bool(false),
                                    Field: cribl.String("iusto"),
                                    Headers: cribl.String("excepturi"),
                                    Name: cribl.String("Mrs. Sophie Smith MD"),
                                    Type: "perferendis",
                                    Value: cribl.String("ipsam"),
                                },
                            },
                        },
                        Libscope: &shared.AppscopeConfigLibscope{
                            Commanddir: cribl.String("repellendus"),
                            Configevent: cribl.Bool(false),
                            Log: &shared.AppscopeConfigLibscopeLog{
                                Level: shared.AppscopeConfigLibscopeLogLevelNone.ToPointer(),
                                Transport: &shared.AppscopeTransport{
                                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                    Host: cribl.String("odit"),
                                    Path: cribl.String("at"),
                                    Port: cribl.Int64(870088),
                                    TLS: &shared.AppscopeTransportTLS{
                                        Cacertpath: cribl.String("maiores"),
                                        Enable: cribl.Bool(false),
                                        Validateserver: cribl.Bool(false),
                                    },
                                    Type: cribl.String("molestiae"),
                                },
                            },
                            Summaryperiod: cribl.Int64(799159),
                        },
                        Metric: &shared.AppscopeConfigMetric{
                            Enable: false,
                            Format: shared.AppscopeConfigMetricFormat{
                                Statsdmaxlen: cribl.Int64(800911),
                                Statsdprefix: cribl.String("esse"),
                                Type: cribl.String("totam"),
                                Verbosity: cribl.Int64(780529),
                            },
                            Transport: shared.AppscopeTransport{
                                Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                                Host: cribl.String("dicta"),
                                Path: cribl.String("nam"),
                                Port: cribl.Int64(639921),
                                TLS: &shared.AppscopeTransportTLS{
                                    Cacertpath: cribl.String("occaecati"),
                                    Enable: cribl.Bool(false),
                                    Validateserver: cribl.Bool(false),
                                },
                                Type: cribl.String("fugit"),
                            },
                            Watch: []string{
                                "deleniti",
                            },
                        },
                        Payload: &shared.AppscopeConfigPayload{
                            Dir: "hic",
                            Enable: false,
                        },
                        Protocol: []shared.AppscopeConfigProtocol{
                            shared.AppscopeConfigProtocol{
                                Binary: false,
                                Detect: false,
                                Len: 758616,
                                Name: "Jack Johns",
                                Payload: false,
                                Regex: "qui",
                            },
                        },
                        Tags: []shared.AppscopeConfigTags{
                            shared.AppscopeConfigTags{
                                Key: "impedit",
                                Value: "cum",
                            },
                        },
                    },
                    Env: cribl.String("esse"),
                    Hostname: cribl.String("dull-mister.com"),
                    Procname: cribl.String("perferendis"),
                    Username: cribl.String("Enrique61"),
                },
            },
            Event: &shared.AppscopeConfigWithCustomEvent{
                Enable: false,
                Format: shared.AppscopeConfigWithCustomEventFormat{
                    Enhancefs: false,
                    Maxeventpersec: 222321,
                },
                Transport: shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                    Host: cribl.String("laboriosam"),
                    Path: cribl.String("hic"),
                    Port: cribl.Int64(902599),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("fuga"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("in"),
                },
                Type: shared.AppscopeConfigWithCustomEventTypeNdjson,
                Watch: []shared.AppscopeConfigWithCustomEventWatch{
                    shared.AppscopeConfigWithCustomEventWatch{
                        Allowbinary: cribl.Bool(false),
                        Enabled: cribl.Bool(false),
                        Field: cribl.String("corporis"),
                        Headers: cribl.String("iste"),
                        Name: cribl.String("Mr. Kerry Predovic"),
                        Type: "est",
                        Value: cribl.String("mollitia"),
                    },
                },
            },
            Libscope: &shared.AppscopeConfigWithCustomLibscope{
                Commanddir: cribl.String("laborum"),
                Configevent: cribl.Bool(false),
                Log: &shared.AppscopeConfigWithCustomLibscopeLog{
                    Level: shared.AppscopeConfigWithCustomLibscopeLogLevelDebug.ToPointer(),
                    Transport: &shared.AppscopeTransport{
                        Buffer: shared.AppscopeTransportBufferLine.ToPointer(),
                        Host: cribl.String("corporis"),
                        Path: cribl.String("explicabo"),
                        Port: cribl.Int64(750686),
                        TLS: &shared.AppscopeTransportTLS{
                            Cacertpath: cribl.String("enim"),
                            Enable: cribl.Bool(false),
                            Validateserver: cribl.Bool(false),
                        },
                        Type: cribl.String("omnis"),
                    },
                },
                Summaryperiod: cribl.Int64(363711),
            },
            Metric: &shared.AppscopeConfigWithCustomMetric{
                Enable: false,
                Format: shared.AppscopeConfigWithCustomMetricFormat{
                    Statsdmaxlen: cribl.Int64(325047),
                    Statsdprefix: cribl.String("excepturi"),
                    Type: cribl.String("accusantium"),
                    Verbosity: cribl.Int64(438601),
                },
                Transport: shared.AppscopeTransport{
                    Buffer: shared.AppscopeTransportBufferFull.ToPointer(),
                    Host: cribl.String("doloribus"),
                    Path: cribl.String("sapiente"),
                    Port: cribl.Int64(102044),
                    TLS: &shared.AppscopeTransportTLS{
                        Cacertpath: cribl.String("mollitia"),
                        Enable: cribl.Bool(false),
                        Validateserver: cribl.Bool(false),
                    },
                    Type: cribl.String("dolorem"),
                },
                Watch: []string{
                    "culpa",
                },
            },
            Payload: &shared.AppscopeConfigWithCustomPayload{
                Dir: "consequuntur",
                Enable: false,
            },
            Protocol: []shared.AppscopeConfigWithCustomProtocol{
                shared.AppscopeConfigWithCustomProtocol{
                    Binary: false,
                    Detect: false,
                    Len: 995300,
                    Name: "Tracy Fritsch",
                    Payload: false,
                    Regex: "molestiae",
                },
            },
            Tags: []shared.AppscopeConfigWithCustomTags{
                shared.AppscopeConfigWithCustomTags{
                    Key: "velit",
                    Value: "error",
                },
            },
        },
        Description: "quia",
        ID: "51aa52c3-f5ad-4019-9a1f-fe78f097b007",
        Lib: shared.CriblLibCribl,
        Tags: cribl.String("maiores"),
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