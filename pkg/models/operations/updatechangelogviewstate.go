// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type UpdateChangelogViewStateResponse struct {
	// a list of ChangelogState objects
	ChangelogStateses *shared.ChangelogStateses
	ContentType       string
	StatusCode        int
	RawResponse       *http.Response
}

func (o *UpdateChangelogViewStateResponse) GetChangelogStateses() *shared.ChangelogStateses {
	if o == nil {
		return nil
	}
	return o.ChangelogStateses
}

func (o *UpdateChangelogViewStateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateChangelogViewStateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateChangelogViewStateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}