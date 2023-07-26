// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetCountFileRequest struct {
	// Group ID
	Group *string `queryParam:"style=form,explode=true,name=group"`
}

func (o *GetCountFileRequest) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

type GetCountFileResponse struct {
	ContentType string
	// a list of any objects
	CountFile   *shared.CountFile
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetCountFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCountFileResponse) GetCountFile() *shared.CountFile {
	if o == nil {
		return nil
	}
	return o.CountFile
}

func (o *GetCountFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCountFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}