// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type UpdateGitSettingsResponse struct {
	ContentType string
	// a list of GitSettings objects
	GitSettings *shared.GitSettings
	StatusCode  int
	RawResponse *http.Response
}

func (o *UpdateGitSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateGitSettingsResponse) GetGitSettings() *shared.GitSettings {
	if o == nil {
		return nil
	}
	return o.GitSettings
}

func (o *UpdateGitSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateGitSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
