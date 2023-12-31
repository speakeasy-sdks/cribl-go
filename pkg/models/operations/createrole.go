// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type CreateRoleResponse struct {
	ContentType string
	// a list of Role objects
	Roles       *shared.Roles
	StatusCode  int
	RawResponse *http.Response
}

func (o *CreateRoleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateRoleResponse) GetRoles() *shared.Roles {
	if o == nil {
		return nil
	}
	return o.Roles
}

func (o *CreateRoleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateRoleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
