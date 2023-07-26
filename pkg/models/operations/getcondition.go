// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetConditionRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetConditionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetConditionResponse struct {
	// a list of Condition objects
	Conditions  *shared.Conditions
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetConditionResponse) GetConditions() *shared.Conditions {
	if o == nil {
		return nil
	}
	return o.Conditions
}

func (o *GetConditionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConditionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConditionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}