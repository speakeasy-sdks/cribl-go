# ExportPackRequest


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ID`                                                      | *string*                                                  | :heavy_check_mark:                                        | Pack name                                                 |
| `Mode`                                                    | *string*                                                  | :heavy_check_mark:                                        | Export mode, one of "merge_safe", "merge", "default_only" |
| `Filename`                                                | **string*                                                 | :heavy_minus_sign:                                        | Filename of the exported Pack                             |