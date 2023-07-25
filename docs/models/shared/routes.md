# Routes

a list of Routes objects


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Comments`                                                     | []map[string]*interface{}*                                     | :heavy_minus_sign:                                             | Comments                                                       |
| `Groups`                                                       | map[string][RoutesGroups](../../models/shared/routesgroups.md) | :heavy_minus_sign:                                             | N/A                                                            |
| `ID`                                                           | **string*                                                      | :heavy_minus_sign:                                             | Routes ID                                                      |
| `Routes`                                                       | []map[string]*interface{}*                                     | :heavy_check_mark:                                             | Pipeline routing rules                                         |