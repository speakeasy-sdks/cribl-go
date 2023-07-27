// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetWorkerEdgeNodesCountRequest struct {
	// Filter expression evaluated against nodes
	FilterExp *string `queryParam:"style=form,explode=true,name=filterExp"`
}

func (o *GetWorkerEdgeNodesCountRequest) GetFilterExp() *string {
	if o == nil {
		return nil
	}
	return o.FilterExp
}

type GetWorkerEdgeNodesCountResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// a list of number objects
	WorkerEdgeNodes *shared.WorkerEdgeNodes
}

func (o *GetWorkerEdgeNodesCountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWorkerEdgeNodesCountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWorkerEdgeNodesCountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetWorkerEdgeNodesCountResponse) GetWorkerEdgeNodes() *shared.WorkerEdgeNodes {
	if o == nil {
		return nil
	}
	return o.WorkerEdgeNodes
}
