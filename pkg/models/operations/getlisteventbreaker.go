// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetListEventBreakerResponse struct {
	ContentType string
	// a list of Event Breaker Ruleset objects
	EventBreakerRulesets *shared.EventBreakerRulesets
	StatusCode           int
	RawResponse          *http.Response
}

func (o *GetListEventBreakerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetListEventBreakerResponse) GetEventBreakerRulesets() *shared.EventBreakerRulesets {
	if o == nil {
		return nil
	}
	return o.EventBreakerRulesets
}

func (o *GetListEventBreakerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetListEventBreakerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
