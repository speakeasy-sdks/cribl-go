// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type PostMasterNodePackageResponse struct {
	ContentType string
	// a list of string objects
	Cribl       *shared.Cribl
	StatusCode  int
	RawResponse *http.Response
}

func (o *PostMasterNodePackageResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostMasterNodePackageResponse) GetCribl() *shared.Cribl {
	if o == nil {
		return nil
	}
	return o.Cribl
}

func (o *PostMasterNodePackageResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostMasterNodePackageResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
