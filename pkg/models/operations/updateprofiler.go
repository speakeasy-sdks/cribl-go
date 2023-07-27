// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type UpdateProfilerRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// ProfilerItem object to be updated
	ProfilerItem *shared.ProfilerItem `request:"mediaType=application/json"`
}

func (o *UpdateProfilerRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateProfilerRequest) GetProfilerItem() *shared.ProfilerItem {
	if o == nil {
		return nil
	}
	return o.ProfilerItem
}

type UpdateProfilerResponse struct {
	ContentType string
	// a list of ProfilerItem objects
	ProfilerItem *shared.ProfilerItem
	StatusCode   int
	RawResponse  *http.Response
}

func (o *UpdateProfilerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateProfilerResponse) GetProfilerItem() *shared.ProfilerItem {
	if o == nil {
		return nil
	}
	return o.ProfilerItem
}

func (o *UpdateProfilerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateProfilerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
