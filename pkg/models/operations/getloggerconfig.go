// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetLoggerConfigRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetLoggerConfigRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetLoggerConfigResponse struct {
	ContentType string
	// a list of LoggerConfig objects
	LoggerConfig *shared.LoggerConfig
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetLoggerConfigResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLoggerConfigResponse) GetLoggerConfig() *shared.LoggerConfig {
	if o == nil {
		return nil
	}
	return o.LoggerConfig
}

func (o *GetLoggerConfigResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLoggerConfigResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
