// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetSavedJobsResponse struct {
	ContentType string
	// a list of SavedJob objects
	SavedJobs   *shared.SavedJobs
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetSavedJobsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSavedJobsResponse) GetSavedJobs() *shared.SavedJobs {
	if o == nil {
		return nil
	}
	return o.SavedJobs
}

func (o *GetSavedJobsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSavedJobsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
