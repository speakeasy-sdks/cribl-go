// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OutputMinioAuthenticationMethod - AWS authentication method
type OutputMinioAuthenticationMethod string

const (
	OutputMinioAuthenticationMethodSecret OutputMinioAuthenticationMethod = "secret"
)

func (e OutputMinioAuthenticationMethod) ToPointer() *OutputMinioAuthenticationMethod {
	return &e
}

func (e *OutputMinioAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		*e = OutputMinioAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioAuthenticationMethod: %v", v)
	}
}

// OutputMinioCompress - Choose data compression format to apply before moving files to final destination.
type OutputMinioCompress string

const (
	OutputMinioCompressNone OutputMinioCompress = "none"
	OutputMinioCompressGzip OutputMinioCompress = "gzip"
)

func (e OutputMinioCompress) ToPointer() *OutputMinioCompress {
	return &e
}

func (e *OutputMinioCompress) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputMinioCompress(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioCompress: %v", v)
	}
}

// OutputMinioDataFormat - Format of the output data.
type OutputMinioDataFormat string

const (
	OutputMinioDataFormatParquet OutputMinioDataFormat = "parquet"
	OutputMinioDataFormatRaw     OutputMinioDataFormat = "raw"
	OutputMinioDataFormatJSON    OutputMinioDataFormat = "json"
)

func (e OutputMinioDataFormat) ToPointer() *OutputMinioDataFormat {
	return &e
}

func (e *OutputMinioDataFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "parquet":
		fallthrough
	case "raw":
		fallthrough
	case "json":
		*e = OutputMinioDataFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioDataFormat: %v", v)
	}
}

// OutputMinioObjectACL - Object ACL to assign to uploaded objects.
type OutputMinioObjectACL string

const (
	OutputMinioObjectACLPrivate                OutputMinioObjectACL = "private"
	OutputMinioObjectACLPublicRead             OutputMinioObjectACL = "public-read"
	OutputMinioObjectACLPublicReadWrite        OutputMinioObjectACL = "public-read-write"
	OutputMinioObjectACLAuthenticatedRead      OutputMinioObjectACL = "authenticated-read"
	OutputMinioObjectACLAwsExecRead            OutputMinioObjectACL = "aws-exec-read"
	OutputMinioObjectACLBucketOwnerRead        OutputMinioObjectACL = "bucket-owner-read"
	OutputMinioObjectACLBucketOwnerFullControl OutputMinioObjectACL = "bucket-owner-full-control"
)

func (e OutputMinioObjectACL) ToPointer() *OutputMinioObjectACL {
	return &e
}

func (e *OutputMinioObjectACL) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "private":
		fallthrough
	case "public-read":
		fallthrough
	case "public-read-write":
		fallthrough
	case "authenticated-read":
		fallthrough
	case "aws-exec-read":
		fallthrough
	case "bucket-owner-read":
		fallthrough
	case "bucket-owner-full-control":
		*e = OutputMinioObjectACL(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioObjectACL: %v", v)
	}
}

// OutputMinioBackpressureBehavior - Whether to block or drop events when all receivers are exerting backpressure.
type OutputMinioBackpressureBehavior string

const (
	OutputMinioBackpressureBehaviorBlock OutputMinioBackpressureBehavior = "block"
	OutputMinioBackpressureBehaviorDrop  OutputMinioBackpressureBehavior = "drop"
)

func (e OutputMinioBackpressureBehavior) ToPointer() *OutputMinioBackpressureBehavior {
	return &e
}

func (e *OutputMinioBackpressureBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputMinioBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioBackpressureBehavior: %v", v)
	}
}

// OutputMinioDataPageVersion - Serialization format of data pages. Note that not all reader implentations support Data page V2.
type OutputMinioDataPageVersion string

const (
	OutputMinioDataPageVersionDataPageV1 OutputMinioDataPageVersion = "DATA_PAGE_V1"
	OutputMinioDataPageVersionDataPageV2 OutputMinioDataPageVersion = "DATA_PAGE_V2"
)

