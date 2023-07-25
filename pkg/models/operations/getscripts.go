// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetScriptsResponse struct {
	ContentType string
	// a list of Script objects
	ScriptLibEntries *shared.ScriptLibEntries
	StatusCode       int
	RawResponse      *http.Response
}

func (o *GetScriptsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetScriptsResponse) GetScriptLibEntries() *shared.ScriptLibEntries {
	if o == nil {
		return nil
	}
	return o.ScriptLibEntries
}

func (o *GetScriptsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetScriptsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
