// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type OutputDevnullType string

const (
	OutputDevnullTypeDevnull OutputDevnullType = "devnull"
)

func (e OutputDevnullType) ToPointer() *OutputDevnullType {
	return &e
}

func (e *OutputDevnullType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "devnull":
		*e = OutputDevnullType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDevnullType: %v", v)
	}
}

type OutputDevnull struct {
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Unique ID for this output
	ID string `json:"id"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string          `json:"systemFields,omitempty"`
	Type         OutputDevnullType `json:"type"`
}

func (o *OutputDevnull) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputDevnull) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *OutputDevnull) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputDevnull) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputDevnull) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputDevnull) GetType() OutputDevnullType {
	if o == nil {
		return OutputDevnullType("")
	}
	return o.Type
}