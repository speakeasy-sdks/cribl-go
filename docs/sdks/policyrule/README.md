# PolicyRule

### Available Operations

* [Create](#create) - Create PolicyRule
* [Delete](#delete) - Delete PolicyRule
* [Get](#get) - Get PolicyRule by ID
* [Update](#update) - Update PolicyRule

## Create

Create PolicyRule

### Example Usage

```go
package main

import(
	"context"
	"log"
	"cribl"
	"cribl/pkg/models/shared"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.PolicyRule.Create(ctx, shared.PolicyRule{
        Args: []string{
            "nemo",
            "dolore",
            "numquam",
        },
        Description: cribl.String("dolor"),
        ID: "9ee22446-0443-4bc1-9418-8c2f56e85da7",
        Template: []string{
            "dolorem",
            "odit",
            "officiis",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PolicyRules != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                              | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ctx`                                                  | [context.Context](https://pkg.go.dev/context#Context)  | :heavy_check_mark:                                     | The context to use for the request.                    |
| `request`                                              | [shared.PolicyRule](../../models/shared/policyrule.md) | :heavy_check_mark:                                     | The request object to use for the request.             |


### Response

**[*operations.CreatePolicyRuleResponse](../../models/operations/createpolicyruleresponse.md), error**


## Delete

Delete PolicyRule

### Example Usage

```go
package main

import(
	"context"
	"log"
	"cribl"
	"cribl/pkg/models/shared"
	"cribl/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.PolicyRule.Delete(ctx, operations.DeletePolicyRuleRequest{
        ID: "abd617c3-b0d5-41a4-8bf0-1bad8706d460",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PolicyRules != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeletePolicyRuleRequest](../../models/operations/deletepolicyrulerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeletePolicyRuleResponse](../../models/operations/deletepolicyruleresponse.md), error**


## Get

Get PolicyRule by ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"cribl"
	"cribl/pkg/models/shared"
	"cribl/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.PolicyRule.Get(ctx, operations.GetPolicyRuleRequest{
        ID: "82bfbdc4-1ff5-4d4e-aae4-fb5cb35d1763",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PolicyRules != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetPolicyRuleRequest](../../models/operations/getpolicyrulerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetPolicyRuleResponse](../../models/operations/getpolicyruleresponse.md), error**


## Update

Update PolicyRule

### Example Usage

```go
package main

import(
	"context"
	"log"
	"cribl"
	"cribl/pkg/models/shared"
	"cribl/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.PolicyRule.Update(ctx, operations.UpdatePolicyRuleRequest{
        PolicyRule: &shared.PolicyRule{
            Args: []string{
                "delectus",
                "sunt",
                "repudiandae",
            },
            Description: cribl.String("facere"),
            ID: "b78359ec-c5cb-4860-b8cd-580ba73810e4",
            Template: []string{
                "necessitatibus",
                "tempora",
                "quaerat",
                "magnam",
            },
        },
        ID: "7297cd3b-1dd3-4bbc-a247-b7684eff5012",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PolicyRules != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdatePolicyRuleRequest](../../models/operations/updatepolicyrulerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdatePolicyRuleResponse](../../models/operations/updatepolicyruleresponse.md), error**

