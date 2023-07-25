// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SavedQuery - New SavedQuery object
type SavedQuery struct {
	Description *string             `json:"description,omitempty"`
	Earliest    *string             `json:"earliest,omitempty"`
	ID          string              `json:"id"`
	Latest      *string             `json:"latest,omitempty"`
	Name        string              `json:"name"`
	Query       string              `json:"query"`
	SampleRate  *int64              `json:"sampleRate,omitempty"`
	Schedule    *SavedQuerySchedule `json:"schedule,omitempty"`
	User        *string             `json:"user,omitempty"`
}

func (o *SavedQuery) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SavedQuery) GetEarliest() *string {
	if o == nil {
		return nil
	}
	return o.Earliest
}

func (o *SavedQuery) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SavedQuery) GetLatest() *string {
	if o == nil {
		return nil
	}
	return o.Latest
}

func (o *SavedQuery) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SavedQuery) GetQuery() string {
	if o == nil {
		return ""
	}
	return o.Query
}

func (o *SavedQuery) GetSampleRate() *int64 {
	if o == nil {
		return nil
	}
	return o.SampleRate
}

func (o *SavedQuery) GetSchedule() *SavedQuerySchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}

func (o *SavedQuery) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}
