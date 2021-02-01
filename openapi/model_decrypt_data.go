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

// DecryptData struct for DecryptData
type DecryptData struct {
	Decrypted string `json:"decrypted"`
}

// NewDecryptData instantiates a new DecryptData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecryptData(decrypted string, ) *DecryptData {
	this := DecryptData{}
	this.Decrypted = decrypted
	return &this
}

// NewDecryptDataWithDefaults instantiates a new DecryptData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecryptDataWithDefaults() *DecryptData {
	this := DecryptData{}
	return &this
}

// GetDecrypted returns the Decrypted field value
func (o *DecryptData) GetDecrypted() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Decrypted
}

// GetDecryptedOk returns a tuple with the Decrypted field value
// and a boolean to check if the value has been set.
func (o *DecryptData) GetDecryptedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Decrypted, true
}

// SetDecrypted sets field value
func (o *DecryptData) SetDecrypted(v string) {
	o.Decrypted = v
}

func (o DecryptData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["decrypted"] = o.Decrypted
	}
	return json.Marshal(toSerialize)
}

type NullableDecryptData struct {
	value *DecryptData
	isSet bool
}

func (v NullableDecryptData) Get() *DecryptData {
	return v.value
}

func (v *NullableDecryptData) Set(val *DecryptData) {
	v.value = val
	v.isSet = true
}

func (v NullableDecryptData) IsSet() bool {
	return v.isSet
}

func (v *NullableDecryptData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecryptData(val *DecryptData) *NullableDecryptData {
	return &NullableDecryptData{value: val, isSet: true}
}

func (v NullableDecryptData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecryptData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


