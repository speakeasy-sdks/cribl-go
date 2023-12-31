// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type OutputSnmpHosts struct {
	// Destination host
	Host string `json:"host"`
	// Destination port, default is 162
	Port int64 `json:"port"`
}

func (o *OutputSnmpHosts) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *OutputSnmpHosts) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

type OutputSnmpType string

const (
	OutputSnmpTypeSnmp OutputSnmpType = "snmp"
)

func (e OutputSnmpType) ToPointer() *OutputSnmpType {
	return &e
}

func (e *OutputSnmpType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "snmp":
		*e = OutputSnmpType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSnmpType: %v", v)
	}
}

type OutputSnmp struct {
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// One or more SNMP destinations to forward traps to
	Hosts []OutputSnmpHosts `json:"hosts"`
	// Unique ID for this output
	ID *string `json:"id,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string       `json:"systemFields,omitempty"`
	Type         OutputSnmpType `json:"type"`
}

func (o *OutputSnmp) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputSnmp) GetHosts() []OutputSnmpHosts {
	if o == nil {
		return []OutputSnmpHosts{}
	}
	return o.Hosts
}

func (o *OutputSnmp) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputSnmp) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputSnmp) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputSnmp) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputSnmp) GetType() OutputSnmpType {
	if o == nil {
		return OutputSnmpType("")
	}
	return o.Type
}
