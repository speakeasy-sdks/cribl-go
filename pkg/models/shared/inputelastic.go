// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// InputElasticAPIVersion - The API version to use for communicating with the server.
type InputElasticAPIVersion string

const (
	InputElasticAPIVersionCustom  InputElasticAPIVersion = "custom"
	InputElasticAPIVersionEight32 InputElasticAPIVersion = "8.3.2"
)

func (e InputElasticAPIVersion) ToPointer() *InputElasticAPIVersion {
	return &e
}

func (e *InputElasticAPIVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "custom":
		fallthrough
	case "8.3.2":
		*e = InputElasticAPIVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputElasticAPIVersion: %v", v)
	}
}

// InputElasticAuthenticationType - Elastic authentication type
type InputElasticAuthenticationType string

const (
	InputElasticAuthenticationTypeAuthTokens        InputElasticAuthenticationType = "authTokens"
	InputElasticAuthenticationTypeBasic             InputElasticAuthenticationType = "basic"
	InputElasticAuthenticationTypeCredentialsSecret InputElasticAuthenticationType = "credentialsSecret"
)

func (e InputElasticAuthenticationType) ToPointer() *InputElasticAuthenticationType {
	return &e
}

func (e *InputElasticAuthenticationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "authTokens":
		fallthrough
	case "basic":
		fallthrough
	case "credentialsSecret":
		*e = InputElasticAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputElasticAuthenticationType: %v", v)
	}
}

type InputElasticConnections struct {
	// Select a Destination.
	Output string `json:"output"`
	// Select Pipeline or Pack. Optional.
	Pipeline *string `json:"pipeline,omitempty"`
}

func (o *InputElasticConnections) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

func (o *InputElasticConnections) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

type InputElasticExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

func (o *InputElasticExtraHTTPHeaders) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *InputElasticExtraHTTPHeaders) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputElasticMetadata struct {
	// Field name
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputElasticMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputElasticMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// InputElasticPqCompression - Codec to use to compress the persisted data.
type InputElasticPqCompression string

const (
	InputElasticPqCompressionNone InputElasticPqCompression = "none"
	InputElasticPqCompressionGzip InputElasticPqCompression = "gzip"
)

func (e InputElasticPqCompression) ToPointer() *InputElasticPqCompression {
	return &e
}

func (e *InputElasticPqCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputElasticPqCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputElasticPqCompression: %v", v)
	}
}

// InputElasticPqMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputElasticPqMode string

const (
	InputElasticPqModeSmart  InputElasticPqMode = "smart"
	InputElasticPqModeAlways InputElasticPqMode = "always"
)

func (e InputElasticPqMode) ToPointer() *InputElasticPqMode {
	return &e
}

func (e *InputElasticPqMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputElasticPqMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputElasticPqMode: %v", v)
	}
}

