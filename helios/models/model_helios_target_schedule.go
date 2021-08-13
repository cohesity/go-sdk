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

// HeliosTargetSchedule Specifies a schedule frequency and schedule unit for copying Snapshots to backup targets.
type HeliosTargetSchedule struct {
	// Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
	Unit NullableString `json:"unit,omitempty"`
	// Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
	Frequency NullableInt32 `json:"frequency,omitempty"`
}

// NewHeliosTargetSchedule instantiates a new HeliosTargetSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosTargetSchedule() *HeliosTargetSchedule {
	this := HeliosTargetSchedule{}
	return &this
}

// NewHeliosTargetScheduleWithDefaults instantiates a new HeliosTargetSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosTargetScheduleWithDefaults() *HeliosTargetSchedule {
	this := HeliosTargetSchedule{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosTargetSchedule) GetUnit() string {
	if o == nil || o.Unit.Get() == nil {
		var ret string
		return ret
	}
	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosTargetSchedule) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// HasUnit returns a boolean if a field has been set.
func (o *HeliosTargetSchedule) HasUnit() bool {
	if o != nil && o.Unit.IsSet() {
		return true
	}

	return false
}

// SetUnit gets a reference to the given NullableString and assigns it to the Unit field.
func (o *HeliosTargetSchedule) SetUnit(v string) {
	o.Unit.Set(&v)
}
// SetUnitNil sets the value for Unit to be an explicit nil
func (o *HeliosTargetSchedule) SetUnitNil() {
	o.Unit.Set(nil)
}

// UnsetUnit ensures that no value is present for Unit, not even an explicit nil
func (o *HeliosTargetSchedule) UnsetUnit() {
	o.Unit.Unset()
}

// GetFrequency returns the Frequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosTargetSchedule) GetFrequency() int32 {
	if o == nil || o.Frequency.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Frequency.Get()
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosTargetSchedule) GetFrequencyOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Frequency.Get(), o.Frequency.IsSet()
}

// HasFrequency returns a boolean if a field has been set.
func (o *HeliosTargetSchedule) HasFrequency() bool {
	if o != nil && o.Frequency.IsSet() {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given NullableInt32 and assigns it to the Frequency field.
func (o *HeliosTargetSchedule) SetFrequency(v int32) {
	o.Frequency.Set(&v)
}
// SetFrequencyNil sets the value for Frequency to be an explicit nil
func (o *HeliosTargetSchedule) SetFrequencyNil() {
	o.Frequency.Set(nil)
}

// UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
func (o *HeliosTargetSchedule) UnsetFrequency() {
	o.Frequency.Unset()
}

func (o HeliosTargetSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Unit.IsSet() {
		toSerialize["unit"] = o.Unit.Get()
	}
	if o.Frequency.IsSet() {
		toSerialize["frequency"] = o.Frequency.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHeliosTargetSchedule struct {
	value *HeliosTargetSchedule
	isSet bool
}

func (v NullableHeliosTargetSchedule) Get() *HeliosTargetSchedule {
	return v.value
}

func (v *NullableHeliosTargetSchedule) Set(val *HeliosTargetSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosTargetSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosTargetSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosTargetSchedule(val *HeliosTargetSchedule) *NullableHeliosTargetSchedule {
	return &NullableHeliosTargetSchedule{value: val, isSet: true}
}

func (v NullableHeliosTargetSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosTargetSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


