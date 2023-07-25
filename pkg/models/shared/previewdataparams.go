// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PreviewDataParamsMode string

const (
	PreviewDataParamsModePipe         PreviewDataParamsMode = "pipe"
	PreviewDataParamsModeRoute        PreviewDataParamsMode = "route"
	PreviewDataParamsModeRouteAndSend PreviewDataParamsMode = "routeAndSend"
)

func (e PreviewDataParamsMode) ToPointer() *PreviewDataParamsMode {
	return &e
}

func (e *PreviewDataParamsMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pipe":
		fallthrough
	case "route":
		fallthrough
	case "routeAndSend":
		*e = PreviewDataParamsMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PreviewDataParamsMode: %v", v)
	}
}

// PreviewDataParams - PreviewDataParams object
type PreviewDataParams struct {
	CPUProfile       *bool                    `json:"cpuProfile,omitempty"`
	Dropped          *bool                    `json:"dropped,omitempty"`
	Events           []map[string]interface{} `json:"events,omitempty"`
	InputID          *string                  `json:"inputId,omitempty"`
	Level            *int64                   `json:"level,omitempty"`
	Memory           *int64                   `json:"memory,omitempty"`
	Mode             PreviewDataParamsMode    `json:"mode"`
	PipelineID       string                   `json:"pipelineId"`
	SampleID         string                   `json:"sampleId"`
	SamplePipelineID *string                  `json:"samplePipelineId,omitempty"`
	Timeout          *int64                   `json:"timeout,omitempty"`
}

func (o *PreviewDataParams) GetCPUProfile() *bool {
	if o == nil {
		return nil
	}
	return o.CPUProfile
}

func (o *PreviewDataParams) GetDropped() *bool {
	if o == nil {
		return nil
	}
	return o.Dropped
}

func (o *PreviewDataParams) GetEvents() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Events
}

func (o *PreviewDataParams) GetInputID() *string {
	if o == nil {
		return nil
	}
	return o.InputID
}

func (o *PreviewDataParams) GetLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.Level
}

func (o *PreviewDataParams) GetMemory() *int64 {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *PreviewDataParams) GetMode() PreviewDataParamsMode {
	if o == nil {
		return PreviewDataParamsMode("")
	}
	return o.Mode
}

func (o *PreviewDataParams) GetPipelineID() string {
	if o == nil {
		return ""
	}
	return o.PipelineID
}

func (o *PreviewDataParams) GetSampleID() string {
	if o == nil {
		return ""
	}
	return o.SampleID
}

func (o *PreviewDataParams) GetSamplePipelineID() *string {
	if o == nil {
		return nil
	}
	return o.SamplePipelineID
}

func (o *PreviewDataParams) GetTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}
