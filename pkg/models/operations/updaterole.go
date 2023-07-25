// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type UpdateRoleRequest struct {
	// Role object to be updated
	Role *shared.Role `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateRoleRequest) GetRole() *shared.Role {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *UpdateRoleRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateRoleResponse struct {
	ContentType string
	// a list of Role objects
	Roles       *shared.Roles
	StatusCode  int
	RawResponse *http.Response
}

func (o *UpdateRoleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRoleResponse) GetRoles() *shared.Roles {
	if o == nil {
		return nil
	}
	return o.Roles
}

func (o *UpdateRoleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRoleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
