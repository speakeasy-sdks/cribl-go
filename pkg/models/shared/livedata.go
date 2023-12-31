// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// LiveData - a list of any objects
type LiveData struct {
	// number of items present in the items array
	Count *int64                   `json:"count,omitempty"`
	Items []map[string]interface{} `json:"items,omitempty"`
}

func (o *LiveData) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *LiveData) GetItems() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Items
}
