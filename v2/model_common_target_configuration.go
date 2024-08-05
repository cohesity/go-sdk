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

// checks if the CommonTargetConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonTargetConfiguration{}

// CommonTargetConfiguration Specifies common parameters required while setting up additional protection target configuration.
type CommonTargetConfiguration struct {
	// Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
	BackupRunType NullableString `json:"backupRunType,omitempty"`
	// Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
	ConfigId NullableString `json:"configId,omitempty"`
	// Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
	CopyOnRunSuccess NullableBool `json:"copyOnRunSuccess,omitempty"`
	LogRetention *LogRetention `json:"logRetention,omitempty"`
	Retention Retention `json:"retention"`
	// Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
	RunTimeouts []CancellationTimeoutParams `json:"runTimeouts,omitempty"`
	Schedule TargetSchedule `json:"schedule"`
}

type _CommonTargetConfiguration CommonTargetConfiguration

// NewCommonTargetConfiguration instantiates a new CommonTargetConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonTargetConfiguration(retention Retention, schedule TargetSchedule) *CommonTargetConfiguration {
	this := CommonTargetConfiguration{}
	this.Retention = retention
	this.Schedule = schedule
	return &this
}

// NewCommonTargetConfigurationWithDefaults instantiates a new CommonTargetConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonTargetConfigurationWithDefaults() *CommonTargetConfiguration {
	this := CommonTargetConfiguration{}
	return &this
}

// GetBackupRunType returns the BackupRunType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTargetConfiguration) GetBackupRunType() string {
	if o == nil || IsNil(o.BackupRunType.Get()) {
		var ret string
		return ret
	}
	return *o.BackupRunType.Get()
}

// GetBackupRunTypeOk returns a tuple with the BackupRunType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTargetConfiguration) GetBackupRunTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupRunType.Get(), o.BackupRunType.IsSet()
}

// HasBackupRunType returns a boolean if a field has been set.
func (o *CommonTargetConfiguration) HasBackupRunType() bool {
	if o != nil && o.BackupRunType.IsSet() {
		return true
	}

	return false
}

// SetBackupRunType gets a reference to the given NullableString and assigns it to the BackupRunType field.
func (o *CommonTargetConfiguration) SetBackupRunType(v string) {
	o.BackupRunType.Set(&v)
}
// SetBackupRunTypeNil sets the value for BackupRunType to be an explicit nil
func (o *CommonTargetConfiguration) SetBackupRunTypeNil() {
	o.BackupRunType.Set(nil)
}

// UnsetBackupRunType ensures that no value is present for BackupRunType, not even an explicit nil
func (o *CommonTargetConfiguration) UnsetBackupRunType() {
	o.BackupRunType.Unset()
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTargetConfiguration) GetConfigId() string {
	if o == nil || IsNil(o.ConfigId.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigId.Get()
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTargetConfiguration) GetConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigId.Get(), o.ConfigId.IsSet()
}

// HasConfigId returns a boolean if a field has been set.
func (o *CommonTargetConfiguration) HasConfigId() bool {
	if o != nil && o.ConfigId.IsSet() {
		return true
	}

	return false
}

// SetConfigId gets a reference to the given NullableString and assigns it to the ConfigId field.
func (o *CommonTargetConfiguration) SetConfigId(v string) {
	o.ConfigId.Set(&v)
}
// SetConfigIdNil sets the value for ConfigId to be an explicit nil
func (o *CommonTargetConfiguration) SetConfigIdNil() {
	o.ConfigId.Set(nil)
}

// UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
func (o *CommonTargetConfiguration) UnsetConfigId() {
	o.ConfigId.Unset()
}

// GetCopyOnRunSuccess returns the CopyOnRunSuccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTargetConfiguration) GetCopyOnRunSuccess() bool {
	if o == nil || IsNil(o.CopyOnRunSuccess.Get()) {
		var ret bool
		return ret
	}
	return *o.CopyOnRunSuccess.Get()
}

// GetCopyOnRunSuccessOk returns a tuple with the CopyOnRunSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTargetConfiguration) GetCopyOnRunSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CopyOnRunSuccess.Get(), o.CopyOnRunSuccess.IsSet()
}

// HasCopyOnRunSuccess returns a boolean if a field has been set.
func (o *CommonTargetConfiguration) HasCopyOnRunSuccess() bool {
	if o != nil && o.CopyOnRunSuccess.IsSet() {
		return true
	}

	return false
}

// SetCopyOnRunSuccess gets a reference to the given NullableBool and assigns it to the CopyOnRunSuccess field.
func (o *CommonTargetConfiguration) SetCopyOnRunSuccess(v bool) {
	o.CopyOnRunSuccess.Set(&v)
}
// SetCopyOnRunSuccessNil sets the value for CopyOnRunSuccess to be an explicit nil
func (o *CommonTargetConfiguration) SetCopyOnRunSuccessNil() {
	o.CopyOnRunSuccess.Set(nil)
}

