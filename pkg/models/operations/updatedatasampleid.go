// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type UpdateDataSampleIDRequest struct {
	// DataSample object to be updated
	RequestBody map[string]interface{} `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateDataSampleIDRequest) GetRequestBody() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpdateDataSampleIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateDataSampleIDResponse struct {
	ContentType string
	// a list of DataSample objects
	DataSamples *shared.DataSamples
	StatusCode  int
	RawResponse *http.Response
}

func (o *UpdateDataSampleIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDataSampleIDResponse) GetDataSamples() *shared.DataSamples {
	if o == nil {
		return nil
	}
	return o.DataSamples
}

func (o *UpdateDataSampleIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDataSampleIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
