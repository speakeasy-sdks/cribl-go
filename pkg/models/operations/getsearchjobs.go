// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetSearchJobsResponse struct {
	ContentType string
	// a list of SearchJob objects
	SearchJobs  *shared.SearchJobs
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetSearchJobsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSearchJobsResponse) GetSearchJobs() *shared.SearchJobs {
	if o == nil {
		return nil
	}
	return o.SearchJobs
}

func (o *GetSearchJobsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSearchJobsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
