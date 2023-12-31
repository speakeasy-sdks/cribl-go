// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputGooglePubsubAuthenticationMethod - Google authentication method. Choose Auto to use environment variables PUBSUB_PROJECT and PUBSUB_CREDENTIALS..
type OutputGooglePubsubAuthenticationMethod string

const (
	OutputGooglePubsubAuthenticationMethodSecret OutputGooglePubsubAuthenticationMethod = "secret"
	OutputGooglePubsubAuthenticationMethodManual OutputGooglePubsubAuthenticationMethod = "manual"
)

func (e OutputGooglePubsubAuthenticationMethod) ToPointer() *OutputGooglePubsubAuthenticationMethod {
	return &e
}

func (e *OutputGooglePubsubAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		fallthrough
	case "manual":
		*e = OutputGooglePubsubAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGooglePubsubAuthenticationMethod: %v", v)
	}
}

// OutputGooglePubsubBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputGooglePubsubBackpressureBehavior string

const (
	OutputGooglePubsubBackpressureBehaviorQueue OutputGooglePubsubBackpressureBehavior = "queue"
	OutputGooglePubsubBackpressureBehaviorDrop  OutputGooglePubsubBackpressureBehavior = "drop"
	OutputGooglePubsubBackpressureBehaviorBlock OutputGooglePubsubBackpressureBehavior = "block"
)

func (e OutputGooglePubsubBackpressureBehavior) ToPointer() *OutputGooglePubsubBackpressureBehavior {
	return &e
}

func (e *OutputGooglePubsubBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputGooglePubsubBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGooglePubsubBackpressureBehavior: %v", v)
	}
}

// OutputGooglePubsubCompression - Codec to use to compress the persisted data.
type OutputGooglePubsubCompression string

const (
	OutputGooglePubsubCompressionNone OutputGooglePubsubCompression = "none"
	OutputGooglePubsubCompressionGzip OutputGooglePubsubCompression = "gzip"
)

func (e OutputGooglePubsubCompression) ToPointer() *OutputGooglePubsubCompression {
	return &e
}

func (e *OutputGooglePubsubCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputGooglePubsubCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGooglePubsubCompression: %v", v)
	}
}

type OutputGooglePubsubPqControls struct {
}

// OutputGooglePubsubQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputGooglePubsubQueueFullBehavior string

const (
	OutputGooglePubsubQueueFullBehaviorBlock OutputGooglePubsubQueueFullBehavior = "block"
	OutputGooglePubsubQueueFullBehaviorDrop  OutputGooglePubsubQueueFullBehavior = "drop"
)

func (e OutputGooglePubsubQueueFullBehavior) ToPointer() *OutputGooglePubsubQueueFullBehavior {
	return &e
}

func (e *OutputGooglePubsubQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputGooglePubsubQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGooglePubsubQueueFullBehavior: %v", v)
	}
}

// OutputGooglePubsubRegion - Region to publish messages to. Select 'default' to allow Google to auto-select the nearest region. When using ordered delivery, the selected region must be allowed by message storage policy.
type OutputGooglePubsubRegion string

