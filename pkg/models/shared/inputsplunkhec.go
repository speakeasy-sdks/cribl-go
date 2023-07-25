// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type InputSplunkHecAuthTokensMetadata struct {
	// Field name
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputSplunkHecAuthTokensMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputSplunkHecAuthTokensMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputSplunkHecAuthTokens struct {
	// Enter the values you want to allow in the HEC event index field at the token level. Supports wildcards. To skip validation, leave blank.
	AllowedIndexesAtToken []string `json:"allowedIndexesAtToken,omitempty"`
	// Optional token description
	Description *string `json:"description,omitempty"`
	// Fields to add to events referencing this token.
	Metadata []InputSplunkHecAuthTokensMetadata `json:"metadata,omitempty"`
	// Shared secret to be provided by any client (Authorization: <token>).
	Token string `json:"token"`
}

func (o *InputSplunkHecAuthTokens) GetAllowedIndexesAtToken() []string {
	if o == nil {
		return nil
	}
	return o.AllowedIndexesAtToken
}

func (o *InputSplunkHecAuthTokens) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *InputSplunkHecAuthTokens) GetMetadata() []InputSplunkHecAuthTokensMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputSplunkHecAuthTokens) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type InputSplunkHecConnections struct {
	// Select a Destination.
	Output string `json:"output"`
	// Select Pipeline or Pack. Optional.
	Pipeline *string `json:"pipeline,omitempty"`
}

func (o *InputSplunkHecConnections) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

func (o *InputSplunkHecConnections) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

type InputSplunkHecMetadata struct {
	// Field name
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputSplunkHecMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputSplunkHecMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// InputSplunkHecPqCompression - Codec to use to compress the persisted data.
type InputSplunkHecPqCompression string

const (
	InputSplunkHecPqCompressionNone InputSplunkHecPqCompression = "none"
	InputSplunkHecPqCompressionGzip InputSplunkHecPqCompression = "gzip"
)

func (e InputSplunkHecPqCompression) ToPointer() *InputSplunkHecPqCompression {
	return &e
}

func (e *InputSplunkHecPqCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputSplunkHecPqCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkHecPqCompression: %v", v)
	}
}

// InputSplunkHecPqMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputSplunkHecPqMode string

const (
	InputSplunkHecPqModeSmart  InputSplunkHecPqMode = "smart"
	InputSplunkHecPqModeAlways InputSplunkHecPqMode = "always"
)

func (e InputSplunkHecPqMode) ToPointer() *InputSplunkHecPqMode {
	return &e
}

func (e *InputSplunkHecPqMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputSplunkHecPqMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkHecPqMode: %v", v)
	}
}

