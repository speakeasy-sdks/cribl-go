// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ListSearchJobMetricsRequest struct {
	// search job id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ListSearchJobMetricsRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ListSearchJobMetricsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// string object
	ListSearchJobMetrics200ApplicationJSONString *string
}

func (o *ListSearchJobMetricsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSearchJobMetricsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSearchJobMetricsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListSearchJobMetricsResponse) GetListSearchJobMetrics200ApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.ListSearchJobMetrics200ApplicationJSONString
}