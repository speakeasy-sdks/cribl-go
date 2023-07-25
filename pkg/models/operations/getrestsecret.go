// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetRestSecretRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetRestSecretRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetRestSecretResponse struct {
	ContentType string
	// a list of RestSecret objects
	RestSecret  *shared.RestSecret
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetRestSecretResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRestSecretResponse) GetRestSecret() *shared.RestSecret {
	if o == nil {
		return nil
	}
	return o.RestSecret
}

func (o *GetRestSecretResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRestSecretResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
