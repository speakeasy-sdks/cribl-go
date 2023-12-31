// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type UpdateRegexLibEntryRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// RegexLibEntry object to be updated
	RegexLibEntry *shared.RegexLibEntry `request:"mediaType=application/json"`
}

func (o *UpdateRegexLibEntryRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateRegexLibEntryRequest) GetRegexLibEntry() *shared.RegexLibEntry {
	if o == nil {
		return nil
	}
	return o.RegexLibEntry
}

type UpdateRegexLibEntryResponse struct {
	ContentType string
	// a list of RegexLibEntry objects
	RegexLibEntries *shared.RegexLibEntries
	StatusCode      int
	RawResponse     *http.Response
}

func (o *UpdateRegexLibEntryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRegexLibEntryResponse) GetRegexLibEntries() *shared.RegexLibEntries {
	if o == nil {
		return nil
	}
	return o.RegexLibEntries
}

func (o *UpdateRegexLibEntryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRegexLibEntryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
