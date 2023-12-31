// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SearchJobType string

const (
	SearchJobTypeStandard        SearchJobType = "standard"
	SearchJobTypeDatatypePreview SearchJobType = "datatypePreview"
	SearchJobTypeScheduled       SearchJobType = "scheduled"
)

func (e SearchJobType) ToPointer() *SearchJobType {
	return &e
}

func (e *SearchJobType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "standard":
		fallthrough
	case "datatypePreview":
		fallthrough
	case "scheduled":
		*e = SearchJobType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SearchJobType: %v", v)
	}
}
