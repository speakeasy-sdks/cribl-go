# TokenMetadata

### Available Operations

* [Post](#post) - Add token and optional metadata to an existing hec input
* [Update](#update) - Update token metadata on existing hec input

## Post

Add token and optional metadata to an existing hec input

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
    res, err := s.TokenMetadata.Post(ctx, operations.PostTokenMetadataRequest{
        AddHecTokenRequest: &shared.AddHecTokenRequest{
            Description: cribl.String("tempora"),
            Metadata: []shared.AddHecTokenRequestMetadata{
                shared.AddHecTokenRequestMetadata{
                    Name: "Clifton Gutkowski",
                    Value: "sapiente",
                },
                shared.AddHecTokenRequestMetadata{
                    Name: "Roy Bogisich",
                    Value: "amet",
                },
                shared.AddHecTokenRequestMetadata{
                    Name: "Jackie Welch",
                    Value: "libero",
                },
                shared.AddHecTokenRequestMetadata{
                    Name: "Timothy Schaden",
                    Value: "recusandae",
                },
            },
            Token: "adipisci",
        },
        ID: "72db1344-ba9f-478a-9c0e-d7aab62e9726",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostTokenMetadata200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.PostTokenMetadataRequest](../../models/operations/posttokenmetadatarequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.PostTokenMetadataResponse](../../models/operations/posttokenmetadataresponse.md), error**


## Update

Update token metadata on existing hec input

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
    res, err := s.TokenMetadata.Update(ctx, operations.UpdateTokenMetadataRequest{
        UpdateHecTokenRequest: &shared.UpdateHecTokenRequest{
            Description: cribl.String("architecto"),
            Metadata: []shared.UpdateHecTokenRequestMetadata{
                shared.UpdateHecTokenRequestMetadata{
                    Name: "James Schmitt",
                    Value: "temporibus",
                },
                shared.UpdateHecTokenRequestMetadata{
                    Name: "Maureen Rohan V",
                    Value: "cupiditate",
                },
                shared.UpdateHecTokenRequestMetadata{
                    Name: "Hannah Harber",
                    Value: "expedita",
                },
                shared.UpdateHecTokenRequestMetadata{
                    Name: "Ruth Walsh",
                    Value: "iusto",
                },
            },
        },
        ID: "12b7a7ab-0344-4b17-9068-8deebef897f3",
        Token: "fugiat",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateTokenMetadata200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateTokenMetadataRequest](../../models/operations/updatetokenmetadatarequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UpdateTokenMetadataResponse](../../models/operations/updatetokenmetadataresponse.md), error**

