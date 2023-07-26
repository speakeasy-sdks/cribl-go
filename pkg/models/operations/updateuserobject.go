// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"cribl/pkg/models/shared"
	"net/http"
)

type UpdateUserObjectRequest struct {
	// User object
	User *shared.User `request:"mediaType=application/json"`
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateUserObjectRequest) GetUser() *shared.User {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *UpdateUserObjectRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateUserObjectResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// a list of User objects
	Users *shared.Users
}

func (o *UpdateUserObjectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateUserObjectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateUserObjectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateUserObjectResponse) GetUsers() *shared.Users {
	if o == nil {
		return nil
	}
	return o.Users
}