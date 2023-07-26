// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetCriblsSettingsResponse struct {
	ContentType string
	// a list of string objects
	GitSettings *shared.GitSettings
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetCriblsSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCriblsSettingsResponse) GetGitSettings() *shared.GitSettings {
	if o == nil {
		return nil
	}
	return o.GitSettings
}

func (o *GetCriblsSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCriblsSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}