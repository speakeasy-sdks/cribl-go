// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type DeleteMappingRulesetRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteMappingRulesetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteMappingRulesetResponse struct {
	ContentType string
	// a list of MappingRuleset objects
	MappingRulesets *shared.MappingRulesets
	StatusCode      int
	RawResponse     *http.Response
}

func (o *DeleteMappingRulesetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteMappingRulesetResponse) GetMappingRulesets() *shared.MappingRulesets {
	if o == nil {
		return nil
	}
	return o.MappingRulesets
}

func (o *DeleteMappingRulesetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteMappingRulesetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
