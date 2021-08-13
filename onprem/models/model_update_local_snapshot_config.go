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

// UpdateLocalSnapshotConfig Specifies the params to perform actions on local snapshot taken by a Protection Group Run.
type UpdateLocalSnapshotConfig struct {
	// Specifies whether to retain the snapshot for legal purpose. If set to true, the snapshots cannot be deleted until the retention period. Note that using this option may cause the Cluster to run out of space. If set to false explicitly, the hold is removed, and the snapshots will expire as specified in the policy of the Protection Group. If this field is not specified, there is no change to the hold of the run. This field can be set only by a User having Data Security Role.
	EnableLegalHold NullableBool `json:"enableLegalHold,omitempty"`
	// Specifies whether to delete the snapshot. When this is set to true, all other params will be ignored.
	DeleteSnapshot NullableBool `json:"deleteSnapshot,omitempty"`
	// Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept until the maximum of the snapshot retention time. During that time, the snapshots cannot be deleted. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.
	DataLock NullableString `json:"dataLock,omitempty"`
	// Specifies number of days to retain the snapshots. If positive, then this value is added to exisiting expiry time thereby increasing  the retention period of the snapshot. Conversly, if this value is negative, then value is subtracted to existing expiry time thereby decreasing the retention period of the snaphot. Here, by this operation if expiry time goes below current time then snapshot is immediately deleted.
	DaysToKeep NullableInt64 `json:"daysToKeep,omitempty"`
}

// NewUpdateLocalSnapshotConfig instantiates a new UpdateLocalSnapshotConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLocalSnapshotConfig() *UpdateLocalSnapshotConfig {
	this := UpdateLocalSnapshotConfig{}
	return &this
}

// NewUpdateLocalSnapshotConfigWithDefaults instantiates a new UpdateLocalSnapshotConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLocalSnapshotConfigWithDefaults() *UpdateLocalSnapshotConfig {
	this := UpdateLocalSnapshotConfig{}
	return &this
}

// GetEnableLegalHold returns the EnableLegalHold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateLocalSnapshotConfig) GetEnableLegalHold() bool {
	if o == nil || o.EnableLegalHold.Get() == nil {
		var ret bool
		return ret
	}
	return *o.EnableLegalHold.Get()
}

// GetEnableLegalHoldOk returns a tuple with the EnableLegalHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateLocalSnapshotConfig) GetEnableLegalHoldOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EnableLegalHold.Get(), o.EnableLegalHold.IsSet()
}

// HasEnableLegalHold returns a boolean if a field has been set.
func (o *UpdateLocalSnapshotConfig) HasEnableLegalHold() bool {
	if o != nil && o.EnableLegalHold.IsSet() {
		return true
	}

	return false
}

// SetEnableLegalHold gets a reference to the given NullableBool and assigns it to the EnableLegalHold field.
func (o *UpdateLocalSnapshotConfig) SetEnableLegalHold(v bool) {
	o.EnableLegalHold.Set(&v)
}
// SetEnableLegalHoldNil sets the value for EnableLegalHold to be an explicit nil
func (o *UpdateLocalSnapshotConfig) SetEnableLegalHoldNil() {
	o.EnableLegalHold.Set(nil)
}

// UnsetEnableLegalHold ensures that no value is present for EnableLegalHold, not even an explicit nil
func (o *UpdateLocalSnapshotConfig) UnsetEnableLegalHold() {
	o.EnableLegalHold.Unset()
}

// GetDeleteSnapshot returns the DeleteSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateLocalSnapshotConfig) GetDeleteSnapshot() bool {
	if o == nil || o.DeleteSnapshot.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DeleteSnapshot.Get()
}

// GetDeleteSnapshotOk returns a tuple with the DeleteSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateLocalSnapshotConfig) GetDeleteSnapshotOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeleteSnapshot.Get(), o.DeleteSnapshot.IsSet()
}

// HasDeleteSnapshot returns a boolean if a field has been set.
func (o *UpdateLocalSnapshotConfig) HasDeleteSnapshot() bool {
	if o != nil && o.DeleteSnapshot.IsSet() {
		return true
	}

	return false
}

