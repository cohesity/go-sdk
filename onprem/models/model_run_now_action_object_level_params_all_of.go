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

// RunNowActionObjectLevelParamsAllOf struct for RunNowActionObjectLevelParamsAllOf
type RunNowActionObjectLevelParamsAllOf struct {
	// If sepcified as true then runNow will only take local snapshot ignoring all other targets such as replication, archivals etc. If not sepcified or specified as false then runNow will follow the policy settings.
	TakeLocalSnapshotOnly NullableBool `json:"takeLocalSnapshotOnly,omitempty"`
	// Specifies the backup type should be used for RunNow action.
	BackupType *string `json:"backupType,omitempty"`
}

// NewRunNowActionObjectLevelParamsAllOf instantiates a new RunNowActionObjectLevelParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunNowActionObjectLevelParamsAllOf() *RunNowActionObjectLevelParamsAllOf {
	this := RunNowActionObjectLevelParamsAllOf{}
	return &this
}

// NewRunNowActionObjectLevelParamsAllOfWithDefaults instantiates a new RunNowActionObjectLevelParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunNowActionObjectLevelParamsAllOfWithDefaults() *RunNowActionObjectLevelParamsAllOf {
	this := RunNowActionObjectLevelParamsAllOf{}
	return &this
}

// GetTakeLocalSnapshotOnly returns the TakeLocalSnapshotOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunNowActionObjectLevelParamsAllOf) GetTakeLocalSnapshotOnly() bool {
	if o == nil || o.TakeLocalSnapshotOnly.Get() == nil {
		var ret bool
		return ret
	}
	return *o.TakeLocalSnapshotOnly.Get()
}

// GetTakeLocalSnapshotOnlyOk returns a tuple with the TakeLocalSnapshotOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunNowActionObjectLevelParamsAllOf) GetTakeLocalSnapshotOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TakeLocalSnapshotOnly.Get(), o.TakeLocalSnapshotOnly.IsSet()
}

// HasTakeLocalSnapshotOnly returns a boolean if a field has been set.
func (o *RunNowActionObjectLevelParamsAllOf) HasTakeLocalSnapshotOnly() bool {
	if o != nil && o.TakeLocalSnapshotOnly.IsSet() {
		return true
	}

	return false
}

// SetTakeLocalSnapshotOnly gets a reference to the given NullableBool and assigns it to the TakeLocalSnapshotOnly field.
func (o *RunNowActionObjectLevelParamsAllOf) SetTakeLocalSnapshotOnly(v bool) {
	o.TakeLocalSnapshotOnly.Set(&v)
}
// SetTakeLocalSnapshotOnlyNil sets the value for TakeLocalSnapshotOnly to be an explicit nil
func (o *RunNowActionObjectLevelParamsAllOf) SetTakeLocalSnapshotOnlyNil() {
	o.TakeLocalSnapshotOnly.Set(nil)
}

// UnsetTakeLocalSnapshotOnly ensures that no value is present for TakeLocalSnapshotOnly, not even an explicit nil
func (o *RunNowActionObjectLevelParamsAllOf) UnsetTakeLocalSnapshotOnly() {
	o.TakeLocalSnapshotOnly.Unset()
}

// GetBackupType returns the BackupType field value if set, zero value otherwise.
func (o *RunNowActionObjectLevelParamsAllOf) GetBackupType() string {
	if o == nil || o.BackupType == nil {
		var ret string
		return ret
	}
	return *o.BackupType
}

// GetBackupTypeOk returns a tuple with the BackupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunNowActionObjectLevelParamsAllOf) GetBackupTypeOk() (*string, bool) {
	if o == nil || o.BackupType == nil {
		return nil, false
	}
	return o.BackupType, true
}

// HasBackupType returns a boolean if a field has been set.
func (o *RunNowActionObjectLevelParamsAllOf) HasBackupType() bool {
	if o != nil && o.BackupType != nil {
		return true
	}

	return false
}

// SetBackupType gets a reference to the given string and assigns it to the BackupType field.
func (o *RunNowActionObjectLevelParamsAllOf) SetBackupType(v string) {
	o.BackupType = &v
}

func (o RunNowActionObjectLevelParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TakeLocalSnapshotOnly.IsSet() {
		toSerialize["takeLocalSnapshotOnly"] = o.TakeLocalSnapshotOnly.Get()
	}
	if o.BackupType != nil {
		toSerialize["backupType"] = o.BackupType
	}
	return json.Marshal(toSerialize)
}

type NullableRunNowActionObjectLevelParamsAllOf struct {
	value *RunNowActionObjectLevelParamsAllOf
	isSet bool
}

func (v NullableRunNowActionObjectLevelParamsAllOf) Get() *RunNowActionObjectLevelParamsAllOf {
	return v.value
}

func (v *NullableRunNowActionObjectLevelParamsAllOf) Set(val *RunNowActionObjectLevelParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRunNowActionObjectLevelParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRunNowActionObjectLevelParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunNowActionObjectLevelParamsAllOf(val *RunNowActionObjectLevelParamsAllOf) *NullableRunNowActionObjectLevelParamsAllOf {
	return &NullableRunNowActionObjectLevelParamsAllOf{value: val, isSet: true}
}

func (v NullableRunNowActionObjectLevelParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunNowActionObjectLevelParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o RunNowActionObjectLevelParamsAllOf) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}