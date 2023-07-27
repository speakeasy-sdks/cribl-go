// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetJobInfosRequest struct {
	// Filter by collector id, e.g. "collectorId=Prometheus-in"
	CollectorID *string `queryParam:"style=form,explode=true,name=collectorId"`
	// Filter by job id, e.g. "id=1608713335.3&id=1608713326.1"
	ID *string `queryParam:"style=form,explode=true,name=id"`
	// Maximum number of items to return
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Pagination offset
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
	// Filter by job run type, one of "adhoc", "scheduled", "system"
	RunType *string `queryParam:"style=form,explode=true,name=runType"`
	// Filter by current job state, e.g. "running"
	State *string `queryParam:"style=form,explode=true,name=state"`
}

func (o *GetJobInfosRequest) GetCollectorID() *string {
	if o == nil {
		return nil
	}
	return o.CollectorID
}

func (o *GetJobInfosRequest) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetJobInfosRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetJobInfosRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetJobInfosRequest) GetRunType() *string {
	if o == nil {
		return nil
	}
	return o.RunType
}

func (o *GetJobInfosRequest) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

type GetJobInfosResponse struct {
	ContentType string
	// a list of JobInfo objects
	JobInfos    *shared.JobInfos
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetJobInfosResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetJobInfosResponse) GetJobInfos() *shared.JobInfos {
	if o == nil {
		return nil
	}
	return o.JobInfos
}

func (o *GetJobInfosResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetJobInfosResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
