// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type DeleteGrokFileRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteGrokFileRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteGrokFileResponse struct {
	ContentType string
	// a list of GrokFile objects
	GrokFile    *shared.GrokFile
	StatusCode  int
	RawResponse *http.Response
}

func (o *DeleteGrokFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteGrokFileResponse) GetGrokFile() *shared.GrokFile {
	if o == nil {
		return nil
	}
	return o.GrokFile
}

func (o *DeleteGrokFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteGrokFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
