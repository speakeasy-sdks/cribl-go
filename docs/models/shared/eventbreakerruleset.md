# EventBreakerRuleset

New Event Breaker Ruleset object


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Description`                                                                    | **string*                                                                        | :heavy_minus_sign:                                                               | Brief description of this ruleset. Optional.                                     |
| `ID`                                                                             | *string*                                                                         | :heavy_check_mark:                                                               | N/A                                                                              |
| `Lib`                                                                            | [*EventBreakerRulesetLibrary](../../models/shared/eventbreakerrulesetlibrary.md) | :heavy_minus_sign:                                                               | N/A                                                                              |
| `MinRawLength`                                                                   | **int64*                                                                         | :heavy_minus_sign:                                                               | Threshold number of characters in _raw to determine which rule to use.           |
| `Rules`                                                                          | [][EventBreakerRulesetRules](../../models/shared/eventbreakerrulesetrules.md)    | :heavy_minus_sign:                                                               | List of rules. Evaluated in order, top down.                                     |
| `Tags`                                                                           | **string*                                                                        | :heavy_minus_sign:                                                               | One or more tags related to this ruleset. Optional.                              |