func (e OutputMinioDataPageVersion) ToPointer() *OutputMinioDataPageVersion {
	return &e
}

func (e *OutputMinioDataPageVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DATA_PAGE_V1":
		fallthrough
	case "DATA_PAGE_V2":
		*e = OutputMinioDataPageVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioDataPageVersion: %v", v)
	}
}

// OutputMinioParquetVersion - Determines which data types are supported and how they are represented.
type OutputMinioParquetVersion string

const (
	OutputMinioParquetVersionParquet10 OutputMinioParquetVersion = "PARQUET_1_0"
	OutputMinioParquetVersionParquet24 OutputMinioParquetVersion = "PARQUET_2_4"
	OutputMinioParquetVersionParquet26 OutputMinioParquetVersion = "PARQUET_2_6"
)

func (e OutputMinioParquetVersion) ToPointer() *OutputMinioParquetVersion {
	return &e
}

func (e *OutputMinioParquetVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PARQUET_1_0":
		fallthrough
	case "PARQUET_2_4":
		fallthrough
	case "PARQUET_2_6":
		*e = OutputMinioParquetVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioParquetVersion: %v", v)
	}
}

// OutputMinioRegion - Region where the MinIO service/cluster is located
type OutputMinioRegion string

const (
	OutputMinioRegionUsEast1      OutputMinioRegion = "us-east-1"
	OutputMinioRegionUsEast2      OutputMinioRegion = "us-east-2"
	OutputMinioRegionUsWest1      OutputMinioRegion = "us-west-1"
	OutputMinioRegionUsWest2      OutputMinioRegion = "us-west-2"
	OutputMinioRegionAfSouth1     OutputMinioRegion = "af-south-1"
	OutputMinioRegionCaCentral1   OutputMinioRegion = "ca-central-1"
	OutputMinioRegionEuWest1      OutputMinioRegion = "eu-west-1"
	OutputMinioRegionEuCentral1   OutputMinioRegion = "eu-central-1"
	OutputMinioRegionEuWest2      OutputMinioRegion = "eu-west-2"
	OutputMinioRegionEuSouth1     OutputMinioRegion = "eu-south-1"
	OutputMinioRegionEuWest3      OutputMinioRegion = "eu-west-3"
	OutputMinioRegionEuNorth1     OutputMinioRegion = "eu-north-1"
	OutputMinioRegionApEast1      OutputMinioRegion = "ap-east-1"
	OutputMinioRegionApNortheast1 OutputMinioRegion = "ap-northeast-1"
	OutputMinioRegionApNortheast2 OutputMinioRegion = "ap-northeast-2"
	OutputMinioRegionApSoutheast1 OutputMinioRegion = "ap-southeast-1"
	OutputMinioRegionApSoutheast2 OutputMinioRegion = "ap-southeast-2"
	OutputMinioRegionApSouth1     OutputMinioRegion = "ap-south-1"
	OutputMinioRegionMeSouth1     OutputMinioRegion = "me-south-1"
	OutputMinioRegionSaEast1      OutputMinioRegion = "sa-east-1"
	OutputMinioRegionUsGovEast1   OutputMinioRegion = "us-gov-east-1"
	OutputMinioRegionUsGovWest1   OutputMinioRegion = "us-gov-west-1"
)

func (e OutputMinioRegion) ToPointer() *OutputMinioRegion {
	return &e
}

func (e *OutputMinioRegion) UnmarshalJSON(data []byte) error {
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
		*e = OutputMinioRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioRegion: %v", v)
	}
}

// OutputMinioServerSideEncryption - Server-side encryption for uploaded objects.
type OutputMinioServerSideEncryption string

const (
	OutputMinioServerSideEncryptionAes256 OutputMinioServerSideEncryption = "AES256"
)

func (e OutputMinioServerSideEncryption) ToPointer() *OutputMinioServerSideEncryption {
	return &e
}

func (e *OutputMinioServerSideEncryption) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AES256":
		*e = OutputMinioServerSideEncryption(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioServerSideEncryption: %v", v)
	}
}

// OutputMinioSignatureVersion - Signature version to use for signing MinIO requests.
type OutputMinioSignatureVersion string

