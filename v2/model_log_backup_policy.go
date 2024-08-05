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

// checks if the LogBackupPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogBackupPolicy{}

// LogBackupPolicy Specifies log backup settings for a Protection Group.
type LogBackupPolicy struct {
	Retention Retention `json:"retention"`
	Schedule LogSchedule `json:"schedule"`
}

type _LogBackupPolicy LogBackupPolicy

// NewLogBackupPolicy instantiates a new LogBackupPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogBackupPolicy(retention Retention, schedule LogSchedule) *LogBackupPolicy {
	this := LogBackupPolicy{}
	this.Retention = retention
	this.Schedule = schedule
	return &this
}

// NewLogBackupPolicyWithDefaults instantiates a new LogBackupPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogBackupPolicyWithDefaults() *LogBackupPolicy {
	this := LogBackupPolicy{}
	return &this
}

// GetRetention returns the Retention field value
func (o *LogBackupPolicy) GetRetention() Retention {
	if o == nil {
		var ret Retention
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *LogBackupPolicy) GetRetentionOk() (*Retention, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *LogBackupPolicy) SetRetention(v Retention) {
	o.Retention = v
}

// GetSchedule returns the Schedule field value
func (o *LogBackupPolicy) GetSchedule() LogSchedule {
	if o == nil {
		var ret LogSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *LogBackupPolicy) GetScheduleOk() (*LogSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *LogBackupPolicy) SetSchedule(v LogSchedule) {
	o.Schedule = v
}

func (o LogBackupPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogBackupPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["retention"] = o.Retention
	toSerialize["schedule"] = o.Schedule
	return toSerialize, nil
}

func (o *LogBackupPolicy) UnmarshalJSON(data []byte) (err error) {
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

	varLogBackupPolicy := _LogBackupPolicy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLogBackupPolicy)

	if err != nil {
		return err
	}

	*o = LogBackupPolicy(varLogBackupPolicy)

	return err
}

type NullableLogBackupPolicy struct {
	value *LogBackupPolicy
	isSet bool
}

func (v NullableLogBackupPolicy) Get() *LogBackupPolicy {
	return v.value
}

func (v *NullableLogBackupPolicy) Set(val *LogBackupPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableLogBackupPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableLogBackupPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogBackupPolicy(val *LogBackupPolicy) *NullableLogBackupPolicy {
	return &NullableLogBackupPolicy{value: val, isSet: true}
}

func (v NullableLogBackupPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogBackupPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


