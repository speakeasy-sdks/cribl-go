// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputWavefrontAuthenticationMethod - Enter a token directly, or provide a secret referencing a token
type OutputWavefrontAuthenticationMethod string

const (
	OutputWavefrontAuthenticationMethodSecret OutputWavefrontAuthenticationMethod = "secret"
	OutputWavefrontAuthenticationMethodManual OutputWavefrontAuthenticationMethod = "manual"
)

func (e OutputWavefrontAuthenticationMethod) ToPointer() *OutputWavefrontAuthenticationMethod {
	return &e
}

func (e *OutputWavefrontAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		fallthrough
	case "manual":
		*e = OutputWavefrontAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWavefrontAuthenticationMethod: %v", v)
	}
}

type OutputWavefrontExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

func (o *OutputWavefrontExtraHTTPHeaders) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *OutputWavefrontExtraHTTPHeaders) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// OutputWavefrontFailedRequestLoggingMode - Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
type OutputWavefrontFailedRequestLoggingMode string

const (
	OutputWavefrontFailedRequestLoggingModePayload           OutputWavefrontFailedRequestLoggingMode = "payload"
	OutputWavefrontFailedRequestLoggingModePayloadAndHeaders OutputWavefrontFailedRequestLoggingMode = "payloadAndHeaders"
	OutputWavefrontFailedRequestLoggingModeNone              OutputWavefrontFailedRequestLoggingMode = "none"
)

func (e OutputWavefrontFailedRequestLoggingMode) ToPointer() *OutputWavefrontFailedRequestLoggingMode {
	return &e
}

func (e *OutputWavefrontFailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "payload":
		fallthrough
	case "payloadAndHeaders":
		fallthrough
	case "none":
		*e = OutputWavefrontFailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWavefrontFailedRequestLoggingMode: %v", v)
	}
}

// OutputWavefrontBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputWavefrontBackpressureBehavior string

const (
	OutputWavefrontBackpressureBehaviorQueue OutputWavefrontBackpressureBehavior = "queue"
	OutputWavefrontBackpressureBehaviorDrop  OutputWavefrontBackpressureBehavior = "drop"
	OutputWavefrontBackpressureBehaviorBlock OutputWavefrontBackpressureBehavior = "block"
)

func (e OutputWavefrontBackpressureBehavior) ToPointer() *OutputWavefrontBackpressureBehavior {
	return &e
}

func (e *OutputWavefrontBackpressureBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "queue":
		fallthrough
	case "drop":
		fallthrough
	case "block":
		*e = OutputWavefrontBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWavefrontBackpressureBehavior: %v", v)
	}
}

// OutputWavefrontCompression - Codec to use to compress the persisted data.
type OutputWavefrontCompression string

const (
	OutputWavefrontCompressionNone OutputWavefrontCompression = "none"
	OutputWavefrontCompressionGzip OutputWavefrontCompression = "gzip"
)

func (e OutputWavefrontCompression) ToPointer() *OutputWavefrontCompression {
	return &e
}

func (e *OutputWavefrontCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputWavefrontCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWavefrontCompression: %v", v)
	}
}

type OutputWavefrontPqControls struct {
}

// OutputWavefrontQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputWavefrontQueueFullBehavior string

const (
	OutputWavefrontQueueFullBehaviorBlock OutputWavefrontQueueFullBehavior = "block"
	OutputWavefrontQueueFullBehaviorDrop  OutputWavefrontQueueFullBehavior = "drop"
)

func (e OutputWavefrontQueueFullBehavior) ToPointer() *OutputWavefrontQueueFullBehavior {
	return &e
}

func (e *OutputWavefrontQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputWavefrontQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWavefrontQueueFullBehavior: %v", v)
	}
}

type OutputWavefrontType string

const (
	OutputWavefrontTypeWavefront OutputWavefrontType = "wavefront"
)

func (e OutputWavefrontType) ToPointer() *OutputWavefrontType {
	return &e
}

func (e *OutputWavefrontType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "wavefront":
		*e = OutputWavefrontType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWavefrontType: %v", v)
	}
}

