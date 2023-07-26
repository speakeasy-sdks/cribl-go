// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type IAWSKMSServiceConfig struct {
	KmsKeyArn string `json:"kmsKeyArn"`
	Region    string `json:"region"`
}

func (o *IAWSKMSServiceConfig) GetKmsKeyArn() string {
	if o == nil {
		return ""
	}
	return o.KmsKeyArn
}

func (o *IAWSKMSServiceConfig) GetRegion() string {
	if o == nil {
		return ""
	}
	return o.Region
}