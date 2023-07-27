// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetWorkingTreeRequest struct {
	// Group ID
	Group *string `queryParam:"style=form,explode=true,name=group"`
}

func (o *GetWorkingTreeRequest) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

type GetWorkingTreeResponse struct {
	ContentType string
	// a list of GitStatusResult objects
	GitStatusResults *shared.GitStatusResults
	StatusCode       int
	RawResponse      *http.Response
}

func (o *GetWorkingTreeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWorkingTreeResponse) GetGitStatusResults() *shared.GitStatusResults {
	if o == nil {
		return nil
	}
	return o.GitStatusResults
}

func (o *GetWorkingTreeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWorkingTreeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
