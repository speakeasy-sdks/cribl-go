# OutputAzureLogsFailedRequestLoggingMode

Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `OutputAzureLogsFailedRequestLoggingModePayload`           | payload                                                    |
| `OutputAzureLogsFailedRequestLoggingModePayloadAndHeaders` | payloadAndHeaders                                          |
| `OutputAzureLogsFailedRequestLoggingModeNone`              | none                                                       |