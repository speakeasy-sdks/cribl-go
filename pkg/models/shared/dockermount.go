// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DockerMount struct {
	Destination string `json:"Destination"`
	Source      string `json:"Source"`
}

func (o *DockerMount) GetDestination() string {
	if o == nil {
		return ""
	}
	return o.Destination
}

func (o *DockerMount) GetSource() string {
	if o == nil {
		return ""
	}
	return o.Source
}
