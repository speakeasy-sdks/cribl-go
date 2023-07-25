# SearchConfig


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `Datasets`                                          | []*string*                                          | :heavy_check_mark:                                  | N/A                                                 |
| `HasSendOperator`                                   | *bool*                                              | :heavy_check_mark:                                  | N/A                                                 |
| `OrderedFieldNames`                                 | []*string*                                          | :heavy_check_mark:                                  | N/A                                                 |
| `SearchTerms`                                       | [][SearchTerm](../../models/shared/searchterm.md)   | :heavy_check_mark:                                  | N/A                                                 |
| `SortFields`                                        | [][SortByField](../../models/shared/sortbyfield.md) | :heavy_minus_sign:                                  | N/A                                                 |
| `SuppressEventMarking`                              | *bool*                                              | :heavy_check_mark:                                  | N/A                                                 |
| `UseFormattedVisualization`                         | *bool*                                              | :heavy_check_mark:                                  | N/A                                                 |