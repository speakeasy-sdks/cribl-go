// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpgradeResults - a list of UpgradeResult objects
type UpgradeResults struct {
	// number of items present in the items array
	Count *int64 `json:"count,omitempty"`
}

func (o *UpgradeResults) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}