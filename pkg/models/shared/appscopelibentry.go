// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AppscopeLibEntry - New AppscopeLibEntry object
type AppscopeLibEntry struct {
	Config      AppscopeConfigWithCustom `json:"config"`
	Description string                   `json:"description"`
	ID          string                   `json:"id"`
	Lib         CriblLib                 `json:"lib"`
	Tags        *string                  `json:"tags,omitempty"`
}

func (o *AppscopeLibEntry) GetConfig() AppscopeConfigWithCustom {
	if o == nil {
		return AppscopeConfigWithCustom{}
	}
	return o.Config
}

func (o *AppscopeLibEntry) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *AppscopeLibEntry) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AppscopeLibEntry) GetLib() CriblLib {
	if o == nil {
		return CriblLib("")
	}
	return o.Lib
}

func (o *AppscopeLibEntry) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}