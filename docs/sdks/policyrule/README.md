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
            "commodi",
            "veniam",
            "debitis",
            "doloremque",
        },
        Description: cribl.String("esse"),
        ID: "6cc7abf6-16ea-45c7-9641-934b90f2e09d",
        Template: []string{
            "natus",
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
        ID: "d2fc2f9e-2e10-4594-8b93-5d237a72f908",
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
        ID: "49d6aed4-aecb-4753-bcd9-222c9ff57491",
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
                "id",
                "harum",
                "sapiente",
            },
            Description: cribl.String("laborum"),
            ID: "2e761f0c-a4d4-456e-b103-1e6899f0c200",
            Template: []string{
                "saepe",
            },
        },
        ID: "22cd55cc-0584-4a18-8d76-d971fc820c65",
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

