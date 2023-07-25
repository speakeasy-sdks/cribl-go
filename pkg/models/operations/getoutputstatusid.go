// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetOutputStatusIDRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetOutputStatusIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetOutputStatusIDResponse struct {
	ContentType string
	// a list of OutputStatus objects
	OutputStatuses *shared.OutputStatuses
	StatusCode     int
	RawResponse    *http.Response
}

func (o *GetOutputStatusIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetOutputStatusIDResponse) GetOutputStatuses() *shared.OutputStatuses {
	if o == nil {
		return nil
	}
	return o.OutputStatuses
}

func (o *GetOutputStatusIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetOutputStatusIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
