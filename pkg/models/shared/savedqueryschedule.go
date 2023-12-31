// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SavedQuerySchedule struct {
	CronSchedule string `json:"cronSchedule"`
	Enabled      bool   `json:"enabled"`
	KeepLastN    int64  `json:"keepLastN"`
	Tz           string `json:"tz"`
}

func (o *SavedQuerySchedule) GetCronSchedule() string {
	if o == nil {
		return ""
	}
	return o.CronSchedule
}

func (o *SavedQuerySchedule) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *SavedQuerySchedule) GetKeepLastN() int64 {
	if o == nil {
		return 0
	}
	return o.KeepLastN
}

func (o *SavedQuerySchedule) GetTz() string {
	if o == nil {
		return ""
	}
	return o.Tz
}
