// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type ListSearchLimitsResponse struct {
	ContentType string
	// a list of SearchSettings objects
	SearchSettingses *shared.SearchSettingses
	StatusCode       int
	RawResponse      *http.Response
}

func (o *ListSearchLimitsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSearchLimitsResponse) GetSearchSettingses() *shared.SearchSettingses {
	if o == nil {
		return nil
	}
	return o.SearchSettingses
}

func (o *ListSearchLimitsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSearchLimitsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
