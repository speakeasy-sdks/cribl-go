// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ProfilerItems - a list of ProfilerItem objects
type ProfilerItems struct {
	// number of items present in the items array
	Count *int64         `json:"count,omitempty"`
	Items []ProfilerItem `json:"items,omitempty"`
}

func (o *ProfilerItems) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ProfilerItems) GetItems() []ProfilerItem {
	if o == nil {
		return nil
	}
	return o.Items
}
