# OutputInfluxdbBackpressureBehavior

Whether to block, drop, or queue events when all receivers are exerting backpressure.


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `OutputInfluxdbBackpressureBehaviorQueue` | queue                                     |
| `OutputInfluxdbBackpressureBehaviorDrop`  | drop                                      |
| `OutputInfluxdbBackpressureBehaviorBlock` | block                                     |