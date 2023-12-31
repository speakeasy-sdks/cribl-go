// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ClientRoleEntries - a list of string objects
type ClientRoleEntries struct {
	// number of items present in the items array
	Count *int64   `json:"count,omitempty"`
	Items []string `json:"items,omitempty"`
}

func (o *ClientRoleEntries) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ClientRoleEntries) GetItems() []string {
	if o == nil {
		return nil
	}
	return o.Items
}
