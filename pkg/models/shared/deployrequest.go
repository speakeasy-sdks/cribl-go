// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DeployRequest - DeployRequest object
type DeployRequest struct {
	Version string `json:"version"`
}

func (o *DeployRequest) GetVersion() string {
	if o == nil {
		return ""
	}
	return o.Version
}
