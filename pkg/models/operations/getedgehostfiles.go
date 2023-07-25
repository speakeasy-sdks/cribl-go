// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetEdgeHostFilesResponse struct {
	ContentType string
	// a list of EdgeFileInspectResponse objects
	EdgeHostFiles *shared.EdgeHostFiles
	StatusCode    int
	RawResponse   *http.Response
}

func (o *GetEdgeHostFilesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEdgeHostFilesResponse) GetEdgeHostFiles() *shared.EdgeHostFiles {
	if o == nil {
		return nil
	}
	return o.EdgeHostFiles
}

func (o *GetEdgeHostFilesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEdgeHostFilesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
