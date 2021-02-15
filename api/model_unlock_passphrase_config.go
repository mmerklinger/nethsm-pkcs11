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

// UnlockPassphraseConfig struct for UnlockPassphraseConfig
type UnlockPassphraseConfig struct {
	Passphrase           string `json:"passphrase"`
	AdditionalProperties map[string]interface{}
}

type _UnlockPassphraseConfig UnlockPassphraseConfig

// NewUnlockPassphraseConfig instantiates a new UnlockPassphraseConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnlockPassphraseConfig(passphrase string) *UnlockPassphraseConfig {
	this := UnlockPassphraseConfig{}
	this.Passphrase = passphrase
	return &this
}

// NewUnlockPassphraseConfigWithDefaults instantiates a new UnlockPassphraseConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnlockPassphraseConfigWithDefaults() *UnlockPassphraseConfig {
	this := UnlockPassphraseConfig{}
	return &this
}

// GetPassphrase returns the Passphrase field value
func (o *UnlockPassphraseConfig) GetPassphrase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value
// and a boolean to check if the value has been set.
func (o *UnlockPassphraseConfig) GetPassphraseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Passphrase, true
}

// SetPassphrase sets field value
func (o *UnlockPassphraseConfig) SetPassphrase(v string) {
	o.Passphrase = v
}

func (o UnlockPassphraseConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["passphrase"] = o.Passphrase
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UnlockPassphraseConfig) UnmarshalJSON(bytes []byte) (err error) {
	varUnlockPassphraseConfig := _UnlockPassphraseConfig{}

	if err = json.Unmarshal(bytes, &varUnlockPassphraseConfig); err == nil {
		*o = UnlockPassphraseConfig(varUnlockPassphraseConfig)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "passphrase")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUnlockPassphraseConfig struct {
	value *UnlockPassphraseConfig
	isSet bool
}

func (v NullableUnlockPassphraseConfig) Get() *UnlockPassphraseConfig {
	return v.value
}

func (v *NullableUnlockPassphraseConfig) Set(val *UnlockPassphraseConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUnlockPassphraseConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUnlockPassphraseConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnlockPassphraseConfig(val *UnlockPassphraseConfig) *NullableUnlockPassphraseConfig {
	return &NullableUnlockPassphraseConfig{value: val, isSet: true}
}

func (v NullableUnlockPassphraseConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnlockPassphraseConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
