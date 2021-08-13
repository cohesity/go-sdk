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

// ReplicationConfig Specifies settings for copying Snapshots to Remote Clusters. This also specifies the retention policy that should be applied to Snapshots after they have been copied to the specified target.
type ReplicationConfig struct {
	Schedule TargetSchedule `json:"schedule"`
	Retention Retention `json:"retention"`
	// Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
	CopyOnRunSuccess NullableBool `json:"copyOnRunSuccess,omitempty"`
	// Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
	ConfigId NullableString `json:"configId,omitempty"`
	// Specifies the type of target to which replication need to be performed.
	TargetType string `json:"targetType"`
	RemoteTargetConfig *RemoteTargetConfig `json:"remoteTargetConfig,omitempty"`
	AwsTargetConfig *AWSTargetConfig `json:"awsTargetConfig,omitempty"`
	AzureTargetConfig *AzureTargetConfig `json:"azureTargetConfig,omitempty"`
}

// NewReplicationConfig instantiates a new ReplicationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicationConfig(schedule TargetSchedule, retention Retention, targetType string) *ReplicationConfig {
	this := ReplicationConfig{}
	this.Schedule = schedule
	this.Retention = retention
	this.TargetType = targetType
	return &this
}

// NewReplicationConfigWithDefaults instantiates a new ReplicationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationConfigWithDefaults() *ReplicationConfig {
	this := ReplicationConfig{}
	return &this
}

// GetSchedule returns the Schedule field value
func (o *ReplicationConfig) GetSchedule() TargetSchedule {
	if o == nil {
		var ret TargetSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *ReplicationConfig) GetScheduleOk() (*TargetSchedule, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *ReplicationConfig) SetSchedule(v TargetSchedule) {
	o.Schedule = v
}

// GetRetention returns the Retention field value
func (o *ReplicationConfig) GetRetention() Retention {
	if o == nil {
		var ret Retention
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *ReplicationConfig) GetRetentionOk() (*Retention, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *ReplicationConfig) SetRetention(v Retention) {
	o.Retention = v
}

// GetCopyOnRunSuccess returns the CopyOnRunSuccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplicationConfig) GetCopyOnRunSuccess() bool {
	if o == nil || o.CopyOnRunSuccess.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CopyOnRunSuccess.Get()
}

// GetCopyOnRunSuccessOk returns a tuple with the CopyOnRunSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicationConfig) GetCopyOnRunSuccessOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CopyOnRunSuccess.Get(), o.CopyOnRunSuccess.IsSet()
}

// HasCopyOnRunSuccess returns a boolean if a field has been set.
func (o *ReplicationConfig) HasCopyOnRunSuccess() bool {
	if o != nil && o.CopyOnRunSuccess.IsSet() {
		return true
	}

	return false
}

// SetCopyOnRunSuccess gets a reference to the given NullableBool and assigns it to the CopyOnRunSuccess field.
func (o *ReplicationConfig) SetCopyOnRunSuccess(v bool) {
	o.CopyOnRunSuccess.Set(&v)
}
// SetCopyOnRunSuccessNil sets the value for CopyOnRunSuccess to be an explicit nil
func (o *ReplicationConfig) SetCopyOnRunSuccessNil() {
	o.CopyOnRunSuccess.Set(nil)
}

// UnsetCopyOnRunSuccess ensures that no value is present for CopyOnRunSuccess, not even an explicit nil
func (o *ReplicationConfig) UnsetCopyOnRunSuccess() {
	o.CopyOnRunSuccess.Unset()
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReplicationConfig) GetConfigId() string {
	if o == nil || o.ConfigId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigId.Get()
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReplicationConfig) GetConfigIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConfigId.Get(), o.ConfigId.IsSet()
}

// HasConfigId returns a boolean if a field has been set.
func (o *ReplicationConfig) HasConfigId() bool {
	if o != nil && o.ConfigId.IsSet() {
		return true
	}

	return false
}

// SetConfigId gets a reference to the given NullableString and assigns it to the ConfigId field.
func (o *ReplicationConfig) SetConfigId(v string) {
	o.ConfigId.Set(&v)
}
// SetConfigIdNil sets the value for ConfigId to be an explicit nil
func (o *ReplicationConfig) SetConfigIdNil() {
	o.ConfigId.Set(nil)
}

// UnsetConfigId ensures that no value is present for ConfigId, not even an explicit nil
func (o *ReplicationConfig) UnsetConfigId() {
	o.ConfigId.Unset()
}

// GetTargetType returns the TargetType field value
func (o *ReplicationConfig) GetTargetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value
// and a boolean to check if the value has been set.
func (o *ReplicationConfig) GetTargetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetType, true
}

