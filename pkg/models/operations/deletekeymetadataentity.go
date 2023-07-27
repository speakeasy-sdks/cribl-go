// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type DeleteKeyMetadataEntityRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteKeyMetadataEntityRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteKeyMetadataEntityResponse struct {
	ContentType string
	// a list of KeyMetadataEntity objects
	KeyMetadataEntities *shared.KeyMetadataEntities
	StatusCode          int
	RawResponse         *http.Response
}

func (o *DeleteKeyMetadataEntityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteKeyMetadataEntityResponse) GetKeyMetadataEntities() *shared.KeyMetadataEntities {
	if o == nil {
		return nil
	}
	return o.KeyMetadataEntities
}

func (o *DeleteKeyMetadataEntityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteKeyMetadataEntityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
