// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetSchemaIDRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetSchemaIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetSchemaIDResponse struct {
	ContentType string
	// a list of Schema objects
	SchemaLibEntries *shared.SchemaLibEntries
	StatusCode       int
	RawResponse      *http.Response
}

func (o *GetSchemaIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSchemaIDResponse) GetSchemaLibEntries() *shared.SchemaLibEntries {
	if o == nil {
		return nil
	}
	return o.SchemaLibEntries
}

func (o *GetSchemaIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSchemaIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
