// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GethSearchJobMetricsRequest struct {
	// search job id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GethSearchJobMetricsRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GethSearchJobMetricsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// string object
	GethSearchJobMetrics200ApplicationJSONString *string
}

func (o *GethSearchJobMetricsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GethSearchJobMetricsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GethSearchJobMetricsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GethSearchJobMetricsResponse) GetGethSearchJobMetrics200ApplicationJSONString() *string {
	if o == nil {
		return nil
	}
	return o.GethSearchJobMetrics200ApplicationJSONString
}