type InputSplunkHecPq struct {
	// The number of events to send downstream before committing that Stream has read them.
	CommitFrequency *int64 `json:"commitFrequency,omitempty"`
	// Codec to use to compress the persisted data.
	Compress *InputSplunkHecPqCompression `json:"compress,omitempty"`
	// The maximum number of events to hold in memory before writing the events to disk.
	MaxBufferSize *int64 `json:"maxBufferSize,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	MaxFileSize *string `json:"maxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `json:"maxSize,omitempty"`
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputSplunkHecPqMode `json:"mode,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>.
	Path *string `json:"path,omitempty"`
}

func (o *InputSplunkHecPq) GetCommitFrequency() *int64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputSplunkHecPq) GetCompress() *InputSplunkHecPqCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *InputSplunkHecPq) GetMaxBufferSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputSplunkHecPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputSplunkHecPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputSplunkHecPq) GetMode() *InputSplunkHecPqMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputSplunkHecPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

// InputSplunkHecTLSSettingsServerSideMaximumTLSVersion - Maximum TLS version to accept from connections.
type InputSplunkHecTLSSettingsServerSideMaximumTLSVersion string

const (
	InputSplunkHecTLSSettingsServerSideMaximumTLSVersionTlSv1  InputSplunkHecTLSSettingsServerSideMaximumTLSVersion = "TLSv1"
	InputSplunkHecTLSSettingsServerSideMaximumTLSVersionTlSv11 InputSplunkHecTLSSettingsServerSideMaximumTLSVersion = "TLSv1.1"
	InputSplunkHecTLSSettingsServerSideMaximumTLSVersionTlSv12 InputSplunkHecTLSSettingsServerSideMaximumTLSVersion = "TLSv1.2"
	InputSplunkHecTLSSettingsServerSideMaximumTLSVersionTlSv13 InputSplunkHecTLSSettingsServerSideMaximumTLSVersion = "TLSv1.3"
)

func (e InputSplunkHecTLSSettingsServerSideMaximumTLSVersion) ToPointer() *InputSplunkHecTLSSettingsServerSideMaximumTLSVersion {
	return &e
}

func (e *InputSplunkHecTLSSettingsServerSideMaximumTLSVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSv1":
		fallthrough
	case "TLSv1.1":
		fallthrough
	case "TLSv1.2":
		fallthrough
	case "TLSv1.3":
		*e = InputSplunkHecTLSSettingsServerSideMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkHecTLSSettingsServerSideMaximumTLSVersion: %v", v)
	}
}

// InputSplunkHecTLSSettingsServerSideMinimumTLSVersion - Minimum TLS version to accept from connections.
type InputSplunkHecTLSSettingsServerSideMinimumTLSVersion string

const (
	InputSplunkHecTLSSettingsServerSideMinimumTLSVersionTlSv1  InputSplunkHecTLSSettingsServerSideMinimumTLSVersion = "TLSv1"
	InputSplunkHecTLSSettingsServerSideMinimumTLSVersionTlSv11 InputSplunkHecTLSSettingsServerSideMinimumTLSVersion = "TLSv1.1"
	InputSplunkHecTLSSettingsServerSideMinimumTLSVersionTlSv12 InputSplunkHecTLSSettingsServerSideMinimumTLSVersion = "TLSv1.2"
	InputSplunkHecTLSSettingsServerSideMinimumTLSVersionTlSv13 InputSplunkHecTLSSettingsServerSideMinimumTLSVersion = "TLSv1.3"
)

func (e InputSplunkHecTLSSettingsServerSideMinimumTLSVersion) ToPointer() *InputSplunkHecTLSSettingsServerSideMinimumTLSVersion {
	return &e
}

func (e *InputSplunkHecTLSSettingsServerSideMinimumTLSVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSv1":
		fallthrough
	case "TLSv1.1":
		fallthrough
	case "TLSv1.2":
		fallthrough
	case "TLSv1.3":
		*e = InputSplunkHecTLSSettingsServerSideMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkHecTLSSettingsServerSideMinimumTLSVersion: %v", v)
	}
}

type InputSplunkHecTLSSettingsServerSide struct {
	// Path on server containing CA certificates to use. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Path on server containing certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// The name of the predefined certificate.
	CertificateName *string     `json:"certificateName,omitempty"`
	CommonNameRegex interface{} `json:"commonNameRegex,omitempty"`
	Disabled        *bool       `json:"disabled,omitempty"`
	// Maximum TLS version to accept from connections.
	MaxVersion *InputSplunkHecTLSSettingsServerSideMaximumTLSVersion `json:"maxVersion,omitempty"`
	// Minimum TLS version to accept from connections.
	MinVersion *InputSplunkHecTLSSettingsServerSideMinimumTLSVersion `json:"minVersion,omitempty"`
	// Passphrase to use to decrypt private key.
	Passphrase *string `json:"passphrase,omitempty"`
	// Path on server containing the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath        *string     `json:"privKeyPath,omitempty"`
	RejectUnauthorized interface{} `json:"rejectUnauthorized,omitempty"`
	// Whether to require clients to present their certificates. Used to perform client authentication using SSL certs.
	RequestCert *bool `json:"requestCert,omitempty"`
}

func (o *InputSplunkHecTLSSettingsServerSide) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *InputSplunkHecTLSSettingsServerSide) GetCertPath() *string {
	if o == nil {
		return nil
	}
	return o.CertPath
}

func (o *InputSplunkHecTLSSettingsServerSide) GetCertificateName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateName
}

func (o *InputSplunkHecTLSSettingsServerSide) GetCommonNameRegex() interface{} {
	if o == nil {
		return nil
	}
	return o.CommonNameRegex
}

func (o *InputSplunkHecTLSSettingsServerSide) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputSplunkHecTLSSettingsServerSide) GetMaxVersion() *InputSplunkHecTLSSettingsServerSideMaximumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MaxVersion
}

func (o *InputSplunkHecTLSSettingsServerSide) GetMinVersion() *InputSplunkHecTLSSettingsServerSideMinimumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MinVersion
}

func (o *InputSplunkHecTLSSettingsServerSide) GetPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.Passphrase
}

func (o *InputSplunkHecTLSSettingsServerSide) GetPrivKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.PrivKeyPath
}

func (o *InputSplunkHecTLSSettingsServerSide) GetRejectUnauthorized() interface{} {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *InputSplunkHecTLSSettingsServerSide) GetRequestCert() *bool {
	if o == nil {
		return nil
	}
	return o.RequestCert
}

type InputSplunkHecType string

const (
	InputSplunkHecTypeSplunkHec InputSplunkHecType = "splunk_hec"
)

func (e InputSplunkHecType) ToPointer() *InputSplunkHecType {
	return &e
}

func (e *InputSplunkHecType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "splunk_hec":
		*e = InputSplunkHecType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkHecType: %v", v)
	}
}

type InputSplunkHec struct {
	// Optionally, list HTTP headers that @{product} will send to allowed origins as "Access-Control-Allow-Headers" in a CORS preflight response. Use "*" to allow all headers.
	AccessControlAllowHeaders []string `json:"accessControlAllowHeaders,omitempty"`
	// Optionally, list HTTP origins to which @{product} should send CORS (cross-origin resource sharing) Access-Control-Allow-* headers. Supports wildcards.
	AccessControlAllowOrigin []string `json:"accessControlAllowOrigin,omitempty"`
	// How often request activity is logged at the `info` level. A value of 1 would log every request, 10 every 10th request, etc.
	ActivityLogSampleRate *int64 `json:"activityLogSampleRate,omitempty"`
	// List values allowed in HEC event index field. Leave blank to skip validation. Supports wildcards. The values here can expand index validation at the token level.
	AllowedIndexes []string `json:"allowedIndexes,omitempty"`
	// Shared secrets to be provided by any client (Authorization: <token>). If empty, unauthed access is permitted
	AuthTokens []InputSplunkHecAuthTokens `json:"authTokens,omitempty"`
	// A list of event breaking rulesets that will be applied, in order, to the input data stream.
	BreakerRulesets []string `json:"breakerRulesets,omitempty"`
	// Toggle this to Yes to add request headers to events, in the __headers field.
	CaptureHeaders *bool `json:"captureHeaders,omitempty"`
	// Direct connections to Destinations, optionally via a Pipeline or a Pack.
	Connections []InputSplunkHecConnections `json:"connections,omitempty"`
	// Enable/disable this input
	Disabled *bool `json:"disabled,omitempty"`
	// Enable if the connection is proxied by a device that supports Proxy Protocol V1 or V2.
	EnableProxyHeader *bool `json:"enableProxyHeader,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Address to bind on. Defaults to 0.0.0.0 (all addresses).
	Host string `json:"host"`
	// Unique ID for this input
	ID *string `json:"id,omitempty"`
	// After the last response is sent, @{product} will wait this long for additional data before closing the socket connection. Minimum 1 sec.; maximum 600 sec. (10 min.).
	KeepAliveTimeout *int64 `json:"keepAliveTimeout,omitempty"`
	// Maximum number of active requests per Worker Process. Use 0 for unlimited.
	MaxActiveReq *int64 `json:"maxActiveReq,omitempty"`
	// Fields to add to every event. May be overridden by fields added at the token or request level.
	Metadata []InputSplunkHecMetadata `json:"metadata,omitempty"`
	// Pipeline to process data from this Source before sending it through the Routes.
	Pipeline *string `json:"pipeline,omitempty"`
	// Port to listen on.
	Port int64             `json:"port"`
	Pq   *InputSplunkHecPq `json:"pq,omitempty"`
	// For details on Persistent Queues, see: [https://docs.cribl.io/stream/persistent-queues](https://docs.cribl.io/stream/persistent-queues)
	PqEnabled *bool `json:"pqEnabled,omitempty"`
	// How long to wait for an incoming request to complete before aborting it. Use 0 to disable.
	RequestTimeout *int64 `json:"requestTimeout,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool `json:"sendToRoutes,omitempty"`
	// How long @{product} should wait before assuming that an inactive socket has timed out. To wait forever, set to 0.
	SocketTimeout *int64 `json:"socketTimeout,omitempty"`
	// Absolute path on which to listen for the Splunk HTTP Event Collector API requests. This input supports the /event, /raw and /s2s endpoints.
	SplunkHecAPI string `json:"splunkHecAPI"`
	// Whether to enable Splunk HEC acknowledgements
	SplunkHecAcks *bool `json:"splunkHecAcks,omitempty"`
	// The amount of time (in milliseconds) the Event Breaker will wait for new data to be sent to a specific channel, before flushing the data stream out, as-is, to the Pipelines.
	StaleChannelFlushMs *int64 `json:"staleChannelFlushMs,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string                             `json:"streamtags,omitempty"`
	TLS        *InputSplunkHecTLSSettingsServerSide `json:"tls,omitempty"`
	Type       *InputSplunkHecType                  `json:"type,omitempty"`
	// Enables Event Breakers to determine events' time zone from UF-provided metadata, when TZ can't be inferred from the raw event. Toggle to 'No' to disable this fallback.
	UseFwdTimezone *bool `json:"useFwdTimezone,omitempty"`
}

