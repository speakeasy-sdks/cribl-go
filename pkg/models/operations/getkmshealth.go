// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetKMSHealthResponse struct {
	ContentType string
	// a list of IKMSHealth objects
	KMSHealth   *shared.KMSHealth
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetKMSHealthResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetKMSHealthResponse) GetKMSHealth() *shared.KMSHealth {
	if o == nil {
		return nil
	}
	return o.KMSHealth
}

func (o *GetKMSHealthResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetKMSHealthResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
