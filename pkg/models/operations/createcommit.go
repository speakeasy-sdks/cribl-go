// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type CreateCommitResponse struct {
	ContentType string
	// a list of GitCommitSummary objects
	GitCommit   *shared.GitCommit
	StatusCode  int
	RawResponse *http.Response
}

func (o *CreateCommitResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCommitResponse) GetGitCommit() *shared.GitCommit {
	if o == nil {
		return nil
	}
	return o.GitCommit
}

func (o *CreateCommitResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCommitResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
