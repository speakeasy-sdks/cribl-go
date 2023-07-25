# Container


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Containerd`                                             | [*ContainerdInfo](../../models/shared/containerdinfo.md) | :heavy_minus_sign:                                       | N/A                                                      |
| `Created`                                                | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `Docker`                                                 | [*DockerInfo](../../models/shared/dockerinfo.md)         | :heavy_minus_sign:                                       | N/A                                                      |
| `ID`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `Image`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `Ips`                                                    | []*string*                                               | :heavy_minus_sign:                                       | N/A                                                      |
| `Name`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `Ports`                                                  | [][ContainerPort](../../models/shared/containerport.md)  | :heavy_minus_sign:                                       | N/A                                                      |
| `Status`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `Type`                                                   | [ContainerType](../../models/shared/containertype.md)    | :heavy_check_mark:                                       | N/A                                                      |