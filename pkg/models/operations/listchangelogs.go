// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type ListChangelogsResponse struct {
	// a list of ChangelogState objects
	ChangelogStates *shared.ChangelogStates
	ContentType     string
	StatusCode      int
	RawResponse     *http.Response
}

func (o *ListChangelogsResponse) GetChangelogStates() *shared.ChangelogStates {
	if o == nil {
		return nil
	}
	return o.ChangelogStates
}

func (o *ListChangelogsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListChangelogsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListChangelogsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}