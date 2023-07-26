# OutputOpenTelemetryBackpressureBehavior

Whether to block, drop, or queue events when all receivers are exerting backpressure.


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `OutputOpenTelemetryBackpressureBehaviorQueue` | queue                                          |
| `OutputOpenTelemetryBackpressureBehaviorDrop`  | drop                                           |
| `OutputOpenTelemetryBackpressureBehaviorBlock` | block                                          |