// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type ListProcessesResponse struct {
	ContentType string
	// a list of ProcessEntry objects
	ProcessEntries *shared.ProcessEntries
	StatusCode     int
	RawResponse    *http.Response
}

func (o *ListProcessesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListProcessesResponse) GetProcessEntries() *shared.ProcessEntries {
	if o == nil {
		return nil
	}
	return o.ProcessEntries
}

func (o *ListProcessesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListProcessesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}