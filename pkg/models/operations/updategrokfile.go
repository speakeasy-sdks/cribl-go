// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type UpdateGrokFileRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// GrokFile object to be updated
	GrokFile *shared.GrokFile `request:"mediaType=application/json"`
}

func (o *UpdateGrokFileRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateGrokFileRequest) GetGrokFile() *shared.GrokFile {
	if o == nil {
		return nil
	}
	return o.GrokFile
}

type UpdateGrokFileResponse struct {
	ContentType string
	// a list of GrokFile objects
	GrokFile    *shared.GrokFile
	StatusCode  int
	RawResponse *http.Response
}

func (o *UpdateGrokFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateGrokFileResponse) GetGrokFile() *shared.GrokFile {
	if o == nil {
		return nil
	}
	return o.GrokFile
}

func (o *UpdateGrokFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateGrokFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
