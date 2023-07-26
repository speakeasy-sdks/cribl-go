// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RedisCacheLimits struct {
	ClientTrackingMechanism *string `json:"clientTrackingMechanism,omitempty"`
	EnableServerAssist      *bool   `json:"enableServerAssist,omitempty"`
	ID                      string  `json:"id"`
	KeyTTLSecs              *int64  `json:"keyTTLSecs,omitempty"`
	MaxCacheSize            *int64  `json:"maxCacheSize,omitempty"`
	MaxNumKeys              *int64  `json:"maxNumKeys,omitempty"`
	ServicePeriodSecs       *int64  `json:"servicePeriodSecs,omitempty"`
}

func (o *RedisCacheLimits) GetClientTrackingMechanism() *string {
	if o == nil {
		return nil
	}
	return o.ClientTrackingMechanism
}

func (o *RedisCacheLimits) GetEnableServerAssist() *bool {
	if o == nil {
		return nil
	}
	return o.EnableServerAssist
}

func (o *RedisCacheLimits) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RedisCacheLimits) GetKeyTTLSecs() *int64 {
	if o == nil {
		return nil
	}
	return o.KeyTTLSecs
}

func (o *RedisCacheLimits) GetMaxCacheSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxCacheSize
}

func (o *RedisCacheLimits) GetMaxNumKeys() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxNumKeys
}

func (o *RedisCacheLimits) GetServicePeriodSecs() *int64 {
	if o == nil {
		return nil
	}
	return o.ServicePeriodSecs
}