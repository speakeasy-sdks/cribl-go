# OutputLokiFailedRequestLoggingMode

Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `OutputLokiFailedRequestLoggingModePayload`           | payload                                               |
| `OutputLokiFailedRequestLoggingModePayloadAndHeaders` | payloadAndHeaders                                     |
| `OutputLokiFailedRequestLoggingModeNone`              | none                                                  |