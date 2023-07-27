// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetProcessRunningDetailRequest struct {
	// Unique ID
	Pid string `pathParam:"style=simple,explode=false,name=pid"`
}

func (o *GetProcessRunningDetailRequest) GetPid() string {
	if o == nil {
		return ""
	}
	return o.Pid
}

type GetProcessRunningDetailResponse struct {
	ContentType string
	// a list of Process objects
	Processes   *shared.Processes
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetProcessRunningDetailResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetProcessRunningDetailResponse) GetProcesses() *shared.Processes {
	if o == nil {
		return nil
	}
	return o.Processes
}

func (o *GetProcessRunningDetailResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetProcessRunningDetailResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