// SetTargetType sets field value
func (o *ReplicationConfig) SetTargetType(v string) {
	o.TargetType = v
}

// GetRemoteTargetConfig returns the RemoteTargetConfig field value if set, zero value otherwise.
func (o *ReplicationConfig) GetRemoteTargetConfig() RemoteTargetConfig {
	if o == nil || o.RemoteTargetConfig == nil {
		var ret RemoteTargetConfig
		return ret
	}
	return *o.RemoteTargetConfig
}

// GetRemoteTargetConfigOk returns a tuple with the RemoteTargetConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationConfig) GetRemoteTargetConfigOk() (*RemoteTargetConfig, bool) {
	if o == nil || o.RemoteTargetConfig == nil {
		return nil, false
	}
	return o.RemoteTargetConfig, true
}

// HasRemoteTargetConfig returns a boolean if a field has been set.
func (o *ReplicationConfig) HasRemoteTargetConfig() bool {
	if o != nil && o.RemoteTargetConfig != nil {
		return true
	}

	return false
}

// SetRemoteTargetConfig gets a reference to the given RemoteTargetConfig and assigns it to the RemoteTargetConfig field.
func (o *ReplicationConfig) SetRemoteTargetConfig(v RemoteTargetConfig) {
	o.RemoteTargetConfig = &v
}

// GetAwsTargetConfig returns the AwsTargetConfig field value if set, zero value otherwise.
func (o *ReplicationConfig) GetAwsTargetConfig() AWSTargetConfig {
	if o == nil || o.AwsTargetConfig == nil {
		var ret AWSTargetConfig
		return ret
	}
	return *o.AwsTargetConfig
}

// GetAwsTargetConfigOk returns a tuple with the AwsTargetConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationConfig) GetAwsTargetConfigOk() (*AWSTargetConfig, bool) {
	if o == nil || o.AwsTargetConfig == nil {
		return nil, false
	}
	return o.AwsTargetConfig, true
}

// HasAwsTargetConfig returns a boolean if a field has been set.
func (o *ReplicationConfig) HasAwsTargetConfig() bool {
	if o != nil && o.AwsTargetConfig != nil {
		return true
	}

	return false
}

// SetAwsTargetConfig gets a reference to the given AWSTargetConfig and assigns it to the AwsTargetConfig field.
func (o *ReplicationConfig) SetAwsTargetConfig(v AWSTargetConfig) {
	o.AwsTargetConfig = &v
}

// GetAzureTargetConfig returns the AzureTargetConfig field value if set, zero value otherwise.
func (o *ReplicationConfig) GetAzureTargetConfig() AzureTargetConfig {
	if o == nil || o.AzureTargetConfig == nil {
		var ret AzureTargetConfig
		return ret
	}
	return *o.AzureTargetConfig
}

// GetAzureTargetConfigOk returns a tuple with the AzureTargetConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationConfig) GetAzureTargetConfigOk() (*AzureTargetConfig, bool) {
	if o == nil || o.AzureTargetConfig == nil {
		return nil, false
	}
	return o.AzureTargetConfig, true
}

// HasAzureTargetConfig returns a boolean if a field has been set.
func (o *ReplicationConfig) HasAzureTargetConfig() bool {
	if o != nil && o.AzureTargetConfig != nil {
		return true
	}

	return false
}

// SetAzureTargetConfig gets a reference to the given AzureTargetConfig and assigns it to the AzureTargetConfig field.
func (o *ReplicationConfig) SetAzureTargetConfig(v AzureTargetConfig) {
	o.AzureTargetConfig = &v
}

func (o ReplicationConfig) MarshalJSON() ([]byte, error) {
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
		toSerialize["targetType"] = o.TargetType
	}
	if o.RemoteTargetConfig != nil {
		toSerialize["remoteTargetConfig"] = o.RemoteTargetConfig
	}
	if o.AwsTargetConfig != nil {
		toSerialize["awsTargetConfig"] = o.AwsTargetConfig
	}
	if o.AzureTargetConfig != nil {
		toSerialize["azureTargetConfig"] = o.AzureTargetConfig
	}
	return json.Marshal(toSerialize)
}

type NullableReplicationConfig struct {
	value *ReplicationConfig
	isSet bool
}

func (v NullableReplicationConfig) Get() *ReplicationConfig {
	return v.value
}

func (v *NullableReplicationConfig) Set(val *ReplicationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicationConfig(val *ReplicationConfig) *NullableReplicationConfig {
	return &NullableReplicationConfig{value: val, isSet: true}
}

func (v NullableReplicationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o ReplicationConfig) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}