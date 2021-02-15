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
	"fmt"
)

// DecryptMode the model 'DecryptMode'
type DecryptMode string

// List of DecryptMode
const (
	DECRYPTMODE_RAW         DecryptMode = "RAW"
	DECRYPTMODE_PKCS1       DecryptMode = "PKCS1"
	DECRYPTMODE_OAEP_MD5    DecryptMode = "OAEP_MD5"
	DECRYPTMODE_OAEP_SHA1   DecryptMode = "OAEP_SHA1"
	DECRYPTMODE_OAEP_SHA224 DecryptMode = "OAEP_SHA224"
	DECRYPTMODE_OAEP_SHA256 DecryptMode = "OAEP_SHA256"
	DECRYPTMODE_OAEP_SHA384 DecryptMode = "OAEP_SHA384"
	DECRYPTMODE_OAEP_SHA512 DecryptMode = "OAEP_SHA512"
)

func (v *DecryptMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DecryptMode(value)
	for _, existing := range []DecryptMode{"RAW", "PKCS1", "OAEP_MD5", "OAEP_SHA1", "OAEP_SHA224", "OAEP_SHA256", "OAEP_SHA384", "OAEP_SHA512"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DecryptMode", value)
}

// Ptr returns reference to DecryptMode value
func (v DecryptMode) Ptr() *DecryptMode {
	return &v
}

type NullableDecryptMode struct {
	value *DecryptMode
	isSet bool
}

func (v NullableDecryptMode) Get() *DecryptMode {
	return v.value
}

func (v *NullableDecryptMode) Set(val *DecryptMode) {
	v.value = val
	v.isSet = true
}

func (v NullableDecryptMode) IsSet() bool {
	return v.isSet
}

func (v *NullableDecryptMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecryptMode(val *DecryptMode) *NullableDecryptMode {
	return &NullableDecryptMode{value: val, isSet: true}
}

func (v NullableDecryptMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecryptMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
