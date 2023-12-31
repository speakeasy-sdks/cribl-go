// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type UpdateEventBreakerRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Event Breaker Ruleset object to be updated
	EventBreakerRuleset *shared.EventBreakerRuleset `request:"mediaType=application/json"`
}

func (o *UpdateEventBreakerRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateEventBreakerRequest) GetEventBreakerRuleset() *shared.EventBreakerRuleset {
	if o == nil {
		return nil
	}
	return o.EventBreakerRuleset
}

type UpdateEventBreakerResponse struct {
	ContentType string
	// a list of Event Breaker Ruleset objects
	EventBreakerRulesets *shared.EventBreakerRulesets
	StatusCode           int
	RawResponse          *http.Response
}

func (o *UpdateEventBreakerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateEventBreakerResponse) GetEventBreakerRulesets() *shared.EventBreakerRulesets {
	if o == nil {
		return nil
	}
	return o.EventBreakerRulesets
}

func (o *UpdateEventBreakerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateEventBreakerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
