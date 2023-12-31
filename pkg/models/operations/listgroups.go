// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type ListGroupsRequest struct {
	// additional fields to add to results: git.commit, git.localChanges, git.log
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// filter to specific product: "stream" or "edge"
	Product *string `queryParam:"style=form,explode=true,name=product"`
}

func (o *ListGroupsRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListGroupsRequest) GetProduct() *string {
	if o == nil {
		return nil
	}
	return o.Product
}

type ListGroupsResponse struct {
	// a list of ConfigGroup objects
	ConfigGroups *shared.ConfigGroups
	ContentType  string
	StatusCode   int
	RawResponse  *http.Response
}

func (o *ListGroupsResponse) GetConfigGroups() *shared.ConfigGroups {
	if o == nil {
		return nil
	}
	return o.ConfigGroups
}

func (o *ListGroupsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListGroupsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListGroupsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
