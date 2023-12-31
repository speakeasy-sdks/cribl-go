// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetRegexLibEntryIDRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetRegexLibEntryIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetRegexLibEntryIDResponse struct {
	ContentType string
	// a list of RegexLibEntry objects
	RegexLibEntries *shared.RegexLibEntries
	StatusCode      int
	RawResponse     *http.Response
}

func (o *GetRegexLibEntryIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRegexLibEntryIDResponse) GetRegexLibEntries() *shared.RegexLibEntries {
	if o == nil {
		return nil
	}
	return o.RegexLibEntries
}

func (o *GetRegexLibEntryIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRegexLibEntryIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
