// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type UpdateRouteObjectRequest struct {
	// There is only one route entity and it should be accessed with id: default.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Routes object
	Routes *shared.Routes `request:"mediaType=application/json"`
}

func (o *UpdateRouteObjectRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateRouteObjectRequest) GetRoutes() *shared.Routes {
	if o == nil {
		return nil
	}
	return o.Routes
}

type UpdateRouteObjectResponse struct {
	ContentType string
	// a list of Routes objects
	Routes      *shared.Routes
	StatusCode  int
	RawResponse *http.Response
}

func (o *UpdateRouteObjectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRouteObjectResponse) GetRoutes() *shared.Routes {
	if o == nil {
		return nil
	}
	return o.Routes
}

func (o *UpdateRouteObjectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRouteObjectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
