// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetTextualDiffRequest struct {
	// Commit hash (default is HEAD)
	Commit *string `queryParam:"style=form,explode=true,name=commit"`
	// Group ID
	Group *string `queryParam:"style=form,explode=true,name=group"`
}

func (o *GetTextualDiffRequest) GetCommit() *string {
	if o == nil {
		return nil
	}
	return o.Commit
}

func (o *GetTextualDiffRequest) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

type GetTextualDiffResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// a list of any objects
	TextualDiff *shared.TextualDiff
}

func (o *GetTextualDiffResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTextualDiffResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTextualDiffResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTextualDiffResponse) GetTextualDiff() *shared.TextualDiff {
	if o == nil {
		return nil
	}
	return o.TextualDiff
}
