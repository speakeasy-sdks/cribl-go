// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DockerGraphDriver struct {
	Data *DockerGraphDriverData `json:"Data,omitempty"`
}

func (o *DockerGraphDriver) GetData() *DockerGraphDriverData {
	if o == nil {
		return nil
	}
	return o.Data
}
