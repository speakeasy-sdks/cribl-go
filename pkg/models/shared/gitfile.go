// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GitFile struct {
	Children []GitFile `json:"children,omitempty"`
	Name     string    `json:"name"`
	State    *string   `json:"state,omitempty"`
}

func (o *GitFile) GetChildren() []GitFile {
	if o == nil {
		return nil
	}
	return o.Children
}

func (o *GitFile) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GitFile) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}
