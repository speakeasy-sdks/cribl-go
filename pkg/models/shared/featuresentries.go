// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// FeaturesEntries - a list of FeaturesEntry objects
type FeaturesEntries struct {
	// number of items present in the items array
	Count *int64          `json:"count,omitempty"`
	Items []FeaturesEntry `json:"items,omitempty"`
}

func (o *FeaturesEntries) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *FeaturesEntries) GetItems() []FeaturesEntry {
	if o == nil {
		return nil
	}
	return o.Items
}
