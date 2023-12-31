// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type UpdateKeyMetadataEntityRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// KeyMetadataEntity object to be updated
	KeyMetadataEntity *shared.KeyMetadataEntity `request:"mediaType=application/json"`
}

func (o *UpdateKeyMetadataEntityRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateKeyMetadataEntityRequest) GetKeyMetadataEntity() *shared.KeyMetadataEntity {
	if o == nil {
		return nil
	}
	return o.KeyMetadataEntity
}

type UpdateKeyMetadataEntityResponse struct {
	ContentType string
	// a list of KeyMetadataEntity objects
	KeyMetadataEntities *shared.KeyMetadataEntities
	StatusCode          int
	RawResponse         *http.Response
}

func (o *UpdateKeyMetadataEntityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateKeyMetadataEntityResponse) GetKeyMetadataEntities() *shared.KeyMetadataEntities {
	if o == nil {
		return nil
	}
	return o.KeyMetadataEntities
}

func (o *UpdateKeyMetadataEntityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateKeyMetadataEntityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
