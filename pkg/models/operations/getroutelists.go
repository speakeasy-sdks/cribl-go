// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetRouteListsResponse struct {
	ContentType string
	// a list of Routes objects
	Routes      *shared.Routes
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetRouteListsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRouteListsResponse) GetRoutes() *shared.Routes {
	if o == nil {
		return nil
	}
	return o.Routes
}

func (o *GetRouteListsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRouteListsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}