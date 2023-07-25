# InputAppscopeFilterAllow2


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Arg`                                                                              | *string*                                                                           | :heavy_check_mark:                                                                 | Specify a string to substring-match against process command-line.                  |
| `Config`                                                                           | *string*                                                                           | :heavy_check_mark:                                                                 | Choose a config to apply to processes that match the process name and/or argument. |
| `Procname`                                                                         | **string*                                                                          | :heavy_minus_sign:                                                                 | Specify the name of a process or family of processes.                              |