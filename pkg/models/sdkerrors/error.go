// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

// Error - Unauthorized
type Error struct {
	RawResponse *http.Response `json:"-"`
	// Error message
	Message *string `json:"message,omitempty"`
}

var _ error = &Error{}

func (e *Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
