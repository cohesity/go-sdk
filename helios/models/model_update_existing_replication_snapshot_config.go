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

// UpdateExistingReplicationSnapshotConfig Specifies the configuration about updating an existing Replication Snapshot Run.
type UpdateExistingReplicationSnapshotConfig struct {
	// Specifies the cluster id of the replication cluster.
	Id int64 `json:"id"`
	// Specifies whether to retain the snapshot for legal purpose. If set to true, the snapshots cannot be deleted until the retention period. Note that using this option may cause the Cluster to run out of space. If set to false explicitly, the hold is removed, and the snapshots will expire as specified in the policy of the Protection Group. If this field is not specified, there is no change to the hold of the run. This field can be set only by a User having Data Security Role.
	EnableLegalHold NullableBool `json:"enableLegalHold,omitempty"`
	// Specifies whether to delete the snapshot. When this is set to true, all other params will be ignored.
	DeleteSnapshot NullableBool `json:"deleteSnapshot,omitempty"`
	// Specifies whether to retry the replication operation in case if earlier attempt failed. If not specified or set to false, replication is not retried.
	Resync NullableBool `json:"resync,omitempty"`
	// Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept until the maximum of the snapshot retention time. During that time, the snapshots cannot be deleted. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.
	DataLock NullableString `json:"dataLock,omitempty"`
	// Specifies number of days to retain the snapshots. If positive, then this value is added to exisiting expiry time thereby increasing  the retention period of the snapshot. Conversly, if this value is negative, then value is subtracted to existing expiry time thereby decreasing the retention period of the snaphot. Here, by this operation if expiry time goes below current time then snapshot is immediately deleted.
	DaysToKeep NullableInt64 `json:"daysToKeep,omitempty"`
}

// NewUpdateExistingReplicationSnapshotConfig instantiates a new UpdateExistingReplicationSnapshotConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateExistingReplicationSnapshotConfig(id int64) *UpdateExistingReplicationSnapshotConfig {
	this := UpdateExistingReplicationSnapshotConfig{}
	this.Id = id
	return &this
}

// NewUpdateExistingReplicationSnapshotConfigWithDefaults instantiates a new UpdateExistingReplicationSnapshotConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateExistingReplicationSnapshotConfigWithDefaults() *UpdateExistingReplicationSnapshotConfig {
	this := UpdateExistingReplicationSnapshotConfig{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateExistingReplicationSnapshotConfig) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateExistingReplicationSnapshotConfig) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateExistingReplicationSnapshotConfig) SetId(v int64) {
	o.Id = v
}

// GetEnableLegalHold returns the EnableLegalHold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateExistingReplicationSnapshotConfig) GetEnableLegalHold() bool {
	if o == nil || o.EnableLegalHold.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EnableLegalHold.Get()
}

// GetEnableLegalHoldOk returns a tuple with the EnableLegalHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateExistingReplicationSnapshotConfig) GetEnableLegalHoldOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EnableLegalHold.Get(), o.EnableLegalHold.IsSet()
}

// HasEnableLegalHold returns a boolean if a field has been set.
func (o *UpdateExistingReplicationSnapshotConfig) HasEnableLegalHold() bool {
	if o != nil && o.EnableLegalHold.IsSet() {
		return true
	}

	return false
}

// SetEnableLegalHold gets a reference to the given NullableBool and assigns it to the EnableLegalHold field.
func (o *UpdateExistingReplicationSnapshotConfig) SetEnableLegalHold(v bool) {
	o.EnableLegalHold.Set(&v)
}
// SetEnableLegalHoldNil sets the value for EnableLegalHold to be an explicit nil
func (o *UpdateExistingReplicationSnapshotConfig) SetEnableLegalHoldNil() {
	o.EnableLegalHold.Set(nil)
}

// UnsetEnableLegalHold ensures that no value is present for EnableLegalHold, not even an explicit nil
func (o *UpdateExistingReplicationSnapshotConfig) UnsetEnableLegalHold() {
	o.EnableLegalHold.Unset()
}

// GetDeleteSnapshot returns the DeleteSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateExistingReplicationSnapshotConfig) GetDeleteSnapshot() bool {
	if o == nil || o.DeleteSnapshot.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DeleteSnapshot.Get()
}

// GetDeleteSnapshotOk returns a tuple with the DeleteSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateExistingReplicationSnapshotConfig) GetDeleteSnapshotOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeleteSnapshot.Get(), o.DeleteSnapshot.IsSet()
}

// HasDeleteSnapshot returns a boolean if a field has been set.
func (o *UpdateExistingReplicationSnapshotConfig) HasDeleteSnapshot() bool {
	if o != nil && o.DeleteSnapshot.IsSet() {
		return true
	}

	return false
}

// SetDeleteSnapshot gets a reference to the given NullableBool and assigns it to the DeleteSnapshot field.
func (o *UpdateExistingReplicationSnapshotConfig) SetDeleteSnapshot(v bool) {
	o.DeleteSnapshot.Set(&v)
}
// SetDeleteSnapshotNil sets the value for DeleteSnapshot to be an explicit nil
func (o *UpdateExistingReplicationSnapshotConfig) SetDeleteSnapshotNil() {
	o.DeleteSnapshot.Set(nil)
}

// UnsetDeleteSnapshot ensures that no value is present for DeleteSnapshot, not even an explicit nil
func (o *UpdateExistingReplicationSnapshotConfig) UnsetDeleteSnapshot() {
	o.DeleteSnapshot.Unset()
}

