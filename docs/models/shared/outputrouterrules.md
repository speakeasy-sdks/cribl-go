# OutputRouterRules


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Description`                                                                    | **string*                                                                        | :heavy_minus_sign:                                                               | Description of this rule's purpose                                               |
| `Filter`                                                                         | *string*                                                                         | :heavy_check_mark:                                                               | JavaScript expression to select events to send to output                         |
| `Final`                                                                          | **bool*                                                                          | :heavy_minus_sign:                                                               | Flag to control whether to stop the event from being checked against other rules |
| `Output`                                                                         | *string*                                                                         | :heavy_check_mark:                                                               | Output to send matching events to                                                |