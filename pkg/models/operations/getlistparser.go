// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetListParserResponse struct {
	ContentType string
	// a list of Parser objects
	ParserLibEntries *shared.ParserLibEntries
	StatusCode       int
	RawResponse      *http.Response
}

func (o *GetListParserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetListParserResponse) GetParserLibEntries() *shared.ParserLibEntries {
	if o == nil {
		return nil
	}
	return o.ParserLibEntries
}

func (o *GetListParserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetListParserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}