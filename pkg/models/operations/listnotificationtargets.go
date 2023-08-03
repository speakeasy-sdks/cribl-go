// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type ListNotificationTargetsResponse struct {
	ContentType string
	// a list of NotificationTarget objects
	NotificationTargets *shared.NotificationTargets
	StatusCode          int
	RawResponse         *http.Response
}

func (o *ListNotificationTargetsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListNotificationTargetsResponse) GetNotificationTargets() *shared.NotificationTargets {
	if o == nil {
		return nil
	}
	return o.NotificationTargets
}

func (o *ListNotificationTargetsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListNotificationTargetsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}