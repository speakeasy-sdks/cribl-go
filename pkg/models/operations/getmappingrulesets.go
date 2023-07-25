// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetMappingRulesetsResponse struct {
	ContentType string
	// a list of MappingRuleset objects
	MappingRulesets *shared.MappingRulesets
	StatusCode      int
	RawResponse     *http.Response
}

func (o *GetMappingRulesetsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMappingRulesetsResponse) GetMappingRulesets() *shared.MappingRulesets {
	if o == nil {
		return nil
	}
	return o.MappingRulesets
}

func (o *GetMappingRulesetsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMappingRulesetsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
