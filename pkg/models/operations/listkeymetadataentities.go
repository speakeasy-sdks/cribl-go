// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type ListKeyMetadataEntitiesResponse struct {
	ContentType string
	// a list of KeyMetadataEntity objects
	KeyMetadataEntities *shared.KeyMetadataEntities
	StatusCode          int
	RawResponse         *http.Response
}

func (o *ListKeyMetadataEntitiesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListKeyMetadataEntitiesResponse) GetKeyMetadataEntities() *shared.KeyMetadataEntities {
	if o == nil {
		return nil
	}
	return o.KeyMetadataEntities
}

func (o *ListKeyMetadataEntitiesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListKeyMetadataEntitiesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
