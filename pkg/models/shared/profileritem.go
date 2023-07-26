// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ProfilerItem - New ProfilerItem object
type ProfilerItem struct {
	CreateTime *int64  `json:"createTime,omitempty"`
	ID         string  `json:"id"`
	Size       *int64  `json:"size,omitempty"`
	WorkerID   *string `json:"workerId,omitempty"`
}

func (o *ProfilerItem) GetCreateTime() *int64 {
	if o == nil {
		return nil
	}
	return o.CreateTime
}

func (o *ProfilerItem) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ProfilerItem) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *ProfilerItem) GetWorkerID() *string {
	if o == nil {
		return nil
	}
	return o.WorkerID
}
