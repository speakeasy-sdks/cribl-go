// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetCriblMetadataResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Service Provider metadata
	GetCriblMetadata200TextXMLString *string
}

func (o *GetCriblMetadataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCriblMetadataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCriblMetadataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCriblMetadataResponse) GetGetCriblMetadata200TextXMLString() *string {
	if o == nil {
		return nil
	}
	return o.GetCriblMetadata200TextXMLString
}
