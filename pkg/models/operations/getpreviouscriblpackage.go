// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetPreviousCriblPackageRequest struct {
	// Name of the file to be downloaded
	File string `pathParam:"style=simple,explode=false,name=file"`
}

func (o *GetPreviousCriblPackageRequest) GetFile() string {
	if o == nil {
		return ""
	}
	return o.File
}

type GetPreviousCriblPackageResponse struct {
	ContentType string
	// a list of any objects
	CriblPackage *shared.CriblPackage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetPreviousCriblPackageResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPreviousCriblPackageResponse) GetCriblPackage() *shared.CriblPackage {
	if o == nil {
		return nil
	}
	return o.CriblPackage
}

func (o *GetPreviousCriblPackageResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPreviousCriblPackageResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
