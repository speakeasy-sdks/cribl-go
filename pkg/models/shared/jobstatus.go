// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobStatus struct {
	Reason map[string]interface{} `json:"reason,omitempty"`
	State  State                  `json:"state"`
}

func (o *JobStatus) GetReason() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Reason
}

func (o *JobStatus) GetState() State {
	if o == nil {
		return State("")
	}
	return o.State
}
