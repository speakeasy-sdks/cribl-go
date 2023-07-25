# GlobalVar

New Global Variable object


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Description`                                         | **string*                                             | :heavy_minus_sign:                                    | Brief description of this variable. Optional.         |
| `ID`                                                  | *string*                                              | :heavy_check_mark:                                    | Global variable name.                                 |
| `Lib`                                                 | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `Tags`                                                | **string*                                             | :heavy_minus_sign:                                    | One or more tags related to this variable. Optional.  |
| `Type`                                                | [GlobalVarType](../../models/shared/globalvartype.md) | :heavy_check_mark:                                    | Type of variable.                                     |
| `Value`                                               | *string*                                              | :heavy_check_mark:                                    | Value of variable                                     |