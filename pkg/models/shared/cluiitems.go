// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CluiItems - a list of CluiItem objects
type CluiItems struct {
	// number of items present in the items array
	Count *int64     `json:"count,omitempty"`
	Items []CluiItem `json:"items,omitempty"`
}

func (o *CluiItems) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *CluiItems) GetItems() []CluiItem {
	if o == nil {
		return nil
	}
	return o.Items
}
