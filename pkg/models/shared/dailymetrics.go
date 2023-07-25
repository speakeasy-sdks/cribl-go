// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DailyMetrics - a list of DailyUsageMetrics objects
type DailyMetrics struct {
	// number of items present in the items array
	Count *int64              `json:"count,omitempty"`
	Items []DailyUsageMetrics `json:"items,omitempty"`
}

func (o *DailyMetrics) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *DailyMetrics) GetItems() []DailyUsageMetrics {
	if o == nil {
		return nil
	}
	return o.Items
}
