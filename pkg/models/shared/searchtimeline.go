// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SearchTimeline - SearchTimeline object
type SearchTimeline struct {
	Buckets         []Bucket `json:"buckets"`
	TotalEventCount int64    `json:"totalEventCount"`
}

func (o *SearchTimeline) GetBuckets() []Bucket {
	if o == nil {
		return []Bucket{}
	}
	return o.Buckets
}

func (o *SearchTimeline) GetTotalEventCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalEventCount
}
