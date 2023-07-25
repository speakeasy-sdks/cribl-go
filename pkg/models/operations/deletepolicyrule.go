// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type DeletePolicyRuleRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeletePolicyRuleRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeletePolicyRuleResponse struct {
	ContentType string
	// a list of PolicyRule objects
	PolicyRules *shared.PolicyRules
	StatusCode  int
	RawResponse *http.Response
}

func (o *DeletePolicyRuleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeletePolicyRuleResponse) GetPolicyRules() *shared.PolicyRules {
	if o == nil {
		return nil
	}
	return o.PolicyRules
}

func (o *DeletePolicyRuleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeletePolicyRuleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
