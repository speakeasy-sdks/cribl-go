// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// EdgeMetadatas - a list of EdgeMetadata objects
type EdgeMetadatas struct {
	// number of items present in the items array
	Count *int64 `json:"count,omitempty"`
}

func (o *EdgeMetadatas) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}