// SetDeleteSnapshot gets a reference to the given NullableBool and assigns it to the DeleteSnapshot field.
func (o *UpdateLocalSnapshotConfig) SetDeleteSnapshot(v bool) {
	o.DeleteSnapshot.Set(&v)
}
// SetDeleteSnapshotNil sets the value for DeleteSnapshot to be an explicit nil
func (o *UpdateLocalSnapshotConfig) SetDeleteSnapshotNil() {
	o.DeleteSnapshot.Set(nil)
}

// UnsetDeleteSnapshot ensures that no value is present for DeleteSnapshot, not even an explicit nil
func (o *UpdateLocalSnapshotConfig) UnsetDeleteSnapshot() {
	o.DeleteSnapshot.Unset()
}

// GetDataLock returns the DataLock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateLocalSnapshotConfig) GetDataLock() string {
	if o == nil || o.DataLock.Get() == nil {
		var ret string
		return ret
	}
	return *o.DataLock.Get()
}

// GetDataLockOk returns a tuple with the DataLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateLocalSnapshotConfig) GetDataLockOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DataLock.Get(), o.DataLock.IsSet()
}

// HasDataLock returns a boolean if a field has been set.
func (o *UpdateLocalSnapshotConfig) HasDataLock() bool {
	if o != nil && o.DataLock.IsSet() {
		return true
	}

	return false
}

// SetDataLock gets a reference to the given NullableString and assigns it to the DataLock field.
func (o *UpdateLocalSnapshotConfig) SetDataLock(v string) {
	o.DataLock.Set(&v)
}
// SetDataLockNil sets the value for DataLock to be an explicit nil
func (o *UpdateLocalSnapshotConfig) SetDataLockNil() {
	o.DataLock.Set(nil)
}

// UnsetDataLock ensures that no value is present for DataLock, not even an explicit nil
func (o *UpdateLocalSnapshotConfig) UnsetDataLock() {
	o.DataLock.Unset()
}

// GetDaysToKeep returns the DaysToKeep field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateLocalSnapshotConfig) GetDaysToKeep() int64 {
	if o == nil || o.DaysToKeep.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DaysToKeep.Get()
}

// GetDaysToKeepOk returns a tuple with the DaysToKeep field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateLocalSnapshotConfig) GetDaysToKeepOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DaysToKeep.Get(), o.DaysToKeep.IsSet()
}

// HasDaysToKeep returns a boolean if a field has been set.
func (o *UpdateLocalSnapshotConfig) HasDaysToKeep() bool {
	if o != nil && o.DaysToKeep.IsSet() {
		return true
	}

	return false
}

// SetDaysToKeep gets a reference to the given NullableInt64 and assigns it to the DaysToKeep field.
func (o *UpdateLocalSnapshotConfig) SetDaysToKeep(v int64) {
	o.DaysToKeep.Set(&v)
}
// SetDaysToKeepNil sets the value for DaysToKeep to be an explicit nil
func (o *UpdateLocalSnapshotConfig) SetDaysToKeepNil() {
	o.DaysToKeep.Set(nil)
}

// UnsetDaysToKeep ensures that no value is present for DaysToKeep, not even an explicit nil
func (o *UpdateLocalSnapshotConfig) UnsetDaysToKeep() {
	o.DaysToKeep.Unset()
}

func (o UpdateLocalSnapshotConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnableLegalHold.IsSet() {
		toSerialize["enableLegalHold"] = o.EnableLegalHold.Get()
	}
	if o.DeleteSnapshot.IsSet() {
		toSerialize["deleteSnapshot"] = o.DeleteSnapshot.Get()
	}
	if o.DataLock.IsSet() {
		toSerialize["dataLock"] = o.DataLock.Get()
	}
	if o.DaysToKeep.IsSet() {
		toSerialize["daysToKeep"] = o.DaysToKeep.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateLocalSnapshotConfig struct {
	value *UpdateLocalSnapshotConfig
	isSet bool
}

func (v NullableUpdateLocalSnapshotConfig) Get() *UpdateLocalSnapshotConfig {
	return v.value
}

func (v *NullableUpdateLocalSnapshotConfig) Set(val *UpdateLocalSnapshotConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLocalSnapshotConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLocalSnapshotConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLocalSnapshotConfig(val *UpdateLocalSnapshotConfig) *NullableUpdateLocalSnapshotConfig {
	return &NullableUpdateLocalSnapshotConfig{value: val, isSet: true}
}

func (v NullableUpdateLocalSnapshotConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLocalSnapshotConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o UpdateLocalSnapshotConfig) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}