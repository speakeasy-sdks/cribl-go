# Output402PrometheusAuth


## Fields

| Field                                                                                                                                                                                                     | Type                                                                                                                                                                                                      | Required                                                                                                                                                                                                  | Description                                                                                                                                                                                               |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `AuthType`                                                                                                                                                                                                | [*Output402PrometheusAuthAuthenticationType](../../models/shared/output402prometheusauthauthenticationtype.md)                                                                                            | :heavy_minus_sign:                                                                                                                                                                                        | The authentication method to use for the HTTP requests                                                                                                                                                    |
| `CredentialsSecret`                                                                                                                                                                                       | **string*                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                        | Select (or create) a secret that references your credentials                                                                                                                                              |
| `Password`                                                                                                                                                                                                | **string*                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                        | Password (a.k.a API key in Grafana Cloud domain) for authentication                                                                                                                                       |
| `TextSecret`                                                                                                                                                                                              | **string*                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                        | Select (or create) a stored text secret                                                                                                                                                                   |
| `Token`                                                                                                                                                                                                   | **string*                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                        | Bearer token to include in the authorization header. In Grafana Cloud, this is generally built by concatenating the username and the API key, separated by a colon. E.g.: <your-username>:<your-api-key>. |
| `Username`                                                                                                                                                                                                | **string*                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                        | Username for authentication                                                                                                                                                                               |