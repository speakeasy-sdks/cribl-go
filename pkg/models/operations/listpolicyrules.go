// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type ListPolicyRulesResponse struct {
	ContentType string
	// a list of PolicyRule objects
	PolicyRules *shared.PolicyRules
	StatusCode  int
	RawResponse *http.Response
}

func (o *ListPolicyRulesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPolicyRulesResponse) GetPolicyRules() *shared.PolicyRules {
	if o == nil {
		return nil
	}
	return o.PolicyRules
}

func (o *ListPolicyRulesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPolicyRulesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
