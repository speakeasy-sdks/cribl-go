// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Conditions - a list of Condition objects
type Conditions struct {
	// number of items present in the items array
	Count *int64      `json:"count,omitempty"`
	Items []Condition `json:"items,omitempty"`
}

func (o *Conditions) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *Conditions) GetItems() []Condition {
	if o == nil {
		return nil
	}
	return o.Items
}
