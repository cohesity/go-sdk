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

// HeliosArchivalConfig Specifies settings for copying Snapshots External Targets (such as AWS or Tape). This also specifies the retention policy that should be applied to Snapshots after they have been copied to the specified target.
type HeliosArchivalConfig struct {
	Schedule *HeliosTargetSchedule `json:"schedule,omitempty"`
	Retention *HeliosRetention `json:"retention,omitempty"`
	// Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
	CopyOnRunSuccess NullableBool `json:"copyOnRunSuccess,omitempty"`
	// Specifies the unique identifier for the target getting added. This field need to be passed only when helios policies are updated.
	ConfigId NullableString `json:"configId,omitempty"`
}

// NewHeliosArchivalConfig instantiates a new HeliosArchivalConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeliosArchivalConfig() *HeliosArchivalConfig {
	this := HeliosArchivalConfig{}
	return &this
}

// NewHeliosArchivalConfigWithDefaults instantiates a new HeliosArchivalConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeliosArchivalConfigWithDefaults() *HeliosArchivalConfig {
	this := HeliosArchivalConfig{}
	return &this
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *HeliosArchivalConfig) GetSchedule() HeliosTargetSchedule {
	if o == nil || o.Schedule == nil {
		var ret HeliosTargetSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosArchivalConfig) GetScheduleOk() (*HeliosTargetSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *HeliosArchivalConfig) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given HeliosTargetSchedule and assigns it to the Schedule field.
func (o *HeliosArchivalConfig) SetSchedule(v HeliosTargetSchedule) {
	o.Schedule = &v
}

// GetRetention returns the Retention field value if set, zero value otherwise.
func (o *HeliosArchivalConfig) GetRetention() HeliosRetention {
	if o == nil || o.Retention == nil {
		var ret HeliosRetention
		return ret
	}
	return *o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeliosArchivalConfig) GetRetentionOk() (*HeliosRetention, bool) {
	if o == nil || o.Retention == nil {
		return nil, false
	}
	return o.Retention, true
}

// HasRetention returns a boolean if a field has been set.
func (o *HeliosArchivalConfig) HasRetention() bool {
	if o != nil && o.Retention != nil {
		return true
	}

	return false
}

// SetRetention gets a reference to the given HeliosRetention and assigns it to the Retention field.
func (o *HeliosArchivalConfig) SetRetention(v HeliosRetention) {
	o.Retention = &v
}

// GetCopyOnRunSuccess returns the CopyOnRunSuccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosArchivalConfig) GetCopyOnRunSuccess() bool {
	if o == nil || o.CopyOnRunSuccess.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CopyOnRunSuccess.Get()
}

// GetCopyOnRunSuccessOk returns a tuple with the CopyOnRunSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosArchivalConfig) GetCopyOnRunSuccessOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CopyOnRunSuccess.Get(), o.CopyOnRunSuccess.IsSet()
}

// HasCopyOnRunSuccess returns a boolean if a field has been set.
func (o *HeliosArchivalConfig) HasCopyOnRunSuccess() bool {
	if o != nil && o.CopyOnRunSuccess.IsSet() {
		return true
	}

	return false
}

// SetCopyOnRunSuccess gets a reference to the given NullableBool and assigns it to the CopyOnRunSuccess field.
func (o *HeliosArchivalConfig) SetCopyOnRunSuccess(v bool) {
	o.CopyOnRunSuccess.Set(&v)
}
// SetCopyOnRunSuccessNil sets the value for CopyOnRunSuccess to be an explicit nil
func (o *HeliosArchivalConfig) SetCopyOnRunSuccessNil() {
	o.CopyOnRunSuccess.Set(nil)
}

// UnsetCopyOnRunSuccess ensures that no value is present for CopyOnRunSuccess, not even an explicit nil
func (o *HeliosArchivalConfig) UnsetCopyOnRunSuccess() {
	o.CopyOnRunSuccess.Unset()
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeliosArchivalConfig) GetConfigId() string {
	if o == nil || o.ConfigId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigId.Get()
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeliosArchivalConfig) GetConfigIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConfigId.Get(), o.ConfigId.IsSet()
}

// HasConfigId returns a boolean if a field has been set.
func (o *HeliosArchivalConfig) HasConfigId() bool {
	if o != nil && o.ConfigId.IsSet() {
		return true
	}

	return false
}

// SetConfigId gets a reference to the given NullableString and assigns it to the ConfigId field.
func (o *HeliosArchivalConfig) SetConfigId(v string) {
	o.ConfigId.Set(&v)
}
// SetConfigIdNil sets the value for ConfigId to be an explicit nil
func (o *HeliosArchivalConfig) SetConfigIdNil() {
	o.ConfigId.Set(nil)
}

// UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
func (o *HeliosArchivalConfig) UnsetConfigId() {
	o.ConfigId.Unset()
}

func (o HeliosArchivalConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.Retention != nil {
		toSerialize["retention"] = o.Retention
	}
	if o.CopyOnRunSuccess.IsSet() {
		toSerialize["copyOnRunSuccess"] = o.CopyOnRunSuccess.Get()
	}
	if o.ConfigId.IsSet() {
		toSerialize["configId"] = o.ConfigId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHeliosArchivalConfig struct {
	value *HeliosArchivalConfig
	isSet bool
}

func (v NullableHeliosArchivalConfig) Get() *HeliosArchivalConfig {
	return v.value
}

func (v *NullableHeliosArchivalConfig) Set(val *HeliosArchivalConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHeliosArchivalConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHeliosArchivalConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeliosArchivalConfig(val *HeliosArchivalConfig) *NullableHeliosArchivalConfig {
	return &NullableHeliosArchivalConfig{value: val, isSet: true}
}

func (v NullableHeliosArchivalConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeliosArchivalConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