const (
	OutputGooglePubsubRegionPubsubGoogleapisCom                       OutputGooglePubsubRegion = "pubsub.googleapis.com"
	OutputGooglePubsubRegionUsEast1PubsubGoogleapisCom                OutputGooglePubsubRegion = "us-east1-pubsub.googleapis.com"
	OutputGooglePubsubRegionUsEast4PubsubGoogleapisCom                OutputGooglePubsubRegion = "us-east4-pubsub.googleapis.com"
	OutputGooglePubsubRegionUsCentral1PubsubGoogleapisCom             OutputGooglePubsubRegion = "us-central1-pubsub.googleapis.com"
	OutputGooglePubsubRegionUsCentral2PubsubGoogleapisCom             OutputGooglePubsubRegion = "us-central2-pubsub.googleapis.com"
	OutputGooglePubsubRegionUsWest1PubsubGoogleapisCom                OutputGooglePubsubRegion = "us-west1-pubsub.googleapis.com"
	OutputGooglePubsubRegionUsWest2PubsubGoogleapisCom                OutputGooglePubsubRegion = "us-west2-pubsub.googleapis.com"
	OutputGooglePubsubRegionUsWest3PubsubGoogleapisCom                OutputGooglePubsubRegion = "us-west3-pubsub.googleapis.com"
	OutputGooglePubsubRegionSouthamericaEast1PubsubGoogleapisCom      OutputGooglePubsubRegion = "southamerica-east1-pubsub.googleapis.com"
	OutputGooglePubsubRegionNorthamericaNortheast1PubsubGoogleapisCom OutputGooglePubsubRegion = "northamerica-northeast1-pubsub.googleapis.com"
	OutputGooglePubsubRegionEuropeWest6PubsubGoogleapisCom            OutputGooglePubsubRegion = "europe-west6-pubsub.googleapis.com"
	OutputGooglePubsubRegionEuropeWest4PubsubGoogleapisCom            OutputGooglePubsubRegion = "europe-west4-pubsub.googleapis.com"
	OutputGooglePubsubRegionEuropeWest3PubsubGoogleapisCom            OutputGooglePubsubRegion = "europe-west3-pubsub.googleapis.com"
	OutputGooglePubsubRegionEuropeWest2PubsubGoogleapisCom            OutputGooglePubsubRegion = "europe-west2-pubsub.googleapis.com"
	OutputGooglePubsubRegionEuropeWest1PubsubGoogleapisCom            OutputGooglePubsubRegion = "europe-west1-pubsub.googleapis.com"
	OutputGooglePubsubRegionEuropeNorth1PubsubGoogleapisCom           OutputGooglePubsubRegion = "europe-north1-pubsub.googleapis.com"
	OutputGooglePubsubRegionAustraliaSoutheast1PubsubGoogleapisCom    OutputGooglePubsubRegion = "australia-southeast1-pubsub.googleapis.com"
	OutputGooglePubsubRegionAsiaSoutheast1PubsubGoogleapisCom         OutputGooglePubsubRegion = "asia-southeast1-pubsub.googleapis.com"
	OutputGooglePubsubRegionAsiaSouth1PubsubGoogleapisCom             OutputGooglePubsubRegion = "asia-south1-pubsub.googleapis.com"
	OutputGooglePubsubRegionAsiaNortheast3PubsubGoogleapisCom         OutputGooglePubsubRegion = "asia-northeast3-pubsub.googleapis.com"
	OutputGooglePubsubRegionAsiaNortheast2PubsubGoogleapisCom         OutputGooglePubsubRegion = "asia-northeast2-pubsub.googleapis.com"
	OutputGooglePubsubRegionAsiaNortheast1PubsubGoogleapisCom         OutputGooglePubsubRegion = "asia-northeast1-pubsub.googleapis.com"
	OutputGooglePubsubRegionAsiaEast2PubsubGoogleapisCom              OutputGooglePubsubRegion = "asia-east2-pubsub.googleapis.com"
	OutputGooglePubsubRegionAsiaEast1PubsubGoogleapisCom              OutputGooglePubsubRegion = "asia-east1-pubsub.googleapis.com"
)

func (e OutputGooglePubsubRegion) ToPointer() *OutputGooglePubsubRegion {
	return &e
}

func (e *OutputGooglePubsubRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pubsub.googleapis.com":
		fallthrough
	case "us-east1-pubsub.googleapis.com":
		fallthrough
	case "us-east4-pubsub.googleapis.com":
		fallthrough
	case "us-central1-pubsub.googleapis.com":
		fallthrough
	case "us-central2-pubsub.googleapis.com":
		fallthrough
	case "us-west1-pubsub.googleapis.com":
		fallthrough
	case "us-west2-pubsub.googleapis.com":
		fallthrough
	case "us-west3-pubsub.googleapis.com":
		fallthrough
	case "southamerica-east1-pubsub.googleapis.com":
		fallthrough
	case "northamerica-northeast1-pubsub.googleapis.com":
		fallthrough
	case "europe-west6-pubsub.googleapis.com":
		fallthrough
	case "europe-west4-pubsub.googleapis.com":
		fallthrough
	case "europe-west3-pubsub.googleapis.com":
		fallthrough
	case "europe-west2-pubsub.googleapis.com":
		fallthrough
	case "europe-west1-pubsub.googleapis.com":
		fallthrough
	case "europe-north1-pubsub.googleapis.com":
		fallthrough
	case "australia-southeast1-pubsub.googleapis.com":
		fallthrough
	case "asia-southeast1-pubsub.googleapis.com":
		fallthrough
	case "asia-south1-pubsub.googleapis.com":
		fallthrough
	case "asia-northeast3-pubsub.googleapis.com":
		fallthrough
	case "asia-northeast2-pubsub.googleapis.com":
		fallthrough
	case "asia-northeast1-pubsub.googleapis.com":
		fallthrough
	case "asia-east2-pubsub.googleapis.com":
		fallthrough
	case "asia-east1-pubsub.googleapis.com":
		*e = OutputGooglePubsubRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGooglePubsubRegion: %v", v)
	}
}

type OutputGooglePubsubType string

const (
	OutputGooglePubsubTypeGooglePubsub OutputGooglePubsubType = "google_pubsub"
)

func (e OutputGooglePubsubType) ToPointer() *OutputGooglePubsubType {
	return &e
}

func (e *OutputGooglePubsubType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "google_pubsub":
		*e = OutputGooglePubsubType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputGooglePubsubType: %v", v)
	}
}

