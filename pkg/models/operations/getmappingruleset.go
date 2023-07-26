// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetMappingRulesetRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMappingRulesetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetMappingRulesetResponse struct {
	ContentType string
	// a list of MappingRuleset objects
	MappingRulesets *shared.MappingRulesets
	StatusCode      int
	RawResponse     *http.Response
}

func (o *GetMappingRulesetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMappingRulesetResponse) GetMappingRulesets() *shared.MappingRulesets {
	if o == nil {
		return nil
	}
	return o.MappingRulesets
}

func (o *GetMappingRulesetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMappingRulesetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}