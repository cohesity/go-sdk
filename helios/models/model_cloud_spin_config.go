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

// CloudSpinConfig Specifies settings for copying Snapshots to Cloud. This also specifies the retention policy that should be applied to Snapshots after they have been copied to Cloud.
type CloudSpinConfig struct {
	Schedule TargetSchedule `json:"schedule"`
	Retention Retention `json:"retention"`
	// Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
	CopyOnRunSuccess NullableBool `json:"copyOnRunSuccess,omitempty"`
	// Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
	ConfigId NullableString `json:"configId,omitempty"`
	Target CloudSpinTarget `json:"target"`
}

// NewCloudSpinConfig instantiates a new CloudSpinConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSpinConfig(schedule TargetSchedule, retention Retention, target CloudSpinTarget) *CloudSpinConfig {
	this := CloudSpinConfig{}
	this.Schedule = schedule
	this.Retention = retention
	this.Target = target
	return &this
}

// NewCloudSpinConfigWithDefaults instantiates a new CloudSpinConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSpinConfigWithDefaults() *CloudSpinConfig {
	this := CloudSpinConfig{}
	return &this
}

// GetSchedule returns the Schedule field value
func (o *CloudSpinConfig) GetSchedule() TargetSchedule {
	if o == nil {
		var ret TargetSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *CloudSpinConfig) GetScheduleOk() (*TargetSchedule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *CloudSpinConfig) SetSchedule(v TargetSchedule) {
	o.Schedule = v
}

// GetRetention returns the Retention field value
func (o *CloudSpinConfig) GetRetention() Retention {
	if o == nil {
		var ret Retention
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *CloudSpinConfig) GetRetentionOk() (*Retention, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *CloudSpinConfig) SetRetention(v Retention) {
	o.Retention = v
}

// GetCopyOnRunSuccess returns the CopyOnRunSuccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudSpinConfig) GetCopyOnRunSuccess() bool {
	if o == nil || o.CopyOnRunSuccess.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CopyOnRunSuccess.Get()
}

// GetCopyOnRunSuccessOk returns a tuple with the CopyOnRunSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudSpinConfig) GetCopyOnRunSuccessOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CopyOnRunSuccess.Get(), o.CopyOnRunSuccess.IsSet()
}

// HasCopyOnRunSuccess returns a boolean if a field has been set.
func (o *CloudSpinConfig) HasCopyOnRunSuccess() bool {
	if o != nil && o.CopyOnRunSuccess.IsSet() {
		return true
	}

	return false
}

// SetCopyOnRunSuccess gets a reference to the given NullableBool and assigns it to the CopyOnRunSuccess field.
func (o *CloudSpinConfig) SetCopyOnRunSuccess(v bool) {
	o.CopyOnRunSuccess.Set(&v)
}
// SetCopyOnRunSuccessNil sets the value for CopyOnRunSuccess to be an explicit nil
func (o *CloudSpinConfig) SetCopyOnRunSuccessNil() {
	o.CopyOnRunSuccess.Set(nil)
}

// UnsetCopyOnRunSuccess ensures that no value is present for CopyOnRunSuccess, not even an explicit nil
func (o *CloudSpinConfig) UnsetCopyOnRunSuccess() {
	o.CopyOnRunSuccess.Unset()
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudSpinConfig) GetConfigId() string {
	if o == nil || o.ConfigId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigId.Get()
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudSpinConfig) GetConfigIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConfigId.Get(), o.ConfigId.IsSet()
}

// HasConfigId returns a boolean if a field has been set.
func (o *CloudSpinConfig) HasConfigId() bool {
	if o != nil && o.ConfigId.IsSet() {
		return true
	}

	return false
}

// SetConfigId gets a reference to the given NullableString and assigns it to the ConfigId field.
func (o *CloudSpinConfig) SetConfigId(v string) {
	o.ConfigId.Set(&v)
}
// SetConfigIdNil sets the value for ConfigId to be an explicit nil
func (o *CloudSpinConfig) SetConfigIdNil() {
	o.ConfigId.Set(nil)
}

// UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
func (o *CloudSpinConfig) UnsetConfigId() {
	o.ConfigId.Unset()
}

// GetTarget returns the Target field value
func (o *CloudSpinConfig) GetTarget() CloudSpinTarget {
	if o == nil {
		var ret CloudSpinTarget
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *CloudSpinConfig) GetTargetOk() (*CloudSpinTarget, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *CloudSpinConfig) SetTarget(v CloudSpinTarget) {
	o.Target = v
}

func (o CloudSpinConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	if true {
		toSerialize["retention"] = o.Retention
	}
	if o.CopyOnRunSuccess.IsSet() {
		toSerialize["copyOnRunSuccess"] = o.CopyOnRunSuccess.Get()
	}
	if o.ConfigId.IsSet() {
		toSerialize["configId"] = o.ConfigId.Get()
	}
	if true {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableCloudSpinConfig struct {
	value *CloudSpinConfig
	isSet bool
}

func (v NullableCloudSpinConfig) Get() *CloudSpinConfig {
	return v.value
}

func (v *NullableCloudSpinConfig) Set(val *CloudSpinConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSpinConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSpinConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSpinConfig(val *CloudSpinConfig) *NullableCloudSpinConfig {
	return &NullableCloudSpinConfig{value: val, isSet: true}
}

func (v NullableCloudSpinConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSpinConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


