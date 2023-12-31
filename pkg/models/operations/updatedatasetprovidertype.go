// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type UpdateDatasetProviderTypeRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// DatasetProviderType object to be updated
	DatasetProviderType *shared.DatasetProviderType `request:"mediaType=application/json"`
}

func (o *UpdateDatasetProviderTypeRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateDatasetProviderTypeRequest) GetDatasetProviderType() *shared.DatasetProviderType {
	if o == nil {
		return nil
	}
	return o.DatasetProviderType
}

type UpdateDatasetProviderTypeResponse struct {
	ContentType string
	// a list of DatasetProviderType objects
	DatasetProviderType *shared.DatasetProviderType
	StatusCode          int
	RawResponse         *http.Response
}

func (o *UpdateDatasetProviderTypeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDatasetProviderTypeResponse) GetDatasetProviderType() *shared.DatasetProviderType {
	if o == nil {
		return nil
	}
	return o.DatasetProviderType
}

func (o *UpdateDatasetProviderTypeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDatasetProviderTypeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
