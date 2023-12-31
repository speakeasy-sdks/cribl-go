// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type DeleteDataSampleIDRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteDataSampleIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteDataSampleIDResponse struct {
	ContentType string
	// a list of DataSample objects
	DataSamples *shared.DataSamples
	StatusCode  int
	RawResponse *http.Response
}

func (o *DeleteDataSampleIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteDataSampleIDResponse) GetDataSamples() *shared.DataSamples {
	if o == nil {
		return nil
	}
	return o.DataSamples
}

func (o *DeleteDataSampleIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteDataSampleIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
