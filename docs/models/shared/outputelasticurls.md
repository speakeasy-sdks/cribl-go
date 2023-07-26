# OutputElasticUrls


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `URL`                                                                      | *string*                                                                   | :heavy_check_mark:                                                         | URL to an Elastic node to send events to – e.g., http://elastic:9200/_bulk |
| `Weight`                                                                   | **int64*                                                                   | :heavy_minus_sign:                                                         | The weight to use for load-balancing purposes.                             |