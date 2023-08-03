// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type ListAuthorizationsResponse struct {
	// a list of AuthPolicyEntry objects
	AuthPolicyEntries *shared.AuthPolicyEntries
	ContentType       string
	StatusCode        int
	RawResponse       *http.Response
}

func (o *ListAuthorizationsResponse) GetAuthPolicyEntries() *shared.AuthPolicyEntries {
	if o == nil {
		return nil
	}
	return o.AuthPolicyEntries
}

func (o *ListAuthorizationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAuthorizationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAuthorizationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}