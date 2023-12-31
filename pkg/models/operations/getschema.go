// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetSchemaRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetSchemaRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetSchemaResponse struct {
	ContentType string
	// a list of Schema objects
	SchemaLibEntry map[string]interface{}
	StatusCode     int
	RawResponse    *http.Response
}

func (o *GetSchemaResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSchemaResponse) GetSchemaLibEntry() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.SchemaLibEntry
}

func (o *GetSchemaResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSchemaResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
