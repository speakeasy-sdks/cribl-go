// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type PostDataSampleResponse struct {
	ContentType string
	// a list of DataSample objects
	DataSamples *shared.DataSamples
	StatusCode  int
	RawResponse *http.Response
}

func (o *PostDataSampleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostDataSampleResponse) GetDataSamples() *shared.DataSamples {
	if o == nil {
		return nil
	}
	return o.DataSamples
}

func (o *PostDataSampleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostDataSampleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
