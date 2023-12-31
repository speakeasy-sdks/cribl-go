// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type ListBytesRequest struct {
	// The path to the file to sample
	Path string `queryParam:"style=form,explode=true,name=path"`
	// The number of bytes to return;   this value could be constrained by system limits.
	BytesRequested *int64 `queryParam:"style=form,explode=true,name=bytesRequested"`
}

func (o *ListBytesRequest) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *ListBytesRequest) GetBytesRequested() *int64 {
	if o == nil {
		return nil
	}
	return o.BytesRequested
}

type ListBytesResponse struct {
	ContentType string
	// a list of SampleFile objects
	SampleFiles *shared.SampleFiles
	StatusCode  int
	RawResponse *http.Response
}

func (o *ListBytesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListBytesResponse) GetSampleFiles() *shared.SampleFiles {
	if o == nil {
		return nil
	}
	return o.SampleFiles
}

func (o *ListBytesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListBytesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
