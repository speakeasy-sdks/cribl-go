// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputDatasetAuthenticationMethod - Enter API key directly, or select a stored secret
type OutputDatasetAuthenticationMethod string

const (
	OutputDatasetAuthenticationMethodSecret OutputDatasetAuthenticationMethod = "secret"
	OutputDatasetAuthenticationMethodManual OutputDatasetAuthenticationMethod = "manual"
)

func (e OutputDatasetAuthenticationMethod) ToPointer() *OutputDatasetAuthenticationMethod {
	return &e
}

func (e *OutputDatasetAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		fallthrough
	case "manual":
		*e = OutputDatasetAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDatasetAuthenticationMethod: %v", v)
	}
}

// OutputDatasetSeverity - Default value for event severity. If the `sev` or `__severity` fields are set on an event, the first one matching will override this value.
type OutputDatasetSeverity string

const (
	OutputDatasetSeverityFinest  OutputDatasetSeverity = "finest"
	OutputDatasetSeverityFiner   OutputDatasetSeverity = "finer"
	OutputDatasetSeverityFine    OutputDatasetSeverity = "fine"
	OutputDatasetSeverityInfo    OutputDatasetSeverity = "info"
	OutputDatasetSeverityWarning OutputDatasetSeverity = "warning"
	OutputDatasetSeverityError   OutputDatasetSeverity = "error"
	OutputDatasetSeverityFatal   OutputDatasetSeverity = "fatal"
)

func (e OutputDatasetSeverity) ToPointer() *OutputDatasetSeverity {
	return &e
}

func (e *OutputDatasetSeverity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "finest":
		fallthrough
	case "finer":
		fallthrough
	case "fine":
		fallthrough
	case "info":
		fallthrough
	case "warning":
		fallthrough
	case "error":
		fallthrough
	case "fatal":
		*e = OutputDatasetSeverity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDatasetSeverity: %v", v)
	}
}

type OutputDatasetExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

func (o *OutputDatasetExtraHTTPHeaders) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *OutputDatasetExtraHTTPHeaders) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// OutputDatasetFailedRequestLoggingMode - Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
type OutputDatasetFailedRequestLoggingMode string

const (
	OutputDatasetFailedRequestLoggingModePayload           OutputDatasetFailedRequestLoggingMode = "payload"
	OutputDatasetFailedRequestLoggingModePayloadAndHeaders OutputDatasetFailedRequestLoggingMode = "payloadAndHeaders"
	OutputDatasetFailedRequestLoggingModeNone              OutputDatasetFailedRequestLoggingMode = "none"
)

func (e OutputDatasetFailedRequestLoggingMode) ToPointer() *OutputDatasetFailedRequestLoggingMode {
	return &e
}

func (e *OutputDatasetFailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputDatasetFailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDatasetFailedRequestLoggingMode: %v", v)
	}
}

// OutputDatasetBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputDatasetBackpressureBehavior string

const (
	OutputDatasetBackpressureBehaviorQueue OutputDatasetBackpressureBehavior = "queue"
	OutputDatasetBackpressureBehaviorDrop  OutputDatasetBackpressureBehavior = "drop"
	OutputDatasetBackpressureBehaviorBlock OutputDatasetBackpressureBehavior = "block"
)

func (e OutputDatasetBackpressureBehavior) ToPointer() *OutputDatasetBackpressureBehavior {
	return &e
}

func (e *OutputDatasetBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputDatasetBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDatasetBackpressureBehavior: %v", v)
	}
}

// OutputDatasetCompression - Codec to use to compress the persisted data.
type OutputDatasetCompression string

const (
	OutputDatasetCompressionNone OutputDatasetCompression = "none"
	OutputDatasetCompressionGzip OutputDatasetCompression = "gzip"
)

func (e OutputDatasetCompression) ToPointer() *OutputDatasetCompression {
	return &e
}

func (e *OutputDatasetCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputDatasetCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDatasetCompression: %v", v)
	}
}

type OutputDatasetPqControls struct {
}

// OutputDatasetQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputDatasetQueueFullBehavior string

const (
	OutputDatasetQueueFullBehaviorBlock OutputDatasetQueueFullBehavior = "block"
	OutputDatasetQueueFullBehaviorDrop  OutputDatasetQueueFullBehavior = "drop"
)

