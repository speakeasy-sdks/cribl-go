// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type ResumeJobRequest struct {
	// Job instance id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ResumeJobRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ResumeJobResponse struct {
	ContentType string
	// a list of any objects
	JobResume   *shared.JobResume
	StatusCode  int
	RawResponse *http.Response
}

func (o *ResumeJobResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ResumeJobResponse) GetJobResume() *shared.JobResume {
	if o == nil {
		return nil
	}
	return o.JobResume
}

func (o *ResumeJobResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ResumeJobResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
