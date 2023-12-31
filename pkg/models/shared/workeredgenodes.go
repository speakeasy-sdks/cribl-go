// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// WorkerEdgeNodes - a list of number objects
type WorkerEdgeNodes struct {
	// number of items present in the items array
	Count *int64  `json:"count,omitempty"`
	Items []int64 `json:"items,omitempty"`
}

func (o *WorkerEdgeNodes) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *WorkerEdgeNodes) GetItems() []int64 {
	if o == nil {
		return nil
	}
	return o.Items
}
