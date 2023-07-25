// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ServicesLimits struct {
	Connections   CommonServiceLimitConfigs `json:"connections"`
	ID            string                    `json:"id"`
	Metrics       CommonServiceLimitConfigs `json:"metrics"`
	Notifications CommonServiceLimitConfigs `json:"notifications"`
}

func (o *ServicesLimits) GetConnections() CommonServiceLimitConfigs {
	if o == nil {
		return CommonServiceLimitConfigs{}
	}
	return o.Connections
}

func (o *ServicesLimits) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ServicesLimits) GetMetrics() CommonServiceLimitConfigs {
	if o == nil {
		return CommonServiceLimitConfigs{}
	}
	return o.Metrics
}

func (o *ServicesLimits) GetNotifications() CommonServiceLimitConfigs {
	if o == nil {
		return CommonServiceLimitConfigs{}
	}
	return o.Notifications
}
