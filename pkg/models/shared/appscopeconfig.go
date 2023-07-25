// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AppscopeConfigCribl struct {
	Authtoken               *string            `json:"authtoken,omitempty"`
	Enable                  *bool              `json:"enable,omitempty"`
	Transport               *AppscopeTransport `json:"transport,omitempty"`
	UseScopeSourceTransport *bool              `json:"useScopeSourceTransport,omitempty"`
}

func (o *AppscopeConfigCribl) GetAuthtoken() *string {
	if o == nil {
		return nil
	}
	return o.Authtoken
}

func (o *AppscopeConfigCribl) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

func (o *AppscopeConfigCribl) GetTransport() *AppscopeTransport {
	if o == nil {
		return nil
	}
	return o.Transport
}

func (o *AppscopeConfigCribl) GetUseScopeSourceTransport() *bool {
	if o == nil {
		return nil
	}
	return o.UseScopeSourceTransport
}

type AppscopeConfigEventFormat struct {
	Enhancefs      bool  `json:"enhancefs"`
	Maxeventpersec int64 `json:"maxeventpersec"`
}

func (o *AppscopeConfigEventFormat) GetEnhancefs() bool {
	if o == nil {
		return false
	}
	return o.Enhancefs
}

func (o *AppscopeConfigEventFormat) GetMaxeventpersec() int64 {
	if o == nil {
		return 0
	}
	return o.Maxeventpersec
}

type AppscopeConfigEventType string

const (
	AppscopeConfigEventTypeNdjson AppscopeConfigEventType = "ndjson"
)

func (e AppscopeConfigEventType) ToPointer() *AppscopeConfigEventType {
	return &e
}

func (e *AppscopeConfigEventType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ndjson":
		*e = AppscopeConfigEventType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppscopeConfigEventType: %v", v)
	}
}

