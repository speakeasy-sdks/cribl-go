// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type CreateDatasetProviderTypeResponse struct {
	ContentType string
	// a list of DatasetProviderType objects
	DatasetProviderType *shared.DatasetProviderType
	StatusCode          int
	RawResponse         *http.Response
}

func (o *CreateDatasetProviderTypeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateDatasetProviderTypeResponse) GetDatasetProviderType() *shared.DatasetProviderType {
	if o == nil {
		return nil
	}
	return o.DatasetProviderType
}

func (o *CreateDatasetProviderTypeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateDatasetProviderTypeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
