// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetSearchDocResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// string object
	GetSearchDoc200ApplicationJSONString *string
}

func (o *GetSearchDocResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSearchDocResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSearchDocResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSearchDocResponse) GetGetSearchDoc200ApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.GetSearchDoc200ApplicationJSONString
}
