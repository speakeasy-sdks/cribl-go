// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type PostTestDatabaseConnectionResponse struct {
	ContentType string
	// a list of DatabaseConnectionTestResult objects
	DatabaseConnectionTestResults *shared.DatabaseConnectionTestResults
	StatusCode                    int
	RawResponse                   *http.Response
}

func (o *PostTestDatabaseConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostTestDatabaseConnectionResponse) GetDatabaseConnectionTestResults() *shared.DatabaseConnectionTestResults {
	if o == nil {
		return nil
	}
	return o.DatabaseConnectionTestResults
}

func (o *PostTestDatabaseConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostTestDatabaseConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
