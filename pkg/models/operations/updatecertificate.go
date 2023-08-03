// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type UpdateCertificateRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Certificate object to be updated
	Certificate *shared.Certificate `request:"mediaType=application/json"`
}

func (o *UpdateCertificateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateCertificateRequest) GetCertificate() *shared.Certificate {
	if o == nil {
		return nil
	}
	return o.Certificate
}

type UpdateCertificateResponse struct {
	// a list of Certificate objects
	Certificate *shared.Certificate
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *UpdateCertificateResponse) GetCertificate() *shared.Certificate {
	if o == nil {
		return nil
	}
	return o.Certificate
}

func (o *UpdateCertificateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCertificateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCertificateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
