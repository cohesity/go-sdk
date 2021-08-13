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

// FailoverRunConfiguration Specifies the configuration required for execting special run as a part of failover workflow. This special run is triggered during palnned failover to sync the source cluster to replication cluster with minimum possible delta. Please note that if this object is passed then this special run will ignore the other archivals and retention settings.
type FailoverRunConfiguration struct {
	// Specifies the replication cluster Id where planned run will replicate objects.
	ReplicationClusterId NullableInt64 `json:"replicationClusterId"`
	// Specifies the list of all local entity ids of all the objects being failed from the source cluster.
	Objects []FailoverObject `json:"objects"`
	// Specifies the active protection group id on the source cluster from where the objects are being failed over.
	ProtectionGroupId NullableString `json:"protectionGroupId"`
	// Specifies the type of the backup run to be triggered by this request. If this is not set defaults to incremental backup.
	RunType *string `json:"runType,omitempty"`
	// If failover is initiated by view based orchastrator, then this field specifies the local view id of source cluster which is being failed over.
	ViewId NullableInt64 `json:"viewId,omitempty"`
	// If set to true, other ongoing runs backing up the same set of entities being failed over will be initiated for cancellation. Non conflicting run operations such as replications to other clusters, archivals will not be cancelled. If set to false, then new run will wait for all the pending operations to finish normally before scheduling a new backup/replication.
	CancelNonFailoverRuns NullableBool `json:"cancelNonFailoverRuns,omitempty"`
	// If this is set to true then unless failover operation is completed, all the next runs will be pasued.
	PauseNextRuns NullableBool `json:"pauseNextRuns,omitempty"`
}

// NewFailoverRunConfiguration instantiates a new FailoverRunConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailoverRunConfiguration(replicationClusterId NullableInt64, objects []FailoverObject, protectionGroupId NullableString) *FailoverRunConfiguration {
	this := FailoverRunConfiguration{}
	this.ReplicationClusterId = replicationClusterId
	this.Objects = objects
	this.ProtectionGroupId = protectionGroupId
	return &this
}

// NewFailoverRunConfigurationWithDefaults instantiates a new FailoverRunConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailoverRunConfigurationWithDefaults() *FailoverRunConfiguration {
	this := FailoverRunConfiguration{}
	return &this
}

// GetReplicationClusterId returns the ReplicationClusterId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *FailoverRunConfiguration) GetReplicationClusterId() int64 {
	if o == nil || o.ReplicationClusterId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.ReplicationClusterId.Get()
}

// GetReplicationClusterIdOk returns a tuple with the ReplicationClusterId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FailoverRunConfiguration) GetReplicationClusterIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReplicationClusterId.Get(), o.ReplicationClusterId.IsSet()
}

// SetReplicationClusterId sets field value
func (o *FailoverRunConfiguration) SetReplicationClusterId(v int64) {
	o.ReplicationClusterId.Set(&v)
}

// GetObjects returns the Objects field value
// If the value is explicit nil, the zero value for []FailoverObject will be returned
func (o *FailoverRunConfiguration) GetObjects() []FailoverObject {
	if o == nil {
		var ret []FailoverObject
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FailoverRunConfiguration) GetObjectsOk() (*[]FailoverObject, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value
func (o *FailoverRunConfiguration) SetObjects(v []FailoverObject) {
	o.Objects = v
}

// GetProtectionGroupId returns the ProtectionGroupId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FailoverRunConfiguration) GetProtectionGroupId() string {
	if o == nil || o.ProtectionGroupId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProtectionGroupId.Get()
}

// GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FailoverRunConfiguration) GetProtectionGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProtectionGroupId.Get(), o.ProtectionGroupId.IsSet()
}

// SetProtectionGroupId sets field value
func (o *FailoverRunConfiguration) SetProtectionGroupId(v string) {
	o.ProtectionGroupId.Set(&v)
}

// GetRunType returns the RunType field value if set, zero value otherwise.
func (o *FailoverRunConfiguration) GetRunType() string {
	if o == nil || o.RunType == nil {
		var ret string
		return ret
	}
	return *o.RunType
}

// GetRunTypeOk returns a tuple with the RunType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailoverRunConfiguration) GetRunTypeOk() (*string, bool) {
	if o == nil || o.RunType == nil {
		return nil, false
	}
	return o.RunType, true
}

// HasRunType returns a boolean if a field has been set.
func (o *FailoverRunConfiguration) HasRunType() bool {
	if o != nil && o.RunType != nil {
		return true
	}

	return false
}

// SetRunType gets a reference to the given string and assigns it to the RunType field.
func (o *FailoverRunConfiguration) SetRunType(v string) {
	o.RunType = &v
}

// GetViewId returns the ViewId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FailoverRunConfiguration) GetViewId() int64 {
	if o == nil || o.ViewId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ViewId.Get()
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FailoverRunConfiguration) GetViewIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ViewId.Get(), o.ViewId.IsSet()
}