type InputElasticPq struct {
	// The number of events to send downstream before committing that Stream has read them.
	CommitFrequency *int64 `json:"commitFrequency,omitempty"`
	// Codec to use to compress the persisted data.
	Compress *InputElasticPqCompression `json:"compress,omitempty"`
	// The maximum number of events to hold in memory before writing the events to disk.
	MaxBufferSize *int64 `json:"maxBufferSize,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	MaxFileSize *string `json:"maxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `json:"maxSize,omitempty"`
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputElasticPqMode `json:"mode,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>.
	Path *string `json:"path,omitempty"`
}

func (o *InputElasticPq) GetCommitFrequency() *int64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputElasticPq) GetCompress() *InputElasticPqCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *InputElasticPq) GetMaxBufferSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputElasticPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputElasticPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputElasticPq) GetMode() *InputElasticPqMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputElasticPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

// InputElasticProxyModeAuthenticationMethod - Enter credentials directly, or select a stored secret
type InputElasticProxyModeAuthenticationMethod string

const (
	InputElasticProxyModeAuthenticationMethodNone   InputElasticProxyModeAuthenticationMethod = "none"
	InputElasticProxyModeAuthenticationMethodManual InputElasticProxyModeAuthenticationMethod = "manual"
	InputElasticProxyModeAuthenticationMethodSecret InputElasticProxyModeAuthenticationMethod = "secret"
)

func (e InputElasticProxyModeAuthenticationMethod) ToPointer() *InputElasticProxyModeAuthenticationMethod {
	return &e
}

func (e *InputElasticProxyModeAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "manual":
		fallthrough
	case "secret":
		*e = InputElasticProxyModeAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputElasticProxyModeAuthenticationMethod: %v", v)
	}
}

type InputElasticProxyMode struct {
	// Enter credentials directly, or select a stored secret
	AuthType *InputElasticProxyModeAuthenticationMethod `json:"authType,omitempty"`
	// Enable proxying of non-bulk API requests to an external Elastic server. Enable this only if you understand the implications; see docs for more details.
	Enabled bool `json:"enabled"`
	// List of headers to remove from the request to proxy
	RemoveHeaders []string `json:"removeHeaders,omitempty"`
	// Amount of time, in seconds, to wait for a proxy request to complete before aborting it.
	TimeoutSec *int64 `json:"timeoutSec,omitempty"`
	// URL of the Elastic server to proxy non-bulk requests to, e.g., http://elastic:9200
	URL *string `json:"url,omitempty"`
}

func (o *InputElasticProxyMode) GetAuthType() *InputElasticProxyModeAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *InputElasticProxyMode) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *InputElasticProxyMode) GetRemoveHeaders() []string {
	if o == nil {
		return nil
	}
	return o.RemoveHeaders
}

func (o *InputElasticProxyMode) GetTimeoutSec() *int64 {
	if o == nil {
		return nil
	}
	return o.TimeoutSec
}

func (o *InputElasticProxyMode) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// InputElasticTLSSettingsServerSideMaximumTLSVersion - Maximum TLS version to accept from connections.
type InputElasticTLSSettingsServerSideMaximumTLSVersion string

const (
	InputElasticTLSSettingsServerSideMaximumTLSVersionTlSv1  InputElasticTLSSettingsServerSideMaximumTLSVersion = "TLSv1"
	InputElasticTLSSettingsServerSideMaximumTLSVersionTlSv11 InputElasticTLSSettingsServerSideMaximumTLSVersion = "TLSv1.1"
	InputElasticTLSSettingsServerSideMaximumTLSVersionTlSv12 InputElasticTLSSettingsServerSideMaximumTLSVersion = "TLSv1.2"
	InputElasticTLSSettingsServerSideMaximumTLSVersionTlSv13 InputElasticTLSSettingsServerSideMaximumTLSVersion = "TLSv1.3"
)

func (e InputElasticTLSSettingsServerSideMaximumTLSVersion) ToPointer() *InputElasticTLSSettingsServerSideMaximumTLSVersion {
	return &e
}

func (e *InputElasticTLSSettingsServerSideMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputElasticTLSSettingsServerSideMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputElasticTLSSettingsServerSideMaximumTLSVersion: %v", v)
	}
}

// InputElasticTLSSettingsServerSideMinimumTLSVersion - Minimum TLS version to accept from connections.
type InputElasticTLSSettingsServerSideMinimumTLSVersion string

const (
	InputElasticTLSSettingsServerSideMinimumTLSVersionTlSv1  InputElasticTLSSettingsServerSideMinimumTLSVersion = "TLSv1"
	InputElasticTLSSettingsServerSideMinimumTLSVersionTlSv11 InputElasticTLSSettingsServerSideMinimumTLSVersion = "TLSv1.1"
	InputElasticTLSSettingsServerSideMinimumTLSVersionTlSv12 InputElasticTLSSettingsServerSideMinimumTLSVersion = "TLSv1.2"
	InputElasticTLSSettingsServerSideMinimumTLSVersionTlSv13 InputElasticTLSSettingsServerSideMinimumTLSVersion = "TLSv1.3"
)

func (e InputElasticTLSSettingsServerSideMinimumTLSVersion) ToPointer() *InputElasticTLSSettingsServerSideMinimumTLSVersion {
	return &e
}

func (e *InputElasticTLSSettingsServerSideMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputElasticTLSSettingsServerSideMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputElasticTLSSettingsServerSideMinimumTLSVersion: %v", v)
	}
}

type InputElasticTLSSettingsServerSide struct {
	// Path on server containing CA certificates to use. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Path on server containing certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// The name of the predefined certificate.
	CertificateName *string     `json:"certificateName,omitempty"`
	CommonNameRegex interface{} `json:"commonNameRegex,omitempty"`
	Disabled        *bool       `json:"disabled,omitempty"`
	// Maximum TLS version to accept from connections.
	MaxVersion *InputElasticTLSSettingsServerSideMaximumTLSVersion `json:"maxVersion,omitempty"`
	// Minimum TLS version to accept from connections.
	MinVersion *InputElasticTLSSettingsServerSideMinimumTLSVersion `json:"minVersion,omitempty"`
	// Passphrase to use to decrypt private key.
	Passphrase *string `json:"passphrase,omitempty"`
	// Path on server containing the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath        *string     `json:"privKeyPath,omitempty"`
	RejectUnauthorized interface{} `json:"rejectUnauthorized,omitempty"`
	// Whether to require clients to present their certificates. Used to perform client authentication using SSL certs.
	RequestCert *bool `json:"requestCert,omitempty"`
}

func (o *InputElasticTLSSettingsServerSide) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *InputElasticTLSSettingsServerSide) GetCertPath() *string {
	if o == nil {
		return nil
	}
	return o.CertPath
}

func (o *InputElasticTLSSettingsServerSide) GetCertificateName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateName
}

func (o *InputElasticTLSSettingsServerSide) GetCommonNameRegex() interface{} {
	if o == nil {
		return nil
	}
	return o.CommonNameRegex
}

func (o *InputElasticTLSSettingsServerSide) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputElasticTLSSettingsServerSide) GetMaxVersion() *InputElasticTLSSettingsServerSideMaximumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MaxVersion
}

func (o *InputElasticTLSSettingsServerSide) GetMinVersion() *InputElasticTLSSettingsServerSideMinimumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MinVersion
}

func (o *InputElasticTLSSettingsServerSide) GetPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.Passphrase
}

func (o *InputElasticTLSSettingsServerSide) GetPrivKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.PrivKeyPath
}

func (o *InputElasticTLSSettingsServerSide) GetRejectUnauthorized() interface{} {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *InputElasticTLSSettingsServerSide) GetRequestCert() *bool {
	if o == nil {
		return nil
	}
	return o.RequestCert
}

type InputElasticType string

const (
	InputElasticTypeElastic InputElasticType = "elastic"
)

func (e InputElasticType) ToPointer() *InputElasticType {
	return &e
}

func (e *InputElasticType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "elastic":
		*e = InputElasticType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputElasticType: %v", v)
	}
}

type InputElastic struct {
	// How often request activity is logged at the `info` level. A value of 1 would log every request, 10 every 10th request, etc.
	ActivityLogSampleRate *int64 `json:"activityLogSampleRate,omitempty"`
	// The API version to use for communicating with the server.
	APIVersion *InputElasticAPIVersion `json:"apiVersion,omitempty"`
	// Bearer tokens to include in the authorization header
	AuthTokens []string `json:"authTokens,omitempty"`
	// Elastic authentication type
	AuthType *InputElasticAuthenticationType `json:"authType,omitempty"`
	// Toggle this to Yes to add request headers to events, in the __headers field.
	CaptureHeaders *bool `json:"captureHeaders,omitempty"`
	// Direct connections to Destinations, optionally via a Pipeline or a Pack.
	Connections []InputElasticConnections `json:"connections,omitempty"`
	// Select (or create) a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
	// Custom version information to respond to requests
	CustomAPIVersion *string `json:"customAPIVersion,omitempty"`
	// Enable/disable this input
	Disabled *bool `json:"disabled,omitempty"`
	// Absolute path on which to listen for Elasticsearch API requests. Defaults to /. _bulk will be appended automatically, e.g., /myPath becomes /myPath/_bulk. Requests can then be made to either /myPath/_bulk or /myPath/<myIndexName>/_bulk. Other entries are faked as success.
	ElasticAPI string `json:"elasticAPI"`
	// Enable if the connection is proxied by a device that supports Proxy Protocol V1 or V2.
	EnableProxyHeader *bool `json:"enableProxyHeader,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Headers to add to all events.
	ExtraHTTPHeaders []InputElasticExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// Address to bind on. Defaults to 0.0.0.0 (all addresses).
	Host string `json:"host"`
	// Unique ID for this input
	ID *string `json:"id,omitempty"`
	// After the last response is sent, @{product} will wait this long for additional data before closing the socket connection. Minimum 1 sec.; maximum 600 sec. (10 min.).
	KeepAliveTimeout *int64 `json:"keepAliveTimeout,omitempty"`
	// Maximum number of active requests per Worker Process. Use 0 for unlimited.
	MaxActiveReq *int64 `json:"maxActiveReq,omitempty"`
	// Fields to add to events from this input.
	Metadata []InputElasticMetadata `json:"metadata,omitempty"`
	// Password for Basic authentication
	Password *string `json:"password,omitempty"`
	// Pipeline to process data from this Source before sending it through the Routes.
	Pipeline *string `json:"pipeline,omitempty"`
	// Port to listen on.
	Port int64           `json:"port"`
	Pq   *InputElasticPq `json:"pq,omitempty"`
	// For details on Persistent Queues, see: [https://docs.cribl.io/stream/persistent-queues](https://docs.cribl.io/stream/persistent-queues)
	PqEnabled *bool                  `json:"pqEnabled,omitempty"`
	ProxyMode *InputElasticProxyMode `json:"proxyMode,omitempty"`
	// How long to wait for an incoming request to complete before aborting it. Use 0 to disable.
	RequestTimeout *int64 `json:"requestTimeout,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool `json:"sendToRoutes,omitempty"`
	// How long @{product} should wait before assuming that an inactive socket has timed out. To wait forever, set to 0.
	SocketTimeout *int64 `json:"socketTimeout,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string                           `json:"streamtags,omitempty"`
	TLS        *InputElasticTLSSettingsServerSide `json:"tls,omitempty"`
	Type       *InputElasticType                  `json:"type,omitempty"`
	// Username for Basic authentication
	Username *string `json:"username,omitempty"`
}

