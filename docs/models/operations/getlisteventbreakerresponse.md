# GetListEventBreakerResponse


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ContentType`                                                               | *string*                                                                    | :heavy_check_mark:                                                          | N/A                                                                         |
| `EventBreakerRulesets`                                                      | [*shared.EventBreakerRulesets](../../models/shared/eventbreakerrulesets.md) | :heavy_minus_sign:                                                          | a list of Event Breaker Ruleset objects                                     |
| `StatusCode`                                                                | *int*                                                                       | :heavy_check_mark:                                                          | N/A                                                                         |
| `RawResponse`                                                               | [*http.Response](https://pkg.go.dev/net/http#Response)                      | :heavy_minus_sign:                                                          | N/A                                                                         |