func (e OutputDatasetQueueFullBehavior) ToPointer() *OutputDatasetQueueFullBehavior {
	return &e
}

func (e *OutputDatasetQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputDatasetQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDatasetQueueFullBehavior: %v", v)
	}
}

// OutputDatasetDataSetSite - DataSet site to which events should be sent
type OutputDatasetDataSetSite string

const (
	OutputDatasetDataSetSiteUs     OutputDatasetDataSetSite = "us"
	OutputDatasetDataSetSiteEu     OutputDatasetDataSetSite = "eu"
	OutputDatasetDataSetSiteCustom OutputDatasetDataSetSite = "custom"
)

func (e OutputDatasetDataSetSite) ToPointer() *OutputDatasetDataSetSite {
	return &e
}

func (e *OutputDatasetDataSetSite) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "us":
		fallthrough
	case "eu":
		fallthrough
	case "custom":
		*e = OutputDatasetDataSetSite(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDatasetDataSetSite: %v", v)
	}
}

type OutputDatasetType string

const (
	OutputDatasetTypeDataset OutputDatasetType = "dataset"
)

func (e OutputDatasetType) ToPointer() *OutputDatasetType {
	return &e
}

func (e *OutputDatasetType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dataset":
		*e = OutputDatasetType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputDatasetType: %v", v)
	}
}

type OutputDataset struct {
	// A 'Log Write Access' API key for the DataSet account
	APIKey *string `json:"apiKey,omitempty"`
	// Enter API key directly, or select a stored secret
	AuthType *OutputDatasetAuthenticationMethod `json:"authType,omitempty"`
	// Whether to compress the payload body before sending.
	Compress *bool `json:"compress,omitempty"`
	// Maximum number of ongoing requests before blocking.
	Concurrency *int64  `json:"concurrency,omitempty"`
	CustomURL   *string `json:"customUrl,omitempty"`
	// Default value for event severity. If the `sev` or `__severity` fields are set on an event, the first one matching will override this value.
	DefaultSeverity *OutputDatasetSeverity `json:"defaultSeverity,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Fields to exclude from the event if the Message field is either unspecified or refers to an object. Ignored if the Message field is a string. If empty, we send all non-internal fields.
	ExcludeFields []string `json:"excludeFields,omitempty"`
	// Headers to add to all events.
	ExtraHTTPHeaders []OutputDatasetExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// Determines which data should be logged when a request fails. Defaults to None.  All headers are redacted by default, except those listed under `Safe Headers`.
	FailedRequestLoggingMode *OutputDatasetFailedRequestLoggingMode `json:"failedRequestLoggingMode,omitempty"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Max body size.
	FlushPeriodSec *int64 `json:"flushPeriodSec,omitempty"`
	// Unique ID for this output
	ID string `json:"id"`
	// Max number of events to include in the request body. Default is 0 (unlimited).
	MaxPayloadEvents *int64 `json:"maxPayloadEvents,omitempty"`
	// Maximum size, in KB, of the request body.
	MaxPayloadSizeKB *int64 `json:"maxPayloadSizeKB,omitempty"`
	// Name of the event field that contains the message or attributes to send. If not specified, all of the event's non-internal fields will be sent as attributes.
	MessageField *string `json:"messageField,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputDatasetBackpressureBehavior `json:"onBackpressure,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputDatasetCompression `json:"pqCompress,omitempty"`
	PqControls *OutputDatasetPqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputDatasetQueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool `json:"pqStrictOrdering,omitempty"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another trusted CA (e.g., the system's CA). Defaults to Yes.
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// List of headers that are safe to log in plain text.
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// Name of the event field that contains the `serverHost` identifier. If not specified, defaults to `cribl_<outputId>`.
	ServerHostField *string `json:"serverHostField,omitempty"`
	// DataSet site to which events should be sent
	Site *OutputDatasetDataSetSite `json:"site,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string `json:"systemFields,omitempty"`
	// Select (or create) a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// Amount of time, in seconds, to wait for a request to complete before aborting it.
	TimeoutSec *int64 `json:"timeoutSec,omitempty"`
	// Name of the event field that contains the timestamp. If not specified, defaults to `ts`, `_time`, or `Date.now()`, in that order.
	TimestampField *string           `json:"timestampField,omitempty"`
	Type           OutputDatasetType `json:"type"`
	// Enable to use round-robin DNS lookup. When a DNS server returns multiple addresses, this will cause Stream to cycle through them in the order returned.
	UseRoundRobinDNS *bool `json:"useRoundRobinDns,omitempty"`
}

