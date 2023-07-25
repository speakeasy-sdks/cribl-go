// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Role - New Role object
type Role struct {
	ID     string   `json:"id"`
	Policy []string `json:"policy"`
}

func (o *Role) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Role) GetPolicy() []string {
	if o == nil {
		return []string{}
	}
	return o.Policy
}
