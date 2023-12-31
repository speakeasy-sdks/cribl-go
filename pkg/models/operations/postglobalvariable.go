// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type PostGlobalVariableResponse struct {
	ContentType string
	// a list of Global Variable objects
	GlobalVars  *shared.GlobalVars
	StatusCode  int
	RawResponse *http.Response
}

func (o *PostGlobalVariableResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostGlobalVariableResponse) GetGlobalVars() *shared.GlobalVars {
	if o == nil {
		return nil
	}
	return o.GlobalVars
}

func (o *PostGlobalVariableResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostGlobalVariableResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