const (
	OutputMinioSignatureVersionV2 OutputMinioSignatureVersion = "v2"
	OutputMinioSignatureVersionV4 OutputMinioSignatureVersion = "v4"
)

func (e OutputMinioSignatureVersion) ToPointer() *OutputMinioSignatureVersion {
	return &e
}

func (e *OutputMinioSignatureVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "v2":
		fallthrough
	case "v4":
		*e = OutputMinioSignatureVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioSignatureVersion: %v", v)
	}
}

// OutputMinioStorageClass - Storage class to select for uploaded objects.
type OutputMinioStorageClass string

const (
	OutputMinioStorageClassStandard          OutputMinioStorageClass = "STANDARD"
	OutputMinioStorageClassReducedRedundancy OutputMinioStorageClass = "REDUCED_REDUNDANCY"
)

func (e OutputMinioStorageClass) ToPointer() *OutputMinioStorageClass {
	return &e
}

func (e *OutputMinioStorageClass) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STANDARD":
		fallthrough
	case "REDUCED_REDUNDANCY":
		*e = OutputMinioStorageClass(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioStorageClass: %v", v)
	}
}

type OutputMinioType string

const (
	OutputMinioTypeMinio OutputMinioType = "minio"
)

func (e OutputMinioType) ToPointer() *OutputMinioType {
	return &e
}

func (e *OutputMinioType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "minio":
		*e = OutputMinioType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputMinioType: %v", v)
	}
}

