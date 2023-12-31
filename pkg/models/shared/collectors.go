// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Collectors - a list of Collector objects
type Collectors struct {
	// number of items present in the items array
	Count *int64      `json:"count,omitempty"`
	Items []Collector `json:"items,omitempty"`
}

func (o *Collectors) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *Collectors) GetItems() []Collector {
	if o == nil {
		return nil
	}
	return o.Items
}
