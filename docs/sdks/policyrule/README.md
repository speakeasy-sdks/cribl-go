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
    res, err := s.PolicyRule.Create(ctx, shared.PolicyRule{
        Args: []string{
            "quas",
        },
        Description: cribl.String("blanditiis"),
        ID: "b1c4ee2c-8c6c-4e61-9fee-b1c7cbdb6eec",
        Template: []string{
            "quaerat",
            "ipsum",
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
    id := "ducimus"

    ctx := context.Background()
    res, err := s.PolicyRule.Delete(ctx, id)
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
    id := "laudantium"

    ctx := context.Background()
    res, err := s.PolicyRule.Get(ctx, id)
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
    id := "rerum"
    policyRule := &shared.PolicyRule{
        Args: []string{
            "odit",
            "ad",
            "sequi",
        },
        Description: cribl.String("beatae"),
        ID: "7747dc91-5ad2-4caf-9dd6-723dc0f5ae2f",
        Template: []string{
            "officia",
        },
    }

    ctx := context.Background()
    res, err := s.PolicyRule.Update(ctx, id, policyRule)
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

