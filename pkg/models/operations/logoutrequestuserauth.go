// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type LogoutRequestUserAuthResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Logout success
	Success *shared.Success
}

func (o *LogoutRequestUserAuthResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LogoutRequestUserAuthResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LogoutRequestUserAuthResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *LogoutRequestUserAuthResponse) GetSuccess() *shared.Success {
	if o == nil {
		return nil
	}
	return o.Success
}