type OutputWavefront struct {
	// Enter a token directly, or provide a secret referencing a token
	AuthType *OutputWavefrontAuthenticationMethod `json:"authType,omitempty"`
	// Whether to compress the payload body before sending.
	Compress *bool `json:"compress,omitempty"`
	// Maximum number of ongoing requests before blocking.
	Concurrency *int64 `json:"concurrency,omitempty"`
	// WaveFront domain name, e.g. "longboard"
	Domain string `json:"domain"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Headers to add to all events.
	ExtraHTTPHeaders []OutputWavefrontExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
	FailedRequestLoggingMode *OutputWavefrontFailedRequestLoggingMode `json:"failedRequestLoggingMode,omitempty"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Max body size.
	FlushPeriodSec *int64 `json:"flushPeriodSec,omitempty"`
	// Unique ID for this output
	ID *string `json:"id,omitempty"`
	// Max number of events to include in the request body. Default is 0 (unlimited).
	MaxPayloadEvents *int64 `json:"maxPayloadEvents,omitempty"`
	// Maximum size, in KB, of the request body.
	MaxPayloadSizeKB *int64 `json:"maxPayloadSizeKB,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputWavefrontBackpressureBehavior `json:"onBackpressure,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputWavefrontCompression `json:"pqCompress,omitempty"`
	PqControls *OutputWavefrontPqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputWavefrontQueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool `json:"pqStrictOrdering,omitempty"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another trusted CA (e.g., the system's CA). Defaults to Yes.
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// List of headers that are safe to log in plain text.
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string `json:"systemFields,omitempty"`
	// Select (or create) a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// Amount of time, in seconds, to wait for a request to complete before aborting it.
	TimeoutSec *int64 `json:"timeoutSec,omitempty"`
	// WaveFront API authentication token (see [here](https://docs.wavefront.com/wavefront_api.html#generating-an-api-token))
	Token *string             `json:"token,omitempty"`
	Type  OutputWavefrontType `json:"type"`
	// Enable to use round-robin DNS lookup. When a DNS server returns multiple addresses, this will cause Stream to cycle through them in the order returned.
	UseRoundRobinDNS *bool `json:"useRoundRobinDns,omitempty"`
}

func (o *OutputWavefront) GetAuthType() *OutputWavefrontAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *OutputWavefront) GetCompress() *bool {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *OutputWavefront) GetConcurrency() *int64 {
	if o == nil {
		return nil
	}
	return o.Concurrency
}

func (o *OutputWavefront) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *OutputWavefront) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputWavefront) GetExtraHTTPHeaders() []OutputWavefrontExtraHTTPHeaders {
	if o == nil {
		return nil
	}
	return o.ExtraHTTPHeaders
}

func (o *OutputWavefront) GetFailedRequestLoggingMode() *OutputWavefrontFailedRequestLoggingMode {
	if o == nil {
		return nil
	}
	return o.FailedRequestLoggingMode
}

func (o *OutputWavefront) GetFlushPeriodSec() *int64 {
	if o == nil {
		return nil
	}
	return o.FlushPeriodSec
}

func (o *OutputWavefront) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputWavefront) GetMaxPayloadEvents() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadEvents
}

func (o *OutputWavefront) GetMaxPayloadSizeKB() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadSizeKB
}

func (o *OutputWavefront) GetOnBackpressure() *OutputWavefrontBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputWavefront) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputWavefront) GetPqCompress() *OutputWavefrontCompression {
	if o == nil {
		return nil
	}
	return o.PqCompress
}

func (o *OutputWavefront) GetPqControls() *OutputWavefrontPqControls {
	if o == nil {
		return nil
	}
	return o.PqControls
}

func (o *OutputWavefront) GetPqMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxFileSize
}

func (o *OutputWavefront) GetPqMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxSize
}

func (o *OutputWavefront) GetPqOnBackpressure() *OutputWavefrontQueueFullBehavior {
	if o == nil {
		return nil
	}
	return o.PqOnBackpressure
}

func (o *OutputWavefront) GetPqPath() *string {
	if o == nil {
		return nil
	}
	return o.PqPath
}

func (o *OutputWavefront) GetPqStrictOrdering() *bool {
	if o == nil {
		return nil
	}
	return o.PqStrictOrdering
}

func (o *OutputWavefront) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputWavefront) GetSafeHeaders() []string {
	if o == nil {
		return nil
	}
	return o.SafeHeaders
}

func (o *OutputWavefront) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputWavefront) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputWavefront) GetTextSecret() *string {
	if o == nil {
		return nil
	}
	return o.TextSecret
}

func (o *OutputWavefront) GetTimeoutSec() *int64 {
	if o == nil {
		return nil
	}
	return o.TimeoutSec
}

func (o *OutputWavefront) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *OutputWavefront) GetType() OutputWavefrontType {
	if o == nil {
		return OutputWavefrontType("")
	}
	return o.Type
}

func (o *OutputWavefront) GetUseRoundRobinDNS() *bool {
	if o == nil {
		return nil
	}
	return o.UseRoundRobinDNS
}