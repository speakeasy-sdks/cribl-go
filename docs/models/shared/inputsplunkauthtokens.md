# InputSplunkAuthTokens


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `Description`                                                                                  | **string*                                                                                      | :heavy_minus_sign:                                                                             | Optional token description                                                                     |
| `Token`                                                                                        | *string*                                                                                       | :heavy_check_mark:                                                                             | Shared secrets to be provided by any Splunk forwarder. If empty, unauthed access is permitted. |