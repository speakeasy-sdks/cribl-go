// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Branches - a list of any objects
type Branches struct {
	// number of items present in the items array
	Count *int64                   `json:"count,omitempty"`
	Items []map[string]interface{} `json:"items,omitempty"`
}

func (o *Branches) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *Branches) GetItems() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Items
}
