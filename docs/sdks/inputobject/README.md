# InputObject

### Available Operations

* [Get](#get) - Get a list of Input objects
* [Post](#post) - Create Input

## Get

Get a list of Input objects

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
    res, err := s.InputObject.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Inputs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetInputObjectResponse](../../models/operations/getinputobjectresponse.md), error**


## Post

Create Input

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
    res, err := s.InputObject.Post(ctx, shared.InputKinesis{
        AssumeRoleArn: cribl.String("modi"),
        AssumeRoleExternalID: cribl.String("cumque"),
        AvoidDuplicates: cribl.Bool(false),
        AwsAPIKey: cribl.String("ipsam"),
        AwsAuthenticationMethod: shared.InputKinesisAuthenticationMethodManual.ToPointer(),
        AwsSecret: cribl.String("ipsum"),
        AwsSecretKey: cribl.String("accusamus"),
        Connections: []shared.InputKinesisConnections{
            shared.InputKinesisConnections{
                Output: "quasi",
                Pipeline: cribl.String("fugit"),
            },
            shared.InputKinesisConnections{
                Output: "quo",
                Pipeline: cribl.String("temporibus"),
            },
            shared.InputKinesisConnections{
                Output: "mollitia",
                Pipeline: cribl.String("id"),
            },
            shared.InputKinesisConnections{
                Output: "quibusdam",
                Pipeline: cribl.String("ipsa"),
            },
        },
        Disabled: cribl.Bool(false),
        EnableAssumeRole: cribl.Bool(false),
        Endpoint: cribl.String("accusamus"),
        Environment: cribl.String("placeat"),
        GetRecordsLimit: cribl.Int64(464878),
        GetRecordsLimitTotal: cribl.Int64(627756),
        ID: cribl.String("fedbd80d-f448-4a47-b939-0c58880983da"),
        LoadBalancingAlgorithm: shared.InputKinesisShardLoadBalancingRoundRobin.ToPointer(),
        Metadata: []shared.InputKinesisMetadata{
            shared.InputKinesisMetadata{
                Name: "Clay Will",
                Value: "voluptatibus",
            },
            shared.InputKinesisMetadata{
                Name: "Levi Mohr",
                Value: "voluptatibus",
            },
            shared.InputKinesisMetadata{
                Name: "Lillie Moen",
                Value: "modi",
            },
            shared.InputKinesisMetadata{
                Name: "Mike Heaney",
                Value: "non",
            },
        },
        PayloadFormat: shared.InputKinesisRecordDataFormatLine.ToPointer(),
        Pipeline: cribl.String("fugiat"),
        Pq: &shared.InputKinesisPq{
            CommitFrequency: cribl.Int64(743938),
            Compress: shared.InputKinesisPqCompressionNone.ToPointer(),
            MaxBufferSize: cribl.Int64(967047),
            MaxFileSize: cribl.String("labore"),
            MaxSize: cribl.String("vero"),
            Mode: shared.InputKinesisPqModeSmart.ToPointer(),
            Path: cribl.String("quas"),
        },
        PqEnabled: cribl.Bool(false),
        Region: shared.InputKinesisRegionUsWest1,
        RejectUnauthorized: cribl.Bool(false),
        ReuseConnections: cribl.Bool(false),
        SendToRoutes: cribl.Bool(false),
        ServiceInterval: cribl.Int64(101284),
        ShardExpr: cribl.String("praesentium"),
        ShardIteratorType: shared.InputKinesisShardIteratorStartTrimHorizon.ToPointer(),
        SignatureVersion: shared.InputKinesisSignatureVersionV4.ToPointer(),
        StreamName: "enim",
        Streamtags: []string{
            "laudantium",
            "modi",
        },
        Type: shared.InputKinesisTypeKinesis.ToPointer(),
        VerifyKPLCheckSums: cribl.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Inputs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [interface{}](../../models//.md)                      | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostInputObjectResponse](../../models/operations/postinputobjectresponse.md), error**

