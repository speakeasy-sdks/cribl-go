// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetRouteListIDRequest struct {
	// There is only one route entity and it should be accessed with id: default.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetRouteListIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetRouteListIDResponse struct {
	ContentType string
	// a list of Routes objects
	Routes      *shared.Routes
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetRouteListIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRouteListIDResponse) GetRoutes() *shared.Routes {
	if o == nil {
		return nil
	}
	return o.Routes
}

func (o *GetRouteListIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRouteListIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}