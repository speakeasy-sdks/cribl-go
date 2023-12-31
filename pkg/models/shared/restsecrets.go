// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RestSecrets - a list of RestSecret objects
type RestSecrets struct {
	// number of items present in the items array
	Count *int64       `json:"count,omitempty"`
	Items []RestSecret `json:"items,omitempty"`
}

func (o *RestSecrets) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *RestSecrets) GetItems() []RestSecret {
	if o == nil {
		return nil
	}
	return o.Items
}
