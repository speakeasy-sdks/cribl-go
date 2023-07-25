// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetProfilersResponse struct {
	ContentType string
	// a list of ProfilerItem objects
	ProfilerItems *shared.ProfilerItems
	StatusCode    int
	RawResponse   *http.Response
}

func (o *GetProfilersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetProfilersResponse) GetProfilerItems() *shared.ProfilerItems {
	if o == nil {
		return nil
	}
	return o.ProfilerItems
}

func (o *GetProfilersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetProfilersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
