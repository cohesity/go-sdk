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

// HeliosMinuteSchedule Specifies settings that define a schedule for a Protection Group runs to start after certain number of minutes.
type HeliosMinuteSchedule struct {
	// Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. If selected unit is 'Weeks' or 'Months' then frequency will only be applied if policy type is DMaas.
	Frequency NullableInt64 `json:"frequency"`
}

// NewHeliosMinuteSchedule instantiates a new HeliosMinuteSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosMinuteSchedule(frequency NullableInt64) *HeliosMinuteSchedule {
	this := HeliosMinuteSchedule{}
	this.Frequency = frequency
	return &this
}

// NewHeliosMinuteScheduleWithDefaults instantiates a new HeliosMinuteSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosMinuteScheduleWithDefaults() *HeliosMinuteSchedule {
	this := HeliosMinuteSchedule{}
	return &this
}

// GetFrequency returns the Frequency field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *HeliosMinuteSchedule) GetFrequency() int64 {
	if o == nil || o.Frequency.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Frequency.Get()
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosMinuteSchedule) GetFrequencyOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Frequency.Get(), o.Frequency.IsSet()
}

// SetFrequency sets field value
func (o *HeliosMinuteSchedule) SetFrequency(v int64) {
	o.Frequency.Set(&v)
}

func (o HeliosMinuteSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["frequency"] = o.Frequency.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHeliosMinuteSchedule struct {
	value *HeliosMinuteSchedule
	isSet bool
}

func (v NullableHeliosMinuteSchedule) Get() *HeliosMinuteSchedule {
	return v.value
}

func (v *NullableHeliosMinuteSchedule) Set(val *HeliosMinuteSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosMinuteSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosMinuteSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosMinuteSchedule(val *HeliosMinuteSchedule) *NullableHeliosMinuteSchedule {
	return &NullableHeliosMinuteSchedule{value: val, isSet: true}
}

func (v NullableHeliosMinuteSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosMinuteSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


