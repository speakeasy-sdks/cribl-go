# Policies

### Available Operations

* [Create](#create) - Create PolicyRule
* [Delete](#delete) - Delete PolicyRule
* [Get](#get) - Get PolicyRule by ID
* [ListPolicyRules](#listpolicyrules) - Get a list of PolicyRule objects
* [Update](#update) - Update PolicyRule

## Create

Create PolicyRule

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Policies.Create(ctx, shared.PolicyRule{
        Args: []string{
            "cum",
            "id",
        },
        Description: cribl.String("possimus"),
        ID: "2553819b-474b-40ed-a0e5-6248fff639a9",
        Template: []string{
            "ipsa",
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
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )
    id := "mollitia"

    ctx := context.Background()
    res, err := s.Policies.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.PolicyRules != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


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
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )
    id := "nam"

    ctx := context.Background()
    res, err := s.Policies.Get(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.PolicyRules != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique ID                                             |


### Response

**[*operations.GetPolicyRuleResponse](../../models/operations/getpolicyruleresponse.md), error**


## ListPolicyRules

Get a list of PolicyRule objects

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Policies.ListPolicyRules(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.PolicyRules != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListPolicyRulesResponse](../../models/operations/listpolicyrulesresponse.md), error**


## Update

Update PolicyRule

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/cribl-go"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"github.com/speakeasy-sdks/cribl-go/pkg/models/operations"
)

func main() {
    s := cribl.New(
        cribl.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )
    id := "assumenda"
    policyRule := &shared.PolicyRule{
        Args: []string{
            "fuga",
            "tempore",
            "commodi",
            "fugit",
        },
        Description: cribl.String("suscipit"),
        ID: "76696e1e-c002-421b-b35d-89acb3ecfda8",
        Template: []string{
            "sit",
            "quo",
            "veniam",
            "aliquam",
        },
    }

    ctx := context.Background()
    res, err := s.Policies.Update(ctx, id, policyRule)
    if err != nil {
        log.Fatal(err)
    }

    if res.PolicyRules != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                               | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ctx`                                                   | [context.Context](https://pkg.go.dev/context#Context)   | :heavy_check_mark:                                      | The context to use for the request.                     |
| `id`                                                    | *string*                                                | :heavy_check_mark:                                      | Unique ID                                               |
| `policyRule`                                            | [*shared.PolicyRule](../../models/shared/policyrule.md) | :heavy_minus_sign:                                      | PolicyRule object to be updated                         |


### Response

**[*operations.UpdatePolicyRuleResponse](../../models/operations/updatepolicyruleresponse.md), error**

