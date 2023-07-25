# OutputKafkaBackpressureBehavior

Whether to block, drop, or queue events when all receivers are exerting backpressure.


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `OutputKafkaBackpressureBehaviorQueue` | queue                                  |
| `OutputKafkaBackpressureBehaviorDrop`  | drop                                   |
| `OutputKafkaBackpressureBehaviorBlock` | block                                  |