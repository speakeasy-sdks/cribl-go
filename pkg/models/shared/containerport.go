// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ContainerPort struct {
	PrivatePort int64 `json:"privatePort"`
	PublicPort  int64 `json:"publicPort"`
}

func (o *ContainerPort) GetPrivatePort() int64 {
	if o == nil {
		return 0
	}
	return o.PrivatePort
}

func (o *ContainerPort) GetPublicPort() int64 {
	if o == nil {
		return 0
	}
	return o.PublicPort
}
