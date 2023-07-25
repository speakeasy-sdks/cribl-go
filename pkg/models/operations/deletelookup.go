// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteLookupRequest struct {
	// Unique ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteLookupRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteLookup200ApplicationJSON2 struct {
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

func (o *DeleteLookup200ApplicationJSON2) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *DeleteLookup200ApplicationJSON2) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *DeleteLookup200ApplicationJSON2) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteLookup200ApplicationJSON2) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *DeleteLookup200ApplicationJSON2) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// DeleteLookup200ApplicationJSON1FileInfo - Uploaded file information
type DeleteLookup200ApplicationJSON1FileInfo struct {
	Filename string `json:"filename"`
}

func (o *DeleteLookup200ApplicationJSON1FileInfo) GetFilename() string {
	if o == nil {
		return ""
	}
	return o.Filename
}

type DeleteLookup200ApplicationJSON1 struct {
	// Brief description of this lookup. Optional.
	Description *string `json:"description,omitempty"`
	// Uploaded file information
	FileInfo *DeleteLookup200ApplicationJSON1FileInfo `json:"fileInfo,omitempty"`
	// Filename with the lookup table. Required.
	ID string `json:"id"`
	// File size. Optional.
	Size *int64 `json:"size,omitempty"`
	// One or more tags related to this lookup. Optional.
	Tags *string `json:"tags,omitempty"`
}

func (o *DeleteLookup200ApplicationJSON1) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *DeleteLookup200ApplicationJSON1) GetFileInfo() *DeleteLookup200ApplicationJSON1FileInfo {
	if o == nil {
		return nil
	}
	return o.FileInfo
}

func (o *DeleteLookup200ApplicationJSON1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteLookup200ApplicationJSON1) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *DeleteLookup200ApplicationJSON1) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

type DeleteLookupResponse struct {
	ContentType string
	// a list of LookupFile objects
	LookupFile  interface{}
	StatusCode  int
	RawResponse *http.Response
}

func (o *DeleteLookupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteLookupResponse) GetLookupFile() interface{} {
	if o == nil {
		return nil
	}
	return o.LookupFile
}

func (o *DeleteLookupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteLookupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