// HasViewId returns a boolean if a field has been set.
func (o *FailoverRunConfiguration) HasViewId() bool {
	if o != nil && o.ViewId.IsSet() {
		return true
	}

	return false
}

// SetViewId gets a reference to the given NullableInt64 and assigns it to the ViewId field.
func (o *FailoverRunConfiguration) SetViewId(v int64) {
	o.ViewId.Set(&v)
}
// SetViewIdNil sets the value for ViewId to be an explicit nil
func (o *FailoverRunConfiguration) SetViewIdNil() {
	o.ViewId.Set(nil)
}

// UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
func (o *FailoverRunConfiguration) UnsetViewId() {
	o.ViewId.Unset()
}

// GetCancelNonFailoverRuns returns the CancelNonFailoverRuns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FailoverRunConfiguration) GetCancelNonFailoverRuns() bool {
	if o == nil || o.CancelNonFailoverRuns.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CancelNonFailoverRuns.Get()
}

// GetCancelNonFailoverRunsOk returns a tuple with the CancelNonFailoverRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FailoverRunConfiguration) GetCancelNonFailoverRunsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CancelNonFailoverRuns.Get(), o.CancelNonFailoverRuns.IsSet()
}

// HasCancelNonFailoverRuns returns a boolean if a field has been set.
func (o *FailoverRunConfiguration) HasCancelNonFailoverRuns() bool {
	if o != nil && o.CancelNonFailoverRuns.IsSet() {
		return true
	}

	return false
}

// SetCancelNonFailoverRuns gets a reference to the given NullableBool and assigns it to the CancelNonFailoverRuns field.
func (o *FailoverRunConfiguration) SetCancelNonFailoverRuns(v bool) {
	o.CancelNonFailoverRuns.Set(&v)
}
// SetCancelNonFailoverRunsNil sets the value for CancelNonFailoverRuns to be an explicit nil
func (o *FailoverRunConfiguration) SetCancelNonFailoverRunsNil() {
	o.CancelNonFailoverRuns.Set(nil)
}

// UnsetCancelNonFailoverRuns ensures that no value is present for CancelNonFailoverRuns, not even an explicit nil
func (o *FailoverRunConfiguration) UnsetCancelNonFailoverRuns() {
	o.CancelNonFailoverRuns.Unset()
}

// GetPauseNextRuns returns the PauseNextRuns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FailoverRunConfiguration) GetPauseNextRuns() bool {
	if o == nil || o.PauseNextRuns.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PauseNextRuns.Get()
}

// GetPauseNextRunsOk returns a tuple with the PauseNextRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FailoverRunConfiguration) GetPauseNextRunsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PauseNextRuns.Get(), o.PauseNextRuns.IsSet()
}

// HasPauseNextRuns returns a boolean if a field has been set.
func (o *FailoverRunConfiguration) HasPauseNextRuns() bool {
	if o != nil && o.PauseNextRuns.IsSet() {
		return true
	}

	return false
}

// SetPauseNextRuns gets a reference to the given NullableBool and assigns it to the PauseNextRuns field.
func (o *FailoverRunConfiguration) SetPauseNextRuns(v bool) {
	o.PauseNextRuns.Set(&v)
}
// SetPauseNextRunsNil sets the value for PauseNextRuns to be an explicit nil
func (o *FailoverRunConfiguration) SetPauseNextRunsNil() {
	o.PauseNextRuns.Set(nil)
}

// UnsetPauseNextRuns ensures that no value is present for PauseNextRuns, not even an explicit nil
func (o *FailoverRunConfiguration) UnsetPauseNextRuns() {
	o.PauseNextRuns.Unset()
}

func (o FailoverRunConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["replicationClusterId"] = o.ReplicationClusterId.Get()
	}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if true {
		toSerialize["protectionGroupId"] = o.ProtectionGroupId.Get()
	}
	if o.RunType != nil {
		toSerialize["runType"] = o.RunType
	}
	if o.ViewId.IsSet() {
		toSerialize["viewId"] = o.ViewId.Get()
	}
	if o.CancelNonFailoverRuns.IsSet() {
		toSerialize["cancelNonFailoverRuns"] = o.CancelNonFailoverRuns.Get()
	}
	if o.PauseNextRuns.IsSet() {
		toSerialize["pauseNextRuns"] = o.PauseNextRuns.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFailoverRunConfiguration struct {
	value *FailoverRunConfiguration
	isSet bool
}

func (v NullableFailoverRunConfiguration) Get() *FailoverRunConfiguration {
	return v.value
}

func (v *NullableFailoverRunConfiguration) Set(val *FailoverRunConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableFailoverRunConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableFailoverRunConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailoverRunConfiguration(val *FailoverRunConfiguration) *NullableFailoverRunConfiguration {
	return &NullableFailoverRunConfiguration{value: val, isSet: true}
}

func (v NullableFailoverRunConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailoverRunConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


