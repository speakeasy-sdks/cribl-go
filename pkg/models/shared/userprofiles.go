// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UserProfiles - a list of UserProfile objects
type UserProfiles struct {
	// number of items present in the items array
	Count *int64        `json:"count,omitempty"`
	Items []UserProfile `json:"items,omitempty"`
}

func (o *UserProfiles) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *UserProfiles) GetItems() []UserProfile {
	if o == nil {
		return nil
	}
	return o.Items
}
