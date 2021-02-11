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
	"time"
)

// TimeConfig struct for TimeConfig
type TimeConfig struct {
	Time time.Time `json:"time"`
	AdditionalProperties map[string]interface{}
}

type _TimeConfig TimeConfig

// NewTimeConfig instantiates a new TimeConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeConfig(time time.Time, ) *TimeConfig {
	this := TimeConfig{}
	this.Time = time
	return &this
}

// NewTimeConfigWithDefaults instantiates a new TimeConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeConfigWithDefaults() *TimeConfig {
	this := TimeConfig{}
	return &this
}

// GetTime returns the Time field value
func (o *TimeConfig) GetTime() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *TimeConfig) GetTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *TimeConfig) SetTime(v time.Time) {
	o.Time = v
}

func (o TimeConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TimeConfig) UnmarshalJSON(bytes []byte) (err error) {
	varTimeConfig := _TimeConfig{}

	if err = json.Unmarshal(bytes, &varTimeConfig); err == nil {
		*o = TimeConfig(varTimeConfig)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTimeConfig struct {
	value *TimeConfig
	isSet bool
}

func (v NullableTimeConfig) Get() *TimeConfig {
	return v.value
}

func (v *NullableTimeConfig) Set(val *TimeConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeConfig(val *TimeConfig) *NullableTimeConfig {
	return &NullableTimeConfig{value: val, isSet: true}
}

func (v NullableTimeConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


