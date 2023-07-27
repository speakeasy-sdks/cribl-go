// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type DeleteRoleRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteRoleRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteRoleResponse struct {
	ContentType string
	// a list of Role objects
	Roles       *shared.Roles
	StatusCode  int
	RawResponse *http.Response
}

func (o *DeleteRoleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteRoleResponse) GetRoles() *shared.Roles {
	if o == nil {
		return nil
	}
	return o.Roles
}

func (o *DeleteRoleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteRoleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
