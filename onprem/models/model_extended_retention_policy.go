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

// ExtendedRetentionPolicy Specifies additional retention policies to apply to backup snapshots.
type ExtendedRetentionPolicy struct {
	Schedule ExtendedRetentionSchedule `json:"schedule"`
	Retention Retention `json:"retention"`
	// The backup run type to which this extended retention applies to. If this is not set, the extended retention will be applicable to all non-log backup types. Currently, the only value that can be set here is Full. 'Regular' indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a Regular schedule captures all the blocks. 'Full' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. 'Log' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. 'System' indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time.
	RunType NullableString `json:"runType,omitempty"`
	// Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated.
	ConfigId NullableString `json:"configId,omitempty"`
}

// NewExtendedRetentionPolicy instantiates a new ExtendedRetentionPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendedRetentionPolicy(schedule ExtendedRetentionSchedule, retention Retention) *ExtendedRetentionPolicy {
	this := ExtendedRetentionPolicy{}
	this.Schedule = schedule
	this.Retention = retention
	return &this
}

// NewExtendedRetentionPolicyWithDefaults instantiates a new ExtendedRetentionPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendedRetentionPolicyWithDefaults() *ExtendedRetentionPolicy {
	this := ExtendedRetentionPolicy{}
	return &this
}

// GetSchedule returns the Schedule field value
func (o *ExtendedRetentionPolicy) GetSchedule() ExtendedRetentionSchedule {
	if o == nil {
		var ret ExtendedRetentionSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *ExtendedRetentionPolicy) GetScheduleOk() (*ExtendedRetentionSchedule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *ExtendedRetentionPolicy) SetSchedule(v ExtendedRetentionSchedule) {
	o.Schedule = v
}

// GetRetention returns the Retention field value
func (o *ExtendedRetentionPolicy) GetRetention() Retention {
	if o == nil {
		var ret Retention
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *ExtendedRetentionPolicy) GetRetentionOk() (*Retention, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *ExtendedRetentionPolicy) SetRetention(v Retention) {
	o.Retention = v
}

// GetRunType returns the RunType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtendedRetentionPolicy) GetRunType() string {
	if o == nil || o.RunType.Get() == nil {
		var ret string
		return ret
	}
	return *o.RunType.Get()
}

// GetRunTypeOk returns a tuple with the RunType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtendedRetentionPolicy) GetRunTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RunType.Get(), o.RunType.IsSet()
}

// HasRunType returns a boolean if a field has been set.
func (o *ExtendedRetentionPolicy) HasRunType() bool {
	if o != nil && o.RunType.IsSet() {
		return true
	}

	return false
}

// SetRunType gets a reference to the given NullableString and assigns it to the RunType field.
func (o *ExtendedRetentionPolicy) SetRunType(v string) {
	o.RunType.Set(&v)
}
// SetRunTypeNil sets the value for RunType to be an explicit nil
func (o *ExtendedRetentionPolicy) SetRunTypeNil() {
	o.RunType.Set(nil)
}

// UnsetRunType ensures that no value is present for RunType, not even an explicit nil
func (o *ExtendedRetentionPolicy) UnsetRunType() {
	o.RunType.Unset()
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtendedRetentionPolicy) GetConfigId() string {
	if o == nil || o.ConfigId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigId.Get()
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtendedRetentionPolicy) GetConfigIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConfigId.Get(), o.ConfigId.IsSet()
}

// HasConfigId returns a boolean if a field has been set.
func (o *ExtendedRetentionPolicy) HasConfigId() bool {
	if o != nil && o.ConfigId.IsSet() {
		return true
	}

	return false
}

// SetConfigId gets a reference to the given NullableString and assigns it to the ConfigId field.
func (o *ExtendedRetentionPolicy) SetConfigId(v string) {
	o.ConfigId.Set(&v)
}
// SetConfigIdNil sets the value for ConfigId to be an explicit nil
func (o *ExtendedRetentionPolicy) SetConfigIdNil() {
	o.ConfigId.Set(nil)
}

// UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
func (o *ExtendedRetentionPolicy) UnsetConfigId() {
	o.ConfigId.Unset()
}

func (o ExtendedRetentionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	if true {
		toSerialize["retention"] = o.Retention
	}
	if o.RunType.IsSet() {
		toSerialize["runType"] = o.RunType.Get()
	}
	if o.ConfigId.IsSet() {
		toSerialize["configId"] = o.ConfigId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableExtendedRetentionPolicy struct {
	value *ExtendedRetentionPolicy
	isSet bool
}

func (v NullableExtendedRetentionPolicy) Get() *ExtendedRetentionPolicy {
	return v.value
}

func (v *NullableExtendedRetentionPolicy) Set(val *ExtendedRetentionPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendedRetentionPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendedRetentionPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendedRetentionPolicy(val *ExtendedRetentionPolicy) *NullableExtendedRetentionPolicy {
	return &NullableExtendedRetentionPolicy{value: val, isSet: true}
}

func (v NullableExtendedRetentionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendedRetentionPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ExtendedRetentionPolicy) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}