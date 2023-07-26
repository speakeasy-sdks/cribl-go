// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetSearchLimitsResponse struct {
	ContentType string
	// a list of SearchSettings objects
	SearchSettingses *shared.SearchSettingses
	StatusCode       int
	RawResponse      *http.Response
}

func (o *GetSearchLimitsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSearchLimitsResponse) GetSearchSettingses() *shared.SearchSettingses {
	if o == nil {
		return nil
	}
	return o.SearchSettingses
}

func (o *GetSearchLimitsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSearchLimitsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}