// UnsetCopyOnRunSuccess ensures that no value is present for CopyOnRunSuccess, not even an explicit nil
func (o *CommonTargetConfiguration) UnsetCopyOnRunSuccess() {
	o.CopyOnRunSuccess.Unset()
}

// GetLogRetention returns the LogRetention field value if set, zero value otherwise.
func (o *CommonTargetConfiguration) GetLogRetention() LogRetention {
	if o == nil || IsNil(o.LogRetention) {
		var ret LogRetention
		return ret
	}
	return *o.LogRetention
}

// GetLogRetentionOk returns a tuple with the LogRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonTargetConfiguration) GetLogRetentionOk() (*LogRetention, bool) {
	if o == nil || IsNil(o.LogRetention) {
		return nil, false
	}
	return o.LogRetention, true
}

// HasLogRetention returns a boolean if a field has been set.
func (o *CommonTargetConfiguration) HasLogRetention() bool {
	if o != nil && !IsNil(o.LogRetention) {
		return true
	}

	return false
}

// SetLogRetention gets a reference to the given LogRetention and assigns it to the LogRetention field.
func (o *CommonTargetConfiguration) SetLogRetention(v LogRetention) {
	o.LogRetention = &v
}

// GetRetention returns the Retention field value
func (o *CommonTargetConfiguration) GetRetention() Retention {
	if o == nil {
		var ret Retention
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *CommonTargetConfiguration) GetRetentionOk() (*Retention, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *CommonTargetConfiguration) SetRetention(v Retention) {
	o.Retention = v
}

// GetRunTimeouts returns the RunTimeouts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonTargetConfiguration) GetRunTimeouts() []CancellationTimeoutParams {
	if o == nil {
		var ret []CancellationTimeoutParams
		return ret
	}
	return o.RunTimeouts
}

// GetRunTimeoutsOk returns a tuple with the RunTimeouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonTargetConfiguration) GetRunTimeoutsOk() ([]CancellationTimeoutParams, bool) {
	if o == nil || IsNil(o.RunTimeouts) {
		return nil, false
	}
	return o.RunTimeouts, true
}

// HasRunTimeouts returns a boolean if a field has been set.
func (o *CommonTargetConfiguration) HasRunTimeouts() bool {
	if o != nil && !IsNil(o.RunTimeouts) {
		return true
	}

	return false
}

// SetRunTimeouts gets a reference to the given []CancellationTimeoutParams and assigns it to the RunTimeouts field.
func (o *CommonTargetConfiguration) SetRunTimeouts(v []CancellationTimeoutParams) {
	o.RunTimeouts = v
}

// GetSchedule returns the Schedule field value
func (o *CommonTargetConfiguration) GetSchedule() TargetSchedule {
	if o == nil {
		var ret TargetSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *CommonTargetConfiguration) GetScheduleOk() (*TargetSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *CommonTargetConfiguration) SetSchedule(v TargetSchedule) {
	o.Schedule = v
}

func (o CommonTargetConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonTargetConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupRunType.IsSet() {
		toSerialize["backupRunType"] = o.BackupRunType.Get()
	}
	if o.ConfigId.IsSet() {
		toSerialize["configId"] = o.ConfigId.Get()
	}
	if o.CopyOnRunSuccess.IsSet() {
		toSerialize["copyOnRunSuccess"] = o.CopyOnRunSuccess.Get()
	}
	if !IsNil(o.LogRetention) {
		toSerialize["logRetention"] = o.LogRetention
	}
	toSerialize["retention"] = o.Retention
	if o.RunTimeouts != nil {
		toSerialize["runTimeouts"] = o.RunTimeouts
	}
	toSerialize["schedule"] = o.Schedule
	return toSerialize, nil
}

func (o *CommonTargetConfiguration) UnmarshalJSON(data []byte) (err error) {
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

	varCommonTargetConfiguration := _CommonTargetConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonTargetConfiguration)

	if err != nil {
		return err
	}

	*o = CommonTargetConfiguration(varCommonTargetConfiguration)

	return err
}

type NullableCommonTargetConfiguration struct {
	value *CommonTargetConfiguration
	isSet bool
}

func (v NullableCommonTargetConfiguration) Get() *CommonTargetConfiguration {
	return v.value
}

func (v *NullableCommonTargetConfiguration) Set(val *CommonTargetConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonTargetConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonTargetConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonTargetConfiguration(val *CommonTargetConfiguration) *NullableCommonTargetConfiguration {
	return &NullableCommonTargetConfiguration{value: val, isSet: true}
}

func (v NullableCommonTargetConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonTargetConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


