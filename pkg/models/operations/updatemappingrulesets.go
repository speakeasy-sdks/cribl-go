// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type UpdateMappingRulesetsRequest struct {
	// MappingRuleset object to be updated
	MappingRuleset *shared.MappingRuleset `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateMappingRulesetsRequest) GetMappingRuleset() *shared.MappingRuleset {
	if o == nil {
		return nil
	}
	return o.MappingRuleset
}

func (o *UpdateMappingRulesetsRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateMappingRulesetsResponse struct {
	ContentType string
	// a list of MappingRuleset objects
	MappingRulesets *shared.MappingRulesets
	StatusCode      int
	RawResponse     *http.Response
}

func (o *UpdateMappingRulesetsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateMappingRulesetsResponse) GetMappingRulesets() *shared.MappingRulesets {
	if o == nil {
		return nil
	}
	return o.MappingRulesets
}

func (o *UpdateMappingRulesetsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateMappingRulesetsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
