// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DatasetProviderTypes - a list of DatasetProviderType objects
type DatasetProviderTypes struct {
	// number of items present in the items array
	Count *int64                `json:"count,omitempty"`
	Items []DatasetProviderType `json:"items,omitempty"`
}

func (o *DatasetProviderTypes) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *DatasetProviderTypes) GetItems() []DatasetProviderType {
	if o == nil {
		return nil
	}
	return o.Items
}