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

// MSSQLVolumeProtectionGroupParamsAllOf struct for MSSQLVolumeProtectionGroupParamsAllOf
type MSSQLVolumeProtectionGroupParamsAllOf struct {
	// Specifies the list of object ids to be protected.
	Objects []MSSQLVolumeProtectionGroupObjectParams `json:"objects"`
	// Specifies whether or to perform incremental backups the first time after a server restarts. By default, a full backup will be performed.
	IncrementalBackupAfterRestart NullableBool `json:"incrementalBackupAfterRestart,omitempty"`
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`
	// Specifies whether to only backup volumes on which the specified databases reside. If not specified (default), all the volumes of the host will be protected.
	BackupDbVolumesOnly NullableBool `json:"backupDbVolumesOnly,omitempty"`
	// Specifies settings which are to be applied to specific host containers in this protection group.
	AdditionalHostParams *[]MSSQLVolumeProtectionGroupHostParams `json:"additionalHostParams,omitempty"`
}

// NewMSSQLVolumeProtectionGroupParamsAllOf instantiates a new MSSQLVolumeProtectionGroupParamsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMSSQLVolumeProtectionGroupParamsAllOf(objects []MSSQLVolumeProtectionGroupObjectParams) *MSSQLVolumeProtectionGroupParamsAllOf {
	this := MSSQLVolumeProtectionGroupParamsAllOf{}
	this.Objects = objects
	return &this
}

// NewMSSQLVolumeProtectionGroupParamsAllOfWithDefaults instantiates a new MSSQLVolumeProtectionGroupParamsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMSSQLVolumeProtectionGroupParamsAllOfWithDefaults() *MSSQLVolumeProtectionGroupParamsAllOf {
	this := MSSQLVolumeProtectionGroupParamsAllOf{}
	return &this
}

// GetObjects returns the Objects field value
// If the value is explicit nil, the zero value for []MSSQLVolumeProtectionGroupObjectParams will be returned
func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetObjects() []MSSQLVolumeProtectionGroupObjectParams {
	if o == nil {
		var ret []MSSQLVolumeProtectionGroupObjectParams
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetObjectsOk() (*[]MSSQLVolumeProtectionGroupObjectParams, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value
func (o *MSSQLVolumeProtectionGroupParamsAllOf) SetObjects(v []MSSQLVolumeProtectionGroupObjectParams) {
	o.Objects = v
}

// GetIncrementalBackupAfterRestart returns the IncrementalBackupAfterRestart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetIncrementalBackupAfterRestart() bool {
	if o == nil || o.IncrementalBackupAfterRestart.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IncrementalBackupAfterRestart.Get()
}

// GetIncrementalBackupAfterRestartOk returns a tuple with the IncrementalBackupAfterRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetIncrementalBackupAfterRestartOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IncrementalBackupAfterRestart.Get(), o.IncrementalBackupAfterRestart.IsSet()
}

// HasIncrementalBackupAfterRestart returns a boolean if a field has been set.
func (o *MSSQLVolumeProtectionGroupParamsAllOf) HasIncrementalBackupAfterRestart() bool {
	if o != nil && o.IncrementalBackupAfterRestart.IsSet() {
		return true
	}

	return false
}

// SetIncrementalBackupAfterRestart gets a reference to the given NullableBool and assigns it to the IncrementalBackupAfterRestart field.
func (o *MSSQLVolumeProtectionGroupParamsAllOf) SetIncrementalBackupAfterRestart(v bool) {
	o.IncrementalBackupAfterRestart.Set(&v)
}
// SetIncrementalBackupAfterRestartNil sets the value for IncrementalBackupAfterRestart to be an explicit nil
func (o *MSSQLVolumeProtectionGroupParamsAllOf) SetIncrementalBackupAfterRestartNil() {
	o.IncrementalBackupAfterRestart.Set(nil)
}

// UnsetIncrementalBackupAfterRestart ensures that no value is present for IncrementalBackupAfterRestart, not even an explicit nil
func (o *MSSQLVolumeProtectionGroupParamsAllOf) UnsetIncrementalBackupAfterRestart() {
	o.IncrementalBackupAfterRestart.Unset()
}

// GetIndexingPolicy returns the IndexingPolicy field value if set, zero value otherwise.
func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetIndexingPolicy() IndexingPolicy {
	if o == nil || o.IndexingPolicy == nil {
		var ret IndexingPolicy
		return ret
	}
	return *o.IndexingPolicy
}

// GetIndexingPolicyOk returns a tuple with the IndexingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetIndexingPolicyOk() (*IndexingPolicy, bool) {
	if o == nil || o.IndexingPolicy == nil {
		return nil, false
	}
	return o.IndexingPolicy, true
}

// HasIndexingPolicy returns a boolean if a field has been set.
func (o *MSSQLVolumeProtectionGroupParamsAllOf) HasIndexingPolicy() bool {
	if o != nil && o.IndexingPolicy != nil {
		return true
	}

	return false
}

// SetIndexingPolicy gets a reference to the given IndexingPolicy and assigns it to the IndexingPolicy field.
func (o *MSSQLVolumeProtectionGroupParamsAllOf) SetIndexingPolicy(v IndexingPolicy) {
	o.IndexingPolicy = &v
}

// GetBackupDbVolumesOnly returns the BackupDbVolumesOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetBackupDbVolumesOnly() bool {
	if o == nil || o.BackupDbVolumesOnly.Get() == nil {
		var ret bool
		return ret
	}
	return *o.BackupDbVolumesOnly.Get()
}

// GetBackupDbVolumesOnlyOk returns a tuple with the BackupDbVolumesOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetBackupDbVolumesOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BackupDbVolumesOnly.Get(), o.BackupDbVolumesOnly.IsSet()
}

// HasBackupDbVolumesOnly returns a boolean if a field has been set.
func (o *MSSQLVolumeProtectionGroupParamsAllOf) HasBackupDbVolumesOnly() bool {
	if o != nil && o.BackupDbVolumesOnly.IsSet() {
		return true
	}

	return false
}

// SetBackupDbVolumesOnly gets a reference to the given NullableBool and assigns it to the BackupDbVolumesOnly field.
func (o *MSSQLVolumeProtectionGroupParamsAllOf) SetBackupDbVolumesOnly(v bool) {
	o.BackupDbVolumesOnly.Set(&v)
}
// SetBackupDbVolumesOnlyNil sets the value for BackupDbVolumesOnly to be an explicit nil
func (o *MSSQLVolumeProtectionGroupParamsAllOf) SetBackupDbVolumesOnlyNil() {
	o.BackupDbVolumesOnly.Set(nil)
}

// UnsetBackupDbVolumesOnly ensures that no value is present for BackupDbVolumesOnly, not even an explicit nil
func (o *MSSQLVolumeProtectionGroupParamsAllOf) UnsetBackupDbVolumesOnly() {
	o.BackupDbVolumesOnly.Unset()
}

// GetAdditionalHostParams returns the AdditionalHostParams field value if set, zero value otherwise.
func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetAdditionalHostParams() []MSSQLVolumeProtectionGroupHostParams {
	if o == nil || o.AdditionalHostParams == nil {
		var ret []MSSQLVolumeProtectionGroupHostParams
		return ret
	}
	return *o.AdditionalHostParams
}

// GetAdditionalHostParamsOk returns a tuple with the AdditionalHostParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetAdditionalHostParamsOk() (*[]MSSQLVolumeProtectionGroupHostParams, bool) {
	if o == nil || o.AdditionalHostParams == nil {
		return nil, false
	}
	return o.AdditionalHostParams, true
}

// HasAdditionalHostParams returns a boolean if a field has been set.
func (o *MSSQLVolumeProtectionGroupParamsAllOf) HasAdditionalHostParams() bool {
	if o != nil && o.AdditionalHostParams != nil {
		return true
	}

	return false
}

// SetAdditionalHostParams gets a reference to the given []MSSQLVolumeProtectionGroupHostParams and assigns it to the AdditionalHostParams field.
func (o *MSSQLVolumeProtectionGroupParamsAllOf) SetAdditionalHostParams(v []MSSQLVolumeProtectionGroupHostParams) {
	o.AdditionalHostParams = &v
}

func (o MSSQLVolumeProtectionGroupParamsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.IncrementalBackupAfterRestart.IsSet() {
		toSerialize["incrementalBackupAfterRestart"] = o.IncrementalBackupAfterRestart.Get()
	}
	if o.IndexingPolicy != nil {
		toSerialize["indexingPolicy"] = o.IndexingPolicy
	}
	if o.BackupDbVolumesOnly.IsSet() {
		toSerialize["backupDbVolumesOnly"] = o.BackupDbVolumesOnly.Get()
	}
	if o.AdditionalHostParams != nil {
		toSerialize["additionalHostParams"] = o.AdditionalHostParams
	}
	return json.Marshal(toSerialize)
}

type NullableMSSQLVolumeProtectionGroupParamsAllOf struct {
	value *MSSQLVolumeProtectionGroupParamsAllOf
	isSet bool
}

func (v NullableMSSQLVolumeProtectionGroupParamsAllOf) Get() *MSSQLVolumeProtectionGroupParamsAllOf {
	return v.value
}

func (v *NullableMSSQLVolumeProtectionGroupParamsAllOf) Set(val *MSSQLVolumeProtectionGroupParamsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMSSQLVolumeProtectionGroupParamsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMSSQLVolumeProtectionGroupParamsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMSSQLVolumeProtectionGroupParamsAllOf(val *MSSQLVolumeProtectionGroupParamsAllOf) *NullableMSSQLVolumeProtectionGroupParamsAllOf {
	return &NullableMSSQLVolumeProtectionGroupParamsAllOf{value: val, isSet: true}
}

func (v NullableMSSQLVolumeProtectionGroupParamsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMSSQLVolumeProtectionGroupParamsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}




func (o MSSQLVolumeProtectionGroupParamsAllOf) Print() {
		if byteArray, err := o.MarshalJSON(); err != nil {
				fmt.Println(err)
		} else {
				fmt.Println(string(byteArray))
		}
}