// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ContainerdMount struct {
	Destination string `json:"destination"`
	Source      string `json:"source"`
}

func (o *ContainerdMount) GetDestination() string {
	if o == nil {
		return ""
	}
	return o.Destination
}

func (o *ContainerdMount) GetSource() string {
	if o == nil {
		return ""
	}
	return o.Source
}