func (o *InputElastic) GetActivityLogSampleRate() *int64 {
	if o == nil {
		return nil
	}
	return o.ActivityLogSampleRate
}

func (o *InputElastic) GetAPIVersion() *InputElasticAPIVersion {
	if o == nil {
		return nil
	}
	return o.APIVersion
}

func (o *InputElastic) GetAuthTokens() []string {
	if o == nil {
		return nil
	}
	return o.AuthTokens
}

func (o *InputElastic) GetAuthType() *InputElasticAuthenticationType {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *InputElastic) GetCaptureHeaders() *bool {
	if o == nil {
		return nil
	}
	return o.CaptureHeaders
}

func (o *InputElastic) GetConnections() []InputElasticConnections {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputElastic) GetCredentialsSecret() *string {
	if o == nil {
		return nil
	}
	return o.CredentialsSecret
}

func (o *InputElastic) GetCustomAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.CustomAPIVersion
}

func (o *InputElastic) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputElastic) GetElasticAPI() string {
	if o == nil {
		return ""
	}
	return o.ElasticAPI
}

func (o *InputElastic) GetEnableProxyHeader() *bool {
	if o == nil {
		return nil
	}
	return o.EnableProxyHeader
}

func (o *InputElastic) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputElastic) GetExtraHTTPHeaders() []InputElasticExtraHTTPHeaders {
	if o == nil {
		return nil
	}
	return o.ExtraHTTPHeaders
}

func (o *InputElastic) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *InputElastic) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InputElastic) GetKeepAliveTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepAliveTimeout
}

func (o *InputElastic) GetMaxActiveReq() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxActiveReq
}

func (o *InputElastic) GetMetadata() []InputElasticMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputElastic) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *InputElastic) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputElastic) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *InputElastic) GetPq() *InputElasticPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputElastic) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputElastic) GetProxyMode() *InputElasticProxyMode {
	if o == nil {
		return nil
	}
	return o.ProxyMode
}

func (o *InputElastic) GetRequestTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestTimeout
}

func (o *InputElastic) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputElastic) GetSocketTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.SocketTimeout
}

func (o *InputElastic) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputElastic) GetTLS() *InputElasticTLSSettingsServerSide {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *InputElastic) GetType() *InputElasticType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *InputElastic) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}
