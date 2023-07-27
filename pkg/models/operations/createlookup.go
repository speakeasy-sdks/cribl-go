// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/cribl-go/pkg/models/shared"
	"net/http"
)

type CreateLookupRequestBody2 struct {
	// File content.
	Content *string `json:"content,omitempty"`
	// Brief description of this lookup. Optional.
	Description *string `json:"description,omitempty"`
	// Filename with the lookup table. Required.
	ID string `json:"id"`
	// File size. Optional.
	Size *int64 `json:"size,omitempty"`
	// One or more tags related to this lookup. Optional.
	Tags *string `json:"tags,omitempty"`
}

func (o *CreateLookupRequestBody2) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *CreateLookupRequestBody2) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateLookupRequestBody2) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateLookupRequestBody2) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *CreateLookupRequestBody2) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// CreateLookupRequestBody1FileInfo - Uploaded file information
type CreateLookupRequestBody1FileInfo struct {
	Filename string `json:"filename"`
}

func (o *CreateLookupRequestBody1FileInfo) GetFilename() string {
	if o == nil {
		return ""
	}
	return o.Filename
}

type CreateLookupRequestBody1 struct {
	// Brief description of this lookup. Optional.
	Description *string `json:"description,omitempty"`
	// Uploaded file information
	FileInfo *CreateLookupRequestBody1FileInfo `json:"fileInfo,omitempty"`
	// Filename with the lookup table. Required.
	ID string `json:"id"`
	// File size. Optional.
	Size *int64 `json:"size,omitempty"`
	// One or more tags related to this lookup. Optional.
	Tags *string `json:"tags,omitempty"`
}

func (o *CreateLookupRequestBody1) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateLookupRequestBody1) GetFileInfo() *CreateLookupRequestBody1FileInfo {
	if o == nil {
		return nil
	}
	return o.FileInfo
}

func (o *CreateLookupRequestBody1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateLookupRequestBody1) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *CreateLookupRequestBody1) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

type CreateLookupResponse struct {
	ContentType string
	// a list of LookupFile objects
	LookupFiles *shared.LookupFiles
	StatusCode  int
	RawResponse *http.Response
}

func (o *CreateLookupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateLookupResponse) GetLookupFiles() *shared.LookupFiles {
	if o == nil {
		return nil
	}
	return o.LookupFiles
}

func (o *CreateLookupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateLookupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
