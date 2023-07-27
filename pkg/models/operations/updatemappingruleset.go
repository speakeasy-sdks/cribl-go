// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type UpdateMappingRulesetRequest struct {
	// MappingRuleset object to be updated
	MappingRuleset *shared.MappingRuleset `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateMappingRulesetRequest) GetMappingRuleset() *shared.MappingRuleset {
	if o == nil {
		return nil
	}
	return o.MappingRuleset
}

func (o *UpdateMappingRulesetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateMappingRulesetResponse struct {
	ContentType string
	// a list of MappingRuleset objects
	MappingRulesets *shared.MappingRulesets
	StatusCode      int
	RawResponse     *http.Response
}

func (o *UpdateMappingRulesetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateMappingRulesetResponse) GetMappingRulesets() *shared.MappingRulesets {
	if o == nil {
		return nil
	}
	return o.MappingRulesets
}

func (o *UpdateMappingRulesetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateMappingRulesetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
