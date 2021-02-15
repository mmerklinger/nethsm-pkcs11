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

// KeyMechanism the model 'KeyMechanism'
type KeyMechanism string

// List of KeyMechanism
const (
	KEYMECHANISM_RSA_DECRYPTION_RAW         KeyMechanism = "RSA_Decryption_RAW"
	KEYMECHANISM_RSA_DECRYPTION_PKCS1       KeyMechanism = "RSA_Decryption_PKCS1"
	KEYMECHANISM_RSA_DECRYPTION_OAEP_MD5    KeyMechanism = "RSA_Decryption_OAEP_MD5"
	KEYMECHANISM_RSA_DECRYPTION_OAEP_SHA1   KeyMechanism = "RSA_Decryption_OAEP_SHA1"
	KEYMECHANISM_RSA_DECRYPTION_OAEP_SHA224 KeyMechanism = "RSA_Decryption_OAEP_SHA224"
	KEYMECHANISM_RSA_DECRYPTION_OAEP_SHA256 KeyMechanism = "RSA_Decryption_OAEP_SHA256"
	KEYMECHANISM_RSA_DECRYPTION_OAEP_SHA384 KeyMechanism = "RSA_Decryption_OAEP_SHA384"
	KEYMECHANISM_RSA_DECRYPTION_OAEP_SHA512 KeyMechanism = "RSA_Decryption_OAEP_SHA512"
	KEYMECHANISM_RSA_SIGNATURE_PKCS1        KeyMechanism = "RSA_Signature_PKCS1"
	KEYMECHANISM_RSA_SIGNATURE_PSS_MD5      KeyMechanism = "RSA_Signature_PSS_MD5"
	KEYMECHANISM_RSA_SIGNATURE_PSS_SHA1     KeyMechanism = "RSA_Signature_PSS_SHA1"
	KEYMECHANISM_RSA_SIGNATURE_PSS_SHA224   KeyMechanism = "RSA_Signature_PSS_SHA224"
	KEYMECHANISM_RSA_SIGNATURE_PSS_SHA256   KeyMechanism = "RSA_Signature_PSS_SHA256"
	KEYMECHANISM_RSA_SIGNATURE_PSS_SHA384   KeyMechanism = "RSA_Signature_PSS_SHA384"
	KEYMECHANISM_RSA_SIGNATURE_PSS_SHA512   KeyMechanism = "RSA_Signature_PSS_SHA512"
	KEYMECHANISM_ED25519_SIGNATURE          KeyMechanism = "ED25519_Signature"
)

func (v *KeyMechanism) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KeyMechanism(value)
	for _, existing := range []KeyMechanism{"RSA_Decryption_RAW", "RSA_Decryption_PKCS1", "RSA_Decryption_OAEP_MD5", "RSA_Decryption_OAEP_SHA1", "RSA_Decryption_OAEP_SHA224", "RSA_Decryption_OAEP_SHA256", "RSA_Decryption_OAEP_SHA384", "RSA_Decryption_OAEP_SHA512", "RSA_Signature_PKCS1", "RSA_Signature_PSS_MD5", "RSA_Signature_PSS_SHA1", "RSA_Signature_PSS_SHA224", "RSA_Signature_PSS_SHA256", "RSA_Signature_PSS_SHA384", "RSA_Signature_PSS_SHA512", "ED25519_Signature"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KeyMechanism", value)
}

// Ptr returns reference to KeyMechanism value
func (v KeyMechanism) Ptr() *KeyMechanism {
	return &v
}

type NullableKeyMechanism struct {
	value *KeyMechanism
	isSet bool
}

func (v NullableKeyMechanism) Get() *KeyMechanism {
	return v.value
}

func (v *NullableKeyMechanism) Set(val *KeyMechanism) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyMechanism) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyMechanism) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyMechanism(val *KeyMechanism) *NullableKeyMechanism {
	return &NullableKeyMechanism{value: val, isSet: true}
}

func (v NullableKeyMechanism) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyMechanism) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
