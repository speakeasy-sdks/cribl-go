// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ExprLibEntries - a list of ExprLibEntry objects
type ExprLibEntries struct {
	// number of items present in the items array
	Count *int64 `json:"count,omitempty"`
}

func (o *ExprLibEntries) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}
