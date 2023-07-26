// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetProcessesResponse struct {
	ContentType string
	// a list of ProcessEntry objects
	ProcessEntries *shared.ProcessEntries
	StatusCode     int
	RawResponse    *http.Response
}

func (o *GetProcessesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetProcessesResponse) GetProcessEntries() *shared.ProcessEntries {
	if o == nil {
		return nil
	}
	return o.ProcessEntries
}

func (o *GetProcessesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetProcessesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}