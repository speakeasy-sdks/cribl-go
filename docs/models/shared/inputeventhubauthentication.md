# InputEventhubAuthentication

Authentication parameters to use when connecting to brokers. Using TLS is highly recommended.


## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `Disabled`                                                                                                   | *bool*                                                                                                       | :heavy_check_mark:                                                                                           | Enable authentication.                                                                                       |
| `Mechanism`                                                                                                  | [*InputEventhubAuthenticationSASLMechanism](../../models/shared/inputeventhubauthenticationsaslmechanism.md) | :heavy_minus_sign:                                                                                           | SASL authentication mechanism to use                                                                         |