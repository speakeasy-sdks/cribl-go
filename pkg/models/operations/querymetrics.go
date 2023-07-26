// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type QueryMetricsRequest struct {
	// earliest time to query against
	Earliest *int64 `queryParam:"style=form,explode=true,name=earliest"`
	// a js expression to apply against the metrics returned (can filter dimensions)
	FilterExpr *string `queryParam:"style=form,explode=true,name=filterExpr"`
	// latest time to query against
	Latest *int64 `queryParam:"style=form,explode=true,name=latest"`
	// can be a regex or an array of metric names
	MetricNameFilter *string `queryParam:"style=form,explode=true,name=metricNameFilter"`
	// buckets in the past to include in the query results
	NumBuckets *int64 `queryParam:"style=form,explode=true,name=numBuckets"`
}

func (o *QueryMetricsRequest) GetEarliest() *int64 {
	if o == nil {
		return nil
	}
	return o.Earliest
}

func (o *QueryMetricsRequest) GetFilterExpr() *string {
	if o == nil {
		return nil
	}
	return o.FilterExpr
}

func (o *QueryMetricsRequest) GetLatest() *int64 {
	if o == nil {
		return nil
	}
	return o.Latest
}

func (o *QueryMetricsRequest) GetMetricNameFilter() *string {
	if o == nil {
		return nil
	}
	return o.MetricNameFilter
}

func (o *QueryMetricsRequest) GetNumBuckets() *int64 {
	if o == nil {
		return nil
	}
	return o.NumBuckets
}

type QueryMetricsResponse struct {
	ContentType string
	// a list of MetricNameInfo objects
	MetricsInfo *shared.MetricsInfo
	StatusCode  int
	RawResponse *http.Response
}

func (o *QueryMetricsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *QueryMetricsResponse) GetMetricsInfo() *shared.MetricsInfo {
	if o == nil {
		return nil
	}
	return o.MetricsInfo
}

func (o *QueryMetricsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *QueryMetricsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}