type AppscopeConfigEventWatch struct {
	Allowbinary *bool   `json:"allowbinary,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
	Field       *string `json:"field,omitempty"`
	Headers     *string `json:"headers,omitempty"`
	Name        *string `json:"name,omitempty"`
	Type        string  `json:"type"`
	Value       *string `json:"value,omitempty"`
}

func (o *AppscopeConfigEventWatch) GetAllowbinary() *bool {
	if o == nil {
		return nil
	}
	return o.Allowbinary
}

func (o *AppscopeConfigEventWatch) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AppscopeConfigEventWatch) GetField() *string {
	if o == nil {
		return nil
	}
	return o.Field
}

func (o *AppscopeConfigEventWatch) GetHeaders() *string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *AppscopeConfigEventWatch) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AppscopeConfigEventWatch) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *AppscopeConfigEventWatch) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type AppscopeConfigEvent struct {
	Enable    bool                       `json:"enable"`
	Format    AppscopeConfigEventFormat  `json:"format"`
	Transport AppscopeTransport          `json:"transport"`
	Type      AppscopeConfigEventType    `json:"type"`
	Watch     []AppscopeConfigEventWatch `json:"watch"`
}

func (o *AppscopeConfigEvent) GetEnable() bool {
	if o == nil {
		return false
	}
	return o.Enable
}

func (o *AppscopeConfigEvent) GetFormat() AppscopeConfigEventFormat {
	if o == nil {
		return AppscopeConfigEventFormat{}
	}
	return o.Format
}

func (o *AppscopeConfigEvent) GetTransport() AppscopeTransport {
	if o == nil {
		return AppscopeTransport{}
	}
	return o.Transport
}

func (o *AppscopeConfigEvent) GetType() AppscopeConfigEventType {
	if o == nil {
		return AppscopeConfigEventType("")
	}
	return o.Type
}

func (o *AppscopeConfigEvent) GetWatch() []AppscopeConfigEventWatch {
	if o == nil {
		return []AppscopeConfigEventWatch{}
	}
	return o.Watch
}

type AppscopeConfigLibscopeLogLevel string

const (
	AppscopeConfigLibscopeLogLevelDebug   AppscopeConfigLibscopeLogLevel = "debug"
	AppscopeConfigLibscopeLogLevelInfo    AppscopeConfigLibscopeLogLevel = "info"
	AppscopeConfigLibscopeLogLevelWarning AppscopeConfigLibscopeLogLevel = "warning"
	AppscopeConfigLibscopeLogLevelError   AppscopeConfigLibscopeLogLevel = "error"
	AppscopeConfigLibscopeLogLevelNone    AppscopeConfigLibscopeLogLevel = "none"
)

func (e AppscopeConfigLibscopeLogLevel) ToPointer() *AppscopeConfigLibscopeLogLevel {
	return &e
}

func (e *AppscopeConfigLibscopeLogLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "debug":
		fallthrough
	case "info":
		fallthrough
	case "warning":
		fallthrough
	case "error":
		fallthrough
	case "none":
		*e = AppscopeConfigLibscopeLogLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppscopeConfigLibscopeLogLevel: %v", v)
	}
}

type AppscopeConfigLibscopeLog struct {
	Level     *AppscopeConfigLibscopeLogLevel `json:"level,omitempty"`
	Transport *AppscopeTransport              `json:"transport,omitempty"`
}

func (o *AppscopeConfigLibscopeLog) GetLevel() *AppscopeConfigLibscopeLogLevel {
	if o == nil {
		return nil
	}
	return o.Level
}

func (o *AppscopeConfigLibscopeLog) GetTransport() *AppscopeTransport {
	if o == nil {
		return nil
	}
	return o.Transport
}

type AppscopeConfigLibscope struct {
	Commanddir    *string                    `json:"commanddir,omitempty"`
	Configevent   *bool                      `json:"configevent,omitempty"`
	Log           *AppscopeConfigLibscopeLog `json:"log,omitempty"`
	Summaryperiod *int64                     `json:"summaryperiod,omitempty"`
}

func (o *AppscopeConfigLibscope) GetCommanddir() *string {
	if o == nil {
		return nil
	}
	return o.Commanddir
}

func (o *AppscopeConfigLibscope) GetConfigevent() *bool {
	if o == nil {
		return nil
	}
	return o.Configevent
}

func (o *AppscopeConfigLibscope) GetLog() *AppscopeConfigLibscopeLog {
	if o == nil {
		return nil
	}
	return o.Log
}

func (o *AppscopeConfigLibscope) GetSummaryperiod() *int64 {
	if o == nil {
		return nil
	}
	return o.Summaryperiod
}

type AppscopeConfigMetricFormat struct {
	Statsdmaxlen *int64  `json:"statsdmaxlen,omitempty"`
	Statsdprefix *string `json:"statsdprefix,omitempty"`
	Type         *string `json:"type,omitempty"`
	Verbosity    *int64  `json:"verbosity,omitempty"`
}

func (o *AppscopeConfigMetricFormat) GetStatsdmaxlen() *int64 {
	if o == nil {
		return nil
	}
	return o.Statsdmaxlen
}

func (o *AppscopeConfigMetricFormat) GetStatsdprefix() *string {
	if o == nil {
		return nil
	}
	return o.Statsdprefix
}

func (o *AppscopeConfigMetricFormat) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AppscopeConfigMetricFormat) GetVerbosity() *int64 {
	if o == nil {
		return nil
	}
	return o.Verbosity
}

type AppscopeConfigMetric struct {
	Enable    bool                       `json:"enable"`
	Format    AppscopeConfigMetricFormat `json:"format"`
	Transport AppscopeTransport          `json:"transport"`
	Watch     []string                   `json:"watch"`
}

func (o *AppscopeConfigMetric) GetEnable() bool {
	if o == nil {
		return false
	}
	return o.Enable
}

func (o *AppscopeConfigMetric) GetFormat() AppscopeConfigMetricFormat {
	if o == nil {
		return AppscopeConfigMetricFormat{}
	}
	return o.Format
}

func (o *AppscopeConfigMetric) GetTransport() AppscopeTransport {
	if o == nil {
		return AppscopeTransport{}
	}
	return o.Transport
}

func (o *AppscopeConfigMetric) GetWatch() []string {
	if o == nil {
		return []string{}
	}
	return o.Watch
}

type AppscopeConfigPayload struct {
	Dir    string `json:"dir"`
	Enable bool   `json:"enable"`
}

func (o *AppscopeConfigPayload) GetDir() string {
	if o == nil {
		return ""
	}
	return o.Dir
}

func (o *AppscopeConfigPayload) GetEnable() bool {
	if o == nil {
		return false
	}
	return o.Enable
}

type AppscopeConfigProtocol struct {
	Binary  bool   `json:"binary"`
	Detect  bool   `json:"detect"`
	Len     int64  `json:"len"`
	Name    string `json:"name"`
	Payload bool   `json:"payload"`
	Regex   string `json:"regex"`
}

func (o *AppscopeConfigProtocol) GetBinary() bool {
	if o == nil {
		return false
	}
	return o.Binary
}

func (o *AppscopeConfigProtocol) GetDetect() bool {
	if o == nil {
		return false
	}
	return o.Detect
}

func (o *AppscopeConfigProtocol) GetLen() int64 {
	if o == nil {
		return 0
	}
	return o.Len
}

func (o *AppscopeConfigProtocol) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AppscopeConfigProtocol) GetPayload() bool {
	if o == nil {
		return false
	}
	return o.Payload
}

func (o *AppscopeConfigProtocol) GetRegex() string {
	if o == nil {
		return ""
	}
	return o.Regex
}

type AppscopeConfigTags struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (o *AppscopeConfigTags) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *AppscopeConfigTags) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type AppscopeConfig struct {
	Cribl    *AppscopeConfigCribl     `json:"cribl,omitempty"`
	Event    *AppscopeConfigEvent     `json:"event,omitempty"`
	Libscope *AppscopeConfigLibscope  `json:"libscope,omitempty"`
	Metric   *AppscopeConfigMetric    `json:"metric,omitempty"`
	Payload  *AppscopeConfigPayload   `json:"payload,omitempty"`
	Protocol []AppscopeConfigProtocol `json:"protocol,omitempty"`
	Tags     []AppscopeConfigTags     `json:"tags,omitempty"`
}

func (o *AppscopeConfig) GetCribl() *AppscopeConfigCribl {
	if o == nil {
		return nil
	}
	return o.Cribl
}

func (o *AppscopeConfig) GetEvent() *AppscopeConfigEvent {
	if o == nil {
		return nil
	}
	return o.Event
}

func (o *AppscopeConfig) GetLibscope() *AppscopeConfigLibscope {
	if o == nil {
		return nil
	}
	return o.Libscope
}

func (o *AppscopeConfig) GetMetric() *AppscopeConfigMetric {
	if o == nil {
		return nil
	}
	return o.Metric
}

func (o *AppscopeConfig) GetPayload() *AppscopeConfigPayload {
	if o == nil {
		return nil
	}
	return o.Payload
}

func (o *AppscopeConfig) GetProtocol() []AppscopeConfigProtocol {
	if o == nil {
		return nil
	}
	return o.Protocol
}

func (o *AppscopeConfig) GetTags() []AppscopeConfigTags {
	if o == nil {
		return nil
	}
	return o.Tags
}
