// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Users - a list of User objects
type Users struct {
	// number of items present in the items array
	Count *int64 `json:"count,omitempty"`
	Items []User `json:"items,omitempty"`
}

func (o *Users) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *Users) GetItems() []User {
	if o == nil {
		return nil
	}
	return o.Items
}