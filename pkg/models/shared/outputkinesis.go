// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputKinesisAuthenticationMethod - AWS authentication method. Choose Auto to use IAM roles.
type OutputKinesisAuthenticationMethod string

const (
	OutputKinesisAuthenticationMethodSecret OutputKinesisAuthenticationMethod = "secret"
	OutputKinesisAuthenticationMethodManual OutputKinesisAuthenticationMethod = "manual"
)

func (e OutputKinesisAuthenticationMethod) ToPointer() *OutputKinesisAuthenticationMethod {
	return &e
}

func (e *OutputKinesisAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		fallthrough
	case "manual":
		*e = OutputKinesisAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKinesisAuthenticationMethod: %v", v)
	}
}

// OutputKinesisBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputKinesisBackpressureBehavior string

const (
	OutputKinesisBackpressureBehaviorQueue OutputKinesisBackpressureBehavior = "queue"
	OutputKinesisBackpressureBehaviorDrop  OutputKinesisBackpressureBehavior = "drop"
	OutputKinesisBackpressureBehaviorBlock OutputKinesisBackpressureBehavior = "block"
)

func (e OutputKinesisBackpressureBehavior) ToPointer() *OutputKinesisBackpressureBehavior {
	return &e
}

func (e *OutputKinesisBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputKinesisBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKinesisBackpressureBehavior: %v", v)
	}
}

// OutputKinesisCompression - Codec to use to compress the persisted data.
type OutputKinesisCompression string

const (
	OutputKinesisCompressionNone OutputKinesisCompression = "none"
	OutputKinesisCompressionGzip OutputKinesisCompression = "gzip"
)

func (e OutputKinesisCompression) ToPointer() *OutputKinesisCompression {
	return &e
}

func (e *OutputKinesisCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputKinesisCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKinesisCompression: %v", v)
	}
}

type OutputKinesisPqControls struct {
}

// OutputKinesisQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputKinesisQueueFullBehavior string

const (
	OutputKinesisQueueFullBehaviorBlock OutputKinesisQueueFullBehavior = "block"
	OutputKinesisQueueFullBehaviorDrop  OutputKinesisQueueFullBehavior = "drop"
)

func (e OutputKinesisQueueFullBehavior) ToPointer() *OutputKinesisQueueFullBehavior {
	return &e
}

func (e *OutputKinesisQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputKinesisQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKinesisQueueFullBehavior: %v", v)
	}
}

// OutputKinesisRegion - Region where the Kinesis stream is located
type OutputKinesisRegion string

const (
	OutputKinesisRegionUsEast1      OutputKinesisRegion = "us-east-1"
	OutputKinesisRegionUsEast2      OutputKinesisRegion = "us-east-2"
	OutputKinesisRegionUsWest1      OutputKinesisRegion = "us-west-1"
	OutputKinesisRegionUsWest2      OutputKinesisRegion = "us-west-2"
	OutputKinesisRegionAfSouth1     OutputKinesisRegion = "af-south-1"
	OutputKinesisRegionCaCentral1   OutputKinesisRegion = "ca-central-1"
	OutputKinesisRegionEuWest1      OutputKinesisRegion = "eu-west-1"
	OutputKinesisRegionEuCentral1   OutputKinesisRegion = "eu-central-1"
	OutputKinesisRegionEuWest2      OutputKinesisRegion = "eu-west-2"
	OutputKinesisRegionEuSouth1     OutputKinesisRegion = "eu-south-1"
	OutputKinesisRegionEuWest3      OutputKinesisRegion = "eu-west-3"
	OutputKinesisRegionEuNorth1     OutputKinesisRegion = "eu-north-1"
	OutputKinesisRegionApEast1      OutputKinesisRegion = "ap-east-1"
	OutputKinesisRegionApNortheast1 OutputKinesisRegion = "ap-northeast-1"
	OutputKinesisRegionApNortheast2 OutputKinesisRegion = "ap-northeast-2"
	OutputKinesisRegionApSoutheast1 OutputKinesisRegion = "ap-southeast-1"
	OutputKinesisRegionApSoutheast2 OutputKinesisRegion = "ap-southeast-2"
	OutputKinesisRegionApSouth1     OutputKinesisRegion = "ap-south-1"
	OutputKinesisRegionMeSouth1     OutputKinesisRegion = "me-south-1"
	OutputKinesisRegionSaEast1      OutputKinesisRegion = "sa-east-1"
	OutputKinesisRegionUsGovEast1   OutputKinesisRegion = "us-gov-east-1"
	OutputKinesisRegionUsGovWest1   OutputKinesisRegion = "us-gov-west-1"
)

