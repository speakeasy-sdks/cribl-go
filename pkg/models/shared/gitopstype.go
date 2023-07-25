// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type GitOpsType string

const (
	GitOpsTypeNone GitOpsType = "none"
	GitOpsTypePull GitOpsType = "pull"
	GitOpsTypePush GitOpsType = "push"
)

func (e GitOpsType) ToPointer() *GitOpsType {
	return &e
}

func (e *GitOpsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "pull":
		fallthrough
	case "push":
		*e = GitOpsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GitOpsType: %v", v)
	}
}
