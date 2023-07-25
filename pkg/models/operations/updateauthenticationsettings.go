// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type UpdateAuthenticationSettingsResponse struct {
	// a list of AuthConfig objects
	AuthConfigs *shared.AuthConfigs
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *UpdateAuthenticationSettingsResponse) GetAuthConfigs() *shared.AuthConfigs {
	if o == nil {
		return nil
	}
	return o.AuthConfigs
}

func (o *UpdateAuthenticationSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAuthenticationSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAuthenticationSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