func (e OutputKinesisRegion) ToPointer() *OutputKinesisRegion {
	return &e
}

func (e *OutputKinesisRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "us-east-1":
		fallthrough
	case "us-east-2":
		fallthrough
	case "us-west-1":
		fallthrough
	case "us-west-2":
		fallthrough
	case "af-south-1":
		fallthrough
	case "ca-central-1":
		fallthrough
	case "eu-west-1":
		fallthrough
	case "eu-central-1":
		fallthrough
	case "eu-west-2":
		fallthrough
	case "eu-south-1":
		fallthrough
	case "eu-west-3":
		fallthrough
	case "eu-north-1":
		fallthrough
	case "ap-east-1":
		fallthrough
	case "ap-northeast-1":
		fallthrough
	case "ap-northeast-2":
		fallthrough
	case "ap-southeast-1":
		fallthrough
	case "ap-southeast-2":
		fallthrough
	case "ap-south-1":
		fallthrough
	case "me-south-1":
		fallthrough
	case "sa-east-1":
		fallthrough
	case "us-gov-east-1":
		fallthrough
	case "us-gov-west-1":
		*e = OutputKinesisRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKinesisRegion: %v", v)
	}
}

// OutputKinesisSignatureVersion - Signature version to use for signing Kinesis stream requests.
type OutputKinesisSignatureVersion string

const (
	OutputKinesisSignatureVersionV2 OutputKinesisSignatureVersion = "v2"
	OutputKinesisSignatureVersionV4 OutputKinesisSignatureVersion = "v4"
)

func (e OutputKinesisSignatureVersion) ToPointer() *OutputKinesisSignatureVersion {
	return &e
}

func (e *OutputKinesisSignatureVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "v2":
		fallthrough
	case "v4":
		*e = OutputKinesisSignatureVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKinesisSignatureVersion: %v", v)
	}
}

type OutputKinesisType string

const (
	OutputKinesisTypeKinesis OutputKinesisType = "kinesis"
)

func (e OutputKinesisType) ToPointer() *OutputKinesisType {
	return &e
}

func (e *OutputKinesisType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "kinesis":
		*e = OutputKinesisType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputKinesisType: %v", v)
	}
}

type OutputKinesis struct {
	// Amazon Resource Name (ARN) of the role to assume
	AssumeRoleArn *string `json:"assumeRoleArn,omitempty"`
	// External ID to use when assuming role
	AssumeRoleExternalID *string `json:"assumeRoleExternalId,omitempty"`
	// Access key
	AwsAPIKey *string `json:"awsApiKey,omitempty"`
	// AWS authentication method. Choose Auto to use IAM roles.
	AwsAuthenticationMethod *OutputKinesisAuthenticationMethod `json:"awsAuthenticationMethod,omitempty"`
	// Select (or create) a stored secret that references your access key and secret key.
	AwsSecret *string `json:"awsSecret,omitempty"`
	// Secret key
	AwsSecretKey *string `json:"awsSecretKey,omitempty"`
	// Maximum number of ongoing put requests before blocking.
	Concurrency *int64 `json:"concurrency,omitempty"`
	// Use Assume Role credentials to access Kinesis stream
	EnableAssumeRole *bool `json:"enableAssumeRole,omitempty"`
	// Kinesis stream service endpoint. If empty, defaults to AWS' Region-specific endpoint. Otherwise, it must point to Kinesis stream-compatible endpoint.
	Endpoint *string `json:"endpoint,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Max record size.
	FlushPeriodSec *int64 `json:"flushPeriodSec,omitempty"`
	// Unique ID for this output
	ID *string `json:"id,omitempty"`
	// Maximum size (KB) of each individual record before compression. For non-compressible data 1MB is the max recommended size
	MaxRecordSizeKB *int64 `json:"maxRecordSizeKB,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputKinesisBackpressureBehavior `json:"onBackpressure,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputKinesisCompression `json:"pqCompress,omitempty"`
	PqControls *OutputKinesisPqControls  `json:"pqControls,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `json:"pqMaxFileSize,omitempty"`
	// The maximum amount of disk space the queue is allowed to consume. Once reached, the system stops queueing and applies the fallback Queue-full behavior. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `json:"pqMaxSize,omitempty"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputKinesisQueueFullBehavior `json:"pqOnBackpressure,omitempty"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `json:"pqPath,omitempty"`
	// Toggle this off to forward new events to receiver(s) before queue is flushed. Otherwise, default drain behavior is FIFO (first in, first out).
	PqStrictOrdering *bool `json:"pqStrictOrdering,omitempty"`
	// Region where the Kinesis stream is located
	Region OutputKinesisRegion `json:"region"`
	// Whether to reject certificates that cannot be verified against a valid CA (e.g., self-signed certificates).
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// Whether to reuse connections between requests, which can improve performance.
	ReuseConnections *bool `json:"reuseConnections,omitempty"`
	// Signature version to use for signing Kinesis stream requests.
	SignatureVersion *OutputKinesisSignatureVersion `json:"signatureVersion,omitempty"`
	// Kinesis stream name to send events to.
	StreamName string `json:"streamName"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string           `json:"systemFields,omitempty"`
	Type         *OutputKinesisType `json:"type,omitempty"`
}

