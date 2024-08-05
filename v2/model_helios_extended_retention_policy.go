/*
Cohesity REST API

Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// checks if the HeliosExtendedRetentionPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeliosExtendedRetentionPolicy{}

// HeliosExtendedRetentionPolicy Specifies additional retention policies to apply to backup snapshots.
type HeliosExtendedRetentionPolicy struct {
	// Specifies the unique identifier for the extedned  retention getting added. This field should only be set if policy is getting updated.
	ConfigId NullableString `json:"configId,omitempty"`
	Retention *HeliosRetention `json:"retention,omitempty"`
	// The backup run type to which this extended retention applies to. If this is not set, the extended retention will be applicable to all non-log backup types. Currently, the only value that can be set here is Full. 'Regular' indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a Regular schedule captures all the blocks. 'Full' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. 'Log' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. 'System' indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time.
	RunType NullableString `json:"runType,omitempty"`
	Schedule *HeliosExtendedRetentionSchedule `json:"schedule,omitempty"`
}

// NewHeliosExtendedRetentionPolicy instantiates a new HeliosExtendedRetentionPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosExtendedRetentionPolicy() *HeliosExtendedRetentionPolicy {
	this := HeliosExtendedRetentionPolicy{}
	return &this
}

// NewHeliosExtendedRetentionPolicyWithDefaults instantiates a new HeliosExtendedRetentionPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosExtendedRetentionPolicyWithDefaults() *HeliosExtendedRetentionPolicy {
	this := HeliosExtendedRetentionPolicy{}
	return &this
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosExtendedRetentionPolicy) GetConfigId() string {
	if o == nil || IsNil(o.ConfigId.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigId.Get()
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosExtendedRetentionPolicy) GetConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigId.Get(), o.ConfigId.IsSet()
}

// HasConfigId returns a boolean if a field has been set.
func (o *HeliosExtendedRetentionPolicy) HasConfigId() bool {
	if o != nil && o.ConfigId.IsSet() {
		return true
	}

	return false
}

// SetConfigId gets a reference to the given NullableString and assigns it to the ConfigId field.
func (o *HeliosExtendedRetentionPolicy) SetConfigId(v string) {
	o.ConfigId.Set(&v)
}
// SetConfigIdNil sets the value for ConfigId to be an explicit nil
func (o *HeliosExtendedRetentionPolicy) SetConfigIdNil() {
	o.ConfigId.Set(nil)
}

// UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
func (o *HeliosExtendedRetentionPolicy) UnsetConfigId() {
	o.ConfigId.Unset()
}

// GetRetention returns the Retention field value if set, zero value otherwise.
func (o *HeliosExtendedRetentionPolicy) GetRetention() HeliosRetention {
	if o == nil || IsNil(o.Retention) {
		var ret HeliosRetention
		return ret
	}
	return *o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosExtendedRetentionPolicy) GetRetentionOk() (*HeliosRetention, bool) {
	if o == nil || IsNil(o.Retention) {
		return nil, false
	}
	return o.Retention, true
}

// HasRetention returns a boolean if a field has been set.
func (o *HeliosExtendedRetentionPolicy) HasRetention() bool {
	if o != nil && !IsNil(o.Retention) {
		return true
	}

	return false
}

// SetRetention gets a reference to the given HeliosRetention and assigns it to the Retention field.
func (o *HeliosExtendedRetentionPolicy) SetRetention(v HeliosRetention) {
	o.Retention = &v
}

// GetRunType returns the RunType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosExtendedRetentionPolicy) GetRunType() string {
	if o == nil || IsNil(o.RunType.Get()) {
		var ret string
		return ret
	}
	return *o.RunType.Get()
}

// GetRunTypeOk returns a tuple with the RunType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosExtendedRetentionPolicy) GetRunTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunType.Get(), o.RunType.IsSet()
}

// HasRunType returns a boolean if a field has been set.
func (o *HeliosExtendedRetentionPolicy) HasRunType() bool {
	if o != nil && o.RunType.IsSet() {
		return true
	}

	return false
}

// SetRunType gets a reference to the given NullableString and assigns it to the RunType field.
func (o *HeliosExtendedRetentionPolicy) SetRunType(v string) {
	o.RunType.Set(&v)
}
// SetRunTypeNil sets the value for RunType to be an explicit nil
func (o *HeliosExtendedRetentionPolicy) SetRunTypeNil() {
	o.RunType.Set(nil)
}

// UnsetRunType ensures that no value is present for RunType, not even an explicit nil
func (o *HeliosExtendedRetentionPolicy) UnsetRunType() {
	o.RunType.Unset()
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *HeliosExtendedRetentionPolicy) GetSchedule() HeliosExtendedRetentionSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret HeliosExtendedRetentionSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosExtendedRetentionPolicy) GetScheduleOk() (*HeliosExtendedRetentionSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *HeliosExtendedRetentionPolicy) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given HeliosExtendedRetentionSchedule and assigns it to the Schedule field.
func (o *HeliosExtendedRetentionPolicy) SetSchedule(v HeliosExtendedRetentionSchedule) {
	o.Schedule = &v
}

func (o HeliosExtendedRetentionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeliosExtendedRetentionPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfigId.IsSet() {
		toSerialize["configId"] = o.ConfigId.Get()
	}
	if !IsNil(o.Retention) {
		toSerialize["retention"] = o.Retention
	}
	if o.RunType.IsSet() {
		toSerialize["runType"] = o.RunType.Get()
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	return toSerialize, nil
}

type NullableHeliosExtendedRetentionPolicy struct {
	value *HeliosExtendedRetentionPolicy
	isSet bool
}

func (v NullableHeliosExtendedRetentionPolicy) Get() *HeliosExtendedRetentionPolicy {
	return v.value
}

func (v *NullableHeliosExtendedRetentionPolicy) Set(val *HeliosExtendedRetentionPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosExtendedRetentionPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosExtendedRetentionPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosExtendedRetentionPolicy(val *HeliosExtendedRetentionPolicy) *NullableHeliosExtendedRetentionPolicy {
	return &NullableHeliosExtendedRetentionPolicy{value: val, isSet: true}
}

func (v NullableHeliosExtendedRetentionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosExtendedRetentionPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


