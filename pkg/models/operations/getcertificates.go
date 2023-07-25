// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetCertificatesResponse struct {
	// a list of Certificate objects
	Certificates *shared.Certificates
	ContentType  string
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetCertificatesResponse) GetCertificates() *shared.Certificates {
	if o == nil {
		return nil
	}
	return o.Certificates
}

func (o *GetCertificatesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCertificatesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCertificatesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
