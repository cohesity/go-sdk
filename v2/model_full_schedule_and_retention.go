/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FullScheduleAndRetention type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FullScheduleAndRetention{}

// FullScheduleAndRetention Specifies the settings to schedule the full backup and retention for each schedule.
type FullScheduleAndRetention struct {
	Retention Retention `json:"retention"`
	Schedule FullSchedule `json:"schedule"`
}

type _FullScheduleAndRetention FullScheduleAndRetention

// NewFullScheduleAndRetention instantiates a new FullScheduleAndRetention object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullScheduleAndRetention(retention Retention, schedule FullSchedule) *FullScheduleAndRetention {
	this := FullScheduleAndRetention{}
	this.Retention = retention
	this.Schedule = schedule
	return &this
}

// NewFullScheduleAndRetentionWithDefaults instantiates a new FullScheduleAndRetention object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullScheduleAndRetentionWithDefaults() *FullScheduleAndRetention {
	this := FullScheduleAndRetention{}
	return &this
}

// GetRetention returns the Retention field value
func (o *FullScheduleAndRetention) GetRetention() Retention {
	if o == nil {
		var ret Retention
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *FullScheduleAndRetention) GetRetentionOk() (*Retention, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *FullScheduleAndRetention) SetRetention(v Retention) {
	o.Retention = v
}

// GetSchedule returns the Schedule field value
func (o *FullScheduleAndRetention) GetSchedule() FullSchedule {
	if o == nil {
		var ret FullSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *FullScheduleAndRetention) GetScheduleOk() (*FullSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *FullScheduleAndRetention) SetSchedule(v FullSchedule) {
	o.Schedule = v
}

func (o FullScheduleAndRetention) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FullScheduleAndRetention) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["retention"] = o.Retention
	toSerialize["schedule"] = o.Schedule
	return toSerialize, nil
}

func (o *FullScheduleAndRetention) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"retention",
		"schedule",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFullScheduleAndRetention := _FullScheduleAndRetention{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFullScheduleAndRetention)

	if err != nil {
		return err
	}

	*o = FullScheduleAndRetention(varFullScheduleAndRetention)

	return err
}

type NullableFullScheduleAndRetention struct {
	value *FullScheduleAndRetention
	isSet bool
}

func (v NullableFullScheduleAndRetention) Get() *FullScheduleAndRetention {
	return v.value
}

func (v *NullableFullScheduleAndRetention) Set(val *FullScheduleAndRetention) {
	v.value = val
	v.isSet = true
}

func (v NullableFullScheduleAndRetention) IsSet() bool {
	return v.isSet
}

func (v *NullableFullScheduleAndRetention) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullScheduleAndRetention(val *FullScheduleAndRetention) *NullableFullScheduleAndRetention {
	return &NullableFullScheduleAndRetention{value: val, isSet: true}
}

func (v NullableFullScheduleAndRetention) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullScheduleAndRetention) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


