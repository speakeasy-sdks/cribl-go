// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type PostExecuteDistributedUpgradeRequest struct {
	// distributedUpgrade object
	DistributedUpgradeRequest *shared.DistributedUpgradeRequest `request:"mediaType=application/json"`
	// Group to upgrade
	Group string `pathParam:"style=simple,explode=false,name=group"`
}

func (o *PostExecuteDistributedUpgradeRequest) GetDistributedUpgradeRequest() *shared.DistributedUpgradeRequest {
	if o == nil {
		return nil
	}
	return o.DistributedUpgradeRequest
}

func (o *PostExecuteDistributedUpgradeRequest) GetGroup() string {
	if o == nil {
		return ""
	}
	return o.Group
}

type PostExecuteDistributedUpgradeResponse struct {
	ContentType string
	// a list of any objects
	CriblPackage *shared.CriblPackage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *PostExecuteDistributedUpgradeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostExecuteDistributedUpgradeResponse) GetCriblPackage() *shared.CriblPackage {
	if o == nil {
		return nil
	}
	return o.CriblPackage
}

func (o *PostExecuteDistributedUpgradeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostExecuteDistributedUpgradeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}