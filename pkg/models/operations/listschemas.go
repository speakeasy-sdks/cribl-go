// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type ListSchemasResponse struct {
	ContentType string
	// a list of Schema objects
	SchemaLibEntries *shared.SchemaLibEntries
	StatusCode       int
	RawResponse      *http.Response
}

func (o *ListSchemasResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSchemasResponse) GetSchemaLibEntries() *shared.SchemaLibEntries {
	if o == nil {
		return nil
	}
	return o.SchemaLibEntries
}

func (o *ListSchemasResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSchemasResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
