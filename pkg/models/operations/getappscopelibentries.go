// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type GetAppscopeLibEntriesResponse struct {
	// a list of AppscopeLibEntry objects
	AppScopeLibEntries *shared.AppScopeLibEntries
	ContentType        string
	StatusCode         int
	RawResponse        *http.Response
}

func (o *GetAppscopeLibEntriesResponse) GetAppScopeLibEntries() *shared.AppScopeLibEntries {
	if o == nil {
		return nil
	}
	return o.AppScopeLibEntries
}

func (o *GetAppscopeLibEntriesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAppscopeLibEntriesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAppscopeLibEntriesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
