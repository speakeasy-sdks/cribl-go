// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetConfigGroupRequest struct {
	// query string additional fields to add to results: git.commit, git.localChanges, git.log
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetConfigGroupRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetConfigGroupRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetConfigGroupResponse struct {
	// a list of ConfigGroup objects
	ConfigGroup *shared.ConfigGroup
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetConfigGroupResponse) GetConfigGroup() *shared.ConfigGroup {
	if o == nil {
		return nil
	}
	return o.ConfigGroup
}

func (o *GetConfigGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConfigGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConfigGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
