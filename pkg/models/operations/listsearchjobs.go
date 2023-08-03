// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type ListSearchJobsResponse struct {
	ContentType string
	// a list of SearchJob objects
	SearchJobs  *shared.SearchJobs
	StatusCode  int
	RawResponse *http.Response
}

func (o *ListSearchJobsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSearchJobsResponse) GetSearchJobs() *shared.SearchJobs {
	if o == nil {
		return nil
	}
	return o.SearchJobs
}

func (o *ListSearchJobsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSearchJobsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}