type OutputGooglePubsub struct {
	// The maximum number of items the Google API should batch before it sends them to the topic.
	BatchSize *int64 `json:"batchSize,omitempty"`
	// The maximum amount of time, in milliseconds, that the Google API should wait to send a batch (if the Batch size is not reached).
	BatchTimeout *int64 `json:"batchTimeout,omitempty"`
	// If enabled, create topic if it does not exist.
	CreateTopic *bool `json:"createTopic,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Maximum time to wait before sending a batch (when Max batch size is not reached).
	FlushPeriodSec interface{} `json:"flushPeriodSec,omitempty"`
	// Google authentication method. Choose Auto to use environment variables PUBSUB_PROJECT and PUBSUB_CREDENTIALS..
	GoogleAuthMethod *OutputGooglePubsubAuthenticationMethod `json:"googleAuthMethod,omitempty"`
	// Unique ID for this output
	ID *string `json:"id,omitempty"`
	// The maximum number of in-progress API requests before backpressure is applied.
	MaxInProgress *int64 `json:"maxInProgress,omitempty"`
	// Maximum number of queued batches before blocking.
	MaxQueueSize *int64 `json:"maxQueueSize,omitempty"`
	// Maximum size (KB) of batches to send.
	MaxRecordSizeKB *int64 `json:"maxRecordSizeKB,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputGooglePubsubBackpressureBehavior `json:"onBackpressure,omitempty"`
	// If enabled, send events in the order they were added to the queue. For this to work correctly, the process receiving events must have ordering enabled.
	OrderedDelivery *bool `json:"orderedDelivery,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputGooglePubsubCompression `json:"pqCompress,omitempty"`
	PqControls *OutputGooglePubsubPqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputGooglePubsubQueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool `json:"pqStrictOrdering,omitempty"`
	// Region to publish messages to. Select 'default' to allow Google to auto-select the nearest region. When using ordered delivery, the selected region must be allowed by message storage policy.
	Region *OutputGooglePubsubRegion `json:"region,omitempty"`
	// Select (or create) a stored text secret
	Secret *string `json:"secret,omitempty"`
	// Contents of service account credentials (JSON keys) file downloaded from Google Cloud. To upload a file, click the upload button at this field's upper right. As an alternative, you can use environment variables (see [here](https://googleapis.dev/ruby/google-cloud-pubsub/latest/file.AUTHENTICATION.html)).
	ServiceAccountCredentials *string `json:"serviceAccountCredentials,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string `json:"systemFields,omitempty"`
	// ID of the topic to send events to.
	TopicName string                 `json:"topicName"`
	Type      OutputGooglePubsubType `json:"type"`
}

func (o *OutputGooglePubsub) GetBatchSize() *int64 {
	if o == nil {
		return nil
	}
	return o.BatchSize
}

func (o *OutputGooglePubsub) GetBatchTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.BatchTimeout
}

func (o *OutputGooglePubsub) GetCreateTopic() *bool {
	if o == nil {
		return nil
	}
	return o.CreateTopic
}

func (o *OutputGooglePubsub) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputGooglePubsub) GetFlushPeriodSec() interface{} {
	if o == nil {
		return nil
	}
	return o.FlushPeriodSec
}

func (o *OutputGooglePubsub) GetGoogleAuthMethod() *OutputGooglePubsubAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.GoogleAuthMethod
}

func (o *OutputGooglePubsub) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputGooglePubsub) GetMaxInProgress() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxInProgress
}

func (o *OutputGooglePubsub) GetMaxQueueSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxQueueSize
}

func (o *OutputGooglePubsub) GetMaxRecordSizeKB() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxRecordSizeKB
}

func (o *OutputGooglePubsub) GetOnBackpressure() *OutputGooglePubsubBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputGooglePubsub) GetOrderedDelivery() *bool {
	if o == nil {
		return nil
	}
	return o.OrderedDelivery
}

func (o *OutputGooglePubsub) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputGooglePubsub) GetPqCompress() *OutputGooglePubsubCompression {
	if o == nil {
		return nil
	}
	return o.PqCompress
}

func (o *OutputGooglePubsub) GetPqControls() *OutputGooglePubsubPqControls {
	if o == nil {
		return nil
	}
	return o.PqControls
}

func (o *OutputGooglePubsub) GetPqMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxFileSize
}

func (o *OutputGooglePubsub) GetPqMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxSize
}

func (o *OutputGooglePubsub) GetPqOnBackpressure() *OutputGooglePubsubQueueFullBehavior {
	if o == nil {
		return nil
	}
	return o.PqOnBackpressure
}

func (o *OutputGooglePubsub) GetPqPath() *string {
	if o == nil {
		return nil
	}
	return o.PqPath
}

func (o *OutputGooglePubsub) GetPqStrictOrdering() *bool {
	if o == nil {
		return nil
	}
	return o.PqStrictOrdering
}

func (o *OutputGooglePubsub) GetRegion() *OutputGooglePubsubRegion {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *OutputGooglePubsub) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *OutputGooglePubsub) GetServiceAccountCredentials() *string {
	if o == nil {
		return nil
	}
	return o.ServiceAccountCredentials
}

func (o *OutputGooglePubsub) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputGooglePubsub) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputGooglePubsub) GetTopicName() string {
	if o == nil {
		return ""
	}
	return o.TopicName
}

func (o *OutputGooglePubsub) GetType() OutputGooglePubsubType {
	if o == nil {
		return OutputGooglePubsubType("")
	}
	return o.Type
}
