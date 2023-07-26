// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetVersioningResponse struct {
	ContentType string
	// a list of GitInfo objects
	GitInfos    *shared.GitInfos
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetVersioningResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetVersioningResponse) GetGitInfos() *shared.GitInfos {
	if o == nil {
		return nil
	}
	return o.GitInfos
}

func (o *GetVersioningResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetVersioningResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}