// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetNotificationTargetRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetNotificationTargetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetNotificationTargetResponse struct {
	ContentType string
	// a list of NotificationTarget objects
	NotificationTargets *shared.NotificationTargets
	StatusCode          int
	RawResponse         *http.Response
}

func (o *GetNotificationTargetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetNotificationTargetResponse) GetNotificationTargets() *shared.NotificationTargets {
	if o == nil {
		return nil
	}
	return o.NotificationTargets
}

func (o *GetNotificationTargetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetNotificationTargetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}