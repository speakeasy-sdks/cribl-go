// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type GetBulletinMessageRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetBulletinMessageRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetBulletinMessageResponse struct {
	// a list of BulletinMessage objects
	BulletinMessage *shared.BulletinMessage
	ContentType     string
	StatusCode      int
	RawResponse     *http.Response
}

func (o *GetBulletinMessageResponse) GetBulletinMessage() *shared.BulletinMessage {
	if o == nil {
		return nil
	}
	return o.BulletinMessage
}

func (o *GetBulletinMessageResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBulletinMessageResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBulletinMessageResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
