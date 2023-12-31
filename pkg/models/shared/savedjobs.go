// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SavedJobs - a list of SavedJob objects
type SavedJobs struct {
	// number of items present in the items array
	Count *int64     `json:"count,omitempty"`
	Items []SavedJob `json:"items,omitempty"`
}

func (o *SavedJobs) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *SavedJobs) GetItems() []SavedJob {
	if o == nil {
		return nil
	}
	return o.Items
}
