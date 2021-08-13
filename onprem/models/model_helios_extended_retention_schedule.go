/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package onprem
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/onprem/utils"
	"fmt"
)

var _ = NullableBool{}

// HeliosExtendedRetentionSchedule Specifies a schedule frequency and schedule unit for Extended Retentions.
type HeliosExtendedRetentionSchedule struct {
	// Specifies the unit interval for retention of Snapshots. <br>'Runs' means that the Snapshot copy retained after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy retained hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy gets retained daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy is retained weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy is retained monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy is retained yearly at the frequency set in the Frequency.
	Unit NullableString `json:"unit,omitempty"`
	// Specifies a factor to multiply the unit by, to determine the retention schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is retained.
	Frequency NullableInt32 `json:"frequency,omitempty"`
}

// NewHeliosExtendedRetentionSchedule instantiates a new HeliosExtendedRetentionSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosExtendedRetentionSchedule() *HeliosExtendedRetentionSchedule {
	this := HeliosExtendedRetentionSchedule{}
	return &this
}

// NewHeliosExtendedRetentionScheduleWithDefaults instantiates a new HeliosExtendedRetentionSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosExtendedRetentionScheduleWithDefaults() *HeliosExtendedRetentionSchedule {
	this := HeliosExtendedRetentionSchedule{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosExtendedRetentionSchedule) GetUnit() string {
	if o == nil || o.Unit.Get() == nil {
		var ret string
		return ret
	}
	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosExtendedRetentionSchedule) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// HasUnit returns a boolean if a field has been set.
func (o *HeliosExtendedRetentionSchedule) HasUnit() bool {
	if o != nil && o.Unit.IsSet() {
		return true
	}

	return false
}

// SetUnit gets a reference to the given NullableString and assigns it to the Unit field.
func (o *HeliosExtendedRetentionSchedule) SetUnit(v string) {
	o.Unit.Set(&v)
}
// SetUnitNil sets the value for Unit to be an explicit nil
func (o *HeliosExtendedRetentionSchedule) SetUnitNil() {
	o.Unit.Set(nil)
}

// UnsetUnit ensures that no value is present for Unit, not even an explicit nil
func (o *HeliosExtendedRetentionSchedule) UnsetUnit() {
	o.Unit.Unset()
}

// GetFrequency returns the Frequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosExtendedRetentionSchedule) GetFrequency() int32 {
	if o == nil || o.Frequency.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Frequency.Get()
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosExtendedRetentionSchedule) GetFrequencyOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Frequency.Get(), o.Frequency.IsSet()
}

// HasFrequency returns a boolean if a field has been set.
func (o *HeliosExtendedRetentionSchedule) HasFrequency() bool {
	if o != nil && o.Frequency.IsSet() {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given NullableInt32 and assigns it to the Frequency field.
func (o *HeliosExtendedRetentionSchedule) SetFrequency(v int32) {
	o.Frequency.Set(&v)
}
// SetFrequencyNil sets the value for Frequency to be an explicit nil
func (o *HeliosExtendedRetentionSchedule) SetFrequencyNil() {
	o.Frequency.Set(nil)
}

// UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
func (o *HeliosExtendedRetentionSchedule) UnsetFrequency() {
	o.Frequency.Unset()
}

func (o HeliosExtendedRetentionSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Unit.IsSet() {
		toSerialize["unit"] = o.Unit.Get()
	}
	if o.Frequency.IsSet() {
		toSerialize["frequency"] = o.Frequency.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHeliosExtendedRetentionSchedule struct {
	value *HeliosExtendedRetentionSchedule
	isSet bool
}

func (v NullableHeliosExtendedRetentionSchedule) Get() *HeliosExtendedRetentionSchedule {
	return v.value
}

func (v *NullableHeliosExtendedRetentionSchedule) Set(val *HeliosExtendedRetentionSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosExtendedRetentionSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosExtendedRetentionSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosExtendedRetentionSchedule(val *HeliosExtendedRetentionSchedule) *NullableHeliosExtendedRetentionSchedule {
	return &NullableHeliosExtendedRetentionSchedule{value: val, isSet: true}
}

func (v NullableHeliosExtendedRetentionSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosExtendedRetentionSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o HeliosExtendedRetentionSchedule) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}