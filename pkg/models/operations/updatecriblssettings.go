// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type UpdateCriblsSettingsResponse struct {
	ContentType string
	// a list of string objects
	GitSettings *shared.GitSettings
	StatusCode  int
	RawResponse *http.Response
}

func (o *UpdateCriblsSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCriblsSettingsResponse) GetGitSettings() *shared.GitSettings {
	if o == nil {
		return nil
	}
	return o.GitSettings
}

func (o *UpdateCriblsSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCriblsSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}