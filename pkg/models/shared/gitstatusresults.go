// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GitStatusResults - a list of GitStatusResult objects
type GitStatusResults struct {
	// number of items present in the items array
	Count *int64            `json:"count,omitempty"`
	Items []GitStatusResult `json:"items,omitempty"`
}

func (o *GitStatusResults) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GitStatusResults) GetItems() []GitStatusResult {
	if o == nil {
		return nil
	}
	return o.Items
}
