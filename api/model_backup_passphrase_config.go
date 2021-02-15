/*
 * NetHSM
 *
 * All endpoints expect exactly the specified JSON. Additional properties will cause a Bad Request Error (400). All HTTP errors contain a JSON structure with an explanation of type string. All <a href=\"https://tools.ietf.org/html/rfc4648#section-4\">base64</a> encoded values are Big Endian.
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BackupPassphraseConfig struct for BackupPassphraseConfig
type BackupPassphraseConfig struct {
	Passphrase           string `json:"passphrase"`
	AdditionalProperties map[string]interface{}
}

type _BackupPassphraseConfig BackupPassphraseConfig

// NewBackupPassphraseConfig instantiates a new BackupPassphraseConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupPassphraseConfig(passphrase string) *BackupPassphraseConfig {
	this := BackupPassphraseConfig{}
	this.Passphrase = passphrase
	return &this
}

// NewBackupPassphraseConfigWithDefaults instantiates a new BackupPassphraseConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupPassphraseConfigWithDefaults() *BackupPassphraseConfig {
	this := BackupPassphraseConfig{}
	return &this
}

// GetPassphrase returns the Passphrase field value
func (o *BackupPassphraseConfig) GetPassphrase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value
// and a boolean to check if the value has been set.
func (o *BackupPassphraseConfig) GetPassphraseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Passphrase, true
}

// SetPassphrase sets field value
func (o *BackupPassphraseConfig) SetPassphrase(v string) {
	o.Passphrase = v
}

func (o BackupPassphraseConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["passphrase"] = o.Passphrase
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BackupPassphraseConfig) UnmarshalJSON(bytes []byte) (err error) {
	varBackupPassphraseConfig := _BackupPassphraseConfig{}

	if err = json.Unmarshal(bytes, &varBackupPassphraseConfig); err == nil {
		*o = BackupPassphraseConfig(varBackupPassphraseConfig)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "passphrase")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBackupPassphraseConfig struct {
	value *BackupPassphraseConfig
	isSet bool
}

func (v NullableBackupPassphraseConfig) Get() *BackupPassphraseConfig {
	return v.value
}

func (v *NullableBackupPassphraseConfig) Set(val *BackupPassphraseConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupPassphraseConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupPassphraseConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupPassphraseConfig(val *BackupPassphraseConfig) *NullableBackupPassphraseConfig {
	return &NullableBackupPassphraseConfig{value: val, isSet: true}
}

func (v NullableBackupPassphraseConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupPassphraseConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
