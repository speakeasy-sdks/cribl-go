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
    res, err := s.EventBreakerOnData.Post(ctx, shared.DatatypePreviewRequestBody{
        EventBreakerRule: &shared.EventBreakerRule{
            CleanFields: cribl.Bool(false),
            Condition: "eum",
            Delimiter: cribl.String("reprehenderit"),
            DelimiterRegex: cribl.String("voluptatum"),
            Disabled: cribl.Bool(false),
            EscapeChar: cribl.String("blanditiis"),
            EventBreakerRegex: cribl.String("nihil"),
            Fields: []shared.EventBreakerRuleFields{
                shared.EventBreakerRuleFields{
                    Name: "Doug Littel",
                    Value: "architecto",
                },
                shared.EventBreakerRuleFields{
                    Name: "Lloyd Ledner V",
                    Value: "placeat",
                },
                shared.EventBreakerRuleFields{
                    Name: "Clara Williamson",
                    Value: "officia",
                },
            },
            FieldsLineRegex: cribl.String("natus"),
            HeaderLineRegex: cribl.String("cumque"),
            JSONArrayField: cribl.String("natus"),
            JSONExtractAll: cribl.Bool(false),
            JSONTimeField: cribl.String("quaerat"),
            MaxEventBytes: 985379,
            Name: "Gretchen O'Hara",
            NullFieldVal: cribl.String("enim"),
            Parser: cribl.String("eum"),
            ParserEnabled: cribl.Bool(false),
            QuoteChar: cribl.String("nemo"),
            TimeField: cribl.String("illum"),
            Timestamp: shared.EventBreakerRuleTimestamp{
                Format: cribl.String("nesciunt"),
                Length: cribl.Int64(22331),
                Type: shared.EventBreakerRuleTimestampTypeFormat,
            },
            TimestampAnchorRegex: "minus",
            TimestampEarliest: cribl.String("asperiores"),
            TimestampLatest: cribl.String("recusandae"),
            TimestampTimezone: "voluptates",
            Type: shared.EventBreakerRuleTypeHeader.ToPointer(),
        },
        Input: "dicta",
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

