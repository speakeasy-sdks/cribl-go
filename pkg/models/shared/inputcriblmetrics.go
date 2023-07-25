// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type InputCriblmetricsConnections struct {
	// Select a Destination.
	Output string `json:"output"`
	// Select Pipeline or Pack. Optional.
	Pipeline *string `json:"pipeline,omitempty"`
}

func (o *InputCriblmetricsConnections) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

func (o *InputCriblmetricsConnections) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

type InputCriblmetricsMetadata struct {
	// Field name
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputCriblmetricsMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputCriblmetricsMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// InputCriblmetricsPqCompression - Codec to use to compress the persisted data.
type InputCriblmetricsPqCompression string

const (
	InputCriblmetricsPqCompressionNone InputCriblmetricsPqCompression = "none"
	InputCriblmetricsPqCompressionGzip InputCriblmetricsPqCompression = "gzip"
)

func (e InputCriblmetricsPqCompression) ToPointer() *InputCriblmetricsPqCompression {
	return &e
}

func (e *InputCriblmetricsPqCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputCriblmetricsPqCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputCriblmetricsPqCompression: %v", v)
	}
}

// InputCriblmetricsPqMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputCriblmetricsPqMode string

const (
	InputCriblmetricsPqModeSmart  InputCriblmetricsPqMode = "smart"
	InputCriblmetricsPqModeAlways InputCriblmetricsPqMode = "always"
)

func (e InputCriblmetricsPqMode) ToPointer() *InputCriblmetricsPqMode {
	return &e
}

func (e *InputCriblmetricsPqMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputCriblmetricsPqMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputCriblmetricsPqMode: %v", v)
	}
}

type InputCriblmetricsPq struct {
	// The number of events to send downstream before committing that Stream has read them.
	CommitFrequency *int64 `json:"commitFrequency,omitempty"`
	// Codec to use to compress the persisted data.
	Compress *InputCriblmetricsPqCompression `json:"compress,omitempty"`
	// The maximum number of events to hold in memory before writing the events to disk.
	MaxBufferSize *int64 `json:"maxBufferSize,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	MaxFileSize *string `json:"maxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `json:"maxSize,omitempty"`
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputCriblmetricsPqMode `json:"mode,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>.
	Path *string `json:"path,omitempty"`
}

func (o *InputCriblmetricsPq) GetCommitFrequency() *int64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputCriblmetricsPq) GetCompress() *InputCriblmetricsPqCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *InputCriblmetricsPq) GetMaxBufferSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputCriblmetricsPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputCriblmetricsPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputCriblmetricsPq) GetMode() *InputCriblmetricsPqMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputCriblmetricsPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

type InputCriblmetricsType string

const (
	InputCriblmetricsTypeCriblmetrics InputCriblmetricsType = "criblmetrics"
)

func (e InputCriblmetricsType) ToPointer() *InputCriblmetricsType {
	return &e
}

func (e *InputCriblmetricsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "criblmetrics":
		*e = InputCriblmetricsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputCriblmetricsType: %v", v)
	}
}

type InputCriblmetrics struct {
	// Direct connections to Destinations, optionally via a Pipeline or a Pack.
	Connections []InputCriblmetricsConnections `json:"connections,omitempty"`
	// Enable/disable this input
	Disabled *bool `json:"disabled,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Include granular metrics.  Disabling this will drop the following metrics events: `cribl.logstream.host.(in_bytes,in_events,out_bytes,out_events)`, `cribl.logstream.index.(in_bytes,in_events,out_bytes,out_events)`, `cribl.logstream.source.(in_bytes,in_events,out_bytes,out_events)`, `cribl.logstream.sourcetype.(in_bytes,in_events,out_bytes,out_events)`.
	FullFidelity *bool `json:"fullFidelity,omitempty"`
	// Unique ID for this input
	ID string `json:"id"`
	// Fields to add to events from this input.
	Metadata []InputCriblmetricsMetadata `json:"metadata,omitempty"`
	// Pipeline to process data from this Source before sending it through the Routes.
	Pipeline *string              `json:"pipeline,omitempty"`
	Pq       *InputCriblmetricsPq `json:"pq,omitempty"`
	// For details on Persistent Queues, see: [https://docs.cribl.io/stream/persistent-queues](https://docs.cribl.io/stream/persistent-queues)
	PqEnabled *bool `json:"pqEnabled,omitempty"`
	// A prefix that is applied to the metrics provided by Cribl Stream
	Prefix *string `json:"prefix,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool `json:"sendToRoutes,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string              `json:"streamtags,omitempty"`
	Type       InputCriblmetricsType `json:"type"`
}

func (o *InputCriblmetrics) GetConnections() []InputCriblmetricsConnections {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputCriblmetrics) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputCriblmetrics) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputCriblmetrics) GetFullFidelity() *bool {
	if o == nil {
		return nil
	}
	return o.FullFidelity
}

func (o *InputCriblmetrics) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *InputCriblmetrics) GetMetadata() []InputCriblmetricsMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputCriblmetrics) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputCriblmetrics) GetPq() *InputCriblmetricsPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputCriblmetrics) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputCriblmetrics) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *InputCriblmetrics) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputCriblmetrics) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputCriblmetrics) GetType() InputCriblmetricsType {
	if o == nil {
		return InputCriblmetricsType("")
	}
	return o.Type
}
