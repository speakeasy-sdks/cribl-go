// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CommonServiceLimitConfigs struct {
	MemoryLimit string `json:"memoryLimit"`
}

func (o *CommonServiceLimitConfigs) GetMemoryLimit() string {
	if o == nil {
		return ""
	}
	return o.MemoryLimit
}
