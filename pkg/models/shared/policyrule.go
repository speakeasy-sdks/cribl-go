// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PolicyRule - New PolicyRule object
type PolicyRule struct {
	Args        []string `json:"args,omitempty"`
	Description *string  `json:"description,omitempty"`
	ID          string   `json:"id"`
	Template    []string `json:"template"`
}

func (o *PolicyRule) GetArgs() []string {
	if o == nil {
		return nil
	}
	return o.Args
}

func (o *PolicyRule) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PolicyRule) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PolicyRule) GetTemplate() []string {
	if o == nil {
		return []string{}
	}
	return o.Template
}
