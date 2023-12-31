// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GrokFiles - a list of GrokFile objects
type GrokFiles struct {
	// number of items present in the items array
	Count *int64     `json:"count,omitempty"`
	Items []GrokFile `json:"items,omitempty"`
}

func (o *GrokFiles) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GrokFiles) GetItems() []GrokFile {
	if o == nil {
		return nil
	}
	return o.Items
}
