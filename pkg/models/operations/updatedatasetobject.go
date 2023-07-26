// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UpdateDatasetObjectRequest struct {
	// Dataset object to be updated
	RequestBody interface{} `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateDatasetObjectRequest) GetRequestBody() interface{} {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpdateDatasetObjectRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateDatasetObjectResponse struct {
	ContentType string
	// a list of Dataset objects
	Dataset     interface{}
	StatusCode  int
	RawResponse *http.Response
}

func (o *UpdateDatasetObjectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDatasetObjectResponse) GetDataset() interface{} {
	if o == nil {
		return nil
	}
	return o.Dataset
}

func (o *UpdateDatasetObjectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDatasetObjectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}