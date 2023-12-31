// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetFeatureRequest struct {
	// Feature id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetFeatureRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetFeatureResponse struct {
	ContentType string
	// a list of FeaturesEntry objects
	FeaturesEntry *shared.FeaturesEntry
	StatusCode    int
	RawResponse   *http.Response
}

func (o *GetFeatureResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFeatureResponse) GetFeaturesEntry() *shared.FeaturesEntry {
	if o == nil {
		return nil
	}
	return o.FeaturesEntry
}

func (o *GetFeatureResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFeatureResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
