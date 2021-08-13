/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// FrequencySchedule Specifies settings that define a daily schedule for a Protection Policy.
type FrequencySchedule struct {
	// Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
	Frequency NullableInt64 `json:"frequency"`
}

// NewFrequencySchedule instantiates a new FrequencySchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrequencySchedule(frequency NullableInt64) *FrequencySchedule {
	this := FrequencySchedule{}
	this.Frequency = frequency
	return &this
}

// NewFrequencyScheduleWithDefaults instantiates a new FrequencySchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrequencyScheduleWithDefaults() *FrequencySchedule {
	this := FrequencySchedule{}
	return &this
}

// GetFrequency returns the Frequency field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *FrequencySchedule) GetFrequency() int64 {
	if o == nil || o.Frequency.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Frequency.Get()
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrequencySchedule) GetFrequencyOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Frequency.Get(), o.Frequency.IsSet()
}

// SetFrequency sets field value
func (o *FrequencySchedule) SetFrequency(v int64) {
	o.Frequency.Set(&v)
}

func (o FrequencySchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["frequency"] = o.Frequency.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFrequencySchedule struct {
	value *FrequencySchedule
	isSet bool
}

func (v NullableFrequencySchedule) Get() *FrequencySchedule {
	return v.value
}

func (v *NullableFrequencySchedule) Set(val *FrequencySchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableFrequencySchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableFrequencySchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrequencySchedule(val *FrequencySchedule) *NullableFrequencySchedule {
	return &NullableFrequencySchedule{value: val, isSet: true}
}

func (v NullableFrequencySchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrequencySchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


