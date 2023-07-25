// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type DeleteRegexLibEntryRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteRegexLibEntryRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteRegexLibEntryResponse struct {
	ContentType string
	// a list of RegexLibEntry objects
	RegexLibEntries *shared.RegexLibEntries
	StatusCode      int
	RawResponse     *http.Response
}

func (o *DeleteRegexLibEntryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteRegexLibEntryResponse) GetRegexLibEntries() *shared.RegexLibEntries {
	if o == nil {
		return nil
	}
	return o.RegexLibEntries
}

func (o *DeleteRegexLibEntryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteRegexLibEntryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
