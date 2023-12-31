// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetDatabaseConnectionConfigIDRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetDatabaseConnectionConfigIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetDatabaseConnectionConfigIDResponse struct {
	ContentType string
	// a list of DatabaseConnectionConfig objects
	DatabaseConnectionConfigs *shared.DatabaseConnectionConfigs
	StatusCode                int
	RawResponse               *http.Response
}

func (o *GetDatabaseConnectionConfigIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDatabaseConnectionConfigIDResponse) GetDatabaseConnectionConfigs() *shared.DatabaseConnectionConfigs {
	if o == nil {
		return nil
	}
	return o.DatabaseConnectionConfigs
}

func (o *GetDatabaseConnectionConfigIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDatabaseConnectionConfigIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
