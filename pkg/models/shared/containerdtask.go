// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ContainerdTask struct {
	Process map[string]interface{} `json:"process,omitempty"`
}

func (o *ContainerdTask) GetProcess() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Process
}