type OutputMinio struct {
	// Append output's ID to staging location.
	AddIDToStagePath *bool `json:"addIdToStagePath,omitempty"`
	// Access key. This value can be a constant or a JavaScript expression(e.g., `${C.env.SOME_ACCESS_KEY}`).
	AwsAPIKey *string `json:"awsApiKey,omitempty"`
	// AWS authentication method
	AwsAuthenticationMethod *OutputMinioAuthenticationMethod `json:"awsAuthenticationMethod,omitempty"`
	// Select (or create) a stored secret that references your access key and secret key.
	AwsSecret *string `json:"awsSecret,omitempty"`
	// Secret key. This value can be a constant or a JavaScript expression(e.g., `${C.env.SOME_SECRET}`).
	AwsSecretKey *string `json:"awsSecretKey,omitempty"`
	// JavaScript expression to define the output filename prefix (can be constant).
	BaseFileName *string `json:"baseFileName,omitempty"`
	// Name of the destination MinIO bucket. This value can be a constant or a JavaScript expression that can only be evaluated at init time. E.g. referencing a Global Variable: `myBucket-${C.vars.myVar}`.
	Bucket string `json:"bucket"`
	// Choose data compression format to apply before moving files to final destination.
	Compress *OutputMinioCompress `json:"compress,omitempty"`
	// Root directory to prepend to path before uploading. Enter a constant, or a JS expression enclosed in quotes or backticks.
	DestPath string `json:"destPath"`
	// How often (secs) to clean-up empty directories when 'Remove Staging Dirs' is enabled.
	EmptyDirCleanupSec *int64 `json:"emptyDirCleanupSec,omitempty"`
	// MinIO service url (e.g. http://minioHost:9000)
	Endpoint string `json:"endpoint"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// JavaScript expression to define the output filename suffix (can be constant).  The `__format` variable refers to the value of the `Data format` field (`json` or `raw`).  The `__compression` field refers to the kind of compression being used (`none` or `gzip`)
	FileNameSuffix *string `json:"fileNameSuffix,omitempty"`
	// Format of the output data.
	Format *OutputMinioDataFormat `json:"format,omitempty"`
	// Unique ID for this output
	ID *string `json:"id,omitempty"`
	// Maximum number of parts to upload in parallel per file. Minimum part size is 5MB.
	MaxConcurrentFileParts *int64 `json:"maxConcurrentFileParts,omitempty"`
	// Maximum amount of time to keep inactive files open. Files open for longer than this will be closed and moved to final output location.
	MaxFileIdleTimeSec *int64 `json:"maxFileIdleTimeSec,omitempty"`
	// Maximum amount of time to write to a file. Files open for longer than this will be closed and moved to final output location.
	MaxFileOpenTimeSec *int64 `json:"maxFileOpenTimeSec,omitempty"`
	// Maximum uncompressed output file size. Files of this size will be closed and moved to final output location.
	MaxFileSizeMB *int64 `json:"maxFileSizeMB,omitempty"`
	// Maximum number of files to keep open concurrently. When exceeded, @{product} will close the oldest open files and move them to the final output location.
	MaxOpenFiles *int64 `json:"maxOpenFiles,omitempty"`
	// Object ACL to assign to uploaded objects.
	ObjectACL *OutputMinioObjectACL `json:"objectACL,omitempty"`
	// Whether to block or drop events when all receivers are exerting backpressure.
	OnBackpressure *OutputMinioBackpressureBehavior `json:"onBackpressure,omitempty"`
	// Serialization format of data pages. Note that not all reader implentations support Data page V2.
	ParquetDataPageVersion *OutputMinioDataPageVersion `json:"parquetDataPageVersion,omitempty"`
	// Ideal memory size for page segments. E.g., 1MB or 128MB. Generally, lower values improve reading speed, while higher values improve compression. Imposes a target, not a strict limit; the final size of a row group may be larger or smaller.
	ParquetPageSize *string `json:"parquetPageSize,omitempty"`
	// Ideal memory size for row group segments. E.g., 128MB or 1GB. Affects memory use when writing. Imposes a target, not a strict limit; the final size of a row group may be larger or smaller.
	ParquetRowGroupSize *string `json:"parquetRowGroupSize,omitempty"`
	// Determines which data types are supported and how they are represented.
	ParquetVersion *OutputMinioParquetVersion `json:"parquetVersion,omitempty"`
	// JS expression defining how files are partitioned and organized. Default is date-based. If blank, Stream will fall back to the event's __partition field value – if present – otherwise to each location's root directory.
	PartitionExpr *string `json:"partitionExpr,omitempty"`
	// Pipeline to process data before sending out to this output.
	Pipeline *string `json:"pipeline,omitempty"`
	// Region where the MinIO service/cluster is located
	Region *OutputMinioRegion `json:"region,omitempty"`
	// Whether to reject certificates that cannot be verified against a valid CA (e.g., self-signed certificates).
	RejectUnauthorized *bool `json:"rejectUnauthorized,omitempty"`
	// Remove empty staging directories after moving files.
	RemoveEmptyDirs *bool `json:"removeEmptyDirs,omitempty"`
	// Whether to reuse connections between requests, which can improve performance.
	ReuseConnections *bool `json:"reuseConnections,omitempty"`
	// Server-side encryption for uploaded objects.
	ServerSideEncryption *OutputMinioServerSideEncryption `json:"serverSideEncryption,omitempty"`
	// To log rows that @{product} skips due to data mismatch, first set logging to Debug, then toggle this on. Logs up to 20 unique rows.
	ShouldLogInvalidRows *bool `json:"shouldLogInvalidRows,omitempty"`
	// Signature version to use for signing MinIO requests.
	SignatureVersion *OutputMinioSignatureVersion `json:"signatureVersion,omitempty"`
	Spacer           *string                      `json:"spacer,omitempty"`
	// Filesystem location in which to buffer files, before compressing and moving to final destination. Use performant stable storage.
	StagePath string `json:"stagePath"`
	// Storage class to select for uploaded objects.
	StorageClass *OutputMinioStorageClass `json:"storageClass,omitempty"`
	// Add tags for filtering and grouping in @{product}.
	Streamtags []string `json:"streamtags,omitempty"`
	// Set of fields to automatically add to events using this output. E.g.: cribl_pipe, c*. Wildcards supported.
	SystemFields []string         `json:"systemFields,omitempty"`
	Type         *OutputMinioType `json:"type,omitempty"`
}

func (o *OutputMinio) GetAddIDToStagePath() *bool {
	if o == nil {
		return nil
	}
	return o.AddIDToStagePath
}

func (o *OutputMinio) GetAwsAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsAPIKey
}

func (o *OutputMinio) GetAwsAuthenticationMethod() *OutputMinioAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AwsAuthenticationMethod
}

func (o *OutputMinio) GetAwsSecret() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecret
}

func (o *OutputMinio) GetAwsSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecretKey
}

func (o *OutputMinio) GetBaseFileName() *string {
	if o == nil {
		return nil
	}
	return o.BaseFileName
}

func (o *OutputMinio) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *OutputMinio) GetCompress() *OutputMinioCompress {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *OutputMinio) GetDestPath() string {
	if o == nil {
		return ""
	}
	return o.DestPath
}

func (o *OutputMinio) GetEmptyDirCleanupSec() *int64 {
	if o == nil {
		return nil
	}
	return o.EmptyDirCleanupSec
}

func (o *OutputMinio) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *OutputMinio) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputMinio) GetFileNameSuffix() *string {
	if o == nil {
		return nil
	}
	return o.FileNameSuffix
}

func (o *OutputMinio) GetFormat() *OutputMinioDataFormat {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *OutputMinio) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputMinio) GetMaxConcurrentFileParts() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxConcurrentFileParts
}

func (o *OutputMinio) GetMaxFileIdleTimeSec() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxFileIdleTimeSec
}

func (o *OutputMinio) GetMaxFileOpenTimeSec() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxFileOpenTimeSec
}

func (o *OutputMinio) GetMaxFileSizeMB() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxFileSizeMB
}

func (o *OutputMinio) GetMaxOpenFiles() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxOpenFiles
}

func (o *OutputMinio) GetObjectACL() *OutputMinioObjectACL {
	if o == nil {
		return nil
	}
	return o.ObjectACL
}

func (o *OutputMinio) GetOnBackpressure() *OutputMinioBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputMinio) GetParquetDataPageVersion() *OutputMinioDataPageVersion {
	if o == nil {
		return nil
	}
	return o.ParquetDataPageVersion
}

func (o *OutputMinio) GetParquetPageSize() *string {
	if o == nil {
		return nil
	}
	return o.ParquetPageSize
}

func (o *OutputMinio) GetParquetRowGroupSize() *string {
	if o == nil {
		return nil
	}
	return o.ParquetRowGroupSize
}

func (o *OutputMinio) GetParquetVersion() *OutputMinioParquetVersion {
	if o == nil {
		return nil
	}
	return o.ParquetVersion
}

func (o *OutputMinio) GetPartitionExpr() *string {
	if o == nil {
		return nil
	}
	return o.PartitionExpr
}

func (o *OutputMinio) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputMinio) GetRegion() *OutputMinioRegion {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *OutputMinio) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputMinio) GetRemoveEmptyDirs() *bool {
	if o == nil {
		return nil
	}
	return o.RemoveEmptyDirs
}

func (o *OutputMinio) GetReuseConnections() *bool {
	if o == nil {
		return nil
	}
	return o.ReuseConnections
}

func (o *OutputMinio) GetServerSideEncryption() *OutputMinioServerSideEncryption {
	if o == nil {
		return nil
	}
	return o.ServerSideEncryption
}

func (o *OutputMinio) GetShouldLogInvalidRows() *bool {
	if o == nil {
		return nil
	}
	return o.ShouldLogInvalidRows
}

func (o *OutputMinio) GetSignatureVersion() *OutputMinioSignatureVersion {
	if o == nil {
		return nil
	}
	return o.SignatureVersion
}

func (o *OutputMinio) GetSpacer() *string {
	if o == nil {
		return nil
	}
	return o.Spacer
}

func (o *OutputMinio) GetStagePath() string {
	if o == nil {
		return ""
	}
	return o.StagePath
}

func (o *OutputMinio) GetStorageClass() *OutputMinioStorageClass {
	if o == nil {
		return nil
	}
	return o.StorageClass
}

func (o *OutputMinio) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputMinio) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputMinio) GetType() *OutputMinioType {
	if o == nil {
		return nil
	}
	return o.Type
}
