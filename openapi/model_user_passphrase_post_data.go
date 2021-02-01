/*
 * NetHSM
 *
 * All endpoints expect exactly the specified JSON. Additional properties will cause a Bad Request Error (400). All HTTP errors contain a JSON structure with an explanation of type string. All <a href=\"https://tools.ietf.org/html/rfc4648#section-4\">base64</a> encoded values are Big Endian.
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UserPassphrasePostData struct for UserPassphrasePostData
type UserPassphrasePostData struct {
	Passphrase string `json:"passphrase"`
}

// NewUserPassphrasePostData instantiates a new UserPassphrasePostData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPassphrasePostData(passphrase string, ) *UserPassphrasePostData {
	this := UserPassphrasePostData{}
	this.Passphrase = passphrase
	return &this
}

// NewUserPassphrasePostDataWithDefaults instantiates a new UserPassphrasePostData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPassphrasePostDataWithDefaults() *UserPassphrasePostData {
	this := UserPassphrasePostData{}
	return &this
}

// GetPassphrase returns the Passphrase field value
func (o *UserPassphrasePostData) GetPassphrase() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value
// and a boolean to check if the value has been set.
func (o *UserPassphrasePostData) GetPassphraseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Passphrase, true
}

// SetPassphrase sets field value
func (o *UserPassphrasePostData) SetPassphrase(v string) {
	o.Passphrase = v
}

func (o UserPassphrasePostData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["passphrase"] = o.Passphrase
	}
	return json.Marshal(toSerialize)
}

type NullableUserPassphrasePostData struct {
	value *UserPassphrasePostData
	isSet bool
}

func (v NullableUserPassphrasePostData) Get() *UserPassphrasePostData {
	return v.value
}

func (v *NullableUserPassphrasePostData) Set(val *UserPassphrasePostData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPassphrasePostData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPassphrasePostData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPassphrasePostData(val *UserPassphrasePostData) *NullableUserPassphrasePostData {
	return &NullableUserPassphrasePostData{value: val, isSet: true}
}

func (v NullableUserPassphrasePostData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPassphrasePostData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


