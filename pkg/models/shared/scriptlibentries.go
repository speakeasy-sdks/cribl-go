// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ScriptLibEntries - a list of Script objects
type ScriptLibEntries struct {
	// number of items present in the items array
	Count *int64                   `json:"count,omitempty"`
	Items []map[string]interface{} `json:"items,omitempty"`
}

func (o *ScriptLibEntries) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ScriptLibEntries) GetItems() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Items
}
