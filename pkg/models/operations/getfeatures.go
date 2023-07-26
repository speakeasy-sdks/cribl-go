// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetFeaturesResponse struct {
	ContentType string
	// a list of FeaturesEntry objects
	FeaturesEntries *shared.FeaturesEntries
	StatusCode      int
	RawResponse     *http.Response
}

func (o *GetFeaturesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFeaturesResponse) GetFeaturesEntries() *shared.FeaturesEntries {
	if o == nil {
		return nil
	}
	return o.FeaturesEntries
}

func (o *GetFeaturesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFeaturesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