func (o *InputSplunkHec) GetAccessControlAllowHeaders() []string {
	if o == nil {
		return nil
	}
	return o.AccessControlAllowHeaders
}

func (o *InputSplunkHec) GetAccessControlAllowOrigin() []string {
	if o == nil {
		return nil
	}
	return o.AccessControlAllowOrigin
}

func (o *InputSplunkHec) GetActivityLogSampleRate() *int64 {
	if o == nil {
		return nil
	}
	return o.ActivityLogSampleRate
}

func (o *InputSplunkHec) GetAllowedIndexes() []string {
	if o == nil {
		return nil
	}
	return o.AllowedIndexes
}

func (o *InputSplunkHec) GetAuthTokens() []InputSplunkHecAuthTokens {
	if o == nil {
		return nil
	}
	return o.AuthTokens
}

func (o *InputSplunkHec) GetBreakerRulesets() []string {
	if o == nil {
		return nil
	}
	return o.BreakerRulesets
}

func (o *InputSplunkHec) GetCaptureHeaders() *bool {
	if o == nil {
		return nil
	}
	return o.CaptureHeaders
}

func (o *InputSplunkHec) GetConnections() []InputSplunkHecConnections {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputSplunkHec) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputSplunkHec) GetEnableProxyHeader() *bool {
	if o == nil {
		return nil
	}
	return o.EnableProxyHeader
}

