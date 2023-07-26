// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ChangelogState struct {
	LastViewedCurrent *string `json:"lastViewedCurrent,omitempty"`
	LastViewedUpgrade *string `json:"lastViewedUpgrade,omitempty"`
}

func (o *ChangelogState) GetLastViewedCurrent() *string {
	if o == nil {
		return nil
	}
	return o.LastViewedCurrent
}

func (o *ChangelogState) GetLastViewedUpgrade() *string {
	if o == nil {
		return nil
	}
	return o.LastViewedUpgrade
}
