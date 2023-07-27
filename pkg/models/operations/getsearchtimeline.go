// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetSearchTimelineRequest struct {
	// search job id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetSearchTimelineRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetSearchTimelineResponse struct {
	ContentType string
	// SearchTimeline object
	SearchTimeline *shared.SearchTimeline
	StatusCode     int
	RawResponse    *http.Response
}

func (o *GetSearchTimelineResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSearchTimelineResponse) GetSearchTimeline() *shared.SearchTimeline {
	if o == nil {
		return nil
	}
	return o.SearchTimeline
}

func (o *GetSearchTimelineResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSearchTimelineResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
