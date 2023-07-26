// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type UpdateConfigGroupRequest struct {
	// ConfigGroup object to be updated
	ConfigGroup *shared.ConfigGroup `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateConfigGroupRequest) GetConfigGroup() *shared.ConfigGroup {
	if o == nil {
		return nil
	}
	return o.ConfigGroup
}

func (o *UpdateConfigGroupRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateConfigGroupResponse struct {
	// a list of ConfigGroup objects
	ConfigGroup *shared.ConfigGroup
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *UpdateConfigGroupResponse) GetConfigGroup() *shared.ConfigGroup {
	if o == nil {
		return nil
	}
	return o.ConfigGroup
}

func (o *UpdateConfigGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateConfigGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateConfigGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
