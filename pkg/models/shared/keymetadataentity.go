// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type KeyMetadataEntityEncryptionAlgorithm string

const (
	KeyMetadataEntityEncryptionAlgorithmAes256Cbc KeyMetadataEntityEncryptionAlgorithm = "aes-256-cbc"
	KeyMetadataEntityEncryptionAlgorithmAes256Gcm KeyMetadataEntityEncryptionAlgorithm = "aes-256-gcm"
)

func (e KeyMetadataEntityEncryptionAlgorithm) ToPointer() *KeyMetadataEntityEncryptionAlgorithm {
	return &e
}

func (e *KeyMetadataEntityEncryptionAlgorithm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "aes-256-cbc":
		fallthrough
	case "aes-256-gcm":
		*e = KeyMetadataEntityEncryptionAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KeyMetadataEntityEncryptionAlgorithm: %v", v)
	}
}

type KeyMetadataEntityKMSForThisKey string

const (
	KeyMetadataEntityKMSForThisKeyLocal KeyMetadataEntityKMSForThisKey = "local"
)

func (e KeyMetadataEntityKMSForThisKey) ToPointer() *KeyMetadataEntityKMSForThisKey {
	return &e
}

func (e *KeyMetadataEntityKMSForThisKey) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "local":
		*e = KeyMetadataEntityKMSForThisKey(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KeyMetadataEntityKMSForThisKey: %v", v)
	}
}

// KeyMetadataEntity - New KeyMetadataEntity object
type KeyMetadataEntity struct {
	Algorithm   KeyMetadataEntityEncryptionAlgorithm `json:"algorithm"`
	CipherKey   *string                              `json:"cipherKey,omitempty"`
	Created     *int64                               `json:"created,omitempty"`
	Description *string                              `json:"description,omitempty"`
	Expires     *int64                               `json:"expires,omitempty"`
	KeyID       string                               `json:"keyId"`
	Keyclass    int64                                `json:"keyclass"`
	Kms         KeyMetadataEntityKMSForThisKey       `json:"kms"`
	PlainKey    *string                              `json:"plainKey,omitempty"`
	// Seed encryption with a [nonce](https://en.wikipedia.org/wiki/Cryptographic_nonce) to make the key more random and unique. Must be toggled on with the aes-256-gcm algorithm.
	UseIV *bool `json:"useIV,omitempty"`
}

func (o *KeyMetadataEntity) GetAlgorithm() KeyMetadataEntityEncryptionAlgorithm {
	if o == nil {
		return KeyMetadataEntityEncryptionAlgorithm("")
	}
	return o.Algorithm
}

func (o *KeyMetadataEntity) GetCipherKey() *string {
	if o == nil {
		return nil
	}
	return o.CipherKey
}

func (o *KeyMetadataEntity) GetCreated() *int64 {
	if o == nil {
		return nil
	}
	return o.Created
}

func (o *KeyMetadataEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *KeyMetadataEntity) GetExpires() *int64 {
	if o == nil {
		return nil
	}
	return o.Expires
}

func (o *KeyMetadataEntity) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *KeyMetadataEntity) GetKeyclass() int64 {
	if o == nil {
		return 0
	}
	return o.Keyclass
}

func (o *KeyMetadataEntity) GetKms() KeyMetadataEntityKMSForThisKey {
	if o == nil {
		return KeyMetadataEntityKMSForThisKey("")
	}
	return o.Kms
}

func (o *KeyMetadataEntity) GetPlainKey() *string {
	if o == nil {
		return nil
	}
	return o.PlainKey
}

func (o *KeyMetadataEntity) GetUseIV() *bool {
	if o == nil {
		return nil
	}
	return o.UseIV
}