func (o *OutputKinesis) GetAssumeRoleArn() *string {
	if o == nil {
		return nil
	}
	return o.AssumeRoleArn
}

func (o *OutputKinesis) GetAssumeRoleExternalID() *string {
	if o == nil {
		return nil
	}
	return o.AssumeRoleExternalID
}

func (o *OutputKinesis) GetAwsAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsAPIKey
}

func (o *OutputKinesis) GetAwsAuthenticationMethod() *OutputKinesisAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AwsAuthenticationMethod
}

func (o *OutputKinesis) GetAwsSecret() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecret
}

func (o *OutputKinesis) GetAwsSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecretKey
}

func (o *OutputKinesis) GetConcurrency() *int64 {
	if o == nil {
		return nil
	}
	return o.Concurrency
}

func (o *OutputKinesis) GetEnableAssumeRole() *bool {
	if o == nil {
		return nil
	}
	return o.EnableAssumeRole
}

func (o *OutputKinesis) GetEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.Endpoint
}

func (o *OutputKinesis) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputKinesis) GetFlushPeriodSec() *int64 {
	if o == nil {
		return nil
	}
	return o.FlushPeriodSec
}

func (o *OutputKinesis) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputKinesis) GetMaxRecordSizeKB() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxRecordSizeKB
}

func (o *OutputKinesis) GetOnBackpressure() *OutputKinesisBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputKinesis) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputKinesis) GetPqCompress() *OutputKinesisCompression {
	if o == nil {
		return nil
	}
	return o.PqCompress
}

func (o *OutputKinesis) GetPqControls() *OutputKinesisPqControls {
	if o == nil {
		return nil
	}
	return o.PqControls
}

func (o *OutputKinesis) GetPqMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxFileSize
}

func (o *OutputKinesis) GetPqMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxSize
}

func (o *OutputKinesis) GetPqOnBackpressure() *OutputKinesisQueueFullBehavior {
	if o == nil {
		return nil
	}
	return o.PqOnBackpressure
}

func (o *OutputKinesis) GetPqPath() *string {
	if o == nil {
		return nil
	}
	return o.PqPath
}

func (o *OutputKinesis) GetPqStrictOrdering() *bool {
	if o == nil {
		return nil
	}
	return o.PqStrictOrdering
}

func (o *OutputKinesis) GetRegion() OutputKinesisRegion {
	if o == nil {
		return OutputKinesisRegion("")
	}
	return o.Region
}

func (o *OutputKinesis) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputKinesis) GetReuseConnections() *bool {
	if o == nil {
		return nil
	}
	return o.ReuseConnections
}

func (o *OutputKinesis) GetSignatureVersion() *OutputKinesisSignatureVersion {
	if o == nil {
		return nil
	}
	return o.SignatureVersion
}

func (o *OutputKinesis) GetStreamName() string {
	if o == nil {
		return ""
	}
	return o.StreamName
}

func (o *OutputKinesis) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputKinesis) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputKinesis) GetType() *OutputKinesisType {
	if o == nil {
		return nil
	}
	return o.Type
}
