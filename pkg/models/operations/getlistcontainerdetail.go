// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetListContainerDetailResponse struct {
	// a list of Container objects
	Containers  *shared.Containers
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetListContainerDetailResponse) GetContainers() *shared.Containers {
	if o == nil {
		return nil
	}
	return o.Containers
}

func (o *GetListContainerDetailResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetListContainerDetailResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetListContainerDetailResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
