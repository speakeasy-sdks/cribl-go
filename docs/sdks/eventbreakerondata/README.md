# EventBreakerOnData

### Available Operations

* [Post](#post) - Runs an event breaker rule on the specified data

## Post

Runs an event breaker rule on the specified data

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
    res, err := s.EventBreakerOnData.Post(ctx, shared.DatatypePreviewRequestBody{
        EventBreakerRule: &shared.EventBreakerRule{
            CleanFields: cribl.Bool(false),
            Condition: "doloremque",
            Delimiter: cribl.String("quis"),
            DelimiterRegex: cribl.String("veniam"),
            Disabled: cribl.Bool(false),
            EscapeChar: cribl.String("libero"),
            EventBreakerRegex: cribl.String("architecto"),
            Fields: []shared.EventBreakerRuleFields{
                shared.EventBreakerRuleFields{
                    Name: "Sheri Schuppe",
                    Value: "itaque",
                },
                shared.EventBreakerRuleFields{
                    Name: "Ollie Harris",
                    Value: "laudantium",
                },
                shared.EventBreakerRuleFields{
                    Name: "Freda Farrell I",
                    Value: "facilis",
                },
            },
            FieldsLineRegex: cribl.String("tempore"),
            HeaderLineRegex: cribl.String("nisi"),
            JSONArrayField: cribl.String("voluptatibus"),
            JSONExtractAll: cribl.Bool(false),
            JSONTimeField: cribl.String("quaerat"),
            MaxEventBytes: 503748,
            Name: "Charlie Harvey",
            NullFieldVal: cribl.String("minus"),
            Parser: cribl.String("facere"),
            ParserEnabled: cribl.Bool(false),
            QuoteChar: cribl.String("facilis"),
            TimeField: cribl.String("ipsum"),
            Timestamp: shared.EventBreakerRuleTimestamp{
                Format: cribl.String("ad"),
                Length: cribl.Int64(973819),
                Type: shared.EventBreakerRuleTimestampTypeCurrent,
            },
            TimestampAnchorRegex: "consequuntur",
            TimestampEarliest: cribl.String("debitis"),
            TimestampLatest: cribl.String("labore"),
            TimestampTimezone: "rerum",
            Type: shared.EventBreakerRuleTypeJSON.ToPointer(),
        },
        Input: "reprehenderit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PreviewResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [shared.DatatypePreviewRequestBody](../../models/shared/datatypepreviewrequestbody.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PostEventBreakerOnDataResponse](../../models/operations/posteventbreakerondataresponse.md), error**

