# DiagBundle

### Available Operations

* [Delete](#delete) - Remove diag bundle
* [Get](#get) - Returns a diag bundle as a tar.gz archive
* [Send](#send) - Send a diag bundle (tar.gz archive) to Cribl

## Delete

Remove diag bundle

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
    res, err := s.DiagBundle.Delete(ctx, operations.DeleteDiagBundleRequest{
        Path: "corrupti",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RemoveDiagResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteDiagBundleRequest](../../models/operations/deletediagbundlerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeleteDiagBundleResponse](../../models/operations/deletediagbundleresponse.md), error**


## Get

Returns a diag bundle as a tar.gz archive

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
    res, err := s.DiagBundle.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDiagBundle200ApplicationTarPlusGzipBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetDiagBundleResponse](../../models/operations/getdiagbundleresponse.md), error**


## Send

Send a diag bundle (tar.gz archive) to Cribl

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
    res, err := s.DiagBundle.Send(ctx, shared.SendDiagBundle{
        IncludeMetrics: cribl.Bool(false),
        MaxIncludeJobs: cribl.Int64(684799),
        Path: cribl.String("facere"),
        RenameJs: cribl.Bool(false),
        SendToCribl: cribl.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RemoveDiagResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [shared.SendDiagBundle](../../models/shared/senddiagbundle.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |


### Response

**[*operations.SendDiagBundleResponse](../../models/operations/senddiagbundleresponse.md), error**

