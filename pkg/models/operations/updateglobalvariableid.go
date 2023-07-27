// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type UpdateGlobalVariableIDRequest struct {
	// Global Variable object to be updated
	GlobalVar *shared.GlobalVar `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateGlobalVariableIDRequest) GetGlobalVar() *shared.GlobalVar {
	if o == nil {
		return nil
	}
	return o.GlobalVar
}

func (o *UpdateGlobalVariableIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateGlobalVariableIDResponse struct {
	ContentType string
	// a list of Global Variable objects
	GlobalVars  *shared.GlobalVars
	StatusCode  int
	RawResponse *http.Response
}

func (o *UpdateGlobalVariableIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateGlobalVariableIDResponse) GetGlobalVars() *shared.GlobalVars {
	if o == nil {
		return nil
	}
	return o.GlobalVars
}

func (o *UpdateGlobalVariableIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateGlobalVariableIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
