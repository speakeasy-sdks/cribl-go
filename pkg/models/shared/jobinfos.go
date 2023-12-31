// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// JobInfos - a list of JobInfo objects
type JobInfos struct {
	// the pre-limited items in the list of results
	Items []JobInfo `json:"items,omitempty"`
	// number of items present in the items array
	Limit *int64 `json:"limit,omitempty"`
	// pagination offset
	Offset *int64 `json:"offset,omitempty"`
}

func (o *JobInfos) GetItems() []JobInfo {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *JobInfos) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *JobInfos) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}
