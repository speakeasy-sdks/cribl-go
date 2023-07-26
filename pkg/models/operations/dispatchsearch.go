// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type DispatchSearchRequest struct {
	// search job id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DispatchSearchRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DispatchSearchResponse struct {
	ContentType string
	// a list of any objects
	SearchID    *shared.SearchID
	StatusCode  int
	RawResponse *http.Response
}

func (o *DispatchSearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DispatchSearchResponse) GetSearchID() *shared.SearchID {
	if o == nil {
		return nil
	}
	return o.SearchID
}

func (o *DispatchSearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DispatchSearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}