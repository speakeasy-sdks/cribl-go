// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type DeleteDatasetProviderTypeRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteDatasetProviderTypeRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteDatasetProviderTypeResponse struct {
	ContentType string
	// a list of DatasetProviderType objects
	DatasetProviderType *shared.DatasetProviderType
	StatusCode          int
	RawResponse         *http.Response
}

func (o *DeleteDatasetProviderTypeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteDatasetProviderTypeResponse) GetDatasetProviderType() *shared.DatasetProviderType {
	if o == nil {
		return nil
	}
	return o.DatasetProviderType
}

func (o *DeleteDatasetProviderTypeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteDatasetProviderTypeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
