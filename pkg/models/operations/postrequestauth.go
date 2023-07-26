// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

// PostRequestAuthRequestBody - Authentication request object
type PostRequestAuthRequestBody struct {
	RelayState *string `form:"name=RelayState"`
	// Authentication request object
	SAMLResponse *string `form:"name=SAMLResponse"`
}

func (o *PostRequestAuthRequestBody) GetRelayState() *string {
	if o == nil {
		return nil
	}
	return o.RelayState
}

func (o *PostRequestAuthRequestBody) GetSAMLResponse() *string {
	if o == nil {
		return nil
	}
	return o.SAMLResponse
}

type PostRequestAuthResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Authentication success
	Success *shared.Success
}

func (o *PostRequestAuthResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostRequestAuthResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostRequestAuthResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostRequestAuthResponse) GetSuccess() *shared.Success {
	if o == nil {
		return nil
	}
	return o.Success
}