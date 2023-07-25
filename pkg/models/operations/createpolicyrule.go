// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type CreatePolicyRuleResponse struct {
	ContentType string
	// a list of PolicyRule objects
	PolicyRules *shared.PolicyRules
	StatusCode  int
	RawResponse *http.Response
}

func (o *CreatePolicyRuleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePolicyRuleResponse) GetPolicyRules() *shared.PolicyRules {
	if o == nil {
		return nil
	}
	return o.PolicyRules
}

func (o *CreatePolicyRuleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePolicyRuleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