func (o *OutputDataset) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

func (o *OutputDataset) GetAuthType() *OutputDatasetAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *OutputDataset) GetCompress() *bool {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *OutputDataset) GetConcurrency() *int64 {
	if o == nil {
		return nil
	}
	return o.Concurrency
}

func (o *OutputDataset) GetCustomURL() *string {
	if o == nil {
		return nil
	}
	return o.CustomURL
}

func (o *OutputDataset) GetDefaultSeverity() *OutputDatasetSeverity {
	if o == nil {
		return nil
	}
	return o.DefaultSeverity
}

func (o *OutputDataset) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputDataset) GetExcludeFields() []string {
	if o == nil {
		return nil
	}
	return o.ExcludeFields
}

func (o *OutputDataset) GetExtraHTTPHeaders() []OutputDatasetExtraHTTPHeaders {
	if o == nil {
		return nil
	}
	return o.ExtraHTTPHeaders
}

func (o *OutputDataset) GetFailedRequestLoggingMode() *OutputDatasetFailedRequestLoggingMode {
	if o == nil {
		return nil
	}
	return o.FailedRequestLoggingMode
}

func (o *OutputDataset) GetFlushPeriodSec() *int64 {
	if o == nil {
		return nil
	}
	return o.FlushPeriodSec
}

func (o *OutputDataset) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *OutputDataset) GetMaxPayloadEvents() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadEvents
}

func (o *OutputDataset) GetMaxPayloadSizeKB() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadSizeKB
}

func (o *OutputDataset) GetMessageField() *string {
	if o == nil {
		return nil
	}
	return o.MessageField
}

func (o *OutputDataset) GetOnBackpressure() *OutputDatasetBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputDataset) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputDataset) GetPqCompress() *OutputDatasetCompression {
	if o == nil {
		return nil
	}
	return o.PqCompress
}

func (o *OutputDataset) GetPqControls() *OutputDatasetPqControls {
	if o == nil {
		return nil
	}
	return o.PqControls
}

func (o *OutputDataset) GetPqMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxFileSize
}

func (o *OutputDataset) GetPqMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxSize
}

func (o *OutputDataset) GetPqOnBackpressure() *OutputDatasetQueueFullBehavior {
	if o == nil {
		return nil
	}
	return o.PqOnBackpressure
}

func (o *OutputDataset) GetPqPath() *string {
	if o == nil {
		return nil
	}
	return o.PqPath
}

func (o *OutputDataset) GetPqStrictOrdering() *bool {
	if o == nil {
		return nil
	}
	return o.PqStrictOrdering
}

func (o *OutputDataset) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputDataset) GetSafeHeaders() []string {
	if o == nil {
		return nil
	}
	return o.SafeHeaders
}

func (o *OutputDataset) GetServerHostField() *string {
	if o == nil {
		return nil
	}
	return o.ServerHostField
}

func (o *OutputDataset) GetSite() *OutputDatasetDataSetSite {
	if o == nil {
		return nil
	}
	return o.Site
}

func (o *OutputDataset) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputDataset) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputDataset) GetTextSecret() *string {
	if o == nil {
		return nil
	}
	return o.TextSecret
}

func (o *OutputDataset) GetTimeoutSec() *int64 {
	if o == nil {
		return nil
	}
	return o.TimeoutSec
}

func (o *OutputDataset) GetTimestampField() *string {
	if o == nil {
		return nil
	}
	return o.TimestampField
}

func (o *OutputDataset) GetType() OutputDatasetType {
	if o == nil {
		return OutputDatasetType("")
	}
	return o.Type
}

func (o *OutputDataset) GetUseRoundRobinDNS() *bool {
	if o == nil {
		return nil
	}
	return o.UseRoundRobinDNS
}
