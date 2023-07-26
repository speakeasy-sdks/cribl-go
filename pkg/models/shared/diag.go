// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Diag struct {
	ID      string `json:"id"`
	ModTime *int64 `json:"modTime,omitempty"`
	Path    string `json:"path"`
	Size    *int64 `json:"size,omitempty"`
}

func (o *Diag) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Diag) GetModTime() *int64 {
	if o == nil {
		return nil
	}
	return o.ModTime
}

func (o *Diag) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *Diag) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}