func (o *InputSplunkHec) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputSplunkHec) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *InputSplunkHec) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InputSplunkHec) GetKeepAliveTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepAliveTimeout
}

func (o *InputSplunkHec) GetMaxActiveReq() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxActiveReq
}

func (o *InputSplunkHec) GetMetadata() []InputSplunkHecMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputSplunkHec) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputSplunkHec) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *InputSplunkHec) GetPq() *InputSplunkHecPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputSplunkHec) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputSplunkHec) GetRequestTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestTimeout
}

func (o *InputSplunkHec) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputSplunkHec) GetSocketTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.SocketTimeout
}

func (o *InputSplunkHec) GetSplunkHecAPI() string {
	if o == nil {
		return ""
	}
	return o.SplunkHecAPI
}

func (o *InputSplunkHec) GetSplunkHecAcks() *bool {
	if o == nil {
		return nil
	}
	return o.SplunkHecAcks
}

func (o *InputSplunkHec) GetStaleChannelFlushMs() *int64 {
	if o == nil {
		return nil
	}
	return o.StaleChannelFlushMs
}

func (o *InputSplunkHec) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputSplunkHec) GetTLS() *InputSplunkHecTLSSettingsServerSide {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *InputSplunkHec) GetType() *InputSplunkHecType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *InputSplunkHec) GetUseFwdTimezone() *bool {
	if o == nil {
		return nil
	}
	return o.UseFwdTimezone
}