// GetResync returns the Resync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateExistingReplicationSnapshotConfig) GetResync() bool {
	if o == nil || o.Resync.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Resync.Get()
}

// GetResyncOk returns a tuple with the Resync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateExistingReplicationSnapshotConfig) GetResyncOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Resync.Get(), o.Resync.IsSet()
}

// HasResync returns a boolean if a field has been set.
func (o *UpdateExistingReplicationSnapshotConfig) HasResync() bool {
	if o != nil && o.Resync.IsSet() {
		return true
	}

	return false
}

// SetResync gets a reference to the given NullableBool and assigns it to the Resync field.
func (o *UpdateExistingReplicationSnapshotConfig) SetResync(v bool) {
	o.Resync.Set(&v)
}
// SetResyncNil sets the value for Resync to be an explicit nil
func (o *UpdateExistingReplicationSnapshotConfig) SetResyncNil() {
	o.Resync.Set(nil)
}

// UnsetResync ensures that no value is present for Resync, not even an explicit nil
func (o *UpdateExistingReplicationSnapshotConfig) UnsetResync() {
	o.Resync.Unset()
}

// GetDataLock returns the DataLock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateExistingReplicationSnapshotConfig) GetDataLock() string {
	if o == nil || o.DataLock.Get() == nil {
		var ret string
		return ret
	}
	return *o.DataLock.Get()
}

// GetDataLockOk returns a tuple with the DataLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateExistingReplicationSnapshotConfig) GetDataLockOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DataLock.Get(), o.DataLock.IsSet()
}

// HasDataLock returns a boolean if a field has been set.
func (o *UpdateExistingReplicationSnapshotConfig) HasDataLock() bool {
	if o != nil && o.DataLock.IsSet() {
		return true
	}

	return false
}

// SetDataLock gets a reference to the given NullableString and assigns it to the DataLock field.
func (o *UpdateExistingReplicationSnapshotConfig) SetDataLock(v string) {
	o.DataLock.Set(&v)
}
// SetDataLockNil sets the value for DataLock to be an explicit nil
func (o *UpdateExistingReplicationSnapshotConfig) SetDataLockNil() {
	o.DataLock.Set(nil)
}

// UnsetDataLock ensures that no value is present for DataLock, not even an explicit nil
func (o *UpdateExistingReplicationSnapshotConfig) UnsetDataLock() {
	o.DataLock.Unset()
}

// GetDaysToKeep returns the DaysToKeep field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateExistingReplicationSnapshotConfig) GetDaysToKeep() int64 {
	if o == nil || o.DaysToKeep.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DaysToKeep.Get()
}

// GetDaysToKeepOk returns a tuple with the DaysToKeep field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateExistingReplicationSnapshotConfig) GetDaysToKeepOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DaysToKeep.Get(), o.DaysToKeep.IsSet()
}

// HasDaysToKeep returns a boolean if a field has been set.
func (o *UpdateExistingReplicationSnapshotConfig) HasDaysToKeep() bool {
	if o != nil && o.DaysToKeep.IsSet() {
		return true
	}

	return false
}

// SetDaysToKeep gets a reference to the given NullableInt64 and assigns it to the DaysToKeep field.
func (o *UpdateExistingReplicationSnapshotConfig) SetDaysToKeep(v int64) {
	o.DaysToKeep.Set(&v)
}
// SetDaysToKeepNil sets the value for DaysToKeep to be an explicit nil
func (o *UpdateExistingReplicationSnapshotConfig) SetDaysToKeepNil() {
	o.DaysToKeep.Set(nil)
}

// UnsetDaysToKeep ensures that no value is present for DaysToKeep, not even an explicit nil
func (o *UpdateExistingReplicationSnapshotConfig) UnsetDaysToKeep() {
	o.DaysToKeep.Unset()
}

func (o UpdateExistingReplicationSnapshotConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.EnableLegalHold.IsSet() {
		toSerialize["enableLegalHold"] = o.EnableLegalHold.Get()
	}
	if o.DeleteSnapshot.IsSet() {
		toSerialize["deleteSnapshot"] = o.DeleteSnapshot.Get()
	}
	if o.Resync.IsSet() {
		toSerialize["resync"] = o.Resync.Get()
	}
	if o.DataLock.IsSet() {
		toSerialize["dataLock"] = o.DataLock.Get()
	}
	if o.DaysToKeep.IsSet() {
		toSerialize["daysToKeep"] = o.DaysToKeep.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateExistingReplicationSnapshotConfig struct {
	value *UpdateExistingReplicationSnapshotConfig
	isSet bool
}

func (v NullableUpdateExistingReplicationSnapshotConfig) Get() *UpdateExistingReplicationSnapshotConfig {
	return v.value
}

func (v *NullableUpdateExistingReplicationSnapshotConfig) Set(val *UpdateExistingReplicationSnapshotConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateExistingReplicationSnapshotConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateExistingReplicationSnapshotConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateExistingReplicationSnapshotConfig(val *UpdateExistingReplicationSnapshotConfig) *NullableUpdateExistingReplicationSnapshotConfig {
	return &NullableUpdateExistingReplicationSnapshotConfig{value: val, isSet: true}
}

func (v NullableUpdateExistingReplicationSnapshotConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateExistingReplicationSnapshotConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


