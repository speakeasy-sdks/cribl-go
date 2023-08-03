// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type ListCertificatesResponse struct {
	// a list of Certificate objects
	Certificates *shared.Certificates
	ContentType  string
	StatusCode   int
	RawResponse  *http.Response
}

func (o *ListCertificatesResponse) GetCertificates() *shared.Certificates {
	if o == nil {
		return nil
	}
	return o.Certificates
}

func (o *ListCertificatesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListCertificatesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListCertificatesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}