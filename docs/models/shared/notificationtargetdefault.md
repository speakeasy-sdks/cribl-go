# NotificationTargetDefault


## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `DefaultID`                                                                                                | *string*                                                                                                   | :heavy_check_mark:                                                                                         | ID of the default output. This will be used whenever a nonexistent/deleted output is referenced.           |
| `ID`                                                                                                       | *string*                                                                                                   | :heavy_check_mark:                                                                                         | Unique ID for this output                                                                                  |
| `SystemFields`                                                                                             | []*string*                                                                                                 | :heavy_minus_sign:                                                                                         | Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported. |
| `Type`                                                                                                     | [NotificationTargetDefaultType](../../models/shared/notificationtargetdefaulttype.md)                      | :heavy_check_mark:                                                                                         | N/A                                                                                                        |