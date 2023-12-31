// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ListJobResultsRequest struct {
	// Job instance id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ListJobResultsRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ListJobResultsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Get results for a discover job by instance id
	ListJobResults200ApplicationXNdjsonBinaryString []byte
}

func (o *ListJobResultsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListJobResultsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListJobResultsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListJobResultsResponse) GetListJobResults200ApplicationXNdjsonBinaryString() []byte {
	if o == nil {
		return nil
	}
	return o.ListJobResults200ApplicationXNdjsonBinaryString
}
