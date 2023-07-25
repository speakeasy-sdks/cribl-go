// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputRingCompression - Select data compression format. Optional.
type OutputRingCompression string

const (
	OutputRingCompressionNone OutputRingCompression = "none"
	OutputRingCompressionGzip OutputRingCompression = "gzip"
)

func (e OutputRingCompression) ToPointer() *OutputRingCompression {
	return &e
}

func (e *OutputRingCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputRingCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputRingCompression: %v", v)
	}
}

// OutputRingDataFormat - Format of the output data.
type OutputRingDataFormat string

const (
	OutputRingDataFormatJSON OutputRingDataFormat = "json"
	OutputRingDataFormatRaw  OutputRingDataFormat = "raw"
)

func (e OutputRingDataFormat) ToPointer() *OutputRingDataFormat {
	return &e
}

func (e *OutputRingDataFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "raw":
		*e = OutputRingDataFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputRingDataFormat: %v", v)
	}
}

// OutputRingBackpressureBehavior - Whether to block or drop events when all receivers are exerting backpressure.
type OutputRingBackpressureBehavior string

const (
	OutputRingBackpressureBehaviorBlock OutputRingBackpressureBehavior = "block"
	OutputRingBackpressureBehaviorDrop  OutputRingBackpressureBehavior = "drop"
)

func (e OutputRingBackpressureBehavior) ToPointer() *OutputRingBackpressureBehavior {
	return &e
}

func (e *OutputRingBackpressureBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputRingBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputRingBackpressureBehavior: %v", v)
	}
}

type OutputRingType string

const (
	OutputRingTypeRing OutputRingType = "ring"
)

func (e OutputRingType) ToPointer() *OutputRingType {
	return &e
}

func (e *OutputRingType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ring":
		*e = OutputRingType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputRingType: %v", v)
	}
}

type OutputRing struct {
	// Select data compression format. Optional.
	Compress *OutputRingCompression `json:"compress,omitempty"`
	// Path to use to write metrics. Defaults to $CRIBL_HOME/state/<id>
	DestPath *string `json:"destPath,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Format of the output data.
	Format *OutputRingDataFormat `json:"format,omitempty"`
	// Unique ID for this output
	ID string `json:"id"`
	// Maximum disk space allowed to be consumed (e.g., 420MB or 4GB). Once reached, older data will be deleted.
	MaxDataSize *string `json:"maxDataSize,omitempty"`
	// Maximum amount of time to retain data (e.g., 2h or 4d). Once reached, older data will be deleted.
	MaxDataTime *string `json:"maxDataTime,omitempty"`
	// Whether to block or drop events when all receivers are exerting backpressure.
	OnBackpressure *OutputRingBackpressureBehavior `json:"onBackpressure,omitempty"`
	// JS expression to define how files are partitioned and organized. If left blank, Cribl Stream will fallback on event.__partition.
	PartitionExpr *string `json:"partitionExpr,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string       `json:"systemFields,omitempty"`
	Type         OutputRingType `json:"type"`
}

func (o *OutputRing) GetCompress() *OutputRingCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *OutputRing) GetDestPath() *string {
	if o == nil {
		return nil
	}
	return o.DestPath
}

func (o *OutputRing) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputRing) GetFormat() *OutputRingDataFormat {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *OutputRing) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *OutputRing) GetMaxDataSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxDataSize
}

func (o *OutputRing) GetMaxDataTime() *string {
	if o == nil {
		return nil
	}
	return o.MaxDataTime
}

func (o *OutputRing) GetOnBackpressure() *OutputRingBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputRing) GetPartitionExpr() *string {
	if o == nil {
		return nil
	}
	return o.PartitionExpr
}

func (o *OutputRing) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputRing) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputRing) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputRing) GetType() OutputRingType {
	if o == nil {
		return OutputRingType("")
	}
	return o.Type
}
