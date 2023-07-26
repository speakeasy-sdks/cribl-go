// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type DeleteParserIDRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteParserIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteParserIDResponse struct {
	ContentType string
	// a list of Parser objects
	ParserLibEntries *shared.ParserLibEntries
	StatusCode       int
	RawResponse      *http.Response
}

func (o *DeleteParserIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteParserIDResponse) GetParserLibEntries() *shared.ParserLibEntries {
	if o == nil {
		return nil
	}
	return o.ParserLibEntries
}

func (o *DeleteParserIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteParserIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}