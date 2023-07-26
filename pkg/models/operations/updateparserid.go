// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type UpdateParserIDRequest struct {
	// Parser object to be updated
	RequestBody map[string]interface{} `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateParserIDRequest) GetRequestBody() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpdateParserIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateParserIDResponse struct {
	ContentType string
	// a list of Parser objects
	ParserLibEntries *shared.ParserLibEntries
	StatusCode       int
	RawResponse      *http.Response
}

func (o *UpdateParserIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateParserIDResponse) GetParserLibEntries() *shared.ParserLibEntries {
	if o == nil {
		return nil
	}
	return o.ParserLibEntries
}

func (o *UpdateParserIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateParserIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
