// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AuthConfigType string

const (
	AuthConfigTypeSplunk AuthConfigType = "splunk"
	AuthConfigTypeLocal  AuthConfigType = "local"
	AuthConfigTypeLdap   AuthConfigType = "ldap"
	AuthConfigTypeOpenid AuthConfigType = "openid"
	AuthConfigTypeSaml   AuthConfigType = "saml"
)

func (e AuthConfigType) ToPointer() *AuthConfigType {
	return &e
}

func (e *AuthConfigType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "splunk":
		fallthrough
	case "local":
		fallthrough
	case "ldap":
		fallthrough
	case "openid":
		fallthrough
	case "saml":
		*e = AuthConfigType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthConfigType: %v", v)
	}
}

type AuthConfig struct {
	CanChangePassword *bool          `json:"canChangePassword,omitempty"`
	Fallback          bool           `json:"fallback"`
	FallbackBadLogin  bool           `json:"fallbackBadLogin"`
	FilterType        *string        `json:"filter_type,omitempty"`
	Host              string         `json:"host"`
	Port              int64          `json:"port"`
	Ssl               bool           `json:"ssl"`
	Type              AuthConfigType `json:"type"`
}

func (o *AuthConfig) GetCanChangePassword() *bool {
	if o == nil {
		return nil
	}
	return o.CanChangePassword
}

func (o *AuthConfig) GetFallback() bool {
	if o == nil {
		return false
	}
	return o.Fallback
}

func (o *AuthConfig) GetFallbackBadLogin() bool {
	if o == nil {
		return false
	}
	return o.FallbackBadLogin
}

func (o *AuthConfig) GetFilterType() *string {
	if o == nil {
		return nil
	}
	return o.FilterType
}

func (o *AuthConfig) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *AuthConfig) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *AuthConfig) GetSsl() bool {
	if o == nil {
		return false
	}
	return o.Ssl
}

func (o *AuthConfig) GetType() AuthConfigType {
	if o == nil {
		return AuthConfigType("")
	}
	return o.Type
}
