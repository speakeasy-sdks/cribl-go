// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ParserLibEntries - a list of Parser objects
type ParserLibEntries struct {
	// number of items present in the items array
	Count *int64                   `json:"count,omitempty"`
	Items []map[string]interface{} `json:"items,omitempty"`
}

func (o *ParserLibEntries) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ParserLibEntries) GetItems() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Items
}
