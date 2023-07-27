// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetGroupBundleRequest struct {
	// Group ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// DeployRequest object
	DeployRequest *shared.DeployRequest `request:"mediaType=application/json"`
}

func (o *GetGroupBundleRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetGroupBundleRequest) GetDeployRequest() *shared.DeployRequest {
	if o == nil {
		return nil
	}
	return o.DeployRequest
}

type GetGroupBundleResponse struct {
	ContentType string
	// a list of string objects
	GroupBundle *shared.GroupBundle
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetGroupBundleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGroupBundleResponse) GetGroupBundle() *shared.GroupBundle {
	if o == nil {
		return nil
	}
	return o.GroupBundle
}

func (o *GetGroupBundleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGroupBundleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
