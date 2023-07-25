# InputID

### Available Operations

* [Delete](#delete) - Delete Input
* [Get](#get) - Get Input by ID
* [Update](#update) - Update Input

## Delete

Delete Input

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
    res, err := s.InputID.Delete(ctx, operations.DeleteInputIDRequest{
        ID: "67b8e3c8-db03-4408-96d3-64ffd455906d",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteInputIDRequest](../../models/operations/deleteinputidrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.DeleteInputIDResponse](../../models/operations/deleteinputidresponse.md), error**


## Get

Get Input by ID

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
    res, err := s.InputID.Get(ctx, operations.GetInputIDRequest{
        ID: "1263d48e-935c-42c9-a81f-30be3e43202d",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetInputIDRequest](../../models/operations/getinputidrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetInputIDResponse](../../models/operations/getinputidresponse.md), error**


## Update

Update Input

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
    res, err := s.InputID.Update(ctx, operations.UpdateInputIDRequest{
        RequestBody: &shared.InputCriblHTTP{
            ActivityLogSampleRate: cribl.Int64(168976),
            AuthTokens: []string{
                "aliquid",
            },
            CaptureHeaders: cribl.Bool(false),
            Connections: []shared.InputCriblHTTPConnections{
                shared.InputCriblHTTPConnections{
                    Output: "voluptate",
                    Pipeline: cribl.String("vel"),
                },
                shared.InputCriblHTTPConnections{
                    Output: "minima",
                    Pipeline: cribl.String("sit"),
                },
            },
            Disabled: cribl.Bool(false),
            EnableProxyHeader: cribl.Bool(false),
            Environment: cribl.String("vel"),
            Host: "laboriosam",
            ID: cribl.String("41870d9d-21f9-4ad0-b0c4-ecc11a083642"),
            KeepAliveTimeout: cribl.Int64(616183),
            MaxActiveReq: cribl.Int64(32719),
            Metadata: []shared.InputCriblHTTPMetadata{
                shared.InputCriblHTTPMetadata{
                    Name: "Andres Larkin Sr.",
                    Value: "officia",
                },
                shared.InputCriblHTTPMetadata{
                    Name: "Ana Torp",
                    Value: "esse",
                },
            },
            Pipeline: cribl.String("neque"),
            Port: 697591,
            Pq: &shared.InputCriblHTTPPq{
                CommitFrequency: cribl.Int64(788469),
                Compress: shared.InputCriblHTTPPqCompressionGzip.ToPointer(),
                MaxBufferSize: cribl.Int64(273349),
                MaxFileSize: cribl.String("ipsam"),
                MaxSize: cribl.String("officiis"),
                Mode: shared.InputCriblHTTPPqModeSmart.ToPointer(),
                Path: cribl.String("magni"),
            },
            PqEnabled: cribl.Bool(false),
            RequestTimeout: cribl.Int64(29881),
            SendToRoutes: cribl.Bool(false),
            SocketTimeout: cribl.Int64(665960),
            Streamtags: []string{
                "veritatis",
            },
            TLS: &shared.InputCriblHTTPTLSSettingsServerSide{
                CaPath: cribl.String("error"),
                CertPath: cribl.String("voluptatibus"),
                CertificateName: cribl.String("numquam"),
                CommonNameRegex: cribl.String("rerum"),
                Disabled: cribl.Bool(false),
                MaxVersion: shared.InputCriblHTTPTLSSettingsServerSideMaximumTLSVersionTlSv12.ToPointer(),
                MinVersion: shared.InputCriblHTTPTLSSettingsServerSideMinimumTLSVersionTlSv13.ToPointer(),
                Passphrase: cribl.String("earum"),
                PrivKeyPath: cribl.String("excepturi"),
                RejectUnauthorized: cribl.String("numquam"),
                RequestCert: cribl.Bool(false),
            },
            Type: shared.InputCriblHTTPTypeCriblHTTP.ToPointer(),
        },
        ID: "7c9a867b-c424-4266-a581-6ddca8ef51fc",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateInputIDRequest](../../models/operations/updateinputidrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateInputIDResponse](../../models/operations/updateinputidresponse.md), error**

