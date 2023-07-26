// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GrokFile - New GrokFile object
type GrokFile struct {
	Content string  `json:"content"`
	ID      string  `json:"id"`
	Size    int64   `json:"size"`
	Tags    *string `json:"tags,omitempty"`
}

func (o *GrokFile) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *GrokFile) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GrokFile) GetSize() int64 {
	if o == nil {
		return 0
	}
	return o.Size
}

func (o *GrokFile) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}